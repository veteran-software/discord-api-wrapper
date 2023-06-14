/*
 * Copyright (c) 2023. Veteran Software
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
	"net/http"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/veteran-software/discord-api-wrapper/v10/utilities"
)

func TestNewRatelimiter(t *testing.T) {
	tests := []struct {
		name string
		want *RateLimiter
	}{
		{
			name: "Basic",
			want: &RateLimiter{
				buckets: make(map[string]*bucket),
				global:  new(int64),
				customRateLimits: []*customRateLimit{
					{
						suffix:   "//reactions//",
						requests: 1,
						reset:    200 * time.Millisecond,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRatelimiter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRatelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateLimiter_getBucket(t *testing.T) {
	type fields struct {
		Mutex            *sync.Mutex
		global           *int64
		buckets          map[string]*bucket
		customRateLimits []*customRateLimit
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bucket
	}{
		{
			name: "Bucket Exists",
			fields: fields{
				Mutex:  new(sync.Mutex),
				global: new(int64),
				buckets: map[string]*bucket{"123": {
					Key:       "123",
					Remaining: 1,
				}},
				customRateLimits: nil,
			},
			args: args{
				key: "123",
			},
			want: &bucket{
				Key:       "123",
				Remaining: 1,
			},
		},
		{
			name: "Bucket Does Not Exist",
			fields: fields{
				Mutex:  new(sync.Mutex),
				global: new(int64),
				buckets: map[string]*bucket{"123": {
					Key:       "123",
					Remaining: 1,
				}},
				customRateLimits: nil,
			},
			args: args{
				key: "456",
			},
			want: &bucket{
				Key:       "456",
				Remaining: 1,
				global:    new(int64),
			},
		},
		{
			name: "Bucket Does Not Exist; Custom Rate Limit; Suffix Found",
			fields: fields{
				Mutex:  new(sync.Mutex),
				global: new(int64),
				buckets: map[string]*bucket{"123": {
					Key:       "123//monkey//",
					Remaining: 1,
				}},
				customRateLimits: []*customRateLimit{
					{
						suffix:   "//monkey//",
						requests: 1,
						reset:    200 * time.Millisecond,
					},
					{
						suffix:   "//potato//",
						requests: 2,
						reset:    400 * time.Millisecond,
					},
				},
			},
			args: args{
				key: "123//monkey//",
			},
			want: &bucket{
				Key:       "123//monkey//",
				Remaining: 1,
				global:    new(int64),
				customRateLimit: &customRateLimit{
					suffix:   "//monkey//",
					requests: 1,
					reset:    200 * time.Millisecond,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RateLimiter{
				Mutex:            sync.Mutex{},
				global:           tt.fields.global,
				buckets:          tt.fields.buckets,
				customRateLimits: tt.fields.customRateLimits,
			}
			if got := r.getBucket(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateLimiter_getWaitTime(t *testing.T) {
	type fields struct {
		Mutex            *sync.Mutex
		global           *int64
		buckets          map[string]*bucket
		customRateLimits []*customRateLimit
	}
	type args struct {
		b            *bucket
		minRemaining int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		// Not checking the first 2 return points because it would require a bit of a rewrite to get reliable test results
		{
			name: "Last return point",
			fields: fields{
				global: utilities.ToPtr(int64(1)),
			},
			args: args{
				b: &bucket{
					Remaining: 5,
					reset:     time.Now(),
				},
				minRemaining: 4,
			},
			want: 0 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RateLimiter{
				global:           tt.fields.global,
				buckets:          tt.fields.buckets,
				customRateLimits: tt.fields.customRateLimits,
			}
			if got := r.getWaitTime(tt.args.b, tt.args.minRemaining); got != tt.want {
				t.Errorf("getWaitTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateLimiter_lockBucket(t *testing.T) {
	type fields struct {
		Mutex            *sync.Mutex
		global           *int64
		buckets          map[string]*bucket
		customRateLimits []*customRateLimit
	}
	type args struct {
		bucketID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Basic",
			fields: fields{
				Mutex:  &sync.Mutex{},
				global: utilities.ToPtr(time.Now().Add(1 * time.Second).Unix()),
				buckets: map[string]*bucket{
					"123": {
						Remaining: 2,
					},
				},
			},
			args: args{"123"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RateLimiter{
				Mutex:            sync.Mutex{},
				global:           tt.fields.global,
				buckets:          tt.fields.buckets,
				customRateLimits: tt.fields.customRateLimits,
			}

			got := r.lockBucket(tt.args.bucketID)
			state := reflect.ValueOf(&got.Mutex).Elem().FieldByName("state").Int()&1 == 1
			if state != tt.want {
				t.Errorf("lockBucket() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRateLimiter_lockBucketObject(t *testing.T) {
	type fields struct {
		Mutex            *sync.Mutex
		global           *int64
		buckets          map[string]*bucket
		customRateLimits []*customRateLimit
	}
	type args struct {
		b *bucket
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Basic",
			fields: fields{
				Mutex:  &sync.Mutex{},
				global: utilities.ToPtr(time.Now().Add(1 * time.Second).Unix()),
				buckets: map[string]*bucket{
					"123": {
						Remaining: 2,
					},
				},
			},
			args: args{
				&bucket{
					Remaining: 2,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RateLimiter{
				global:           tt.fields.global,
				buckets:          tt.fields.buckets,
				customRateLimits: tt.fields.customRateLimits,
			}
			got := r.lockBucketObject(tt.args.b)
			state := reflect.ValueOf(&got.Mutex).Elem().FieldByName("state").Int()&1 == 1
			if state != tt.want {
				t.Errorf("lockBucket() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_bucket_checkCustomLimit(t *testing.T) {
	type fields struct {
		Remaining int
		reset     time.Time
		lastReset time.Time
	}
	type args struct {
		rl *customRateLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Hit All Conditionals",
			fields: fields{
				Remaining: 0,
				lastReset: time.Now().Add(1 * time.Second),
			},
			args: args{
				rl: &customRateLimit{
					requests: 1,
					reset:    -(2 * time.Second),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bucket{
				Remaining: tt.fields.Remaining,
				reset:     tt.fields.reset,
				lastReset: tt.fields.lastReset,
			}
			if err := b.checkCustomLimit(tt.args.rl); (err != nil) != tt.wantErr {
				t.Errorf("checkCustomLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bucket_checkReset(t *testing.T) {
	type fields struct {
		reset time.Time
	}
	type args struct {
		headers http.Header
		reset   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "ParseTime error",
			fields: fields{},
			args: args{
				headers: map[string][]string{"Date": {"Not A Valid Date"}},
				reset:   "",
			},
			wantErr: true,
		},
		{
			name:   "ParseFloat error",
			fields: fields{},
			args: args{
				headers: map[string][]string{"Date": {"Mon, 12 Jun 2023 03:33:35 GMT"}},
				reset:   "Not A Valid Float",
			},
			wantErr: true,
		},
		{
			name:   "Valid Reset",
			fields: fields{},
			args: args{
				headers: map[string][]string{"Date": {"Mon, 12 Jun 2023 03:33:35 GMT"}},
				reset:   "25",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bucket{
				reset: tt.fields.reset,
			}
			if err := b.checkReset(tt.args.headers, tt.args.reset); (err != nil) != tt.wantErr {
				t.Errorf("checkReset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bucket_checkResetAfter(t *testing.T) {
	type fields struct {
		reset  time.Time
		global *int64
	}
	type args struct {
		resetAfter string
		global     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "ParseFloat error",
			fields: fields{},
			args: args{
				resetAfter: "Not a Valid Float",
				global:     "",
			},
			wantErr: true,
		},
		{
			name:   "Empty global",
			fields: fields{},
			args: args{
				resetAfter: "5",
				global:     "",
			},
			wantErr: false,
		},
		{
			name: "Non-empty global",
			fields: fields{
				global: utilities.ToPtr(int64(5)),
			},
			args: args{
				resetAfter: "5.5",
				global:     "100",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bucket{
				reset:  tt.fields.reset,
				global: tt.fields.global,
			}
			if err := b.checkResetAfter(tt.args.resetAfter, tt.args.global); (err != nil) != tt.wantErr {
				t.Errorf("checkResetAfter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bucket_release(t *testing.T) {
	type fields struct {
		Mutex           *sync.Mutex
		Key             string
		Remaining       int
		reset           time.Time
		global          *int64
		lastReset       time.Time
		customRateLimit *customRateLimit
	}
	type args struct {
		headers http.Header
	}

	headersThree := http.Header{}
	headersThree.Add("X-RateLimit-Reset-After", "Potato")
	headersThree.Add("X-RateLimit-Global", "")

	headersFour := http.Header{}
	headersFour.Add("X-RateLimit-Reset-After", "500")
	headersFour.Add("X-RateLimit-Global", "")

	headersFive := http.Header{}
	headersFive.Add("X-RateLimit-Reset", "500")
	headersFive.Add("X-RateLimit-Global", "")
	headersFive.Add("Date", "Potato")

	headersSix := http.Header{}
	headersSix.Add("X-RateLimit-Reset", "500")
	headersSix.Add("X-RateLimit-Global", "")
	headersSix.Add("Date", "Mon, 12 Jun 2023 03:33:35 GMT")
	headersSix.Add("X-RateLimit-Remaining", "Potato")

	headersSeven := http.Header{}
	headersSeven.Add("X-RateLimit-Reset", "500")
	headersSeven.Add("X-RateLimit-Global", "")
	headersSeven.Add("Date", "Mon, 12 Jun 2023 03:33:35 GMT")
	headersSeven.Add("X-RateLimit-Remaining", "4")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Custom Rate Limit set",
			fields: fields{
				Mutex:           new(sync.Mutex),
				customRateLimit: &customRateLimit{},
			},
			args:    args{headers: map[string][]string{}},
			wantErr: false,
		},
		{
			name: "nil Headers",
			fields: fields{
				Mutex: new(sync.Mutex),
			},
			args:    args{headers: nil},
			wantErr: false,
		},
		{
			name: "3) 'resetAfter' not empty with error",
			fields: fields{
				Mutex: new(sync.Mutex),
			},
			args: args{
				headers: headersThree,
			},
			wantErr: true,
		},
		{
			name: "4) 'resetAfter' not empty no error",
			fields: fields{
				Mutex: new(sync.Mutex),
			},
			args: args{
				headers: headersFour,
			},
			wantErr: false,
		},
		{
			name: "5) 'reset' not empty with error",
			fields: fields{
				Mutex: new(sync.Mutex),
			},
			args: args{
				headers: headersFive,
			},
			wantErr: true,
		},
		{
			name: "7) 'remaining' not empty with error",
			fields: fields{
				Mutex: new(sync.Mutex),
			},
			args: args{
				headers: headersSix,
			},
			wantErr: true,
		},
		{
			name: "7) 'remaining' not empty no error",
			fields: fields{
				Mutex: new(sync.Mutex),
			},
			args: args{
				headers: headersSeven,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bucket{
				Key:             tt.fields.Key,
				Remaining:       tt.fields.Remaining,
				reset:           tt.fields.reset,
				global:          tt.fields.global,
				lastReset:       tt.fields.lastReset,
				customRateLimit: tt.fields.customRateLimit,
			}

			b.Lock()

			if err := b.release(tt.args.headers); (err != nil) != tt.wantErr {
				t.Errorf("release() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
