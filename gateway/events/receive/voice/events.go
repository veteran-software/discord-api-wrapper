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

package voice

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* VOICE */

type (
	// StateUpdate - Sent when someone joins/leaves/moves voice channels. Inner payload is a voice state object.
	StateUpdate api.VoiceState

	// ServerUpdate - Sent when a guild's voice server is updated.
	// This is sent when initially connecting to voice, and when the current voice instance fails over to a new server.
	//
	// A `null` endpoint means that the voice server allocated has gone away and is trying to be reallocated.
	// You should attempt to disconnect from the currently connected voice server, and not attempt to reconnect until a new voice server is allocated.
	ServerUpdate struct {
		Token    string        `json:"token"`
		GuildID  api.Snowflake `json:"guild_id"`
		Endpoint *string       `json:"endpoint"`
	}
)
