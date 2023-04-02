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

package guilds

import (
	"time"

	"github.com/veteran-software/discord-api-wrapper/v10/api"
	"github.com/veteran-software/discord-api-wrapper/v10/gateway/events/receive/presence"
	"github.com/veteran-software/discord-api-wrapper/v10/gateway/events/send"
)

/* GUILDS */

type (

	// GuildCreate - This event can be sent in three different scenarios:
	//
	//    1. When a user is initially connecting, to lazily load and back-fill information for all unavailable guilds sent in the Ready event. Guilds that are unavailable due to an outage will send a Guild Delete event.
	//    2. When a Guild becomes available again to the client.
	//    3. When the current user joins a new Guild.
	//
	// During an outage, the guild object in scenarios 1 and 3 may be marked as unavailable.
	//
	// The inner payload can be:
	//
	//    An available Guild: a guild object with extra fields, as noted below.
	//    An unavailable Guild: an unavailable guild object.
	GuildCreate struct {
		api.Guild
		JoinedAt             time.Time                  `json:"joined_at,omitempty"`              // when this guild was joined at
		Large                bool                       `json:"large,omitempty"`                  // true if this is considered a large guild
		Unavailable          bool                       `json:"unavailable,omitempty"`            // true if this guild is unavailable due to an outage
		MemberCount          int64                      `json:"member_count,omitempty"`           // total number of members in this guild
		VoiceStates          []*api.VoiceState          `json:"voice_states,omitempty"`           // states of members currently in voice channels; lacks the guild_id key
		Members              []*api.GuildMember         `json:"members,omitempty"`                // users in the guild
		Channels             []*api.Channel             `json:"channels,omitempty"`               // channels in the guild
		Threads              []*api.Channel             `json:"threads,omitempty"`                // all active threads in the guild that current user has permission to view
		Presences            []*presence.Update         `json:"presences,omitempty"`              // presences of the members in the guild, will only include non-offline members if the size is greater than large threshold
		StageInstances       []*api.StageInstance       `json:"stage_instances,omitempty"`        // Stage instances in the guild
		GuildScheduledEvents []*api.GuildScheduledEvent `json:"guild_scheduled_events,omitempty"` // the scheduled events in the guild
	}

	// GuildUpdate - Sent when a guild is updated. The inner payload is a guild object.
	GuildUpdate api.Guild

	// GuildDelete - Sent when a guild becomes or was already unavailable due to an outage, or when the user leaves or is removed from a guild.
	// The inner payload is an unavailable guild object.
	//
	// If the unavailable field is not set, the user was removed from the guild.
	GuildDelete api.UnavailableGuild

	// GuildAuditLogEntryCreate - Sent when a guild audit log entry is created.
	// The inner payload is an Audit Log Entry object.
	//
	// This event is only sent to bots with the VIEW_AUDIT_LOG permission.
	GuildAuditLogEntryCreate api.AuditLogEntry

	// GuildBanAdd - Sent when a user is banned from a guild.
	GuildBanAdd struct {
		GuildID api.Snowflake `json:"guild_id"`
		User    api.User      `json:"user"`
	}

	// GuildBanRemove - Sent when a user is unbanned from a guild.
	GuildBanRemove struct {
		GuildID api.Snowflake `json:"guild_id"`
		User    api.User      `json:"user"`
	}

	// GuildEmojisUpdate - Sent when a guild's emojis have been updated.
	GuildEmojisUpdate struct {
		GuildID api.Snowflake `json:"guild_id"`
		Emojis  []api.Emoji   `json:"emojis"`
	}

	// GuildStickersUpdate - Sent when a guild's stickers have been updated.
	GuildStickersUpdate struct {
		GuildID  api.Snowflake `json:"guild_id"`
		Stickers []api.Sticker `json:"stickers"`
	}

	// GuildIntegrationsUpdate - Sent when a guild integration is updated.
	GuildIntegrationsUpdate struct {
		GuildID api.Snowflake `json:"guild_id"`
	}

	// GuildMemberAdd - Sent when a new user joins a guild. The inner payload is a guild member object with an extra `guild_id` key
	GuildMemberAdd struct {
		api.GuildMember
		GuildID api.Snowflake `json:"guild_id"`
	}

	// GuildMemberRemove - Sent when a user is removed from a guild (leave/kick/ban).
	GuildMemberRemove struct {
		GuildID api.Snowflake `json:"guild_id"`
		User    api.User      `json:"user"`
	}

	// GuildMemberUpdate - Sent when a guild member is updated. This will also fire when the user object of a guild member changes.
	GuildMemberUpdate struct {
		GuildID                    api.Snowflake   `json:"guild_id"`
		Roles                      []api.Snowflake `json:"roles"`
		User                       api.User        `json:"user"`
		Nick                       *string         `json:"nick,omitempty"`
		Avatar                     *string         `json:"avatar"`
		JoinedAt                   *time.Time      `json:"joined_at"`
		PremiumSince               *time.Time      `json:"premium_since,omitempty"`
		Deaf                       bool            `json:"deaf,omitempty"`
		Mute                       bool            `json:"mute,omitempty"`
		Pending                    bool            `json:"pending,omitempty"`
		CommunicationDisabledUntil *time.Time      `json:"communication_disabled_until,omitempty"`
	}

	// GuildMemberChunk - Sent in response to GuildRequestMembers. You can use the `chunk_index` and `chunk_count` to calculate how many chunks are left for your request.
	GuildMemberChunk struct {
		GuildID    api.Snowflake       `json:"guild_id"`
		Members    []api.GuildMember   `json:"members"`
		ChunkIndex uint64              `json:"chunk_index"`
		ChunkCount uint64              `json:"chunk_count"`
		NotFound   []api.Snowflake     `json:"not_found,omitempty"`
		Presences  send.PresenceUpdate `json:"presences,omitempty"`
		Nonce      string              `json:"nonce,omitempty"`
	}

	// GuildRoleCreate - Sent when a guild role is created.
	GuildRoleCreate struct {
		GuildID api.Snowflake `json:"guild_id"`
		Role    api.Role      `json:"role"`
	}

	// GuildRoleUpdate - Sent when a guild role is updated.
	GuildRoleUpdate struct {
		GuildID api.Snowflake `json:"guild_id"`
		Role    api.Role      `json:"role"`
	}

	// GuildRoleDelete - Sent when a guild role is deleted.
	GuildRoleDelete struct {
		GuildID api.Snowflake `json:"guild_id"`
		RoleID  api.Snowflake `json:"role_id"`
	}
)
