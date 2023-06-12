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

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(want)
		if err != nil {
			return
		}
		_, _ = w.Write(b.Bytes())
	}))
	defer srv.Close()

	api = srv.URL
	testClient = srv.Client()

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
