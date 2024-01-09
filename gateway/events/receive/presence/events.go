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

package presence

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* PRESENCE */

type (
	Status       string
	ClientStatus struct {
		Desktop string `json:"desktop,omitempty"` // the user's status set for an active desktop (Windows, Linux, Mac) application session
		Mobile  string `json:"mobile,omitempty"`  // the user's status set for an active mobile (iOS, Android) application session
		Web     string `json:"web,omitempty"`     // the user's status set for an active web (browser, bot account) application session
	}

	// ActivityFlag - describes what the payload includes
	ActivityFlag uint

	// ActivityType - The streaming type currently only supports Twitch and YouTube.
	//
	// Only https://twitch.tv/ and https://youtube.com/ urls will work.
	ActivityType uint8
)

//goland:noinspection GoUnusedConst
const (
	Idle    Status = "idle"    // Online
	Dnd     Status = "dnd"     // Do Not Disturb
	Online  Status = "online"  // Online
	Offline Status = "offline" // Offline

	ActivityFlagInstance                 ActivityFlag = 1 << 0
	ActivityFlagJoin                     ActivityFlag = 1 << 1
	ActivityFlagSpectate                 ActivityFlag = 1 << 2
	ActivityFlagJoinRequest              ActivityFlag = 1 << 3
	ActivityFlagSync                     ActivityFlag = 1 << 4
	ActivityFlagPlay                     ActivityFlag = 1 << 5
	ActivityFlagPartyPrivacyFriends      ActivityFlag = 1 << 6
	ActivityFlagPartyPrivacyVoiceChannel ActivityFlag = 1 << 7
	ActivityFlagEmbedded                 ActivityFlag = 1 << 8

	Game      ActivityType = iota // Playing {name}
	Streaming                     // Streaming {details}
	Listening                     // Listening to {name}
	Watching                      // Watching {name}
	Custom                        // {emoji} {name}
	Competing                     // Competing in {name}
)

type (
	// Update - A user's presence is their current state on a guild. This event is sent when a user's presence or info, such as name or avatar, is updated.
	Update struct {
		User         api.User      `json:"user"`          // the user presence is being updated for
		GuildID      api.Snowflake `json:"guild_id"`      // id of the guild
		Status       Status        `json:"status"`        // either "idle", "dnd", "online", or "offline"
		Activities   []*Activity   `json:"activities"`    // user's current activities
		ClientStatus ClientStatus  `json:"client_status"` // user's platform-dependent status
	}

	// Activity - represents a user activity
	Activity struct {
		Name          string             `json:"name"`                     // the activity's name
		Type          ActivityType       `json:"type"`                     // activity type
		URL           *string            `json:"url,omitempty"`            // stream url, is validated when type is 1
		CreatedAt     int64              `json:"created_at"`               // unix timestamp (in milliseconds) of when the activity was added to the user's session
		Timestamps    ActivityTimestamps `json:"timestamps,omitempty"`     // unix timestamps for start and/or end of the game
		ApplicationID api.Snowflake      `json:"application_id,omitempty"` // application id for the game
		Details       *string            `json:"details,omitempty"`        // what the player is currently doing
		State         *string            `json:"state,omitempty"`          // the user's current party status
		Emoji         *api.Emoji         `json:"emoji,omitempty"`          // the emoji used for a custom status
		Party         ActivityParty      `json:"party,omitempty"`          // information for the current party of the player
		Assets        ActivityAssets     `json:"assets,omitempty"`         // images for the presence and their hover texts
		Secrets       ActivitySecrets    `json:"secrets,omitempty"`        // secrets for Rich Presence joining and spectating
		Instance      bool               `json:"instance,omitempty"`       // whether the activity is an instanced game session
		Flags         ActivityFlag       `json:"flags,omitempty"`          // activity flags ORd together, describes what the payload includes
		Buttons       []*ActivityButtons `json:"buttons,omitempty"`        // the custom buttons shown in the Rich Presence (max 2)
	}

	// ActivityTimestamps - start and stop timestamps for an activity
	ActivityTimestamps struct {
		Start int64 `json:"start,omitempty"` // unix time (in milliseconds) of when the activity started
		End   int64 `json:"end,omitempty"`   // unix time (in milliseconds) of when the activity ends
	}

	// ActivityEmoji - representation of an emoji in a custom status
	ActivityEmoji struct {
		Name     string        `json:"name"`               // the name of the emoji
		ID       api.Snowflake `json:"id,omitempty"`       // the id of the emoji
		Animated bool          `json:"animated,omitempty"` // whether this emoji is animated
	}

	// ActivityParty - information for the current party of the player
	ActivityParty struct {
		ID   string    `json:"id,omitempty"`   // the id of the party
		Size [2]uint16 `json:"size,omitempty"` // the id of the party; used to show the party's current and maximum size
	}

	// ActivityAssets - images for the presence and their hover texts
	ActivityAssets struct {
		LargeImage string `json:"large_image,omitempty"` // see https://discord.com/developers/docs/topics/gateway#activity-object-activity-asset-image
		LargeText  string `json:"large_text,omitempty"`  // text displayed when hovering over the large image of the activity
		SmallImage string `json:"small_image,omitempty"` // see https://discord.com/developers/docs/topics/gateway#activity-object-activity-asset-image
		SmallText  string `json:"small_text,omitempty"`  // text displayed when hovering over the small image of the activity
	}

	// ActivitySecrets - secrets for Rich Presence joining and spectating
	ActivitySecrets struct {
		Join     string `json:"join,omitempty"`     // the secret for joining a party
		Spectate string `json:"spectate,omitempty"` // the secret for spectating a game
		Match    string `json:"match,omitempty"`    // the secret for a specific instanced match
	}

	// ActivityButtons - When received over the gateway, the buttons field is an array of strings, which are the button labels.
	//
	//	Bots cannot access a user's activity button URLs.
	//
	// When sending, the buttons field must be an array
	ActivityButtons struct {
		Label string `json:"label"` // the text shown on the button (1-32 characters)
		URL   string `json:"url"`   // the url opened when clicking the button (1-512 characters)
	}

	// TypingStart - Sent when a user starts typing in a channel.
	TypingStart struct {
		ChannelID api.Snowflake   `json:"channel_id"`
		GuildID   api.Snowflake   `json:"guild_id,omitempty"`
		UserID    api.Snowflake   `json:"user_id"`
		Timestamp uint64          `json:"timestamp"`
		Member    api.GuildMember `json:"member,omitempty"`
	}

	// UserUpdate - Sent when properties about the current bot's user change. Inner payload is a user object.
	UserUpdate api.User
)
