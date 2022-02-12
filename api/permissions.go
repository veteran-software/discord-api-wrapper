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
	"strconv"
)

const (
	CreateInstantInvite     int64 = 1 << 0
	KickMembers             int64 = 1 << 1
	BanMembers              int64 = 1 << 2
	Administrator           int64 = 1 << 3
	ManageChannels          int64 = 1 << 4
	ManageGuild             int64 = 1 << 5
	AddReactions            int64 = 1 << 6
	ViewAuditLogs           int64 = 1 << 7
	PrioritySpeaker         int64 = 1 << 8
	Stream                  int64 = 1 << 9
	ViewChannel             int64 = 1 << 10
	SendMessages            int64 = 1 << 11 // Does not allow for sending messages in Threads
	SendTtsMessages         int64 = 1 << 12
	ManageMessages          int64 = 1 << 13
	EmbedLinks              int64 = 1 << 14
	AttachFiles             int64 = 1 << 15
	ReadMessageHistory      int64 = 1 << 16
	MentionEveryone         int64 = 1 << 17
	UseExternalEmojis       int64 = 1 << 18
	ViewGuildInsights       int64 = 1 << 19
	Connect                 int64 = 1 << 20
	Speak                   int64 = 1 << 21
	MuteMembers             int64 = 1 << 22
	DeafenMembers           int64 = 1 << 23
	MoveMembers             int64 = 1 << 24
	UseVoiceActivity        int64 = 1 << 25
	ChangeNickname          int64 = 1 << 26
	ManageNicknames         int64 = 1 << 27
	ManageRoles             int64 = 1 << 28
	ManageWebhooks          int64 = 1 << 29
	ManageEmojis            int64 = 1 << 30
	UseSlashCommands        int64 = 1 << 31
	RequestToSpeak          int64 = 1 << 32
	_                       int64 = 1 << 33
	ManageThreads           int64 = 1 << 34
	CreatePublicThreads     int64 = 1 << 35
	CreatePrivateThreads    int64 = 1 << 36
	UseExternalStickers     int64 = 1 << 37
	SendMessagesInThreads   int64 = 1 << 38
	StartEmbeddedActivities int64 = 1 << 39
	ModerateMembers         int64 = 1 << 40
)

/* HELPER FUNCTIONS */

// HasAdmin checks to see if the bot has admin on the channel in question
func HasAdmin(p int64) bool {
	return p&Administrator == Administrator
}

func CanManageWebhooks(channel *Channel) bool {
	return rawPerms(channel)&ManageWebhooks == ManageWebhooks
}

func CanManageRoles(channel *Channel) bool {
	return rawPerms(channel)&ManageRoles == ManageRoles
}

func CanUseExternalEmojis(channel *Channel) bool {
	return rawPerms(channel)&UseExternalEmojis == UseExternalEmojis
}

func CanMentionEveryone(channel *Channel) bool {
	return rawPerms(channel)&MentionEveryone == MentionEveryone
}

func CanViewChannel(channel *Channel) bool {
	return rawPerms(channel)&ViewChannel == ViewChannel
}

func CanReadMessageHistory(channel *Channel) bool {
	return rawPerms(channel)&ReadMessageHistory == ReadMessageHistory
}

func CanEmbedLinks(channel *Channel) bool {
	return rawPerms(channel)&EmbedLinks == EmbedLinks
}

func CanManageMessages(channel *Channel) bool {
	return rawPerms(channel)&ManageMessages == ManageMessages
}

func CanSendMessages(channel *Channel) bool {
	return rawPerms(channel)&SendMessages == SendMessages
}

func rawPerms(channel *Channel) int64 {
	i, err := strconv.Atoi(channel.Permissions)
	if err != nil {
		return 0
	}

	return int64(i)
}

func CanAnnounce(c *Channel) bool {
	if CanEmbedLinks(c) && CanManageMessages(c) && CanSendMessages(c) && CanReadMessageHistory(c) && CanViewChannel(c) {
		return true
	}

	return false
}

type RoleTags struct {
	BotID             string  `json:"bot_id,omitempty"`
	IntegrationID     string  `json:"integration_id,omitempty"`
	PremiumSubscriber *string `json:"premium_subscriber,omitempty"`
}

type Role struct {
	ID           Snowflake `json:"id"`
	Name         string    `json:"name"`
	Color        int       `json:"color"`
	Hoist        bool      `json:"hoist"`
	Icon         *string   `json:"icon,omitempty"`
	UnicodeEmoji *string   `json:"unicode_emoji,omitempty"`
	Position     int       `json:"position"`
	Permissions  string    `json:"permissions"`
	Managed      bool      `json:"managed"`
	Mentionable  bool      `json:"mentionable"`
	Tags         RoleTags  `json:"tags,omitempty"`
}
