package utilities

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"DiscordAPIWrapper/api"
	"DiscordAPIWrapper/logging"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
)

var (
	initialTimeout        = 500 * time.Millisecond
	maxTimeout            = 25 * time.Second
	exponentFactor        = 2.0
	maximumJitterInterval = 2 * time.Millisecond
	retryCount            = 2

	backoff = heimdall.NewExponentialBackoff(initialTimeout, maxTimeout, exponentFactor, maximumJitterInterval)

	// Create a new retry mechanism with the backoff
	retrier = heimdall.NewRetrier(backoff)

	httpClient = httpclient.NewClient(
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(retryCount),
	)
)

var (
	Rest *RateLimiter
)

func (r *RateLimiter) Request(method, route string, data interface{}, reason *string) (*http.Response, error) {
	return r.requestWithBucketID(method, route, data, strings.SplitN(route, "?", 2)[0], reason)
}

func (r *RateLimiter) requestWithBucketID(method, route string, data interface{}, bucketID string, reason *string) (*http.Response, error) {
	return r.request(method, route, "application/json", data, bucketID, 0, reason)
}

func (r *RateLimiter) request(method, route, contentType string, b interface{}, bucketID string, sequence int, reason *string) (*http.Response, error) {
	if bucketID == "" {
		bucketID = strings.SplitN(route, "?", 2)[0]
	}

	return r.requestWithLockedBucket(method, route, contentType, b, r.lockBucket(bucketID), sequence, reason)
}

func (r *RateLimiter) requestWithLockedBucket(method, route, contentType string, b interface{}, bucket *bucket, sequence int, reason *string) (*http.Response, error) {
	var buffer bytes.Buffer
	if b != nil {
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&b)
		if err != nil {
			_ = bucket.release(nil)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, route, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		_ = bucket.release(nil)
		return nil, err
	}

	req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bot %s", Token))

	if b != nil {
		req.Header.Set(http.CanonicalHeaderKey("Content-Type"), contentType)
	}

	if reason != nil {
		req.Header.Set("X-Audit-Log-Reason", *reason)
	}

	req.Header.Set(http.CanonicalHeaderKey("User-Agent"), api.UserAgent)

	ctx, cancel := context.WithDeadline(req.Context(), time.Now().Add(12*time.Second))
	defer cancel()

	resp, err := httpClient.Do(req.WithContext(ctx))
	if err != nil {
		_ = bucket.release(nil)
		if errors.Is(err, context.Canceled) {
			logging.Warnln("Context cancelled. Deadline was 12 seconds.")
			logging.Warnln("\tRequest was : ", method, " : ", route)
		} else if errors.Is(err, context.DeadlineExceeded) {
			logging.Warnln("Context deadline exceeded.")
			logging.Warnln("\tRequest was : ", method, " : ", route)
		}
		return nil, err
	}

	err = bucket.release(resp.Header)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusBadGateway:
	case http.StatusTooManyRequests:
		logging.Warnln("Rate Limited!")
		logging.Infoln(route)
		var rlr rateLimitResponse
		err = json.NewDecoder(resp.Body).Decode(&rlr)
		if err != nil {
			return nil, err
		}

		time.Sleep(rlr.RetryAfter)

		resp, err = r.requestWithLockedBucket(method, route, contentType, b, r.lockBucketObject(bucket), sequence, reason)
	}

	return resp, nil
}
