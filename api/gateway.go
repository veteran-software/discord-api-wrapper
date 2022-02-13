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

const (
	gatewayVersion        = 9
	gatewayURLQueryString = "?v=" + string(rune(gatewayVersion)) + "&encoding=json"
)

// GatewayPayload - S and T are null when Op is not 0 (Gateway Dispatch Opcode).
type GatewayPayload struct {
	Op int          `json:"op"` // Op - opcode for the payload
	D  *interface{} `json:"d"`  // D - event data
	S  *int         `json:"s"`  // S - sequence number, used for resuming sessions and heartbeats
	T  *string      `json:"t"`  // T - the event name for this payload
}

type GatewayIntents int64

//goland:noinspection GoSnakeCaseUsage
const (
	GUILDS                    GatewayIntents = 1 << 0
	GUILD_MEMBERS             GatewayIntents = 1 << 1
	GUILD_BANS                GatewayIntents = 1 << 2
	GUILD_EMOJIS_AND_STICKERS GatewayIntents = 1 << 3
	GUILD_INTEGRATIONS        GatewayIntents = 1 << 4
	GUILD_WEBHOOKS            GatewayIntents = 1 << 5
	GUILD_INVITES             GatewayIntents = 1 << 6
	GUILD_VOICE_STATES        GatewayIntents = 1 << 7
	GUILD_PRESENCES           GatewayIntents = 1 << 8
	GUILD_MESSAGES            GatewayIntents = 1 << 9
	GUILD_MESSAGE_REACTIONS   GatewayIntents = 1 << 10
	GUILD_MESSAGE_TYPING      GatewayIntents = 1 << 11
	DIRECT_MESSAGES           GatewayIntents = 1 << 12
	DIRECT_MESSAGE_REACTIONS  GatewayIntents = 1 << 13
	DIRECT_MESSAGE_TYPING     GatewayIntents = 1 << 14
	GUILD_SCHEDULE_EVENTS     GatewayIntents = 1 << 16
)

// Identify - Used to trigger the initial handshake with the gateway.
type Identify struct {
	Token          string                `json:"token"`                     // Token - authentication token
	Properties     string                `json:"properties"`                // Properties - IdentifyConnection properties
	Compress       bool                  `json:"compress,omitempty"`        // Compress - whether this connection supports compression of packets
	LargeThreshold int                   `json:"large_threshold,omitempty"` // LargeThreshold - value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	Shard          [2]int                `json:"shard,omitempty"`           // Shard - used for Guild Sharding
	Presence       GatewayPresenceUpdate `json:"presence,omitempty"`        // Presence - presence structure for initial presence information
	Intents        GatewayIntents        `json:"intents"`                   // Intents - the Gateway Intents you wish to receive
}

type IdentifyConnection struct {
	OS      string `json:"$os"`      // OS - your operating system
	Browser string `json:"$browser"` // Browser - your library name
	Device  string `json:"$device"`  // Device - your library name
}

// Resume - Used to replay missed events when a disconnected client resumes.
type Resume struct {
	Token     string `json:"token"`      // Token - session token
	SessionID string `json:"session_id"` // SessionID - session id
	Seq       int    `json:"seq"`        // Seq - last sequence number received
}

// GuildRequestMembers - Used to request all members for a guild or a list of guilds.
//
// When initially connecting, if you don't have the GUILD_PRESENCES Gateway Intent, or if the guild is over 75k members, it will only send members who are in voice, plus the member for you (the connecting user).
//
// Otherwise, if a guild has over large_threshold members (value in the Gateway Identify), it will only send members who are online, have a role, have a nickname, or are in a voice channel, and if it has under large_threshold members, it will send all members.
//
// If a client wishes to receive additional members, they need to explicitly request them via this operation.
//
// The server will send Guild Members Chunk events in response with up to 1000 members per chunk until all members that match the request have been sent.
//
// Due to our privacy and infrastructural concerns with this feature, there are some limitations that apply:
//
//    GUILD_PRESENCES intent is required to set presences = true. Otherwise, it will always be false
//    GUILD_MEMBERS intent is required to request the entire member list—(query=‘’, limit=0<=n)
//    You will be limited to requesting 1 guild_id per request
//    Requesting a prefix (query parameter) will return a maximum of 100 members
//    Requesting user_ids will continue to be limited to returning 100 members
type GuildRequestMembers struct {
	GuildID   Snowflake   `json:"guild_id"`            // GuildID - id of the guild to get members for
	Query     string      `json:"query,omitempty"`     // Query - string that username starts with, or an empty string to return all members
	Limit     int         `json:"limit"`               // Limit - maximum number of members to send matching the query; a limit of 0 can be used with an empty string query to return all members
	Presences bool        `json:"presences,omitempty"` // Presences - used to specify if we want the presences of the matched members
	UserIDs   []Snowflake `json:"user_ids,omitempty"`  // UserIDs - used to specify which users you wish to fetch
	Nonce     string      `json:"nonce,omitempty"`     // Nonce - nonce to identify the Guild Members Chunk response
}

// GatewayVoiceStateUpdate - Sent when a client wants to join, move, or disconnect from a voice channel.
type GatewayVoiceStateUpdate struct {
	GuildID   Snowflake  `json:"guild_id"`   // GuildID - id of the Guild
	ChannelID *Snowflake `json:"channel_id"` // ChannelID - id of the voice channel client wants to join (null if disconnecting)
	SelfMute  bool       `json:"self_mute"`  // SelfMute - is the client muted
	SelfDeaf  bool       `json:"self_deaf"`  // SelfDeaf - is the client deafened
}

// GatewayPresenceUpdate - Sent by the client to indicate a presence or status update.
type GatewayPresenceUpdate struct {
	Since      *int       `json:"since"`      // Since - unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Activities []Activity `json:"activities"` // Activities - the user's activities
	Status     StatusType `json:"status"`     // Status - the user's new StatusType
	Afk        bool       `json:"afk"`        // Afk - whether the client is afk
}

// Hello - Sent on connection to the websocket. Defines the heartbeat interval that the client should heartbeat to.
type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"` // HeartbeatInterval - the interval (in milliseconds) the client should heartbeat with
}

// Ready - The ready event is dispatched when a client has completed the initial handshake with the gateway (for new sessions).
//
// The ready event can be the largest and most complex event the gateway will send, as it contains all the state required for a client to begin interacting with the rest of the platform.
//
// "guilds" are the guilds of which your bot is a member.
//
// They start out as unavailable when you connect to the gateway.
//
// As they become available, your bot will be notified via Guild Create events.
type Ready struct {
	V           int                `json:"v"`
	User        User               `json:"user"`
	Guilds      []UnavailableGuild `json:"guilds"`
	SessionID   string             `json:"session_id"`
	Shard       [2]int             `json:"shard,omitempty"`
	Application Application        `json:"application"`
}

type StatusType string

const (
	StatusTypeOnline       StatusType = "online"
	StatusTypeDoNotDisturb StatusType = "dnd"
	StatusTypeIdle         StatusType = "idle"
	StatusTypeInvisible    StatusType = "invisible"
	StatusTypeOffline      StatusType = "offline"
)

type PresenceStatus string

const (
	Idle    PresenceStatus = "idle"
	Dnd     PresenceStatus = "dnd"
	Online  PresenceStatus = "online"
	Offline PresenceStatus = "offline"
)

type ActivityTypes int8

const (
	Game ActivityTypes = iota
	Streaming
	Listening
	Watching
	Custom
	Competing
)

type ActivityTimestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type ActivityParty struct {
	ID   string  `json:"id,omitempty"`
	Size []int16 `json:"size,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type ActivityFlags int

const (
	Instance    ActivityFlags = 1 << 0
	Join        ActivityFlags = 1 << 1
	Spectate    ActivityFlags = 1 << 2
	JoinRequest ActivityFlags = 1 << 3
	Sync        ActivityFlags = 1 << 4
	Play        ActivityFlags = 1 << 5
)

type ActivityButtons struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Activity struct {
	Name          string             `json:"name"`
	Type          ActivityTypes      `json:"type"`
	URL           *string            `json:"url,omitempty"`
	CreatedAt     int64              `json:"created_at"`
	Timestamps    ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID Snowflake          `json:"application_id,omitempty"`
	Details       *string            `json:"details,omitempty"`
	State         *string            `json:"state,omitempty"`
	Emoji         *Emoji             `json:"emoji,omitempty"`
	Party         ActivityParty      `json:"party,omitempty"`
	Assets        ActivityAssets     `json:"assets,omitempty"`
	Secrets       ActivitySecrets    `json:"secrets,omitempty"`
	Instance      bool               `json:"instance,omitempty"`
	Flags         ActivityFlags      `json:"flags,omitempty"`
	Buttons       []ActivityButtons  `json:"buttons,omitempty"`
}

type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Web     string `json:"web,omitempty"`
}

type PresenceUpdateEvent struct {
	User         User           `json:"user"`
	GuildID      Snowflake      `json:"guild_id"`
	Status       PresenceStatus `json:"status"`
	Activities   []Activity     `json:"activities"`
	ClientStatus ClientStatus   `json:"client_status"`
}
