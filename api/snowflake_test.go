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
	"reflect"
	"testing"
	"time"
)

func TestSnowflake_ParseSnowflake(t *testing.T) {
	tests := []struct {
		name string
		s    Snowflake
		want FormattedSnowflake
	}{
		{
			name: "ValidSnowflake",
			s:    Snowflake("123456789123456"),
			want: FormattedSnowflake{
				Timestamp:         4841881341367,
				InternalWorkerID:  28,
				InternalProcessID: 14,
				Increment:         1820,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ParseSnowflake(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSnowflake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_String(t *testing.T) {
	tests := []struct {
		name string
		s    Snowflake
		want string
	}{
		{
			name: "ValidSnowflake",
			s:    Snowflake("123456789123456"),
			want: "123456789123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_Timestamp(t *testing.T) {
	tests := []struct {
		name string
		s    Snowflake
		want time.Time
	}{
		{
			name: "ValidSnowflake",
			s:    Snowflake("123456789123456"),
			want: time.Unix(0, 4841881341367),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Timestamp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		s    Snowflake
		want string
	}{
		{
			name: "ValidSnowflake",
			s:    Snowflake("123456789123456"),
			want: "110001110010110011110100110101110110110111111000111001110001110010110011110100110101110110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToBinary(); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToSnowflake(t *testing.T) {
	type args struct {
		s string
	}

	a := Snowflake("123456789123456")

	tests := []struct {
		name string
		args args
		want *Snowflake
	}{
		{
			name: "ValidSnowflake",
			args: args{s: "123456789123456"},
			want: &a,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToSnowflake(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToSnowflake() = %v, want %v", got, tt.want)
			}
		})
	}
}
