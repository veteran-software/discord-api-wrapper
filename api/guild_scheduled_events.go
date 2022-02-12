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

import "time"

type GuildScheduledEvent struct {
	ID                 Snowflake                         `json:"id"`
	GuildID            Snowflake                         `json:"guild_id"`
	ChannelID          *Snowflake                        `json:"channel_id"`
	CreatorID          *Snowflake                        `json:"creator_id"`
	Name               string                            `json:"name"`
	Description        string                            `json:"description,omitempty"`
	ScheduledStartTime time.Time                         `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                        `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel   `json:"privacy_level"`
	Status             GuildScheduledEventStatus         `json:"status"`
	EntityType         GuildScheduledEventType           `json:"entity_type"`
	EntityID           *Snowflake                        `json:"entity_id"`
	EntityMetadata     GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            User                              `json:"creator,omitempty"`
	UserCount          int64                             `json:"user_count,omitempty"`
}

type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = iota + 2
)

type GuildScheduledEventType int

const (
	GuildScheduledEventTypeStageInstance GuildScheduledEventType = iota + 1
	GuildScheduledEventTypeVoice
	GuildScheduledEventTypeExternal
)

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = iota + 1
	GuildScheduledEventStatusActive
	GuildScheduledEventStatusCompleted
	GuildScheduledEventStatusCancelled
)

type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"`
}
