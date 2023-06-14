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
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetGateway(t *testing.T) {
	want := &GetGatewayResponse{
		Url: "wss://gateway.discord.gg",
	}

	tests := []struct {
		name    string
		want    *GetGatewayResponse
		wantErr bool
	}{
		{
			name:    "No Error",
			want:    want,
			wantErr: false,
		},
		{
			name:    "Error",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "No Error":
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)

					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(want)
					if err != nil {
						return
					}
					_, _ = w.Write(b.Bytes())
				}))

				api = srv.URL
				testClient = srv.Client()
				defer srv.Close()
			case "Error":
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write(nil)
				}))

				api = srv.URL
				testClient = srv.Client()
				defer srv.Close()
			}

			got, err := GetGateway()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGateway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGateway() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGatewayBot(t *testing.T) {
	want := &GetGatewayBotResponse{
		Url: "wss://gateway.discord.gg",
		SessionStartLimit: struct {
			Total          int `json:"total"`
			Remaining      int `json:"remaining"`
			ResetAfter     int `json:"reset_after"`
			MaxConcurrency int `json:"max_concurrency"`
		}(struct {
			Total          int
			Remaining      int
			ResetAfter     int
			MaxConcurrency int
		}{
			Total: 1000, Remaining: 1000, ResetAfter: 0, MaxConcurrency: 1}),
		Shards: 12,
	}

	tests := []struct {
		name    string
		want    *GetGatewayBotResponse
		wantErr bool
	}{
		{
			name:    "No Error",
			want:    want,
			wantErr: false,
		},
		{
			name:    "Error",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "No Error":
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)

					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(want)
					if err != nil {
						return
					}
					_, _ = w.Write(b.Bytes())
				}))

				api = srv.URL
				testClient = srv.Client()
				defer srv.Close()
			case "Error":
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write(nil)
				}))

				api = srv.URL
				testClient = srv.Client()
				defer srv.Close()
			}

			got, err := GetGatewayBot()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGatewayBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGatewayBot() got = %v, want %v", got, tt.want)
			}
		})
	}
}
