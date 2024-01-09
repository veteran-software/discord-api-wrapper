/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package channels

import (
	"time"

	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* CHANNELS */

type (
	// ChannelCreate - Sent when a new guild channel is created, relevant to the current user. The inner payload is a channel object.
	ChannelCreate api.Channel

	// ChannelUpdate - Sent when a channel is updated.
	//
	// The inner payload is a channel object.
	//
	// This is not sent when the field last_message_id is altered.
	//
	// To keep track of the last_message_id changes, you must listen for Message Create events (or Thread Create events for GUILD_FORUM channels).
	//
	//This event may reference roles or guild members that no longer exist in the guild.
	ChannelUpdate api.Channel

	// ChannelDelete - Sent when a channel relevant to the current user is deleted. The inner payload is a channel object.
	ChannelDelete api.Channel

	// ThreadCreate - Sent when a thread is created, relevant to the current user, or when the current user is added to a thread.
	// The inner payload is a channel object.
	//
	//    When a thread is created, includes an additional newly_created boolean field.
	//    When being added to an existing private thread, includes a thread member object.
	ThreadCreate api.Channel

	// ThreadUpdate - Sent when a thread is updated.
	// The inner payload is a channel object.
	// This is not sent when the field last_message_id is altered.
	// To keep track of the last_message_id changes, you must listen for Message Create events.
	ThreadUpdate api.Channel

	// ThreadDelete - Sent when a thread relevant to the current user is deleted.
	// The inner payload is a subset of the channel object, containing just the `id`, `guild_id`, `parent_id`, and `type` fields.
	ThreadDelete api.Channel

	// ThreadListSync - Sent when the current user gains access to a channel.
	ThreadListSync struct {
		GuildID    api.Snowflake      `json:"guild_id"`              // ID of the guild
		ChannelIDS []api.Snowflake    `json:"channel_ids,omitempty"` // Parent channel IDs whose threads are being synced. If omitted, then threads were synced for the entire guild. This array may contain channel_ids that have no active threads as well, so you know to clear that data.
		Threads    []api.Channel      `json:"threads"`               // All active threads in the given channels that the current user can access
		Members    []api.ThreadMember `json:"members"`               // All thread member objects from the synced threads for the current user, indicating which threads the current user has been added to
	}

	// ThreadMemberUpdate - Sent when the thread member object for the current user is updated.
	// The inner payload is a thread member object with an extra guild_id field.
	// This event is documented for completeness, but unlikely to be used by most bots.
	//
	// For bots, this event largely is just a signal that you are a member of the thread.
	ThreadMemberUpdate struct {
		api.ThreadMember
		GuildID api.Snowflake `json:"guild_id,omitempty"` // 	ID of the guild
	}

	// ThreadMembersUpdate - Sent when anyone is added to or removed from a thread.
	// If the current user does not have the GUILD_MEMBERS Gateway Intent, then this event will only be sent if the current user was added to or removed from the thread.
	ThreadMembersUpdate struct {
		ID           api.Snowflake `json:"id"`
		GuildID      api.Snowflake `json:"guild_id"`
		MemberCount  uint8         `json:"member_count"`
		AddedMembers []struct {
			api.ThreadMember
			GuildMember api.GuildMember `json:"guild_member,omitempty"`
		} `json:"added_members,omitempty"`
		RemovedMemberIDS []api.Snowflake `json:"removed_member_ids,omitempty"`
	}

	// ChannelPinsUpdate - Sent when a message is pinned or unpinned in a text channel.
	// This is not sent when a pinned message is deleted.
	ChannelPinsUpdate struct {
		GuildID          api.Snowflake `json:"guild_id,omitempty"`
		ChannelID        api.Snowflake `json:"channel_id"`
		LastPinTimestamp *time.Time    `json:"last_pin_timestamp,omitempty"`
	}
)
