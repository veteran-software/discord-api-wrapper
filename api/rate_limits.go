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
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type rateLimitResponse struct {
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
	Global     bool    `json:"global"`
}

// customRateLimit holds information for defining a custom rate limit
type customRateLimit struct {
	suffix   string
	requests int
	reset    time.Duration
}

// RateLimiter holds all ratelimit buckets
type RateLimiter struct {
	sync.Mutex

	global           *int64
	buckets          map[string]*bucket
	customRateLimits []*customRateLimit
}

// bucket represents a ratelimit bucket, each bucket gets ratelimited individually (-global ratelimits)
type bucket struct {
	sync.Mutex

	Key             string
	Remaining       int
	reset           time.Time
	global          *int64
	lastReset       time.Time
	customRateLimit *customRateLimit
}

// NewRatelimiter returns a new RateLimiter
//
//goland:noinspection SpellCheckingInspection
func NewRatelimiter() *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]*bucket),
		global:  new(int64),
		customRateLimits: []*customRateLimit{
			{
				suffix:   "//reactions//",
				requests: 1,
				reset:    200 * time.Millisecond,
			},
		},
	}
}

// getBucket retrieves or creates a bucket
func (r *RateLimiter) getBucket(key string) *bucket {
	r.Lock()
	defer r.Unlock()

	if bucket, ok := r.buckets[key]; ok {
		return bucket
	}

	b := &bucket{
		Remaining: 1,
		Key:       key,
		global:    r.global,
	}

	for _, rateLimit := range r.customRateLimits {
		if strings.HasSuffix(b.Key, rateLimit.suffix) {
			b.customRateLimit = rateLimit
			break
		}
	}

	r.buckets[key] = b

	return b
}

// getWaitTime returns the duration you should wait for a bucket
func (r *RateLimiter) getWaitTime(b *bucket, minRemaining int) time.Duration {
	if b.Remaining < minRemaining && b.reset.After(time.Now()) {
		return time.Until(b.reset)
	}

	// Check global rate limits
	sleepTo := time.Unix(0, atomic.LoadInt64(r.global))
	if now := time.Now(); now.Before(sleepTo) {
		return sleepTo.Sub(now)
	}

	return 0
}

// lockBucketObject Locks an already resolved bucket until a request can be made
func (r *RateLimiter) lockBucketObject(b *bucket) *bucket {
	b.Lock()

	if wait := r.getWaitTime(b, 1); wait > 0 {
		time.Sleep(wait)
	}

	b.Remaining--

	return b
}

// lockBucket Locks until a request can be made
func (r *RateLimiter) lockBucket(bucketID string) *bucket {
	return r.lockBucketObject(r.getBucket(bucketID))
}

// release unlocks the bucket and reads the headers to update the buckets ratelimit info and locks up the whole thing in case if there's a global ratelimit.
func (b *bucket) release(headers http.Header) error {
	defer b.Unlock()

	if rl := b.customRateLimit; rl != nil {
		return b.checkCustomLimit(rl)
	}

	if headers == nil {
		return nil
	}

	remaining := headers.Get("X-RateLimit-Remaining")
	reset := headers.Get("X-RateLimit-Reset")
	global := headers.Get("X-RateLimit-Global")
	resetAfter := headers.Get("X-RateLimit-Reset-After")

	// Update global and per bucket reset time if the proper headers are available
	// If global is set, then it will block all buckets until after X-RateLimit-Reset-After
	// If Retry-After without global is provided it will use that for the new reset time since it's more accurate than X-RateLimit-Reset.
	// If Retry-After after is not provided, it will update the reset time from X-RateLimit-Reset
	if resetAfter != "" {
		err := b.checkResetAfter(resetAfter, global)
		if err != nil {
			return err
		}
	} else if reset != "" {
		err := b.checkReset(headers, reset)
		if err != nil {
			return err
		}
	}

	// Update remaining if header is present
	if remaining != "" {
		parsedRemaining, err := strconv.ParseInt(remaining, 10, 32)
		if err != nil {
			return err
		}

		b.Remaining = int(parsedRemaining)
	}

	return nil
}

func (b *bucket) checkCustomLimit(rl *customRateLimit) error {
	if time.Since(b.lastReset) >= rl.reset {
		b.Remaining = rl.requests - 1
		b.lastReset = time.Now()
	}

	if b.Remaining < 1 {
		b.reset = time.Now().Add(rl.reset)
	}

	return nil
}

func (b *bucket) checkReset(headers http.Header, reset string) error {
	discordTime, err := http.ParseTime(headers.Get("Date"))
	if err != nil {
		return err
	}

	unix, err := strconv.ParseFloat(reset, 64)
	if err != nil {
		return err
	}

	whole, frac := math.Modf(unix)
	delta := time.Unix(int64(whole), 0).Add(time.Duration(frac*1000)*time.Millisecond).Sub(discordTime) + (250 * time.Millisecond)
	b.reset = time.Now().Add(delta)
	return nil
}

func (b *bucket) checkResetAfter(resetAfter string, global string) error {
	parsedAfter, err := strconv.ParseFloat(resetAfter, 64)
	if err != nil {
		return err
	}

	whole, frac := math.Modf(parsedAfter)
	resetAt := time.Now().Add(time.Duration(whole) * time.Second).Add(time.Duration(frac*1000) * time.Millisecond)

	if global != "" {
		atomic.StoreInt64(b.global, resetAt.UnixNano())

	} else {
		b.reset = resetAt
	}
	return nil
}
