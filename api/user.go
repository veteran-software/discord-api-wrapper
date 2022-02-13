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

/* USER OBJECT */

/*
User

ID: the user's id

Username: the user's username, not unique across the platform

Discriminator: the user's 4-digit discord-tag

Avatar: the user's avatar hash

Bot: whether the user belongs to an OAuth2 application

System: whether the user is an Official Discord System user (part of the urgent message system)

MfaEnabled: whether the user has two factor enabled on their account

Locale:
*/
type User struct {
	ID            Snowflake `json:"id,omitempty"`
	Username      string    `json:"username,omitempty"`
	Discriminator string    `json:"discriminator,omitempty"`
	Avatar        *string   `json:"avatar"`
	PublicFlags   UserFlags `json:"public_flags,omitempty"`

	// No param below this have been seen in payloads from interactions
	Bot         bool        `json:"bot,omitempty"`
	System      bool        `json:"system,omitempty"`
	MfaEnabled  bool        `json:"mfa_enabled,omitempty"`
	Banner      *string     `json:"banner,omitempty"`
	BannerColor string      `json:"banner_color,omitempty"` // Undocumented as of 10/31/21
	AccentColor *uint       `json:"accent_color,omitempty"`
	Locale      string      `json:"locale,omitempty"`
	Flags       UserFlags   `json:"flags,omitempty"`
	PremiumType PremiumType `json:"premium_type,omitempty"`

	// Below required `email` OAuth2 scope
	Verified bool    `json:"verified,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type UserFlags int

//goland:noinspection SpellCheckingInspection
const (
	FlagsNone             UserFlags = iota
	Staff                 UserFlags = 1 << 0
	Partner               UserFlags = 1 << 1
	HypeSquad             UserFlags = 1 << 2
	BugHunterLevel1       UserFlags = 1 << 3
	HouseBravery          UserFlags = 1 << 6
	HouseBrilliance       UserFlags = 1 << 7
	HouseBalance          UserFlags = 1 << 8
	PremiumEarlySupporter UserFlags = 1 << 9
	TeamPsuedoUser        UserFlags = 1 << 10
	BugHunterLevel2       UserFlags = 1 << 14
	VerifiedBot           UserFlags = 1 << 16
	VerifiedDeveloper     UserFlags = 1 << 17
	CertifiedModerator    UserFlags = 1 << 18
	BotHttpInteractions   UserFlags = 1 << 19
)

type PremiumType int

const (
	None PremiumType = iota
	NitroClassic
	Nitro
)

/* CONNECTION OBJECT */

type Connection struct {
	ID           string                   `json:"id"`
	Name         string                   `json:"name"`
	Type         string                   `json:"type"`
	Revoked      bool                     `json:"revoked,omitempty"`
	Integrations []Integration            `json:"integrations,omitempty"`
	Verified     bool                     `json:"verified"`
	FriendSync   bool                     `json:"friend_sync"`
	ShowActivity bool                     `json:"show_activity"`
	Visibility   ConnectionVisibilityType `json:"visibility"`
}

type ConnectionVisibilityType int

const (
	ConnectionVisibilityTypeNone ConnectionVisibilityType = iota
	ConnectionVisibilityTypeEveryone
)

/* ENDPOINTS */

// GetCurrentUser
//
// Returns the user object of the requester's account.
//
// For OAuth2, this requires the `identify` scope, which will return the object without an email, and optionally the email scope, which returns the object with an email.
func GetCurrentUser() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getCurrentUser, api)
}

/* HELPER METHODS */

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
