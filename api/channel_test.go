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
	"testing"
)

func Test_isTextChannel(t *testing.T) {
	type args struct {
		channel *Channel
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "GuildText",
			args: args{
				channel: &Channel{
					Type: GuildText,
				},
			},
			want: true,
		},
		{
			name: "GuildAnnouncement",
			args: args{
				channel: &Channel{
					Type: GuildAnnouncement,
				},
			},
			want: true,
		},
		{
			name: "GuildAnnouncementThread",
			args: args{
				channel: &Channel{
					Type: GuildAnnouncementThread,
				},
			},
			want: true,
		},
		{
			name: "GuildPublicThread",
			args: args{
				channel: &Channel{
					Type: GuildPublicThread,
				},
			},
			want: true,
		},
		{
			name: "GuildPrivateThread",
			args: args{
				channel: &Channel{
					Type: GuildPrivateThread,
				},
			},
			want: true,
		},
		{
			name: "GuildDirectory",
			args: args{
				channel: &Channel{
					Type: GuildDirectory,
				},
			},
			want: true,
		},
		{
			name: "GuildForum",
			args: args{
				channel: &Channel{
					Type: GuildForum,
				},
			},
			want: true,
		},
		{
			name: "DM",
			args: args{
				channel: &Channel{
					Type: DM,
				},
			},
			want: false,
		},
		{
			name: "GuildVoice",
			args: args{
				channel: &Channel{
					Type: GuildVoice,
				},
			},
			want: false,
		},
		{
			name: "GroupDM",
			args: args{
				channel: &Channel{
					Type: GroupDM,
				},
			},
			want: false,
		},
		{
			name: "GuildCategory",
			args: args{
				channel: &Channel{
					Type: GuildCategory,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTextChannel(tt.args.channel); got != tt.want {
				t.Errorf("isTextChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}
