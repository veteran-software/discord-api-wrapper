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

type Permission uint64

//goland:noinspection GoUnusedConst
const (
	CreateInstantInvite     Permission = 1 << 0  // Allows creation of instant invites
	KickMembers             Permission = 1 << 1  // Allows kicking members
	BanMembers              Permission = 1 << 2  // Allows banning members
	Administrator           Permission = 1 << 3  // Allows all permissions and bypasses channel permission overwrites
	ManageChannels          Permission = 1 << 4  // Allows management and editing of channels
	ManageGuild             Permission = 1 << 5  // Allows management and editing of the guild
	AddReactions            Permission = 1 << 6  // Allows for the addition of reactions to messages
	ViewAuditLog            Permission = 1 << 7  // Allows for viewing of audit logs
	PrioritySpeaker         Permission = 1 << 8  // Allows for using priority speaker in a voice channel
	Stream                  Permission = 1 << 9  // Allows the user to go live
	ViewChannel             Permission = 1 << 10 // Allows guild members to view a channel, which includes reading messages in text channels and joining voice channels
	SendMessages            Permission = 1 << 11 // Allows for sending messages in a channel (does not allow sending messages in threads)
	SendTtsMessages         Permission = 1 << 12 // Allows for sending of /tts messages
	ManageMessages          Permission = 1 << 13 // Allows for deletion of other users messages
	EmbedLinks              Permission = 1 << 14 // Links sent by users with this permission will be auto-embedded
	AttachFiles             Permission = 1 << 15 // Allows for uploading images and files
	ReadMessageHistory      Permission = 1 << 16 // Allows for reading of message history
	MentionEveryone         Permission = 1 << 17 // Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel
	UseExternalEmojis       Permission = 1 << 18 // Allows the usage of custom emojis from other servers
	ViewGuildInsights       Permission = 1 << 19 // Allows for viewing guild insights
	Connect                 Permission = 1 << 20 // Allows for joining of a voice channel
	Speak                   Permission = 1 << 21 // Allows for speaking in a voice channel
	MuteMembers             Permission = 1 << 22 // Allows for muting members in a voice channel
	DeafenMembers           Permission = 1 << 23 // Allows for deafening of members in a voice channel
	MoveMembers             Permission = 1 << 24 // Allows for moving of members between voice channels
	UseVoiceActivity        Permission = 1 << 25 // Allows for using voice-activity-detection in a voice channel
	ChangeNickname          Permission = 1 << 26 // Allows for modification of own nickname
	ManageNicknames         Permission = 1 << 27 // Allows for modification of other users nicknames
	ManageRoles             Permission = 1 << 28 // Allows management and editing of roles
	ManageWebhooks          Permission = 1 << 29 // Allows management and editing of webhooks
	ManageEmojisAndStickers Permission = 1 << 30 // Allows management and editing of emojis and stickers
	UseApplicationCommands  Permission = 1 << 31 // Allows members to use application commands, including slash commands and context menu commands.
	RequestToSpeak          Permission = 1 << 32 // Allows for requesting to speak in stage channels. (This permission is under active development and may be changed or removed.)
	ManageEvents            Permission = 1 << 33 // Allows for creating, editing, and deleting scheduled events
	ManageThreads           Permission = 1 << 34 // Allows for deleting and archiving threads, and viewing all private threads
	CreatePublicThreads     Permission = 1 << 35 // Allows for creating public and announcement threads
	CreatePrivateThreads    Permission = 1 << 36 // Allows for creating private threads
	UseExternalStickers     Permission = 1 << 37 // Allows the usage of custom stickers from other servers
	SendMessagesInThreads   Permission = 1 << 38 // Allows for sending messages in threads
	StartEmbeddedActivities Permission = 1 << 39 // Allows for launching activities (applications with the EMBEDDED flag) in a voice channel
	ModerateMembers         Permission = 1 << 40 // Allows for timing out users to prevent them from sending or reacting to messages in chat and threads, and from speaking in voice and stage channels
)

// HasAdmin checks to see if the bot has admin on the channel in question
func HasAdmin(p Permission) bool {
	return p&Administrator == Administrator
}

// CanManageWebhooks - checks for this permission
func CanManageWebhooks(channel *Channel) bool {
	return rawPerms(channel)&ManageWebhooks == ManageWebhooks
}

// CanManageRoles - checks for this permission
func CanManageRoles(channel *Channel) bool {
	return rawPerms(channel)&ManageRoles == ManageRoles
}

// CanUseExternalEmojis - checks for this permission
func CanUseExternalEmojis(channel *Channel) bool {
	return rawPerms(channel)&UseExternalEmojis == UseExternalEmojis
}

// CanMentionEveryone - checks for this permission
func CanMentionEveryone(channel *Channel) bool {
	return rawPerms(channel)&MentionEveryone == MentionEveryone
}

// CanViewChannel - checks for this permission
func CanViewChannel(channel *Channel) bool {
	return rawPerms(channel)&ViewChannel == ViewChannel
}

// CanReadMessageHistory - checks for this permission
func CanReadMessageHistory(channel *Channel) bool {
	return rawPerms(channel)&ReadMessageHistory == ReadMessageHistory
}

// CanEmbedLinks - checks for this permission
func CanEmbedLinks(channel *Channel) bool {
	return rawPerms(channel)&EmbedLinks == EmbedLinks
}

// CanManageMessages - checks for this permission
func CanManageMessages(channel *Channel) bool {
	return rawPerms(channel)&ManageMessages == ManageMessages
}

// CanSendMessages - checks for this permission
func CanSendMessages(channel *Channel) bool {
	return rawPerms(channel)&SendMessages == SendMessages
}

func rawPerms(channel *Channel) Permission {
	i, err := strconv.Atoi(channel.Permissions)
	if err != nil {
		return 0
	}

	return Permission(i)
}

// CanAnnounce - Deprecated: helper function for checking bas permissions for sending announcements
func CanAnnounce(c *Channel) bool {
	if CanEmbedLinks(c) && CanManageMessages(c) && CanSendMessages(c) && CanReadMessageHistory(c) && CanViewChannel(c) {
		return true
	}

	return false
}

// Role - Roles represent a set of permissions attached to a group of users.
//
// Roles have unique names, colors, and can be "pinned" to the sidebar, causing their members to be listed separately.
//
// Roles are unique per guild, and can have separate permission profiles for the global context (guild) and channel context.
//
// The @everyone role has the same ID as the guild it belongs to.
type Role struct {
	ID           Snowflake  `json:"id"`                      // role id
	Name         string     `json:"name"`                    // role name
	Color        int        `json:"color"`                   // integer representation of hexadecimal color code
	Hoist        bool       `json:"hoist"`                   // if this role is pinned in the user listing
	Icon         *string    `json:"icon,omitempty"`          // role icon hash
	UnicodeEmoji *string    `json:"unicode_emoji,omitempty"` // role unicode emoji
	Position     int        `json:"position"`                // position of this role
	Permissions  Permission `json:"permissions"`             // permission bit set
	Managed      bool       `json:"managed"`                 // whether this role is managed by an integration
	Mentionable  bool       `json:"mentionable"`             // whether this role is mentionable
	Tags         RoleTags   `json:"tags,omitempty"`          // the tags this role has
}

// RoleTags - the tags this Role has
type RoleTags struct {
	BotID             string  `json:"bot_id,omitempty"`             // the id of the bot this role belongs to
	IntegrationID     string  `json:"integration_id,omitempty"`     // the id of the integration this role belongs to
	PremiumSubscriber *string `json:"premium_subscriber,omitempty"` // whether this is the guild's premium subscriber role
}
