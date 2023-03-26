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

package send

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
	"github.com/veteran-software/discord-api-wrapper/v10/gateway"
	"github.com/veteran-software/discord-api-wrapper/v10/gateway/events/receive/presence"
)

// Identify - Used to trigger the initial handshake with the gateway.
type Identify struct {
	Token          string               `json:"token"`                     // authentication token
	Properties     ConnectionProperties `json:"properties"`                // ConnectionProperties properties
	Compress       bool                 `json:"compress,omitempty"`        // whether this connection supports compression of packets
	LargeThreshold int                  `json:"large_threshold,omitempty"` // value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	Shard          [2]int               `json:"shard,omitempty"`           // used for Guild Sharding
	Presence       PresenceUpdate       `json:"presence,omitempty"`        // presence structure for initial presence information
	Intents        gateway.Intents      `json:"intents"`                   // the Gateway Intents you wish to receive
}

// ConnectionProperties - properties
type ConnectionProperties struct {
	OS      string `json:"os"`      // your operating system
	Browser string `json:"browser"` // your library name
	Device  string `json:"device"`  // your library name
}

// Resume - Used to replay missed events when a disconnected client resumes.
type Resume struct {
	Token     string `json:"token"`      // session token
	SessionID string `json:"session_id"` // session id
	Seq       int    `json:"seq"`        // last sequence number received
}

// Heartbeat - Used to maintain an active gateway connection.
//
// Must be sent every heartbeat_interval milliseconds after the Opcode 10 Hello payload is received.
//
// The inner `d` key is the last sequence number — s — received by the client.
//
// If you have not yet received one, send `null`.
type Heartbeat struct {
	Op int `json:"op"`
	D  int `json:"d"`
}

// GuildRequestMembers - Used to request all members for a guild or a list of guilds.
//
// When initially connecting, if you don't have the GuildPresences Gateway Intent, or if the guild is over 75k members, it will only send members who are in voice, plus the member for you (the connecting user).
//
// Otherwise, if a guild has over large_threshold members (value in the Gateway Identify), it will only send members who are online, have a role, have a nickname, or are in a voice channel, and if it has under large_threshold members, it will send all members.
//
// If a client wishes to receive additional members, they need to explicitly request them via this operation.
//
// The server will send Guild Members Chunk events in response with up to 1000 members per chunk until all members that match the request have been sent.
//
// Due to our privacy and infrastructural concerns with this feature, there are some limitations that apply:
//
//	GuildPresences intent is required to set `presences = true`. Otherwise, it will always be false
//	GuildMembers intent is required to request the entire member list—(query=‘’, limit=0<=n)
//	You will be limited to requesting 1 `guild_id` per request
//	Requesting a prefix (`query` parameter) will return a maximum of 100 members
//	Requesting `user_ids` will continue to be limited to returning 100 members
type GuildRequestMembers struct {
	GuildID   api.Snowflake    `json:"guild_id"`            // id of the guild to get members for
	Query     string           `json:"query,omitempty"`     // string that username starts with, or an empty string to return all members
	Limit     int              `json:"limit"`               // maximum number of members to send matching the query; a limit of 0 can be used with an empty string query to return all members
	Presences bool             `json:"presences,omitempty"` // used to specify if we want the presences of the matched members
	UserIDs   []*api.Snowflake `json:"user_ids,omitempty"`  // used to specify which users you wish to fetch
	Nonce     string           `json:"nonce,omitempty"`     // nonce to identify the Guild Members Chunk response
}

// VoiceStateUpdate - Sent when a client wants to join, move, or disconnect from a voice channel.
type VoiceStateUpdate struct {
	GuildID   api.Snowflake  `json:"guild_id"`   // id of the Guild
	ChannelID *api.Snowflake `json:"channel_id"` // id of the voice channel client wants to join (null if disconnecting)
	SelfMute  bool           `json:"self_mute"`  // is the client muted
	SelfDeaf  bool           `json:"self_deaf"`  // is the client deafened
}

// PresenceUpdate - Sent by the client to indicate a presence or status update.
type PresenceUpdate struct {
	Since      *int                 `json:"since"`      // Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Activities []*presence.Activity `json:"activities"` // User's activities
	Status     StatusType           `json:"status"`     // User's new StatusType
	Afk        bool                 `json:"afk"`        // Whether or not the client is afk
}

// StatusType - a user's current activity status
type StatusType string

//goland:noinspection GoUnusedConst
const (
	StatusTypeOnline       StatusType = "online"    // Online
	StatusTypeDoNotDisturb StatusType = "dnd"       // Do Not Disturb
	StatusTypeIdle         StatusType = "idle"      // AFK
	StatusTypeInvisible    StatusType = "invisible" // Invisible and shown as offline
	StatusTypeOffline      StatusType = "offline"   // Offline
)
