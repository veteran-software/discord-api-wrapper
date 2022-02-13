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
	CreateInstantInvite     int64 = 1 << 0  // CreateInstantInvite - Allows creation of instant invites
	KickMembers             int64 = 1 << 1  // KickMembers - Allows kicking members
	BanMembers              int64 = 1 << 2  // BanMembers - Allows banning members
	Administrator           int64 = 1 << 3  // Administrator - Allows all permissions and bypasses channel permission overwrites
	ManageChannels          int64 = 1 << 4  // ManageChannels - Allows management and editing of channels
	ManageGuild             int64 = 1 << 5  // ManageGuild - Allows management and editing of the guild
	AddReactions            int64 = 1 << 6  // AddReactions - Allows for the addition of reactions to messages
	ViewAuditLog            int64 = 1 << 7  // ViewAuditLog - Allows for viewing of audit logs
	PrioritySpeaker         int64 = 1 << 8  // PrioritySpeaker - Allows for using priority speaker in a voice channel
	Stream                  int64 = 1 << 9  // Stream - Allows the user to go live
	ViewChannel             int64 = 1 << 10 // ViewChannel - Allows guild members to view a channel, which includes reading messages in text channels and joining voice channels
	SendMessages            int64 = 1 << 11 // SendMessages - Allows for sending messages in a channel (does not allow sending messages in threads)
	SendTtsMessages         int64 = 1 << 12 // SendTtsMessages - Allows for sending of /tts messages
	ManageMessages          int64 = 1 << 13 // ManageMessages - Allows for deletion of other users messages
	EmbedLinks              int64 = 1 << 14 // EmbedLinks - Links sent by users with this permission will be auto-embedded
	AttachFiles             int64 = 1 << 15 // AttachFiles - Allows for uploading images and files
	ReadMessageHistory      int64 = 1 << 16 // ReadMessageHistory - Allows for reading of message history
	MentionEveryone         int64 = 1 << 17 // MentionEveryone - Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel
	UseExternalEmojis       int64 = 1 << 18 // UseExternalEmojis - Allows the usage of custom emojis from other servers
	ViewGuildInsights       int64 = 1 << 19 // ViewGuildInsights - Allows for viewing guild insights
	Connect                 int64 = 1 << 20 // Connect - Allows for joining of a voice channel
	Speak                   int64 = 1 << 21 // Speak - Allows for speaking in a voice channel
	MuteMembers             int64 = 1 << 22 // MuteMembers - Allows for muting members in a voice channel
	DeafenMembers           int64 = 1 << 23 // DeafenMembers - Allows for deafening of members in a voice channel
	MoveMembers             int64 = 1 << 24 // MoveMembers - Allows for moving of members between voice channels
	UseVoiceActivity        int64 = 1 << 25 // UseVoiceActivity - Allows for using voice-activity-detection in a voice channel
	ChangeNickname          int64 = 1 << 26 // ChangeNickname - Allows for modification of own nickname
	ManageNicknames         int64 = 1 << 27 // ManageNicknames - Allows for modification of other users nicknames
	ManageRoles             int64 = 1 << 28 // ManageRoles - Allows management and editing of roles
	ManageWebhooks          int64 = 1 << 29 // ManageWebhooks - Allows management and editing of webhooks
	ManageEmojisAndStickers int64 = 1 << 30 // ManageEmojisAndStickers - Allows management and editing of emojis and stickers
	UseApplicationCommands  int64 = 1 << 31 // UseApplicationCommands - Allows members to use application commands, including slash commands and context menu commands.
	RequestToSpeak          int64 = 1 << 32 // RequestToSpeak - Allows for requesting to speak in stage channels. (This permission is under active development and may be changed or removed.)
	ManageEvents            int64 = 1 << 33 // ManageEvents - Allows for creating, editing, and deleting scheduled events
	ManageThreads           int64 = 1 << 34 // ManageThreads - Allows for deleting and archiving threads, and viewing all private threads
	CreatePublicThreads     int64 = 1 << 35 // CreatePublicThreads - Allows for creating public and announcement threads
	CreatePrivateThreads    int64 = 1 << 36 // CreatePrivateThreads - Allows for creating private threads
	UseExternalStickers     int64 = 1 << 37 // UseExternalStickers - Allows the usage of custom stickers from other servers
	SendMessagesInThreads   int64 = 1 << 38 // SendMessagesInThreads - Allows for sending messages in threads
	StartEmbeddedActivities int64 = 1 << 39 // StartEmbeddedActivities - Allows for launching activities (applications with the EMBEDDED flag) in a voice channel
	ModerateMembers         int64 = 1 << 40 // ModerateMembers - Allows for timing out users to prevent them from sending or reacting to messages in chat and threads, and from speaking in voice and stage channels
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
