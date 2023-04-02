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

package guild_scheduled_events

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* GUILD SCHEDULED EVENTS */

type (
	// GuildScheduledEventCreate - Sent when a guild scheduled event is created. The inner payload is a guild scheduled event object.
	GuildScheduledEventCreate api.GuildScheduledEvent

	// GuildScheduledEventUpdate - Sent when a guild scheduled event is updated. The inner payload is a guild scheduled event object.
	GuildScheduledEventUpdate api.GuildScheduledEvent

	// GuildScheduledEventDelete - Sent when a guild scheduled event is deleted. The inner payload is a guild scheduled event object.
	GuildScheduledEventDelete api.GuildScheduledEvent

	// GuildScheduledEventUserAdd - Sent when a user has subscribed to a guild scheduled event.
	GuildScheduledEventUserAdd struct {
		GuildScheduledEventID api.Snowflake `json:"guild_scheduled_event_id"`
		UserID                api.Snowflake `json:"user_id"`
		GuildID               api.Snowflake `json:"guild_id"`
	}

	// GuildScheduledEventUserRemove - Sent when a user has unsubscribed from a guild scheduled event.
	GuildScheduledEventUserRemove struct {
		GuildScheduledEventID api.Snowflake `json:"guild_scheduled_event_id"`
		UserID                api.Snowflake `json:"user_id"`
		GuildID               api.Snowflake `json:"guild_id"`
	}
)
