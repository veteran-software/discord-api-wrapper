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
	"time"
)

// GuildScheduledEvent - A representation of a scheduled event in a guild.
//
//	creator_id will be null and creator will not be included for events created before October 25th, 2021, when the concept of creator_id was introduced and tracked.
type GuildScheduledEvent struct {
	ID                 Snowflake                          `json:"id"`                    // the id of the scheduled event
	GuildID            Snowflake                          `json:"guild_id"`              // the guild id which the scheduled event belongs to
	ChannelID          *Snowflake                         `json:"channel_id"`            // the channel id in which the scheduled event will be hosted, or null if scheduled entity type is EXTERNAL
	CreatorID          *Snowflake                         `json:"creator_id"`            // the id of the user that created the scheduled event
	Name               string                             `json:"name"`                  // the name of the scheduled event (1-100 characters)
	Description        *string                            `json:"description,omitempty"` // the description of the scheduled event (1-1000 characters)
	ScheduledStartTime time.Time                          `json:"scheduled_start_time"`  // the time the scheduled event will start
	ScheduledEndTime   *time.Time                         `json:"scheduled_end_time"`    // the time the scheduled event will end, required if entity_type is EXTERNAL
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`         // the privacy level of the scheduled event
	Status             GuildScheduledEventStatus          `json:"status"`                // the status of the scheduled event
	EntityType         GuildScheduledEventType            `json:"entity_type"`           // the type of the scheduled event
	EntityID           *Snowflake                         `json:"entity_id"`             // the id of an entity associated with a guild scheduled event
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`       // additional metadata for the guild scheduled event
	Creator            User                               `json:"creator,omitempty"`     // the user that created the scheduled event
	UserCount          int64                              `json:"user_count,omitempty"`  // the number of users subscribed to the scheduled event
	Image              *string                            `json:"image,omitempty"`       // the cover image hash of the scheduled event
}

// GuildScheduledEventPrivacyLevel - the privacy level of the scheduled event
type GuildScheduledEventPrivacyLevel int

// GuildScheduledEventPrivacyLevelGuildOnly - the scheduled event is only accessible to guild members
//
//goland:noinspection GoUnusedConst
const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = iota + 2 // the scheduled event is only accessible to guild members
)

// GuildScheduledEventType - the type of the scheduled event
type GuildScheduledEventType int

//goland:noinspection GoUnusedConst
const (
	GuildScheduledEventTypeStageInstance GuildScheduledEventType = iota + 1
	GuildScheduledEventTypeVoice
	GuildScheduledEventTypeExternal
)

// GuildScheduledEventStatus - Once status is set to Completed or Cancelled, the status can no longer be updated
type GuildScheduledEventStatus int

//goland:noinspection GoUnusedConst
const (
	Scheduled GuildScheduledEventStatus = iota + 1
	Active
	Completed
	Cancelled
)

// GuildScheduledEventEntityMetadata - required for events with 'entity_type': EXTERNAL
type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"` // location of the event (1-100 characters)
}

// GuildScheduledEventUser - Representation of a user interested in attending an event
type GuildScheduledEventUser struct {
	GuildScheduledEventID Snowflake   `json:"guild_scheduled_event_id"` // the scheduled event id which the user subscribed to
	User                  User        `json:"user"`                     // user which subscribed to an event
	Member                GuildMember `json:"member"`                   // guild member data for this user for the guild which this event belongs to, if any
}
