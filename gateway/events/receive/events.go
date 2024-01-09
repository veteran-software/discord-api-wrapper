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

package receive

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

// Hello - Sent on connection to the websocket. Defines the heartbeat interval that the client should heartbeat to.
type Hello struct {
	HeartbeatInterval int      `json:"heartbeat_interval"` // the interval (in milliseconds) the client should heartbeat with
	Trace             []string `json:"_trace,omitempty"`   // Sent from the gateway but unused & undocumented
}

// Ready - The ready event is dispatched when a client has completed the initial handshake with the gateway (for new sessions).
//
// The ready event can be the largest and most complex event the gateway will send, as it contains all the state required for a client to begin interacting with the rest of the platform.
//
// `guilds` are the guilds of which your bot is a member.
//
// They start out as unavailable when you connect to the gateway.
//
// As they become available, your bot will be notified via Guild Create events.
type Ready struct {
	V                int                     `json:"v"`                  // gateway version
	User             api.User                `json:"user"`               // information about the user including email
	Guilds           []*api.UnavailableGuild `json:"guilds"`             // the guilds the user is in
	SessionID        string                  `json:"session_id"`         // used for resuming connections
	ResumeGatewayURL string                  `json:"resume_gateway_url"` // Gateway URL for resuming connections
	Shard            [2]int                  `json:"shard,omitempty"`    // the shard information associated with this session, if sent when identifying
	Application      api.Application         `json:"application"`        // contains id and flags
}

// Reconnect - The reconnect event is dispatched when a client should reconnect to the gateway (and resume their existing session, if they have one).
//
// This event usually occurs during deploys to migrate sessions gracefully off old hosts.
type Reconnect struct {
	Op int `json:"op"`
	D  any `json:"d"`
}

// InvalidSession - Sent to indicate one of at least three different situations:
//
//	the gateway could not initialize a session after receiving an Opcode 2 Identify
//	the gateway could not resume a previous session after receiving an Opcode 6 Resume
//	the gateway has invalidated an active session and is requesting client action
//
// The inner `d` key is a boolean that indicates whether the session may be resumable.
type InvalidSession struct {
	Op int  `json:"op"`
	D  bool `json:"d"`
}
