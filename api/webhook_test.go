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
	"net/http"
	"testing"
)

const (
	whIDToken                   = "https://discord.com/api/v10/webhooks/905130195520983061/fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"
	whIdTokenMessagesId         = "https://discord.com/api/v10/webhooks/905130195520983061/fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt/messages/148336120936005632"
	whIdTokenMessagesIdThreadId = "https://discord.com/api/v10/webhooks/905130195520983061/fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt/messages/148336120936005632?thread_id=934478965031174194"
)

var w = true

func TestChannelCreateWebhook(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Create Webhook",
			fields: fields{ID: Snowflake("753228874875273256")},
			want:   http.MethodPost,
			want1:  "https://discord.com/api/v10/channels/753228874875273256/webhooks",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Channel{
				ID: tt.fields.ID,
			}
			got, got1 := c.CreateWebhook()
			if got != tt.want {
				t.Errorf("CreateWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestChannelGetChannelWebhooks(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Get Channel Webhooks",
			fields: fields{ID: Snowflake("753228874875273256")},
			want:   http.MethodGet,
			want1:  "https://discord.com/api/v10/channels/753228874875273256/webhooks",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Channel{
				ID: tt.fields.ID,
			}
			got, got1 := c.GetChannelWebhooks()
			if got != tt.want {
				t.Errorf("GetChannelWebhooks() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetChannelWebhooks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGuildGetGuildWebhooks(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Get Guild Webhooks",
			fields: fields{ID: Snowflake("250045505659207699")},
			want:   http.MethodGet,
			want1:  "https://discord.com/api/v10/guilds/250045505659207699/webhooks",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guild{
				ID: tt.fields.ID,
			}
			got, got1 := g.GetGuildWebhooks()
			if got != tt.want {
				t.Errorf("GetGuildWebhooks() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetGuildWebhooks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookDeleteWebhook(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Delete Webhook",
			fields: fields{ID: Snowflake("905130195520983061")},
			want:   http.MethodDelete,
			want1:  "https://discord.com/api/v10/webhooks/905130195520983061",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID: tt.fields.ID,
			}
			got, got1 := w.DeleteWebhook()
			if got != tt.want {
				t.Errorf("DeleteWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DeleteWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookDeleteWebhookMessage(t *testing.T) {
	type fields struct {
		ID    Snowflake
		Token string
	}
	type args struct {
		msgID    Snowflake
		threadID *Snowflake
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		{
			name:   "Delete Webhook Message : Nil Thread ID",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args:   args{msgID: Snowflake("148336120936005632"), threadID: nil},
			want:   http.MethodDelete,
			want1:  whIdTokenMessagesId,
		},
		{
			name:   "Delete Webhook Message : Non-Nil Thread ID",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args:   args{msgID: Snowflake("148336120936005632"), threadID: StringToSnowflake("934478965031174194")},
			want:   http.MethodDelete,
			want1:  whIdTokenMessagesIdThreadId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.DeleteWebhookMessage(tt.args.msgID, tt.args.threadID)
			if got != tt.want {
				t.Errorf("DeleteWebhookMessage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DeleteWebhookMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookDeleteWebhookWithToken(t *testing.T) {
	type fields struct {
		ID    Snowflake
		Token string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Delete Webhook With Token",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			want:   http.MethodDelete,
			want1:  whIDToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.DeleteWebhookWithToken()
			if got != tt.want {
				t.Errorf("DeleteWebhookWithToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DeleteWebhookWithToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookEditWebhookMessage(t *testing.T) {
	type fields struct {
		ID    Snowflake
		Token string
	}
	type args struct {
		msgID    Snowflake
		threadID *Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		{
			name:   "Edit Webhook Message : Nil Thread ID",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args:   args{msgID: Snowflake("148336120936005632"), threadID: nil},
			want:   http.MethodPatch,
			want1:  whIdTokenMessagesId,
		},
		{
			name:   "Edit Webhook Message : Non-Nil Thread ID",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args:   args{msgID: Snowflake("148336120936005632"), threadID: StringToSnowflake("934478965031174194")},
			want:   http.MethodPatch,
			want1:  whIdTokenMessagesIdThreadId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.EditWebhookMessage(tt.args.msgID, tt.args.threadID)
			if got != tt.want {
				t.Errorf("EditWebhookMessage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EditWebhookMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookExecuteGitHubCompatibleWebhook(t *testing.T) {
	tests := setupExecute("github")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.ExecuteGitHubCompatibleWebhook(tt.args.wait, tt.args.threadID)
			if got != tt.want {
				t.Errorf("ExecuteGitHubCompatibleWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExecuteGitHubCompatibleWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func setupExecute(t string) []struct {
	name   string
	fields struct {
		ID    Snowflake
		Token string
	}
	args struct {
		wait     *bool
		threadID *Snowflake
	}
	want  string
	want1 string
} {

	var which string
	switch t {
	case "github":
		fallthrough
	case "slack":
		which = "/" + t
	default:
		which = ""
	}

	return []struct {
		name   string
		fields struct {
			ID    Snowflake
			Token string
		}
		args struct {
			wait     *bool
			threadID *Snowflake
		}
		want  string
		want1 string
	}{
		{
			name: "Execute " + t + " Webhook : Wait : nil; Thread ID : nil",
			fields: struct {
				ID    Snowflake
				Token string
			}{ID: "905130195520983061", Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args: struct {
				wait     *bool
				threadID *Snowflake
			}{wait: nil, threadID: nil},
			want:  http.MethodPost,
			want1: whIDToken + which,
		},
		{
			name: "Execute " + t + " Webhook : Wait : non-nil; Thread ID : non-nil",
			fields: struct {
				ID    Snowflake
				Token string
			}{ID: "905130195520983061", Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args: struct {
				wait     *bool
				threadID *Snowflake
			}{wait: &w, threadID: StringToSnowflake("934478965031174194")},
			want:  http.MethodPost,
			want1: whIDToken + which + "?wait=true&thread_id=934478965031174194",
		},
	}
}

func TestWebhookExecuteSlackCompatibleWebhook(t *testing.T) {
	tests := setupExecute("slack")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.ExecuteSlackCompatibleWebhook(tt.args.wait, tt.args.threadID)
			if got != tt.want {
				t.Errorf("ExecuteSlackCompatibleWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExecuteSlackCompatibleWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookExecuteWebhook(t *testing.T) {
	tests := setupExecute("")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.ExecuteWebhook(tt.args.wait, tt.args.threadID)
			if got != tt.want {
				t.Errorf("ExecuteWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExecuteWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookGetWebhook(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Get Webhook",
			fields: fields{ID: Snowflake("753228874875273256")},
			want:   http.MethodGet,
			want1:  "https://discord.com/api/v10/webhooks/753228874875273256",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID: tt.fields.ID,
			}
			got, got1 := w.GetWebhook()
			if got != tt.want {
				t.Errorf("GetWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookGetWebhookMessage(t *testing.T) {
	type fields struct {
		ID    Snowflake
		Token string
	}
	type args struct {
		msgID    Snowflake
		threadID *Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		{
			name:   "Get Webhook Message : Nil Thread ID",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args:   args{msgID: Snowflake("148336120936005632"), threadID: nil},
			want:   http.MethodGet,
			want1:  whIdTokenMessagesId,
		},
		{
			name:   "Get Webhook Message : Non-Nil Thread ID",
			fields: fields{ID: Snowflake("905130195520983061"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			args:   args{msgID: Snowflake("148336120936005632"), threadID: StringToSnowflake("934478965031174194")},
			want:   http.MethodGet,
			want1:  whIdTokenMessagesIdThreadId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.GetWebhookMessage(tt.args.msgID, tt.args.threadID)
			if got != tt.want {
				t.Errorf("GetWebhookMessage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetWebhookMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookGetWebhookWithToken(t *testing.T) {
	type fields struct {
		ID    Snowflake
		Token string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Get Webhook With Token",
			fields: fields{ID: Snowflake("753228874875273256"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			want:   http.MethodGet,
			want1:  "https://discord.com/api/v10/webhooks/753228874875273256/fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.GetWebhookWithToken()
			if got != tt.want {
				t.Errorf("GetWebhookWithToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetWebhookWithToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookModifyWebhook(t *testing.T) {
	type fields struct {
		ID Snowflake
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Modify Webhook",
			fields: fields{ID: Snowflake("753228874875273256")},
			want:   http.MethodPatch,
			want1:  "https://discord.com/api/v10/webhooks/753228874875273256",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID: tt.fields.ID,
			}
			got, got1 := w.ModifyWebhook()
			if got != tt.want {
				t.Errorf("ModifyWebhook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ModifyWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWebhookModifyWebhookWithToken(t *testing.T) {
	type fields struct {
		ID    Snowflake
		Token string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name:   "Modify Webhook With Token",
			fields: fields{ID: Snowflake("753228874875273256"), Token: "fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt"},
			want:   http.MethodPatch,
			want1:  "https://discord.com/api/v10/webhooks/753228874875273256/fQvqTTtCJVKrBRnUawZG6eFfPJ41A83tmFzTNArt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:    tt.fields.ID,
				Token: tt.fields.Token,
			}
			got, got1 := w.ModifyWebhookWithToken()
			if got != tt.want {
				t.Errorf("ModifyWebhookWithToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ModifyWebhookWithToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
