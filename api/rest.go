/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/veteran-software/discord-api-wrapper/v10/logging"
)

//goland:noinspection SpellCheckingInspection
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
	// Rest - Holds the rate limit buckets
	Rest *RateLimiter
)

func init() {
	Rest = NewRatelimiter()
}

// Request - send an HTTP request with rate limiting
func (r *RateLimiter) Request(method, route string, data interface{}, reason *string) (*http.Response, error) {
	return r.requestWithBucketID(method, route, strings.SplitN(route, "?", 2)[0], data, reason)
}

func (r *RateLimiter) requestWithBucketID(method, route, bucketID string, data interface{}, reason *string) (*http.Response, error) {
	return r.request(method, route, "application/json", data, bucketID, 0, reason)
}

func (r *RateLimiter) request(method, route, contentType string, b interface{}, bucketID string, sequence int, reason *string) (*http.Response, error) {
	if bucketID == "" {
		bucketID = strings.SplitN(route, "?", 2)[0]
	}

	return r.lockedRequest(method, route, contentType, b, r.lockBucket(bucketID), sequence, reason)
}

func processBody(b interface{}, bucket *bucket) (*bytes.Buffer, error) {
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

	return &buffer, nil
}

func (r *RateLimiter) lockedRequest(method, route, contentType string, b interface{}, bucket *bucket, sequence int, reason *string) (*http.Response, error) {
	buffer, err := processBody(b, bucket)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, route, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		_ = bucket.release(nil)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", Token))

	if b != nil {
		req.Header.Set("Content-Type", contentType)
	}

	if reason != nil {
		req.Header.Set("X-Audit-Log-Reason", *reason)
	}

	req.Header.Set("User-Agent", UserAgent)

	ctx, cancel := context.WithDeadline(req.Context(), time.Now().Add(5*time.Second))
	handleContextCancel(ctx, cancel)

	resp, err := httpClient.Do(req.WithContext(ctx))

	if err != nil {
		_ = bucket.release(nil)
		return nil, err
	}

	err = bucket.release(resp.Header)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		logging.Warnln("Rate Limited!")
		logging.Infoln(route)
		logging.Infoln(resp.Status)

		var rlr rateLimitResponse
		err = json.NewDecoder(resp.Body).Decode(&rlr)
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Duration(rlr.RetryAfter * float64(time.Second)))

		return r.lockedRequest(method, route, contentType, b, r.lockBucketObject(bucket), sequence, reason)
	}

	return resp, nil
}

func handleContextCancel(ctx context.Context, cancel context.CancelFunc) {
	go func(ctx context.Context) {
		defer cancel()

		<-ctx.Done()

		switch ctx.Err() {
		case context.DeadlineExceeded:
			logging.Traceln("context timeout exceeded")
		case context.Canceled:
			logging.Traceln("context cancelled; process complete")
		}
	}(ctx)
}

func parseRoute(route string) *url.URL {
	u, err := url.Parse(route)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return u
}

func fireGetRequest(u *url.URL, data *interface{}, reason *string) []byte {
	resp, err := Rest.Request(http.MethodGet, u.String(), data, reason)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Errorln(err)
		return []byte{} // we return an empty byte slice here to avoid nil pointer problems
	}

	return b
}

func firePostRequest(u *url.URL, data interface{}, reason *string) []byte {
	resp, err := Rest.Request(http.MethodPost, u.String(), data, reason)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Errorln(err)
		return []byte{} // we return an empty byte slice here to avoid nil pointer problems
	}

	return b
}

func firePutRequest(u *url.URL, data interface{}, reason *string) []byte {
	resp, err := Rest.Request(http.MethodPut, u.String(), data, reason)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Errorln(err)
		return []byte{} // we return an empty byte slice here to avoid nil pointer problems
	}

	return b
}

func firePatchRequest(u *url.URL, data interface{}, reason *string) []byte {
	resp, err := Rest.Request(http.MethodPatch, u.String(), data, reason)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Errorln(err)
		return []byte{} // we return an empty byte slice here to avoid nil pointer problems
	}

	return b
}

func fireDeleteRequest(u *url.URL, reason *string) error {
	resp, err := Rest.Request(http.MethodDelete, u.String(), nil, reason)
	if err != nil {
		logging.Errorln(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}
