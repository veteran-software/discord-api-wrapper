/*
 * Copyright (c) 2022-2023. Veteran Software
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

//goland:noinspection GoUnusedConst
const (
	gatewayVersion        = 10
	gatewayURLQueryString = "?v=" + string(rune(gatewayVersion)) + "&encoding=json"
)

// GatewayPayload - S and T are null when Op is not 0 (Gateway Dispatch Opcode).
type GatewayPayload struct {
	Op int     `json:"op"` // Gateway opcode, which indicates the payload type
	D  *any    `json:"d"`  // Event data
	S  *int    `json:"s"`  // Sequence number, used for resuming sessions and heartbeats
	T  *string `json:"t"`  // Event name
}

// OpCode
//
// All gateway events in Discord are tagged with an opcode that denotes the payload type.
// Your connection to our gateway may also sometimes close.
// When it does, you will receive a close code that tells you what happened.
type OpCode int

//goland:noinspection GoUnusedConst
const (
	OpDispatch OpCode = iota
	OpHeartbeat
	OpIdentify
	OpPresenceUpdate
	OpVoiceStateUpdate
	OpResume OpCode = iota + 1
	OpReconnect
	OpRequestGuildMembers
	OpInvalidSession
	OpHello
	OpHeartbeatAck
)

// TODO: Gateway Close Event Codes

// GatewayIntents - Maintaining a stateful application can be difficult when it comes to the amount of data you're expected to process, especially at scale.
//
// Gateway Intents are a system to help you lower that computational burden.
//
// When identifying to the gateway, you can specify an intents parameter which allows you to conditionally subscribe to pre-defined "intents", groups of events defined by Discord.
//
// If you do not specify a certain intent, you will not receive any of the gateway events that are batched into that group.
type GatewayIntents uint64

//goland:noinspection GoUnusedConst
const (
	Guilds                      GatewayIntents = 1 << 0
	GuildMembers                GatewayIntents = 1 << 1
	GuildBans                   GatewayIntents = 1 << 2
	GuildEmojisAndStickers      GatewayIntents = 1 << 3
	GuildIntegrations           GatewayIntents = 1 << 4
	GuildWebhooks               GatewayIntents = 1 << 5
	GuildInvites                GatewayIntents = 1 << 6
	GuildVoiceStates            GatewayIntents = 1 << 7
	GuildPresences              GatewayIntents = 1 << 8
	GuildMessages               GatewayIntents = 1 << 9
	GuildMessageReactions       GatewayIntents = 1 << 10
	GuildMessageTyping          GatewayIntents = 1 << 11
	DirectMessages              GatewayIntents = 1 << 12
	DirectMessageReactions      GatewayIntents = 1 << 13
	DirectMessageTyping         GatewayIntents = 1 << 14
	MessageContent              GatewayIntents = 1 << 15
	GuildScheduleEvents         GatewayIntents = 1 << 16
	AutoModerationConfiguration GatewayIntents = 1 << 20
	AutoModerationExecution     GatewayIntents = 1 << 21
)

// Identify - Used to trigger the initial handshake with the gateway.
type Identify struct {
	Token          string                `json:"token"`                     // authentication token
	Properties     IdentifyConnection    `json:"properties"`                // IdentifyConnection properties
	Compress       bool                  `json:"compress,omitempty"`        // whether this connection supports compression of packets
	LargeThreshold int                   `json:"large_threshold,omitempty"` // value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	Shard          [2]int                `json:"shard,omitempty"`           // used for Guild Sharding
	Presence       GatewayPresenceUpdate `json:"presence,omitempty"`        // presence structure for initial presence information
	Intents        GatewayIntents        `json:"intents"`                   // the Gateway Intents you wish to receive
}

// IdentifyConnection - properties
type IdentifyConnection struct {
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
//	GuildPresences intent is required to set presences = true. Otherwise, it will always be false
//	GuildMembers intent is required to request the entire member list—(query=‘’, limit=0<=n)
//	You will be limited to requesting 1 guild_id per request
//	Requesting a prefix (query parameter) will return a maximum of 100 members
//	Requesting user_ids will continue to be limited to returning 100 members
type GuildRequestMembers struct {
	GuildID   Snowflake    `json:"guild_id"`            // id of the guild to get members for
	Query     string       `json:"query,omitempty"`     // string that username starts with, or an empty string to return all members
	Limit     int          `json:"limit"`               // maximum number of members to send matching the query; a limit of 0 can be used with an empty string query to return all members
	Presences bool         `json:"presences,omitempty"` // used to specify if we want the presences of the matched members
	UserIDs   []*Snowflake `json:"user_ids,omitempty"`  // used to specify which users you wish to fetch
	Nonce     string       `json:"nonce,omitempty"`     // nonce to identify the Guild Members Chunk response
}

// GatewayVoiceStateUpdate - Sent when a client wants to join, move, or disconnect from a voice channel.
type GatewayVoiceStateUpdate struct {
	GuildID   Snowflake  `json:"guild_id"`   // id of the Guild
	ChannelID *Snowflake `json:"channel_id"` // id of the voice channel client wants to join (null if disconnecting)
	SelfMute  bool       `json:"self_mute"`  // is the client muted
	SelfDeaf  bool       `json:"self_deaf"`  // is the client deafened
}

// GatewayPresenceUpdate - Sent by the client to indicate a presence or status update.
type GatewayPresenceUpdate struct {
	Since      *int        `json:"since"`      // unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Activities []*Activity `json:"activities"` // the user's activities
	Status     StatusType  `json:"status"`     // the user's new StatusType
	Afk        bool        `json:"afk"`        // whether the client is afk
}

// Hello - Sent on connection to the websocket. Defines the heartbeat interval that the client should heartbeat to.
type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"` // the interval (in milliseconds) the client should heartbeat with
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
	V                int                 `json:"v"`                  // gateway version
	User             User                `json:"user"`               // information about the user including email
	Guilds           []*UnavailableGuild `json:"guilds"`             // the guilds the user is in
	SessionID        string              `json:"session_id"`         // used for resuming connections
	ResumeGatewayURL string              `json:"resume_gateway_url"` // Gateway URL for resuming connections
	Shard            [2]int              `json:"shard,omitempty"`    // the shard information associated with this session, if sent when identifying
	Application      Application         `json:"application"`        // contains id and flags
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

// PresenceStatus - either "idle", "dnd", "online", or "offline"
type PresenceStatus string

//goland:noinspection GoUnusedConst
const (
	Idle    PresenceStatus = "idle"    // Online
	Dnd     PresenceStatus = "dnd"     // Do Not Disturb
	Online  PresenceStatus = "online"  // Online
	Offline PresenceStatus = "offline" // Offline
)

// ActivityType - The streaming type currently only supports Twitch and YouTube.
//
// Only https://twitch.tv/ and https://youtube.com/ urls will work.
type ActivityType uint8

//goland:noinspection GoUnusedConst
const (
	Game      ActivityType = iota // Playing {name}
	Streaming                     // Streaming {details}
	Listening                     // Listening to {name}
	Watching                      // Watching {name}
	Custom                        // {emoji} {name}
	Competing                     // Competing in {name}
)

// ActivityTimestamps - start and stop timestamps for an activity
type ActivityTimestamps struct {
	Start int64 `json:"start,omitempty"` // unix time (in milliseconds) of when the activity started
	End   int64 `json:"end,omitempty"`   // unix time (in milliseconds) of when the activity ends
}

// ActivityEmoji - representation of an emoji in a custom status
type ActivityEmoji struct {
	Name     string    `json:"name"`               // the name of the emoji
	ID       Snowflake `json:"id,omitempty"`       // the id of the emoji
	Animated bool      `json:"animated,omitempty"` // whether this emoji is animated
}

// ActivityParty - information for the current party of the player
type ActivityParty struct {
	ID   string    `json:"id,omitempty"`   // the id of the party
	Size [2]uint16 `json:"size,omitempty"` // the id of the party; used to show the party's current and maximum size
}

// ActivityAssets - images for the presence and their hover texts
type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"` // see https://discord.com/developers/docs/topics/gateway#activity-object-activity-asset-image
	LargeText  string `json:"large_text,omitempty"`  // text displayed when hovering over the large image of the activity
	SmallImage string `json:"small_image,omitempty"` // see https://discord.com/developers/docs/topics/gateway#activity-object-activity-asset-image
	SmallText  string `json:"small_text,omitempty"`  // text displayed when hovering over the small image of the activity
}

// ActivitySecrets - secrets for Rich Presence joining and spectating
type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`     // the secret for joining a party
	Spectate string `json:"spectate,omitempty"` // the secret for spectating a game
	Match    string `json:"match,omitempty"`    // the secret for a specific instanced match
}

// ActivityFlag - describes what the payload includes
type ActivityFlag uint

//goland:noinspection GoUnusedConst
const (
	ActivityFlagInstance                 ActivityFlag = 1 << 0
	ActivityFlagJoin                     ActivityFlag = 1 << 1
	ActivityFlagSpectate                 ActivityFlag = 1 << 2
	ActivityFlagJoinRequest              ActivityFlag = 1 << 3
	ActivityFlagSync                     ActivityFlag = 1 << 4
	ActivityFlagPlay                     ActivityFlag = 1 << 5
	ActivityFlagPartyPrivacyFriends      ActivityFlag = 1 << 6
	ActivityFlagPartyPrivacyVoiceChannel ActivityFlag = 1 << 7
	ActivityFlagEmbedded                 ActivityFlag = 1 << 8
)

// ActivityButtons - When received over the gateway, the buttons field is an array of strings, which are the button labels.
//
//	Bots cannot access a user's activity button URLs.
//
// When sending, the buttons field must be an array
type ActivityButtons struct {
	Label string `json:"label"` // the text shown on the button (1-32 characters)
	URL   string `json:"url"`   // the url opened when clicking the button (1-512 characters)
}

// Activity - represents a user activity
type Activity struct {
	Name          string             `json:"name"`                     // the activity's name
	Type          ActivityType       `json:"type"`                     // activity type
	URL           *string            `json:"url,omitempty"`            // stream url, is validated when type is 1
	CreatedAt     int64              `json:"created_at"`               // unix timestamp (in milliseconds) of when the activity was added to the user's session
	Timestamps    ActivityTimestamps `json:"timestamps,omitempty"`     // unix timestamps for start and/or end of the game
	ApplicationID Snowflake          `json:"application_id,omitempty"` // application id for the game
	Details       *string            `json:"details,omitempty"`        // what the player is currently doing
	State         *string            `json:"state,omitempty"`          // the user's current party status
	Emoji         *Emoji             `json:"emoji,omitempty"`          // the emoji used for a custom status
	Party         ActivityParty      `json:"party,omitempty"`          // information for the current party of the player
	Assets        ActivityAssets     `json:"assets,omitempty"`         // images for the presence and their hover texts
	Secrets       ActivitySecrets    `json:"secrets,omitempty"`        // secrets for Rich Presence joining and spectating
	Instance      bool               `json:"instance,omitempty"`       // whether the activity is an instanced game session
	Flags         ActivityFlag       `json:"flags,omitempty"`          // activity flags ORd together, describes what the payload includes
	Buttons       []*ActivityButtons `json:"buttons,omitempty"`        // the custom buttons shown in the Rich Presence (max 2)
}

// ClientStatus - Active sessions are indicated with an "online", "idle", or "dnd" string per platform.
//
// If a user is offline or invisible, the corresponding field is not present.
type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"` // the user's status set for an active desktop (Windows, Linux, Mac) application session
	Mobile  string `json:"mobile,omitempty"`  // the user's status set for an active mobile (iOS, Android) application session
	Web     string `json:"web,omitempty"`     // the user's status set for an active web (browser, bot account) application session
}

// PresenceUpdateEvent - If you are using Gateway Intents, you must specify the GuildPresences intent in order to receive Presence Update events
//
// A user's presence is their current state on a guild.
//
// This event is sent when a user's presence or info, such as name or avatar, is updated.
//
//	The user object within this event can be partial, the only field which must be sent is the id field, everything else is optional.
//	Along with this limitation, no fields are required, and the types of the fields are not validated.
//	Your client should expect any combination of fields and types within this event.
type PresenceUpdateEvent struct {
	User         User           `json:"user"`          // the user presence is being updated for
	GuildID      Snowflake      `json:"guild_id"`      // id of the guild
	Status       PresenceStatus `json:"status"`        // either "idle", "dnd", "online", or "offline"
	Activities   []*Activity    `json:"activities"`    // user's current activities
	ClientStatus ClientStatus   `json:"client_status"` // user's platform-dependent status
}
