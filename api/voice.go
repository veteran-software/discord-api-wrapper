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
	"fmt"
	"net/http"
	"time"
)

// VoiceState - Used to represent a user's voice connection status.
type VoiceState struct {
	GuildID                 Snowflake   `json:"guild_id,omitempty"`                   // the guild id this voice state is for
	ChannelID               Snowflake   `json:"channel_id,omitempty"`                 // the channel id this user is connected to
	UserID                  Snowflake   `json:"user_id"`                              // the user id this voice state is for
	Member                  GuildMember `json:"member,omitempty"`                     // the guild member this voice state is for
	SessionID               string      `json:"session_id"`                           // the session id for this voice state
	Deaf                    bool        `json:"deaf"`                                 // whether this user is deafened by the server
	Mute                    bool        `json:"mute"`                                 // whether this user is muted by the server
	SelfDeaf                bool        `json:"self_deaf"`                            // whether this user is locally deafened
	SelfMute                bool        `json:"self_mute"`                            // whether this user is locally muted
	SelfStream              bool        `json:"self_stream,omitempty"`                // whether this user is streaming using "Go Live"
	SelfVideo               bool        `json:"self_video"`                           // whether this user's camera is enabled
	Suppress                bool        `json:"suppress"`                             // whether this user is muted by the current user
	RequestToSpeakTimestamp time.Time   `json:"request_to_speak_timestamp,omitempty"` // the time at which the user requested to speak
}

// VoiceRegion - representation of a geographic voice server
type VoiceRegion struct {
	ID         string `json:"id"`         // unique ID for the region
	Name       string `json:"name"`       // name of the region
	Optimal    bool   `json:"optimal"`    // true for a single server that is closest to the current user's client
	Deprecated bool   `json:"deprecated"` // whether this is a deprecated voice region (avoid switching to these)
	Custom     bool   `json:"custom"`     // whether this is a custom voice region (used for events/etc)
}

// ListVoiceRegions - Returns an array of voice region objects that can be used when setting a voice or stage channel's rtc_region.
func ListVoiceRegions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(listVoiceRegions, api)
}
