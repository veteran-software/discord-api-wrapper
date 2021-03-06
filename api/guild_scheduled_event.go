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
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GuildScheduledEvent - A representation of a scheduled event in a guild.
//
//   creator_id will be null and creator will not be included for events created before October 25th, 2021, when the concept of creator_id was introduced and tracked.
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

// ListGuildScheduledEvents - Returns a list of guild scheduled event objects for the given guild.
func (g *Guild) ListGuildScheduledEvents(withUserCount *bool) ([]GuildScheduledEvent, error) {
	u := parseRoute(fmt.Sprintf(listGuildScheduledEvents, api, g.ID.String()))

	q := u.Query()
	if withUserCount != nil {
		q.Set("with_user_count", strconv.FormatBool(*withUserCount))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var guildScheduledEvents []GuildScheduledEvent
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guildScheduledEvents)

	return guildScheduledEvents, err
}

// CreateGuildScheduledEvent - Create a GuildScheduledEvent in the Guild. Returns a guild scheduled event object on success.
//
//    A guild can have a maximum of 100 events with Scheduled or Active status at any time.
//
//    This endpoint supports the `X-Audit-Log-Reason` header.
func (g *Guild) CreateGuildScheduledEvent(payload CreateGuildScheduledEventJSON, reason *string) (*GuildScheduledEvent, error) {
	u := parseRoute(fmt.Sprintf(createGuildScheduledEvent, api, g.ID.String()))

	var guildScheduledEvent *GuildScheduledEvent
	err := json.Unmarshal(firePostRequest(u, payload, reason), &guildScheduledEvent)

	return guildScheduledEvent, err
}

// CreateGuildScheduledEventJSON - JSON payload
type CreateGuildScheduledEventJSON struct {
	ChannelID          Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata     GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Name               string                            `json:"name"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel   `json:"privacy_level"`
	ScheduledStartTime time.Time                         `json:"scheduled_start_time"`
	ScheduledEndTime   time.Time                         `json:"scheduled_end_time,omitempty"`
	Description        string                            `json:"description,omitempty"`
	EntityType         GuildScheduledEventType           `json:"entity_type"`
	Image              string                            `json:"image,omitempty"`
}

// GetGuildScheduledEvent - Get a guild scheduled event. Returns a guild scheduled event object on success.
func (g *Guild) GetGuildScheduledEvent(guildScheduledEventID Snowflake, withUserCount *bool) (*GuildScheduledEvent, error) {
	u := parseRoute(fmt.Sprintf(getGuildScheduledEvent, api, g.ID.String(), guildScheduledEventID.String()))

	q := u.Query()
	if withUserCount != nil {
		q.Set("with_user_count", strconv.FormatBool(*withUserCount))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var guildScheduledEvent *GuildScheduledEvent
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guildScheduledEvent)

	return guildScheduledEvent, err
}

// ModifyGuildScheduledEvent - Modify a guild scheduled event. Returns the modified guild scheduled event object on success.
//
//    To start or end an event, use this endpoint to modify the event's status field.
//
//    This endpoint supports the `X-Audit-Log-Reason` header.
//
//    This endpoint silently discards `entity_metadata` for non-EXTERNAL events.
//
// If updating entity_type to EXTERNAL:
//
//    `channel_id` is required and must be set to null
//    `entity_metadata` with a location field must be provided
//    `scheduled_end_time` must be provided
func (g *Guild) ModifyGuildScheduledEvent(guildScheduledEventID Snowflake, payload ModifyGuildScheduledEventJSON, reason *string) (*GuildScheduledEvent, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildScheduledEvent, api, g.ID.String(), guildScheduledEventID.String()))

	var guildScheduledEvent *GuildScheduledEvent
	err := json.Unmarshal(firePatchRequest(u, payload, reason), &guildScheduledEvent)

	return guildScheduledEvent, err
}

// ModifyGuildScheduledEventJSON - JSON payload
type ModifyGuildScheduledEventJSON struct {
	ChannelID          *Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Name               string                             `json:"name"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	ScheduledStartTime time.Time                          `json:"scheduled_start_time"`
	ScheduledEndTime   time.Time                          `json:"scheduled_end_time,omitempty"`
	Description        *string                            `json:"description,omitempty"`
	EntityType         GuildScheduledEventType            `json:"entity_type"`
	Image              string                             `json:"image,omitempty"`
	Status             GuildScheduledEventStatus          `json:"status,omitempty"`
}

// DeleteGuildScheduledEvent - Delete a guild scheduled event. Returns a 204 on success.
func (g *Guild) DeleteGuildScheduledEvent(guildScheduledEventID Snowflake) error {
	u := parseRoute(fmt.Sprintf(deleteGuildScheduledEvent, api, g.ID.String(), guildScheduledEventID.String()))

	return fireDeleteRequest(u, nil)
}

// GetGuildScheduledEventUsers - Get a list of guild scheduled event users subscribed to a guild scheduled event.
//
// Returns a list of guild scheduled event user objects on success.
//
// GuildMember data, if it exists, is included if the `with_member` query parameter is set.
func (g *Guild) GetGuildScheduledEventUsers(guildScheduledEventID Snowflake, limit *uint64, withMember *bool, before *Snowflake, after *Snowflake) (*GuildScheduledEventUser, error) {
	u := parseRoute(fmt.Sprintf(getGuildScheduledEventUsers, api, g.ID.String(), guildScheduledEventID.String()))

	q := u.Query()
	if limit != nil {
		if *limit > 100 {
			return nil, errors.New("limit must be > 0 && <= 100")
		}
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	if withMember != nil {
		q.Set("with_member", strconv.FormatBool(*withMember))
	}
	if before != nil {
		q.Set("before", before.String())
	}
	if after != nil {
		q.Set("after", after.String())
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var guildScheduledEvent *GuildScheduledEventUser
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guildScheduledEvent)

	return guildScheduledEvent, err
}
