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

// User - Discord enforces the following restrictions for usernames and nicknames:
//
//	Names can contain most valid unicode characters. We limit some zero-width and non-rendering characters.
//	Usernames must be between 2 and 32 characters long.
//	Nicknames must be between 1 and 32 characters long.
//	Names are sanitized and trimmed of leading, trailing, and excessive internal whitespace.
//
// The following restrictions are additionally enforced for usernames:
//
//	Names cannot contain the following substrings: '@', '#', ':', '```', 'discord'.
//	Names cannot be: 'everyone', 'here'.
//
// There are other rules and restrictions not shared here for the sake of spam and abuse mitigation, but the majority of users won't encounter them.
//
// It's important to properly handle all error messages returned by Discord when editing or updating names.
type User struct {
	ID               Snowflake   `json:"id,omitempty"`                // the user's id
	Username         string      `json:"username,omitempty"`          // the user's username, not unique across the platform
	Discriminator    string      `json:"discriminator,omitempty"`     // the user's 4-digit discord-tag
	Avatar           *string     `json:"avatar"`                      // the user's avatar hash
	Bot              bool        `json:"bot,omitempty"`               // whether the user belongs to an OAuth2 application
	System           bool        `json:"system,omitempty"`            // whether the user is an Official Discord System user (part of the urgent message system)
	MfaEnabled       bool        `json:"mfa_enabled,omitempty"`       // whether the user has two factor enabled on their account
	Banner           *string     `json:"banner,omitempty"`            // the user's banner hash
	BannerColor      *string     `json:"banner_color,omitempty"`      // Undocumented as of 10/31/21
	AccentColor      *uint       `json:"accent_color,omitempty"`      // the user's banner color encoded as an integer representation of hexadecimal color code
	Locale           string      `json:"locale,omitempty"`            // the user's chosen language option
	Flags            UserFlags   `json:"flags,omitempty"`             // the flags on a user's account
	PremiumType      PremiumType `json:"premium_type,omitempty"`      // the type of Nitro subscription on a user's account
	PublicFlags      UserFlags   `json:"public_flags,omitempty"`      // the public flags on a user's account
	GlobalName       *string     `json:"global_name,omitempty"`       // UNDOCUMENTED AS OF 3/23/2023
	DisplayName      *string     `json:"display_name,omitempty"`      // UNDOCUMENTED AS OF 3/23/2023
	AvatarDecoration *string     `json:"avatar_decoration,omitempty"` // UNDOCUMENTED AS OF 3/23/2023

	// Below require `email` OAuth2 scope
	Verified bool    `json:"verified,omitempty"` // whether the email on this account has been verified
	Email    *string `json:"email,omitempty"`    // the user's email
}

// UserFlags - public flags on a User account, many display badges on a User profile
type UserFlags uint64

//goland:noinspection SpellCheckingInspection,GoUnusedConst
const (
	FlagsNone             UserFlags = iota    // None
	Staff                 UserFlags = 1 << 0  // Discord Employee
	Partner               UserFlags = 1 << 1  // Partnered Server Owner
	HypeSquad             UserFlags = 1 << 2  // HypeSquad Events Coordinator
	BugHunterLevel1       UserFlags = 1 << 3  // Bug Hunter Level 1
	HouseBravery          UserFlags = 1 << 6  // House Bravery Member
	HouseBrilliance       UserFlags = 1 << 7  // House Brilliance Member
	HouseBalance          UserFlags = 1 << 8  // House Balance Member
	PremiumEarlySupporter UserFlags = 1 << 9  // Early Nitro Supporter
	TeamPsuedoUser        UserFlags = 1 << 10 // User is a team
	BugHunterLevel2       UserFlags = 1 << 14 // Bug Hunter Level 2
	VerifiedBot           UserFlags = 1 << 16 // Verified Bot
	VerifiedDeveloper     UserFlags = 1 << 17 // Early Verified Bot Developer
	CertifiedModerator    UserFlags = 1 << 18 // Discord Certified Moderator
	BotHttpInteractions   UserFlags = 1 << 19 // Bot uses only HTTP interactions and is shown in the online member list
	ActiveDeveloper       UserFlags = 1 << 22 // User is an Active Developer
)

// PremiumType - Premium types denote the level of premium a user has. Visit the Nitro page to learn more about the premium plans we currently offer.
type PremiumType int

//goland:noinspection GoUnusedConst
const (
	None         PremiumType = iota // None
	NitroClassic                    // Nitro Classic
	Nitro                           // Nitro
	NitroBasic                      // Nitro Basic
)

// Connection - The connection object that the user has attached.
type Connection struct {
	ID           string                   `json:"id"`                     // id of the connection account
	Name         string                   `json:"name"`                   // the username of the connection account
	Type         Service                  `json:"type"`                   // the service of the connection (Twitch, YouTube)
	Revoked      bool                     `json:"revoked,omitempty"`      // whether the connection is revoked
	Integrations []*Integration           `json:"integrations,omitempty"` // an array of partial server integrations
	Verified     bool                     `json:"verified"`               // whether the connection is verified
	FriendSync   bool                     `json:"friend_sync"`            // whether friend sync is enabled for this connection
	ShowActivity bool                     `json:"show_activity"`          // whether activities related to this connection will be shown in presence updates
	TwoWayLink   bool                     `json:"two_way_link"`           // whether this connection has a corresponding third party OAuth2 token
	Visibility   ConnectionVisibilityType `json:"visibility"`             // visibility of this connection
}

// Service - the service of the connection
type Service string

//goland:noinspection SpellCheckingInspection,GoUnusedConst
const (
	BattleNet       Service = "battlenet"
	CrunchyRoll     Service = "crunchyroll" // Undocumented as of 03/17/2023
	Ebay            Service = "ebay"
	EpicGames       Service = "epicgames"
	Facebook        Service = "facebook"
	GitHub          Service = "github"
	Instagram       Service = "instagram"
	LeagueOfLegends Service = "leagueoflegends"
	PayPal          Service = "paypal"
	PlayStation     Service = "playstation"
	Reddit          Service = "reddit"
	RiotGames       Service = "riotgames"
	Spotify         Service = "spotify"
	Skype           Service = "skype" // No longer to be added by users
	Steam           Service = "steam"
	TikTok          Service = "tiktok"
	Twitch          Service = "twitch"
	Twitter         Service = "twitter"
	Xbox            Service = "xbox"
	YouTube         Service = "youtube"
)

// ConnectionVisibilityType - visibility of this connection
type ConnectionVisibilityType int

//goland:noinspection GoUnusedConst
const (
	ConnectionVisibilityTypeNone     ConnectionVisibilityType = iota // invisible to everyone except the user themselves
	ConnectionVisibilityTypeEveryone                                 // visible to everyone
)

type ApplicationRoleConnection struct {
	PlatformName     *string                           `json:"platform_name"`
	PlatformUsername *string                           `json:"platform_username"`
	Metadata         ApplicationRoleConnectionMetadata `json:"metadata"`
}
