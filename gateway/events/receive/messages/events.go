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

package messages

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* MESSAGES */

type (
	// MessageCreate - Sent when a message is created. The inner payload is a message object with the following extra fields
	MessageCreate struct {
		api.Message
		GuildID  api.Snowflake   `json:"guild_id,omitempty"`
		Member   api.GuildMember `json:"member,omitempty"`
		Mentions []api.User      `json:"mentions"`
	}

	// MessageUpdate - Sent when a message is updated. The inner payload is a message object with the same extra fields as MessageCreate.
	MessageUpdate MessageCreate

	// MessageDelete - Sent when a message is deleted.
	MessageDelete struct {
		ID        api.Snowflake `json:"id"`
		ChannelID api.Snowflake `json:"channel_id"`
		GuildID   api.Snowflake `json:"guild_id,omitempty"`
	}

	// MessageDeleteBulk - Sent when multiple messages are deleted at once.
	MessageDeleteBulk struct {
		IDS       []api.Snowflake `json:"ids"`
		ChannelID api.Snowflake   `json:"channel_id"`
		GuildID   api.Snowflake   `json:"guild_id,omitempty"`
	}

	// MessageReactionAdd - Sent when a user adds a reaction to a message.
	MessageReactionAdd struct {
		UserID    api.Snowflake   `json:"user_id"`
		ChannelID api.Snowflake   `json:"channel_id"`
		MessageID api.Snowflake   `json:"message_id"`
		GuildID   api.Snowflake   `json:"guild_id,omitempty"`
		Member    api.GuildMember `json:"member,omitempty"`
		Emoji     api.Emoji       `json:"emoji"`
	}

	// MessageReactionRemove - Sent when a user removes a reaction from a message.
	MessageReactionRemove struct {
		UserID    api.Snowflake `json:"user_id"`
		ChannelID api.Snowflake `json:"channel_id"`
		MessageID api.Snowflake `json:"message_id"`
		GuildID   api.Snowflake `json:"guild_id,omitempty"`
		Emoji     api.Emoji     `json:"emoji"`
	}

	// MessageReactionRemoveAll - Sent when a user explicitly removes all reactions from a message.
	MessageReactionRemoveAll struct {
		ChannelID api.Snowflake `json:"channel_id"`
		MessageID api.Snowflake `json:"message_id"`
		GuildID   api.Snowflake `json:"guild_id,omitempty"`
	}

	// MessageReactionRemoveEmoji - Sent when a bot removes all instances of a given emoji from the reactions of a message.
	MessageReactionRemoveEmoji struct {
		ChannelID api.Snowflake `json:"channel_id"`
		GuildID   api.Snowflake `json:"guild_id,omitempty"`
		MessageID api.Snowflake `json:"message_id"`
		Emoji     api.Emoji     `json:"emoji"`
	}
)
