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
	"fmt"
)

// StageInstance - A StageInstance holds information about a live stage.
type StageInstance struct {
	ID                    Snowflake    `json:"id"`                       // The id of this Stage instance
	GuildID               Snowflake    `json:"guild_id"`                 // The guild id of the associated Stage channel
	ChannelID             Snowflake    `json:"channel_id"`               // The id of the associated Stage channel
	Topic                 string       `json:"topic"`                    // The topic of the Stage instance (1-120 characters)
	PrivacyLevel          PrivacyLevel `json:"privacy_level"`            // The privacy level of the Stage instance
	GuildScheduledEventID *Snowflake   `json:"guild_scheduled_event_id"` // The id of the scheduled event for this Stage instance
}

// PrivacyLevel - The privacy level of the Stage instance
type PrivacyLevel int

//goland:noinspection GoUnusedConst
const (
	GuildOnly PrivacyLevel = iota + 2 // The Stage instance is visible to only guild members.
)

// CreateStageInstance - Creates a new Stage instance associated to a Stage channel.
//
// Requires the user to be a moderator of the Stage channel.
//
//    This endpoint supports the X-Audit-Log-Reason header.
//goland:noinspection GoUnusedExportedFunction
func CreateStageInstance(payload CreateStageInstanceJSON, reason *string) (*StageInstance, error) {
	u := parseRoute(fmt.Sprintf(createStageInstance, api))

	var stageInstance *StageInstance
	err := json.Unmarshal(firePostRequest(u, payload, reason), &stageInstance)

	return stageInstance, err
}

// CreateStageInstanceJSON - data to send in the CreateStageInstance payload
type CreateStageInstanceJSON struct {
	ChannelID             Snowflake    `json:"channel_id"`                        // The id of the Stage channel
	Topic                 string       `json:"topic"`                             // The topic of the Stage instance (1-120 characters)
	PrivacyLevel          PrivacyLevel `json:"privacy_level"`                     // The privacy level of the Stage instance (default GuildOnly)
	SendStartNotification bool         `json:"send_start_notification,omitempty"` // Notify @everyone that a Stage instance has started
}

// GetStageInstance - Gets the stage instance associated with the Stage channel, if it exists.
func (s *StageInstance) GetStageInstance() (*StageInstance, error) {
	u := parseRoute(fmt.Sprintf(getStageInstance, api, s.ChannelID.String()))

	var stageInstance *StageInstance
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &stageInstance)

	return stageInstance, err
}

// ModifyStageInstance - Updates fields of an existing Stage instance.
//
//  Requires the user to be a moderator of the Stage channel.
//
//    This endpoint supports the `X-Audit-Log-Reason` header.
func (s *StageInstance) ModifyStageInstance(payload ModifyStageInstanceJSON, reason *string) (*StageInstance, error) {
	u := parseRoute(fmt.Sprintf(modifyStageInstance, api, s.ChannelID.String()))

	var stageInstance *StageInstance
	err := json.Unmarshal(firePatchRequest(u, payload, reason), &stageInstance)

	return stageInstance, err
}

// ModifyStageInstanceJSON - data to send in the ModifyStageInstance payload
type ModifyStageInstanceJSON struct {
	Topic        string       `json:"topic,omitempty"` // The topic of the Stage instance (1-120 characters)
	PrivacyLevel PrivacyLevel `json:"privacy_level"`   // The PrivacyLevel of the Stage instance
}

// DeleteStageInstance - Deletes the Stage instance. Returns `204 No Content`.
//
// Requires the user to be a moderator of the Stage channel.
//
//    This endpoint supports the `X-Audit-Log-Reason` header.
func (s *StageInstance) DeleteStageInstance(reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteStageInstance, api, s.ChannelID.String()))

	return fireDeleteRequest(u, reason)
}
