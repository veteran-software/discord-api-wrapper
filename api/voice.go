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

/* VOICE STATE OBJECT */

type VoiceState struct {
	GuildID                 Snowflake   `json:"guild_id,omitempty"`
	ChannelID               *Snowflake  `json:"channel_id"`
	UserID                  Snowflake   `json:"user_id"`
	Member                  GuildMember `json:"member,omitempty"`
	SessionID               string      `json:"session_id"`
	Deaf                    bool        `json:"deaf"`
	Mute                    bool        `json:"mute"`
	SelfDeaf                bool        `json:"self_deaf"`
	SelfMute                bool        `json:"self_mute"`
	SelfStream              bool        `json:"self_stream,omitempty"`
	SelfVideo               bool        `json:"self_video"`
	Suppress                bool        `json:"suppress"`
	RequestToSpeakTimestamp *time.Time  `json:"request_to_speak_timestamp"`
}

/* VOICE REGION OBJECT */

type VoiceRegion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}

/* ENDPOINTS */

func ListVoiceRegions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(listVoiceRegions, api)
}
