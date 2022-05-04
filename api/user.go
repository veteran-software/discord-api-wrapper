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
	"strconv"
)

// User - Discord enforces the following restrictions for usernames and nicknames:
//
//    Names can contain most valid unicode characters. We limit some zero-width and non-rendering characters.
//    Usernames must be between 2 and 32 characters long.
//    Nicknames must be between 1 and 32 characters long.
//    Names are sanitized and trimmed of leading, trailing, and excessive internal whitespace.
//
// The following restrictions are additionally enforced for usernames:
//
//    Names cannot contain the following substrings: '@', '#', ':', '```'.
//    Names cannot be: 'discordtag', 'everyone', 'here'.
//
// There are other rules and restrictions not shared here for the sake of spam and abuse mitigation, but the majority of users won't encounter them.
//
// It's important to properly handle all error messages returned by Discord when editing or updating names.
type User struct {
	ID            Snowflake   `json:"id,omitempty"`            // the user's id
	Username      string      `json:"username,omitempty"`      // the user's username, not unique across the platform
	Discriminator string      `json:"discriminator,omitempty"` // the user's 4-digit discord-tag
	Avatar        *string     `json:"avatar"`                  // the user's avatar hash
	Bot           bool        `json:"bot,omitempty"`           // whether the user belongs to an OAuth2 application
	System        bool        `json:"system,omitempty"`        // whether the user is an Official Discord System user (part of the urgent message system)
	MfaEnabled    bool        `json:"mfa_enabled,omitempty"`   // whether the user has two factor enabled on their account
	Banner        *string     `json:"banner,omitempty"`        // the user's banner hash
	BannerColor   string      `json:"banner_color,omitempty"`  // Undocumented as of 10/31/21
	AccentColor   *uint       `json:"accent_color,omitempty"`  // the user's banner color encoded as an integer representation of hexadecimal color code
	Locale        string      `json:"locale,omitempty"`        // the user's chosen language option
	Flags         UserFlags   `json:"flags,omitempty"`         // the flags on a user's account
	PremiumType   PremiumType `json:"premium_type,omitempty"`  // the type of Nitro subscription on a user's account
	PublicFlags   UserFlags   `json:"public_flags,omitempty"`  // the public flags on a user's account

	// Below required `email` OAuth2 scope
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
)

// PremiumType - Premium types denote the level of premium a user has. Visit the Nitro page to learn more about the premium plans we currently offer.
type PremiumType int

//goland:noinspection GoUnusedConst
const (
	None         PremiumType = iota // None
	NitroClassic                    // Nitro Classic
	Nitro                           // Nitro
)

// Connection - The connection object that the user has attached.
type Connection struct {
	ID           string                   `json:"id"`                     // id of the connection account
	Name         string                   `json:"name"`                   // the username of the connection account
	Type         string                   `json:"type"`                   // the service of the connection (Twitch, YouTube)
	Revoked      bool                     `json:"revoked,omitempty"`      // whether the connection is revoked
	Integrations []Integration            `json:"integrations,omitempty"` // an array of partial server integrations
	Verified     bool                     `json:"verified"`               // whether the connection is verified
	FriendSync   bool                     `json:"friend_sync"`            // whether friend sync is enabled for this connection
	ShowActivity bool                     `json:"show_activity"`          // whether activities related to this connection will be shown in presence updates
	Visibility   ConnectionVisibilityType `json:"visibility"`             // visibility of this connection
}

// ConnectionVisibilityType - visibility of this connection
type ConnectionVisibilityType int

//goland:noinspection GoUnusedConst
const (
	ConnectionVisibilityTypeNone     ConnectionVisibilityType = iota // invisible to everyone except the user themselves
	ConnectionVisibilityTypeEveryone                                 // visible to everyone
)

// GetCurrentUser - Returns the user object of the requesters account.
//
// For OAuth2, this requires the `identify` scope, which will return the object without an email, and optionally the `email` scope, which returns the object with an email.
//goland:noinspection GoUnusedExportedFunction
func GetCurrentUser() (*User, error) {
	u := parseRoute(fmt.Sprintf(getCurrentUser, api))

	var user *User
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &user)

	return user, err
}

// GetUser - Returns a User object for a given user ID.
//goland:noinspection GoUnusedExportedFunction
func (u *User) GetUser() (*User, error) {
	route := parseRoute(fmt.Sprintf(getUser, api, u.ID.String()))

	var user *User
	err := json.Unmarshal(fireGetRequest(route, nil, nil), &user)

	return user, err
}

// ModifyCurrentUser - Modify the requesters user account settings. Returns a User object on success.
//
//    All parameters to this endpoint are optional.
//goland:noinspection GoUnusedExportedFunction
func ModifyCurrentUser(payload ModifyCurrentUserJSON) (*User, error) {
	u := parseRoute(fmt.Sprintf(modifyCurrentUser, api))

	var user *User
	err := json.Unmarshal(firePatchRequest(u, payload, nil), &user)

	return user, err
}

// ModifyCurrentUserJSON - JSON payload
type ModifyCurrentUserJSON struct {
	Username string  `json:"username,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

// GetCurrentUserGuilds - Returns a list of partial Guild objects the current user is a member of. Requires the `guilds` OAuth2 scope.
//goland:noinspection GoUnusedExportedFunction
func GetCurrentUserGuilds(before *Snowflake, after *Snowflake, limit *uint64) ([]Guild, error) {
	u := parseRoute(fmt.Sprintf(getCurrentUserGuilds, api))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if after != nil {
		q.Set("after", after.String())
	}
	if limit != nil {
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var guilds []Guild
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guilds)

	return guilds, err
}

// GetCurrentUserGuildMember - Returns a GuildMember object for the current user. Requires the `guilds.members.read` OAuth2 scope.
func (g *Guild) GetCurrentUserGuildMember() (*GuildMember, error) {
	u := parseRoute(fmt.Sprintf(getCurrentUserGuildMember, api, g.ID.String()))

	var guildMember *GuildMember
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guildMember)

	return guildMember, err
}

// LeaveGuild - Leave a guild. Returns a 204 empty response on success.
func (g *Guild) LeaveGuild() error {
	u := parseRoute(fmt.Sprintf(leaveGuild, api, g.ID.String()))

	return fireDeleteRequest(u, nil)
}

// CreateDM - Create a new DM Channel with a User. Returns a DM Channel object.
//
//    You should not use this endpoint to DM everyone in a server about something.
//
//    DMs should generally be initiated by a user action.
//
//    If you open a significant amount of DMs too quickly, your bot may be rate limited or blocked from opening new ones.
//goland:noinspection GoUnusedExportedFunction
func CreateDM(payload CreateDmJSON) (*Channel, error) {
	u := parseRoute(fmt.Sprintf(createDM, api))

	var channel *Channel
	err := json.Unmarshal(firePostRequest(u, payload, nil), &channel)

	return channel, err
}

// CreateDmJSON - JSON payload
type CreateDmJSON struct {
	RecipientID Snowflake `json:"recipient_id"` // the recipient to open a DM channel with
}

// CreateGroupDM - Create a new group DM Channel with multiple users. Returns a DM channel object.
//
// This endpoint was intended to be used with the now-deprecated GameBridge SDK.
//
// DMs created with this endpoint will not be shown in the Discord client
//
//    This endpoint is limited to 10 active group DMs.
//goland:noinspection GoUnusedExportedFunction
func CreateGroupDM(payload CreateDmJSON) (*Channel, error) {
	u := parseRoute(fmt.Sprintf(createGroupDM, api))

	var channel *Channel
	err := json.Unmarshal(firePostRequest(u, payload, nil), &channel)

	return channel, err
}

// CreateGroupDmJSON - JSON payload
type CreateGroupDmJSON struct {
	AccessTokens []string             `json:"access_tokens"` // access tokens of users that have granted your app the `gdm.join` scope
	Nicks        map[Snowflake]string `json:"nicks"`         // a dictionary of user ids to their respective nicknames
}

// GetUserConnections - Returns a list of Connection objects. Requires the `connections` OAuth2 scope.
//goland:noinspection GoUnusedExportedFunction
func GetUserConnections() ([]Connection, error) {
	u := parseRoute(fmt.Sprintf(getUserConnections, api))

	var connections []Connection
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &connections)

	return connections, err
}

// GetAvatarUrl - returns a properly formatted avatar url
func (u *User) GetAvatarUrl() string {
	if u.Avatar != nil {
		if PtrStr(u.Avatar)[:2] == "a_" {
			return ImageBaseURL + fmt.Sprintf(getAvatarUrlGif, u.ID, PtrStr(u.Avatar))
		}
	}

	return ImageBaseURL + fmt.Sprintf(getAvatarUrlPng, u.ID, PtrStr(u.Avatar))
}

// GetDefaultUserAvatarUrl - returns the default Discord avatar
func (u *User) GetDefaultUserAvatarUrl() string {
	discriminator, err := strconv.Atoi(u.Discriminator)
	if err != nil {
		return ""
	}

	return ImageBaseURL + fmt.Sprintf(getDefaultUserAvatarUrl, strconv.Itoa(discriminator%5))
}
