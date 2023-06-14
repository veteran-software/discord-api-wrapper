/*
 * Copyright (c) 2022-2023. Veteran Software
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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/veteran-software/nowlive-logging"
)

var (
	// Rest - Holds the rate limit buckets
	Rest       *RateLimiter
	testClient *http.Client
)

type httpData struct {
	route *url.URL

	contentType string
	data        any
	method      string
	reason      *string

	bucket   *bucket
	bucketID string
	sequence int
}

func init() {
	Rest = NewRatelimiter()
}

// Request - send an HTTP request with rate limiting
func (r *RateLimiter) Request(h *httpData) (*http.Response, error) {
	h.bucketID = strings.SplitN(h.route.String(), "?", 2)[0]
	return r.requestWithBucketID(h)
}

func (r *RateLimiter) requestWithBucketID(h *httpData) (*http.Response, error) {
	h.contentType = "application/json"
	h.sequence = 0
	return r.request(h)
}

func (r *RateLimiter) request(h *httpData) (*http.Response, error) {
	if h.bucketID == "" {
		h.bucketID = strings.SplitN(h.route.String(), "?", 2)[0]
	}

	h.bucket = r.lockBucket(h.bucketID)

	return r.lockedRequest(h)
}

func processBody(h *httpData) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	if h.data != nil {
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&h.data)
		if err != nil {
			_ = h.bucket.release(nil)
			return nil, err
		}
	}

	return &buffer, nil
}

func (r *RateLimiter) lockedRequest(h *httpData) (*http.Response, error) {

	buffer, err := processBody(h)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(h.method, h.route.String(), bytes.NewReader(buffer.Bytes()))
	if err != nil {
		_ = h.bucket.release(nil)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", Token))

	if h.data != nil {
		req.Header.Set("Content-Type", h.contentType)
	}

	if h.reason != nil {
		req.Header.Set("X-Audit-Log-Reason", *h.reason)
	}

	req.Header.Set("User-Agent", UserAgent)

	var client *http.Client
	if testClient == nil {
		client = &http.Client{}
	} else {
		client = testClient
	}

	resp, err := client.Do(req)

	if err != nil {
		_ = h.bucket.release(nil)
		return nil, err
	}

	err = h.bucket.release(resp.Header)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		log.Warnln(log.FuncName(), "Rate Limited!")
		log.Infoln(log.FuncName(), h.route)
		log.Infoln(log.FuncName(), resp.Status)

		var rlr rateLimitResponse
		err = json.NewDecoder(resp.Body).Decode(&rlr)
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Duration(rlr.RetryAfter * float64(time.Second)))

		return r.lockedRequest(h)
	case http.StatusBadRequest: // 400
		fallthrough
	case http.StatusUnauthorized: // 401
		fallthrough
	case http.StatusForbidden: // 403
		fallthrough
	case http.StatusNotFound: // 404
		fallthrough
	case http.StatusMethodNotAllowed: // 405
		fallthrough
	case http.StatusBadGateway: // 502
		fallthrough
	case http.StatusServiceUnavailable: // 503
		fallthrough
	case http.StatusGatewayTimeout: // 504
		fallthrough
	case http.StatusInternalServerError: // 500
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New("HTTP Request Error : " + strconv.Itoa(resp.StatusCode) + " : " + resp.Status + " : " + string(b))
	}

	return resp, nil
}

func parseRoute(route string) *url.URL {
	u, err := url.Parse(route)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return nil
	}

	return u
}

func fireGetRequest(h *httpData) ([]byte, error) {
	h.method = http.MethodGet

	resp, err := Rest.Request(h)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return []byte{}, err // we return an empty byte slice here to avoid nil pointer problems
	}

	return b, nil
}

func firePostRequest(h *httpData) ([]byte, error) {
	h.method = http.MethodPost

	resp, err := Rest.Request(h)
	if err != nil {
		// Allow this log to bubble up to the method call
		log.Debugln(log.Discord, log.FuncName(), err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return []byte{}, err // we return an empty byte slice here to avoid nil pointer problems
	}

	return b, nil
}

//goland:noinspection GoUnusedFunction
func firePutRequest(h *httpData) ([]byte, error) {
	h.method = http.MethodPut

	resp, err := Rest.Request(h)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return []byte{}, err // we return an empty byte slice here to avoid nil pointer problems
	}

	return b, nil
}

func firePatchRequest(h *httpData) ([]byte, error) {
	h.method = http.MethodPatch

	resp, err := Rest.Request(h)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return []byte{}, err // we return an empty byte slice here to avoid nil pointer problems
	}

	return b, nil
}

func fireDeleteRequest(h *httpData) error {
	h.method = http.MethodDelete
	resp, err := Rest.Request(h)
	if err != nil {
		log.Errorln(log.FuncName(), err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}
