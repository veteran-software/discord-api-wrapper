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

	"github.com/veteran-software/discord-api-wrapper/v10/utilities"
)

func TestGuild_CreateGuildScheduledEvent(t *testing.T) {
	want := &GuildScheduledEvent{
		ID:      "123456789",
		GuildID: "456789123",
		Name:    "Some Event",
	}

	type fields struct {
		ID Snowflake
	}
	type args struct {
		payload *CreateGuildScheduledEventJSON
		reason  *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GuildScheduledEvent
		wantErr bool
	}{
		{
			name:   "No Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				payload: &CreateGuildScheduledEventJSON{},
				reason:  utilities.ToPtr("Just Because"),
			},
			want:    want,
			wantErr: false,
		},
		{
			name:   "Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				payload: &CreateGuildScheduledEventJSON{},
				reason:  utilities.ToPtr("Just Because"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}

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

			got, err := g.CreateGuildScheduledEvent(tt.args.payload, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGuildScheduledEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGuildScheduledEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuild_DeleteGuildScheduledEvent(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	type args struct {
		guildScheduledEventID *Snowflake
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "No Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("456789123"),
			},
			wantErr: false,
		},
		{
			name:   "Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("456789123"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}

			switch tt.name {
			case "No Error":
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(204)
					_, _ = w.Write(nil)
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

			if err := g.DeleteGuildScheduledEvent(tt.args.guildScheduledEventID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGuildScheduledEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGuild_GetGuildScheduledEvent(t *testing.T) {
	want := &GuildScheduledEvent{
		ID:      "123456789",
		GuildID: "456789123",
		Name:    "Some Event",
	}

	type fields struct {
		ID Snowflake
	}
	type args struct {
		guildScheduledEventID *Snowflake
		withUserCount         *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GuildScheduledEvent
		wantErr bool
	}{
		{
			name:   "No Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("456789123"),
				withUserCount:         utilities.ToPtr(true),
			},
			want:    want,
			wantErr: false,
		},
		{
			name:   "Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("456789123"),
				withUserCount:         utilities.ToPtr(true),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}

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

			got, err := g.GetGuildScheduledEvent(tt.args.guildScheduledEventID, tt.args.withUserCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGuildScheduledEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGuildScheduledEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuild_GetGuildScheduledEventUsers(t *testing.T) {
	want := []*GuildScheduledEventUser{
		{
			GuildScheduledEventID: "123456789",
			User:                  User{ID: "789456123"},
			Member:                GuildMember{Nick: utilities.ToPtr("Kappatato")},
		},
		{
			GuildScheduledEventID: "123456789",
			User:                  User{ID: "456789123"},
			Member:                GuildMember{Nick: utilities.ToPtr("Not Live")},
		},
	}

	type fields struct {
		ID Snowflake
	}
	type args struct {
		guildScheduledEventID *Snowflake
		limit                 *uint64
		withMember            *bool
		before                *Snowflake
		after                 *Snowflake
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*GuildScheduledEventUser
		wantErr bool
	}{
		{
			name:   "No Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("123456789"),
				limit:                 utilities.ToPtr(uint64(5)),
				withMember:            utilities.ToPtr(true),
				before:                StringToSnowflake("741258963"),
				after:                 StringToSnowflake("99874152630"),
			},
			want:    want,
			wantErr: false,
		},
		{
			name:   "Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("456789123"),
				limit:                 utilities.ToPtr(uint64(5)),
				withMember:            utilities.ToPtr(true),
				before:                StringToSnowflake("741258963"),
				after:                 StringToSnowflake("99874152630"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Error (limit too large)",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: StringToSnowflake("456789123"),
				limit:                 utilities.ToPtr(uint64(500)),
				withMember:            utilities.ToPtr(true),
				before:                StringToSnowflake("741258963"),
				after:                 StringToSnowflake("99874152630"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}
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
			case "Error (limit too large)":
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					_, _ = w.Write(nil)
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

			got, err := g.GetGuildScheduledEventUsers(tt.args.guildScheduledEventID,
				tt.args.limit,
				tt.args.withMember,
				tt.args.before,
				tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGuildScheduledEventUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGuildScheduledEventUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuild_ListGuildScheduledEvents(t *testing.T) {
	want := []*GuildScheduledEvent{
		{
			ID:        "123456789",
			Name:      "Potato Event",
			ChannelID: StringToSnowflake("741258963"),
		},
		{
			ID:        "123456789",
			Name:      "Not a Stream Event",
			ChannelID: StringToSnowflake("321789654"),
		},
	}
	type fields struct {
		ID Snowflake
	}
	type args struct {
		withUserCount *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*GuildScheduledEvent
		wantErr bool
	}{
		{
			name:   "No Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				withUserCount: utilities.ToPtr(true),
			},
			want:    want,
			wantErr: false,
		},
		{
			name:   "Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				withUserCount: utilities.ToPtr(true),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}
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

			got, err := g.ListGuildScheduledEvents(tt.args.withUserCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListGuildScheduledEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListGuildScheduledEvents() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuild_ModifyGuildScheduledEvent(t *testing.T) {
	want := &GuildScheduledEvent{
		ID:      "123456789",
		GuildID: "456789123",
		Name:    "Some Event",
	}

	type fields struct {
		ID Snowflake
	}
	type args struct {
		guildScheduledEventID Snowflake
		payload               *ModifyGuildScheduledEventJSON
		reason                *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GuildScheduledEvent
		wantErr bool
	}{
		{
			name:   "No Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: Snowflake("123456789"),
				payload:               &ModifyGuildScheduledEventJSON{},
				reason:                utilities.ToPtr("Just Because"),
			},
			want:    want,
			wantErr: false,
		},
		{
			name:   "Error",
			fields: fields{ID: Snowflake("123456789")},
			args: args{
				guildScheduledEventID: Snowflake("123456789"),
				payload:               &ModifyGuildScheduledEventJSON{},
				reason:                utilities.ToPtr("Just Because"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}
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

			got, err := g.ModifyGuildScheduledEvent(tt.args.guildScheduledEventID, tt.args.payload, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModifyGuildScheduledEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModifyGuildScheduledEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
