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

package invites

import (
	"time"

	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* INVITES */

type (
	// InviteCreate - Sent when a new invite to a channel is created.
	InviteCreate struct {
		ChannelID         api.Snowflake        `json:"channel_id"`
		Code              string               `json:"code"`
		CreatedAt         time.Time            `json:"created_at"`
		GuildID           api.Snowflake        `json:"guild_id,omitempty"`
		Inviter           api.User             `json:"inviter"`
		MaxAge            int64                `json:"max_age"`
		MaxUses           int64                `json:"max_uses"`
		TargetType        api.InviteTargetType `json:"target_type,omitempty"`
		TargetUser        api.User             `json:"target_user,omitempty"`
		TargetApplication api.Application      `json:"target_application,omitempty"`
		Temporary         bool                 `json:"temporary"`
		Uses              int                  `json:"uses"`
	}

	// InviteDelete - Sent when an invite is deleted.
	InviteDelete struct {
		ChannelID api.Snowflake `json:"channel_id"`
		GuildID   api.Snowflake `json:"guild_id,omitempty"`
		Code      string        `json:"code"`
	}
)
