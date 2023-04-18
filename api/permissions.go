/*
 * Copyright (c) 2022-2023. Veteran Software
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

	log "github.com/veteran-software/nowlive-logging"
)

// Permission - Permissions in Discord are a way to limit and grant certain abilities to users.
// A set of base permissions can be configured at the guild level for different roles.
// When these roles are attached to users, they grant or revoke specific privileges within the guild.
// Along with the guild-level permissions, Discord also supports permission overwrites that can be assigned to individual guild roles or guild members on a per-channel basis.
//
// Permissions are stored within a variable-length integer serialized into a string, and are calculated using bitwise operations.
// For example, the permission value 123 will be serialized as "123".
// For long-term stability, we recommend deserializing the permissions using your languages' Big Integer libraries.
// The total permissions integer can be determined by ORing together each individual value, and flags can be checked using AND operations.
type Permission uint64

//goland:noinspection GoUnusedConst
const (
	NoPermissions                    Permission = 0 << 0  // For Vanity Roles that have no other permissions attached to them
	CreateInstantInvite              Permission = 1 << 0  // Allows creation of instant invites
	KickMembers                      Permission = 1 << 1  // Allows kicking members
	BanMembers                       Permission = 1 << 2  // Allows banning members
	Administrator                    Permission = 1 << 3  // Allows all permissions and bypasses channel permission overwrites
	ManageChannels                   Permission = 1 << 4  // Allows management and editing of channels
	ManageGuild                      Permission = 1 << 5  // Allows management and editing of the guild
	AddReactions                     Permission = 1 << 6  // Allows for the addition of reactions to messages
	ViewAuditLog                     Permission = 1 << 7  // Allows for viewing of audit logs
	PrioritySpeaker                  Permission = 1 << 8  // Allows for using priority speaker in a voice channel
	Stream                           Permission = 1 << 9  // Allows the user to go live
	ViewChannel                      Permission = 1 << 10 // Allows guild members to view a channel, which includes reading messages in text channels and joining voice channels
	SendMessages                     Permission = 1 << 11 // Allows for sending messages in a channel (does not allow sending messages in threads)
	SendTtsMessages                  Permission = 1 << 12 // Allows for sending of /tts messages
	ManageMessages                   Permission = 1 << 13 // Allows for deletion of other users messages
	EmbedLinks                       Permission = 1 << 14 // Links sent by users with this permission will be auto-embedded
	AttachFiles                      Permission = 1 << 15 // Allows for uploading images and files
	ReadMessageHistory               Permission = 1 << 16 // Allows for reading of message history
	MentionEveryone                  Permission = 1 << 17 // Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel
	UseExternalEmojis                Permission = 1 << 18 // Allows the usage of custom emojis from other servers
	ViewGuildInsights                Permission = 1 << 19 // Allows for viewing guild insights
	Connect                          Permission = 1 << 20 // Allows for joining of a voice channel
	Speak                            Permission = 1 << 21 // Allows for speaking in a voice channel
	MuteMembers                      Permission = 1 << 22 // Allows for muting members in a voice channel
	DeafenMembers                    Permission = 1 << 23 // Allows for deafening of members in a voice channel
	MoveMembers                      Permission = 1 << 24 // Allows for moving of members between voice channels
	UseVoiceActivity                 Permission = 1 << 25 // Allows for using voice-activity-detection in a voice channel
	ChangeNickname                   Permission = 1 << 26 // Allows for modification of own nickname
	ManageNicknames                  Permission = 1 << 27 // Allows for modification of other users nicknames
	ManageRoles                      Permission = 1 << 28 // Allows management and editing of roles
	ManageWebhooks                   Permission = 1 << 29 // Allows management and editing of webhooks
	ManageGuildExpressions           Permission = 1 << 30 // Allows management and editing of emojis, stickers, and soundboard sounds
	UseApplicationCommands           Permission = 1 << 31 // Allows members to use application commands, including slash commands and context menu commands.
	RequestToSpeak                   Permission = 1 << 32 // Allows for requesting to speak in stage channels. (This permission is under active development and may be changed or removed.)
	ManageEvents                     Permission = 1 << 33 // Allows for creating, editing, and deleting scheduled events
	ManageThreads                    Permission = 1 << 34 // Allows for deleting and archiving threads, and viewing all private threads
	CreatePublicThreads              Permission = 1 << 35 // Allows for creating public and announcement threads
	CreatePrivateThreads             Permission = 1 << 36 // Allows for creating private threads
	UseExternalStickers              Permission = 1 << 37 // Allows the usage of custom stickers from other servers
	SendMessagesInThreads            Permission = 1 << 38 // Allows for sending messages in threads
	UseEmbeddedActivities            Permission = 1 << 39 // Allows for launching activities (applications with the EMBEDDED flag) in a voice channel
	ModerateMembers                  Permission = 1 << 40 // Allows for timing out users to prevent them from sending or reacting to messages in chat and threads, and from speaking in voice and stage channels
	ViewCreatorMonetizationAnalytics Permission = 1 << 41 // Allows for viewing role subscription insights
	UseSoundboard                    Permission = 1 << 42 // Allows for using soundboard in a voice channel
)

/*
Process all relevant permissions and return the correct effective permissions according to Discord's permissions hierarchy

    1. Base permissions given to @everyone are applied at a guild level
    2. Permissions allowed to a user by their roles are applied at a guild level
    3. Overwrites that deny permissions for @everyone are applied at a channel level
    4. Overwrites that allow permissions for @everyone are applied at a channel level
    5. Overwrites that deny permissions for specific roles are applied at a channel level
    6. Overwrites that allow permissions for specific roles are applied at a channel level
    7. Member-specific overwrites that deny permissions are applied at a channel level
    8. Member-specific overwrites that allow permissions are applied at a channel level
*/

// computeBasePermissions = computer the base guild level permissions
func computeBasePermissions(guild *Guild, member *GuildMember) Permission {
	if guild.OwnerID == member.User.ID {
		return Administrator
	}

	var everyone *Role
	for _, role := range guild.Roles {
		if role.ID == guild.ID {
			everyone = role
		}
	}

	if everyone == nil {
		return 0
	}

	permissions := everyone.Permissions

	for _, role := range member.Roles {
		if r, ok := memberHasRole(role, guild.Roles); ok {
			permissions |= r.Permissions
		}
	}

	if permissions&Administrator == Administrator {
		return Administrator
	}

	return permissions
}

func getOverwrite(channel *Channel, id *Snowflake, pType OverwriteType) *Overwrite {
	if channel != nil {
		for _, overwrite := range channel.PermissionOverwrites {
			if overwrite.Type == pType && overwrite.ID == *id {
				return overwrite
			}
		}
	}

	return nil
}

// memberHasRole - Iterate through all the guild roles and return the role when it (if) is found
func memberHasRole(roleID *Snowflake, guildRoles []*Role) (*Role, bool) {
	for _, role := range guildRoles {
		if role.ID == *roleID {
			return role, true
		}
	}

	return nil, false
}

func parseStringPermission(perm string) uint64 {
	p, err := strconv.ParseUint(perm, 10, 64)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return 0
	}

	return p
}

// computeOverwrites - Compute permissions based on the above hierarchy
func computeOverwrites(basePermissions Permission, member *GuildMember, channel *Channel) Permission {
	if basePermissions&Administrator == Administrator {
		return Administrator
	}

	permissions := basePermissions
	overwriteEveryone := getOverwrite(channel, &channel.GuildID, PermissionRole)
	if overwriteEveryone == nil {
		// Do something here
	} else {
		permissions &= ^Permission(parseStringPermission(overwriteEveryone.Deny))
		permissions |= Permission(parseStringPermission(overwriteEveryone.Allow))
	}

	var allow uint64
	var deny uint64
	for _, role := range member.Roles {
		overwrite := getOverwrite(channel, role, PermissionRole)
		if overwrite != nil {
			allow |= parseStringPermission(overwrite.Allow)
			deny |= parseStringPermission(overwrite.Deny)
		}
	}

	permissions &= ^Permission(deny)
	permissions |= ^Permission(allow)

	for _, role := range member.Roles {
		overwrite := getOverwrite(channel, role, PermissionMember)
		if overwrite != nil {
			permissions &= ^Permission(parseStringPermission(overwrite.Deny))
			permissions |= Permission(parseStringPermission(overwrite.Allow))
		}
	}

	return permissions
}

func computePermissions(member *GuildMember, channel *Channel) Permission {
	g := Guild{ID: channel.GuildID}
	guild, err := g.GetGuild(nil)
	if err != nil {
		return 0
	}
	basePerms := computeBasePermissions(guild, member)

	return computeOverwrites(basePerms, member, channel)
}

// CanCreateInstantInvite - Allows creation of instant invites
//
//goland:noinspection GoUnusedExportedFunction
func CanCreateInstantInvite(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&CreateInstantInvite == CreateInstantInvite
	}

	return false
}

// CanKickMembers - Allows kicking members
//
//goland:noinspection GoUnusedExportedFunction
func CanKickMembers(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&KickMembers == KickMembers
}

// CanBanMembers - Allows banning members
//
//goland:noinspection GoUnusedExportedFunction
func CanBanMembers(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&BanMembers == BanMembers
}

// CanAdminister - Allows all permissions and bypasses channel permission overwrites
//
//goland:noinspection GoUnusedExportedFunction
func CanAdminister(member *GuildMember, channel *Channel) bool {
	return computePermissions(member, channel)&Administrator == Administrator
}

// CanManageChannels - Allows management and editing of channels
//
//goland:noinspection GoUnusedExportedFunction
func CanManageChannels(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ManageChannels == ManageChannels
	}

	return false
}

// CanManageGuild - Allows management and editing of the guild
//
//goland:noinspection GoUnusedExportedFunction
func CanManageGuild(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ManageGuild == ManageGuild
}

// CanAddReactions - Allows for the addition of reactions to messages
//
//goland:noinspection GoUnusedExportedFunction
func CanAddReactions(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&AddReactions == AddReactions
	}

	return false
}

// CanViewAuditLog - Allows for viewing of audit logs
//
//goland:noinspection GoUnusedExportedFunction
func CanViewAuditLog(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ViewAuditLog == ViewAuditLog
}

// IsPrioritySpeaker - Allows for using priority speaker in a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func IsPrioritySpeaker(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&PrioritySpeaker == PrioritySpeaker
	}

	return false
}

// CanStream - Allows the user to go live
//
//goland:noinspection GoUnusedExportedFunction
func CanStream(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&Stream == Stream
	}

	return false
}

// CanViewChannel - Allows guild members to view a channel, which includes reading messages in text channels and joining voice channels
//
//goland:noinspection GoUnusedExportedFunction
func CanViewChannel(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ViewChannel == ViewChannel
	}

	return false
}

// CanSendMessages - Allows for sending messages in a channel and creating threads in a forum (does not allow sending messages in threads)
//
//goland:noinspection GoUnusedExportedFunction
func CanSendMessages(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&SendMessages == SendMessages
	}

	return false
}

// CanSendTtsMessages - Allows for sending of /tts messages
//
//goland:noinspection GoUnusedExportedFunction
func CanSendTtsMessages(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&SendTtsMessages == SendTtsMessages
	}

	return false
}

// CanManageMessages - Allows for deletion of other users messages
//
//goland:noinspection GoUnusedExportedFunction
func CanManageMessages(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ManageMessages == ManageMessages
	}

	return false
}

// CanEmbedLinks - Links sent by users with this permission will be auto-embedded
//
//goland:noinspection GoUnusedExportedFunction
func CanEmbedLinks(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&EmbedLinks == EmbedLinks
	}

	return false
}

// CanAttachFiles - Allows for uploading images and files
//
//goland:noinspection GoUnusedExportedFunction
func CanAttachFiles(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&AttachFiles == AttachFiles
	}

	return false
}

// CanReadMessageHistory - Allows for reading of message history
//
//goland:noinspection GoUnusedExportedFunction
func CanReadMessageHistory(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ReadMessageHistory == ReadMessageHistory
	}

	return false
}

// CanMentionEveryone - Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel
//
//goland:noinspection GoUnusedExportedFunction
func CanMentionEveryone(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&MentionEveryone == MentionEveryone
	}

	return false
}

// CanUseExternalEmojis - Allows the usage of custom emojis from other servers
//
//goland:noinspection GoUnusedExportedFunction
func CanUseExternalEmojis(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&UseExternalEmojis == UseExternalEmojis
	}

	return false
}

// CanViewGuildInsights - Allows for viewing guild insights
//
//goland:noinspection GoUnusedExportedFunction
func CanViewGuildInsights(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ViewGuildInsights == ViewGuildInsights
}

// CanConnect - Allows for joining of a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func CanConnect(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&Connect == Connect
	}

	return false
}

// CanSpeak - Allows for speaking in a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func CanSpeak(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&Speak == Speak
	}

	return false
}

// CanMuteMembers - Allows for muting members in a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func CanMuteMembers(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&MuteMembers == MuteMembers
	}

	return false
}

// CanDeafenMembers - Allows for deafening of members in a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func CanDeafenMembers(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&DeafenMembers == DeafenMembers
	}

	return false
}

// CanMoveMembers - Allows for moving of members between voice channels
//
//goland:noinspection GoUnusedExportedFunction
func CanMoveMembers(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&MoveMembers == MoveMembers
	}

	return false
}

// CanUseVoiceActivity - Allows for using voice-activity-detection in a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func CanUseVoiceActivity(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&UseVoiceActivity == UseVoiceActivity
	}

	return false
}

// CanChangeNickname - Allows for modification of own nickname
//
//goland:noinspection GoUnusedExportedFunction
func CanChangeNickname(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ChangeNickname == ChangeNickname
}

// CanManageNicknames - Allows for modification of other users nicknames
//
//goland:noinspection GoUnusedExportedFunction
func CanManageNicknames(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ManageNicknames == ManageNicknames
}

// CanManageRoles - Allows management and editing of roles
//
//goland:noinspection GoUnusedExportedFunction
func CanManageRoles(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ManageRoles == ManageRoles
	}

	return false
}

// CanManageWebhooks - Allows management and editing of webhooks
//
//goland:noinspection GoUnusedExportedFunction
func CanManageWebhooks(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ManageWebhooks == ManageWebhooks
	}

	return false
}

// CanManageGuildExpressions - Allows management and editing of emojis and stickers
//
//goland:noinspection GoUnusedExportedFunction
func CanManageGuildExpressions(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ManageGuildExpressions == ManageGuildExpressions
}

// CanUseApplicationCommands - Allows members to use application commands, including slash commands and context menu commands.
//
//goland:noinspection GoUnusedExportedFunction
func CanUseApplicationCommands(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&UseApplicationCommands == UseApplicationCommands
	}

	return false
}

// CanRequestToSpeak - Allows for requesting to speak in stage channels. (This permission is under active development and may be changed or removed.)
//
//goland:noinspection GoUnusedExportedFunction
func CanRequestToSpeak(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&RequestToSpeak == RequestToSpeak
	}

	return false
}

// CanManageEvents - Allows for creating, editing, and deleting scheduled events
//
//goland:noinspection GoUnusedExportedFunction
func CanManageEvents(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ManageEvents == ManageEvents
	}

	return false
}

// CanManageThreads - Allows for deleting and archiving threads, and viewing all private threads
//
//goland:noinspection GoUnusedExportedFunction
func CanManageThreads(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&ManageThreads == ManageThreads
	}

	return false
}

// CanCreatePublicThreads - Allows for creating public and announcement threads
//
//goland:noinspection GoUnusedExportedFunction
func CanCreatePublicThreads(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&CreatePublicThreads == CreatePublicThreads
	}

	return false
}

// CanCreatePrivateThreads - Allows for creating private threads
//
//goland:noinspection GoUnusedExportedFunction
func CanCreatePrivateThreads(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&CreatePrivateThreads == CreatePrivateThreads
	}

	return false
}

// CanUseExternalStickers - Allows the usage of custom stickers from other servers
//
//goland:noinspection GoUnusedExportedFunction
func CanUseExternalStickers(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) || channel.Type == GuildVoice || channel.Type == GuildStageVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&UseExternalStickers == UseExternalStickers
	}

	return false
}

// CanSendMessagesInThreads - Allows for sending messages in threads
//
//goland:noinspection GoUnusedExportedFunction
func CanSendMessagesInThreads(member *GuildMember, channel *Channel) bool {
	if isTextChannel(channel) {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&SendMessagesInThreads == SendMessagesInThreads
	}

	return false
}

// CanUseEmbeddedActivities - Allows for using Activities (applications with the Embedded flag) in a voice channel
//
//goland:noinspection GoUnusedExportedFunction
func CanUseEmbeddedActivities(member *GuildMember, channel *Channel) bool {
	if channel.Type == GuildVoice {
		permissions := computePermissions(member, channel)

		return permissions&Administrator == Administrator || permissions&UseEmbeddedActivities == UseEmbeddedActivities
	}

	return false
}

// CanModerateMembers - Allows for timing out users to prevent them from sending or reacting to messages in chat and threads, and from speaking in voice and stage channels
//
//goland:noinspection GoUnusedExportedFunction
func CanModerateMembers(member *GuildMember, channel *Channel) bool {
	permissions := computePermissions(member, channel)

	return permissions&Administrator == Administrator || permissions&ModerateMembers == ModerateMembers
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
	BotID             Snowflake `json:"bot_id,omitempty"`             // the id of the bot this role belongs to
	IntegrationID     Snowflake `json:"integration_id,omitempty"`     // the id of the integration this role belongs to
	PremiumSubscriber *string   `json:"premium_subscriber,omitempty"` // whether this is the guild's premium subscriber role
}
