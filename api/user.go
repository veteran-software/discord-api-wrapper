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
	"strconv"
)

/*
User







Locale:
*/
type User struct {
	ID            Snowflake   `json:"id,omitempty"`            // ID: the user's id
	Username      string      `json:"username,omitempty"`      // Username: the user's username, not unique across the platform
	Discriminator string      `json:"discriminator,omitempty"` // Discriminator: the user's 4-digit discord-tag
	Avatar        *string     `json:"avatar"`                  // Avatar: the user's avatar hash
	Bot           bool        `json:"bot,omitempty"`           // Bot: whether the user belongs to an OAuth2 application
	System        bool        `json:"system,omitempty"`        // System: whether the user is an Official Discord System user (part of the urgent message system)
	MfaEnabled    bool        `json:"mfa_enabled,omitempty"`   // MfaEnabled: whether the user has two factor enabled on their account
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

type ConnectionVisibilityType int

//goland:noinspection GoUnusedConst
const (
	ConnectionVisibilityTypeNone     ConnectionVisibilityType = iota // invisible to everyone except the user themselves
	ConnectionVisibilityTypeEveryone                                 // visible to everyone
)

// GetCurrentUser - Returns the user object of the requesters account.
//
// For OAuth2, this requires the `identify` scope, which will return the object without an email, and optionally the email scope, which returns the object with an email.
func GetCurrentUser() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getCurrentUser, api)
}

func (user *User) GetAvatarUrl() string {
	if user.Avatar != nil {
		if PtrStr(user.Avatar)[:2] == "a_" {
			return ImageBaseURL + fmt.Sprintf(getAvatarUrlGif, user.ID, PtrStr(user.Avatar))
		}
	}

	return ImageBaseURL + fmt.Sprintf(getAvatarUrlPng, user.ID, PtrStr(user.Avatar))
}

func (user *User) GetDefaultUserAvatarUrl() string {
	discriminator, err := strconv.Atoi(user.Discriminator)
	if err != nil {
		return ""
	}

	return ImageBaseURL + fmt.Sprintf(getDefaultUserAvatarUrl, strconv.Itoa(discriminator%5))
}
