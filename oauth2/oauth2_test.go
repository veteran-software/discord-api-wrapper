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

package oauth2

import (
	"net/url"
	"reflect"
	"testing"
)

func TestBaseAuthorizationURL(t *testing.T) {
	tests := []struct {
		name string
		want *url.URL
	}{
		{
			name: "No Error",
			want: &url.URL{
				Scheme: "https",
				Host:   "discord.com",
				Path:   "/api/oauth2/authorize",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseAuthorizationURL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseAuthorizationURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenRevocationURL(t *testing.T) {
	tests := []struct {
		name string
		want *url.URL
	}{
		{
			name: "No Error",
			want: &url.URL{
				Scheme: "https",
				Host:   "discord.com",
				Path:   "/api/oauth2/token/revoke",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenRevocationURL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenRevocationURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenURL(t *testing.T) {
	tests := []struct {
		name string
		want *url.URL
	}{
		{
			name: "No Error",
			want: &url.URL{
				Scheme: "https",
				Host:   "discord.com",
				Path:   "/api/oauth2/token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenURL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
