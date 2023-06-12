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
	"strings"
	"testing"
	"time"

	"github.com/veteran-software/discord-api-wrapper/v10/utilities"
)

func TestAuthor_SetIconURL(t *testing.T) {
	type fields struct {
		Name    string
		URL     string
		IconURL *string
	}

	type args struct {
		u *string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Author
	}{
		{
			name:   "nil URL",
			fields: fields{},
			args:   args{nil},
			want:   &Author{},
		},
		{
			// It takes a really fucked up URL to make this fail lol
			name:   "invalid URL",
			fields: fields{},
			args: args{
				u: utilities.ToPtr("\u0009"), // control characters will make this fail
			},
			want: &Author{},
		},
		{
			name:   "valid URL",
			fields: fields{},
			args: args{
				u: utilities.ToPtr("https://nowlivebot.com"),
			},
			want: &Author{
				IconURL: utilities.ToPtr("https://nowlivebot.com"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Author{
				Name:    tt.fields.Name,
				URL:     tt.fields.URL,
				IconURL: tt.fields.IconURL,
			}
			if got := a.SetIconURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIconURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthor_SetName(t *testing.T) {
	type fields struct {
		Name    string
		URL     string
		IconURL *string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Author
	}{
		{
			name:   "Name length that exceeds 256 characters",
			fields: fields{},
			args: args{
				name: strings.Repeat("a", 258),
			},
			want: &Author{
				Name: strings.Repeat("a", 252) + " ...",
			},
		},
		{
			name:   "Name less than 256 characters",
			fields: fields{},
			args: args{
				name: "Kappatato",
			},
			want: &Author{
				Name: "Kappatato",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Author{
				Name:    tt.fields.Name,
				URL:     tt.fields.URL,
				IconURL: tt.fields.IconURL,
			}
			if got := a.SetName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthor_SetURL(t *testing.T) {
	type fields struct {
		Name    string
		URL     string
		IconURL *string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Author
	}{
		{
			// It takes a really fucked up URL to make this fail lol
			name:   "invalid URL",
			fields: fields{},
			args: args{
				u: "\u0009", // control characters will make this fail
			},
			want: &Author{},
		},
		{
			name:   "valid URL",
			fields: fields{},
			args: args{
				u: "https://nowlivebot.com",
			},
			want: &Author{
				URL: "https://nowlivebot.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Author{
				Name:    tt.fields.Name,
				URL:     tt.fields.URL,
				IconURL: tt.fields.IconURL,
			}
			if got := a.SetURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_String(t *testing.T) {
	type fields struct {
		ID   Snowflake
		Type ChannelType
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GuildText",
			fields: fields{
				ID:   "123456789",
				Type: GuildText,
				Name: "GuildText",
			},
			want: "GTC:GuildText(123456789)",
		},
		{
			name: "DM",
			fields: fields{
				ID:   "123456789",
				Type: DM,
				Name: "DM",
			},
			want: "DM:DM(123456789)",
		},
		{
			name: "GroupDM",
			fields: fields{
				ID:   "123456789",
				Type: GroupDM,
				Name: "GroupDM",
			},
			want: "GDM:GroupDM(123456789)",
		},
		{
			name: "GuildAnnouncement",
			fields: fields{
				ID:   "123456789",
				Type: GuildAnnouncement,
				Name: "GuildAnnouncement",
			},
			want: "GNC:GuildAnnouncement(123456789)",
		},
		{
			name: "GuildAnnouncementThread",
			fields: fields{
				ID:   "123456789",
				Type: GuildAnnouncementThread,
				Name: "GuildAnnouncementThread",
			},
			want: "GNT:GuildAnnouncementThread(123456789)",
		},
		{
			name: "GuildPublicThread",
			fields: fields{
				ID:   "123456789",
				Type: GuildPublicThread,
				Name: "GuildPublicThread",
			},
			want: "GPuT:GuildPublicThread(123456789)",
		},
		{
			name: "GuildPrivateThread",
			fields: fields{
				ID:   "123456789",
				Type: GuildPrivateThread,
				Name: "GuildPrivateThread",
			},
			want: "GPrT:GuildPrivateThread(123456789)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Channel{
				ID:   tt.fields.ID,
				Type: tt.fields.Type,
				Name: tt.fields.Name,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_AddField(t *testing.T) {
	type fields struct {
		Fields []*Field
	}
	type args struct {
		name   string
		value  string
		inline bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name: "Empty Name",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "",
				value:  "Value",
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "",
						Value:  "Value",
						Inline: false,
					},
				},
			},
		},
		{
			name: "Standard Name Length",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "Name",
				value:  "Value",
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "Value",
						Inline: false,
					},
				},
			},
		},
		{
			name: "Too Long Name Length",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   strings.Repeat("a", 258),
				value:  "Value",
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   strings.Repeat("a", 252) + " ...",
						Value:  "Value",
						Inline: false,
					},
				},
			},
		},
		{
			name: "Empty Value",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "Name",
				value:  "",
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "",
						Inline: false,
					},
				},
			},
		},
		{
			name: "Standard Value Length",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "Name",
				value:  "Value",
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "Value",
						Inline: false,
					},
				},
			},
		},
		{
			name: "Too Long Value Length",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "Name",
				value:  strings.Repeat("a", 1026),
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  strings.Repeat("a", 1020) + " ...",
						Inline: false,
					},
				},
			},
		},
		{
			name: "Inline true",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "Name",
				value:  "Value",
				inline: true,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "Value",
						Inline: true,
					},
				},
			},
		},
		{
			name: "Inline false",
			fields: fields{
				Fields: []*Field{},
			},
			args: args{
				name:   "Name",
				value:  "Value",
				inline: false,
			},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "Value",
						Inline: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Fields: tt.fields.Fields,
			}
			if got := e.AddField(tt.args.name, tt.args.value, tt.args.inline); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_AddFields(t *testing.T) {
	type fields struct {
		Fields []*Field
	}
	type args struct {
		fields []*Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name:   "No Fields Passed",
			fields: fields{Fields: []*Field{}},
			args:   args{},
			want: &Embed{
				Fields: []*Field{},
			},
		},
		{
			name:   "1 Field Passed",
			fields: fields{Fields: []*Field{}},
			args: args{fields: []*Field{
				{
					Name:   "Name",
					Value:  "Value",
					Inline: true,
				},
			}},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "Value",
						Inline: true,
					},
				},
			},
		},
		{
			name:   "Multiple Field Passed",
			fields: fields{Fields: []*Field{}},
			args: args{fields: []*Field{
				{
					Name:   "Name",
					Value:  "Value",
					Inline: true,
				},
				{
					Name:   "Name 1",
					Value:  "Value 1",
					Inline: true,
				},
			}},
			want: &Embed{
				Fields: []*Field{
					{
						Name:   "Name",
						Value:  "Value",
						Inline: true,
					},
					{
						Name:   "Name 1",
						Value:  "Value 1",
						Inline: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Fields: tt.fields.Fields,
			}
			if got := e.AddFields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_IsValidLength(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Footer      *Footer
		Author      *Author
		Fields      []*Field
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Valid Body; Invalid Field Name",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit+1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: false,
		},
		{
			name: "Valid Body; Invalid Field Value",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit+1),
					},
				},
			},
			want: false,
		},
		{
			name: "Valid Body; Valid Field",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: true,
		},
		{
			name: "Valid Field; Invalid Title",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit+1),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: false,
		},
		{
			name: "Valid Field; Invalid Description",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit+1),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: false,
		},
		{
			name: "Valid Field; Invalid Footer Text",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit+1)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: false,
		},
		{
			name: "Valid Field; Invalid Author name",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit+1)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: false,
		},
		{
			name: "Valid Body; Invalid Field Count",
			fields: fields{
				Title:       strings.Repeat("a", TitleLimit),
				Description: strings.Repeat("a", DescriptionLimit),
				Footer:      &Footer{Text: strings.Repeat("a", FooterTextLimit)},
				Author:      &Author{Name: strings.Repeat("a", AuthorNameLimit)},
				Fields: []*Field{
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
					{
						Name:  strings.Repeat("a", FieldNameLimit-1),
						Value: strings.Repeat("a", FieldValueLimit-1),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Footer:      tt.fields.Footer,
				Author:      tt.fields.Author,
				Fields:      tt.fields.Fields,
			}
			if got := e.IsValidLength(); got != tt.want {
				t.Errorf("IsValidLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetAuthor(t *testing.T) {
	type fields struct {
		Author *Author
	}
	type args struct {
		name    string
		url     string
		iconURL *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name:   "nil Icon URL",
			fields: fields{},
			args: args{
				name:    "Name",
				url:     "URL",
				iconURL: nil,
			},
			want: &Embed{
				Author: &Author{
					Name:    "Name",
					URL:     "URL",
					IconURL: nil,
				},
			},
		},
		{
			name:   "With Icon URL",
			fields: fields{},
			args: args{
				name:    "Name",
				url:     "URL",
				iconURL: utilities.ToPtr("https://discord.com"),
			},
			want: &Embed{
				Author: &Author{
					Name:    "Name",
					URL:     "URL",
					IconURL: utilities.ToPtr("https://discord.com"),
				},
			},
		},
		{
			name:   "Invalid URL",
			fields: fields{},
			args: args{
				name:    "Name",
				url:     "\u0009",
				iconURL: nil,
			},
			want: &Embed{
				Author: &Author{
					Name:    "Name",
					URL:     "",
					IconURL: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Author: tt.fields.Author,
			}
			if got := e.SetAuthor(tt.args.name, tt.args.url, tt.args.iconURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetColor(t *testing.T) {
	type fields struct {
		Color int64
	}
	type args struct {
		c int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name:   "Basic",
			fields: fields{},
			args:   args{c: int64(123456789)},
			want:   &Embed{Color: int64(123456789)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Color: tt.fields.Color,
			}
			if got := e.SetColor(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetDescription(t *testing.T) {
	type fields struct {
		Description string
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name:   "Basic",
			fields: fields{},
			args:   args{description: strings.Repeat("a", DescriptionLimit)},
			want:   &Embed{Description: strings.Repeat("a", DescriptionLimit)},
		},
		{
			name:   "Empty Description",
			fields: fields{},
			args:   args{description: ""},
			want:   &Embed{Description: ""},
		},
		{
			name:   "Too Long Description",
			fields: fields{},
			args:   args{description: strings.Repeat("a", DescriptionLimit+5)},
			want:   &Embed{Description: strings.Repeat("a", DescriptionLimit-4) + " ..."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Description: tt.fields.Description,
			}
			if got := e.SetDescription(tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetFooter(t *testing.T) {
	type fields struct {
		Footer *Footer
	}
	type args struct {
		text    string
		iconURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name:   "Basic",
			fields: fields{},
			args: args{
				text:    "Text Here",
				iconURL: "https://nowlivebot.com",
			},
			want: &Embed{Footer: &Footer{
				Text:    "Text Here",
				IconURL: "https://nowlivebot.com",
			}},
		},
		{
			name:   "Too Long Text",
			fields: fields{},
			args: args{
				text:    strings.Repeat("a", FooterTextLimit+1),
				iconURL: "https://nowlivebot.com",
			},
			want: &Embed{Footer: &Footer{
				Text:    strings.Repeat("a", FooterTextLimit-4) + " ...",
				IconURL: "https://nowlivebot.com",
			}},
		},
		{
			name:   "Invalid URL",
			fields: fields{},
			args: args{
				text:    strings.Repeat("a", FooterTextLimit),
				iconURL: "\u0009",
			},
			want: &Embed{Footer: &Footer{
				Text:    strings.Repeat("a", FooterTextLimit),
				IconURL: "",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Footer: tt.fields.Footer,
			}
			if got := e.SetFooter(tt.args.text, tt.args.iconURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFooter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetImage(t *testing.T) {
	type fields struct {
		Image *Image
	}
	type args struct {
		imageURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name: "Valid URL",
			fields: fields{
				Image: &Image{},
			},
			args: args{
				imageURL: "https://nowlivebot.com",
			},
			want: &Embed{
				Image: &Image{
					URL: "https://nowlivebot.com",
				},
			},
		},
		{
			name: "Invalid URL",
			fields: fields{
				Image: &Image{},
			},
			args: args{
				imageURL: "\u0009",
			},
			want: &Embed{
				Image: &Image{
					URL: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Image: tt.fields.Image,
			}
			if got := e.SetImage(tt.args.imageURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetThumbnail(t *testing.T) {
	type fields struct {
		Thumbnail *Thumbnail
	}
	type args struct {
		thumbnailURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name: "Valid URL",
			fields: fields{
				Thumbnail: &Thumbnail{},
			},
			args: args{
				thumbnailURL: "https://nowlivebot.com",
			},
			want: &Embed{
				Thumbnail: &Thumbnail{
					URL: "https://nowlivebot.com",
				},
			},
		},
		{
			name: "Invalid URL",
			fields: fields{
				Thumbnail: &Thumbnail{},
			},
			args: args{
				thumbnailURL: "\u0009",
			},
			want: &Embed{
				Thumbnail: &Thumbnail{
					URL: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Thumbnail: tt.fields.Thumbnail,
			}
			if got := e.SetThumbnail(tt.args.thumbnailURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetThumbnail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetTimestamp(t *testing.T) {
	type fields struct {
		Timestamp string
	}
	type args struct {
		ts time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name: "Basic",
			fields: fields{
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
			args: args{
				ts: time.Now(),
			},
			want: &Embed{
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Timestamp: tt.fields.Timestamp,
			}
			if got := e.SetTimestamp(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetTitle(t *testing.T) {
	type fields struct {
		Title string
	}
	type args struct {
		title string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name: "Valid title",
			fields: fields{
				Title: "",
			},
			args: args{
				title: strings.Repeat("a", TitleLimit),
			},
			want: &Embed{
				Title: strings.Repeat("a", TitleLimit),
			},
		},
		{
			name: "Too Long",
			fields: fields{
				Title: "",
			},
			args: args{
				title: strings.Repeat("a", TitleLimit+1),
			},
			want: &Embed{
				Title: strings.Repeat("a", TitleLimit-4) + " ...",
			},
		},
		{
			name: "Empty String",
			fields: fields{
				Title: "",
			},
			args: args{
				title: "",
			},
			want: &Embed{
				Title: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				Title: tt.fields.Title,
			}
			if got := e.SetTitle(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_SetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Embed
	}{
		{
			name: "Valid URL",
			fields: fields{
				URL: "",
			},
			args: args{
				u: "https://nowlivebot.com",
			},
			want: &Embed{
				URL: "https://nowlivebot.com",
			},
		},
		{
			name: "Invalid URL",
			fields: fields{
				URL: "",
			},
			args: args{
				u: "\u0009",
			},
			want: &Embed{
				URL: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Embed{
				URL: tt.fields.URL,
			}
			if got := e.SetURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_IsInline(t *testing.T) {
	type fields struct {
		Name   string
		Value  string
		Inline bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "True",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: true,
			},
			want: true,
		},
		{
			name: "False",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				Name:   tt.fields.Name,
				Value:  tt.fields.Value,
				Inline: tt.fields.Inline,
			}
			if got := f.IsInline(); got != tt.want {
				t.Errorf("IsInline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_SetInline(t *testing.T) {
	type fields struct {
		Name   string
		Value  string
		Inline bool
	}
	type args struct {
		inline bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Field
	}{
		{
			name: "Set True",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: false,
			},
			args: args{
				inline: true,
			},
			want: &Field{
				Name:   "",
				Value:  "",
				Inline: true,
			},
		},
		{
			name: "Set False",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: true,
			},
			args: args{
				inline: false,
			},
			want: &Field{
				Name:   "",
				Value:  "",
				Inline: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				Name:   tt.fields.Name,
				Value:  tt.fields.Value,
				Inline: tt.fields.Inline,
			}
			if got := f.SetInline(tt.args.inline); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_SetName(t *testing.T) {
	type fields struct {
		Name   string
		Value  string
		Inline bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Field
	}{
		{
			name: "Valid String",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: false,
			},
			args: args{
				name: strings.Repeat("a", FieldNameLimit),
			},
			want: &Field{
				Name:   strings.Repeat("a", FieldNameLimit),
				Value:  "",
				Inline: false,
			},
		},
		{
			name: "Too Long",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: false,
			},
			args: args{
				name: strings.Repeat("a", FieldNameLimit+1),
			},
			want: &Field{
				Name:   strings.Repeat("a", FieldNameLimit-4) + " ...",
				Value:  "",
				Inline: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				Name:   tt.fields.Name,
				Value:  tt.fields.Value,
				Inline: tt.fields.Inline,
			}
			if got := f.SetName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_SetValue(t *testing.T) {
	type fields struct {
		Name   string
		Value  string
		Inline bool
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Field
	}{
		{
			name: "Valid String",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: false,
			},
			args: args{
				value: strings.Repeat("a", FieldValueLimit),
			},
			want: &Field{
				Name:   "",
				Value:  strings.Repeat("a", FieldValueLimit),
				Inline: false,
			},
		},
		{
			name: "Too Long",
			fields: fields{
				Name:   "",
				Value:  "",
				Inline: false,
			},
			args: args{
				value: strings.Repeat("a", FieldValueLimit+1),
			},
			want: &Field{
				Name:   "",
				Value:  strings.Repeat("a", FieldValueLimit-4) + " ...",
				Inline: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				Name:   tt.fields.Name,
				Value:  tt.fields.Value,
				Inline: tt.fields.Inline,
			}
			if got := f.SetValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFooter_SetIconURL(t *testing.T) {
	type fields struct {
		Text    string
		IconURL string
	}
	type args struct {
		iconURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Footer
	}{
		{
			name: "Valid URL",
			fields: fields{
				Text:    "",
				IconURL: "",
			},
			args: args{
				iconURL: "https://nowlivebot.com",
			},
			want: &Footer{
				Text:    "",
				IconURL: "https://nowlivebot.com",
			},
		},
		{
			name: "Invalid URL",
			fields: fields{
				Text:    "",
				IconURL: "",
			},
			args: args{
				iconURL: "\u0009",
			},
			want: &Footer{
				Text:    "",
				IconURL: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Footer{
				Text:    tt.fields.Text,
				IconURL: tt.fields.IconURL,
			}
			if got := f.SetIconURL(tt.args.iconURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIconURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFooter_SetText(t *testing.T) {
	type fields struct {
		Text    string
		IconURL string
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Footer
	}{
		{
			name: "Valid String",
			fields: fields{
				Text:    "",
				IconURL: "",
			},
			args: args{
				text: strings.Repeat("a", FooterTextLimit),
			},
			want: &Footer{
				Text:    strings.Repeat("a", FooterTextLimit),
				IconURL: "",
			},
		},
		{
			name: "Too Long",
			fields: fields{
				Text:    "",
				IconURL: "",
			},
			args: args{
				text: strings.Repeat("a", FooterTextLimit+1),
			},
			want: &Footer{
				Text:    strings.Repeat("a", FooterTextLimit-4) + " ...",
				IconURL: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Footer{
				Text:    tt.fields.Text,
				IconURL: tt.fields.IconURL,
			}
			if got := f.SetText(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_SetURL(t *testing.T) {
	type fields struct {
		URL    string
		Height int
		Width  int
	}
	type args struct {
		u string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Image
	}{
		{
			name: "Valid URL",
			fields: fields{
				URL: "",
			},
			args: args{
				u: "https://nowlivebot.com",
			},
			want: &Image{
				URL: "https://nowlivebot.com",
			},
		},
		{
			name: "Invalid URL",
			fields: fields{
				URL: "",
			},
			args: args{
				u: "\u0009",
			},
			want: &Image{
				URL: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				URL:    tt.fields.URL,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			if got := i.SetURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmbed(t *testing.T) {
	tests := []struct {
		name string
		want *Embed
	}{
		{
			name: "Basic",
			want: &Embed{
				Type:      RichEmbed,
				Timestamp: time.Now().Format(time.RFC3339),
				Color:     16711680,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmbed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmbed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewField(t *testing.T) {
	tests := []struct {
		name string
		want *Field
	}{
		{
			name: "Basic",
			want: &Field{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThumbnail_SetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Thumbnail
	}{
		{
			name: "Valid URL",
			fields: fields{
				URL: "",
			},
			args: args{
				u: "https://nowlivebot.com",
			},
			want: &Thumbnail{
				URL: "https://nowlivebot.com",
			},
		},
		{
			name: "Invalid URL",
			fields: fields{
				URL: "",
			},
			args: args{
				u: "\u0009",
			},
			want: &Thumbnail{
				URL: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &Thumbnail{
				URL: tt.fields.URL,
			}
			if got := t.SetURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAuthor(t *testing.T) {
	tests := []struct {
		name string
		want *Author
	}{
		{
			name: "Basic",
			want: &Author{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAuthor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newFooter(t *testing.T) {
	tests := []struct {
		name string
		want *Footer
	}{
		{
			name: "Basic",
			want: &Footer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFooter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFooter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newImage(t *testing.T) {
	tests := []struct {
		name string
		want *Image
	}{
		{
			name: "Basic",
			want: &Image{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newImage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newThumbnail(t *testing.T) {
	tests := []struct {
		name string
		want *Thumbnail
	}{
		{
			name: "Basic",
			want: &Thumbnail{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newThumbnail(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newThumbnail() = %v, want %v", got, tt.want)
			}
		})
	}
}

//goland:noinspection SpellCheckingInspection
func TestChannel_getSelfMember(t *testing.T) {
	type fields struct {
		GuildID Snowflake
	}

	want := &GuildMember{
		Avatar:                     nil,
		CommunicationDisabledUntil: nil,
		Flags:                      0,
		JoinedAt:                   time.Now().UTC(),
		Nick:                       nil,
		Pending:                    false,
		PremiumSince:               nil,
		Roles:                      nil,
		User: User{
			ID:               "240729664035880961",
			Username:         "Now Live",
			Avatar:           utilities.ToPtr("fc1d3b807261f224bb566beff75c2327"),
			Discriminator:    "0483",
			PublicFlags:      589824,
			Flags:            589824,
			Bot:              true,
			Banner:           nil,
			AccentColor:      nil,
			GlobalName:       nil,
			AvatarDecoration: nil,
			DisplayName:      nil,
			System:           false,
			MfaEnabled:       false,
			BannerColor:      nil,
		},
		Deaf: false,
		Mute: false,
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
		fields  fields
		want    *GuildMember
		wantErr bool
	}{
		{
			name: "Valid Existing Member",
			fields: fields{
				GuildID: Snowflake("12345678900000"),
			},
			want:    want,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Channel{
				GuildID: tt.fields.GuildID,
			}
			got, err := c.getSelfMember()
			if (err != nil) != tt.wantErr {
				t.Errorf("getSelfMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSelfMember() got = %v, want %v", got, tt.want)
			}
		})
	}

	testClient = nil
}
