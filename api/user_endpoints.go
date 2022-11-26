/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
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
	"github.com/veteran-software/discord-api-wrapper/v10/logging"
	"github.com/veteran-software/discord-api-wrapper/v10/utilities"
	"github.com/vincent-petithory/dataurl"
	"io"
	"net/http"
	"strconv"
)

// GetCurrentUser - Returns the user object of the requesters account.
//
// For OAuth2, this requires the `identify` scope, which will return the object without an email, and optionally the email scope, which returns the object with an email.
//
//goland:noinspection GoUnusedExportedFunction
func GetCurrentUser(scopes []string) *User {
	// we have to have at least the `identify` scope to use this endpoint
	if len(scopes) == 0 || !utilities.Contains(scopes, "identify") {
		return nil
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(getCurrentUser, api), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var user *User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return user
}

// GetUser - Returns a User object for a given user ID.
func (u *User) GetUser() *User {
	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(getUser, api, u.ID), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var user *User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return user
}

// ModifyCurrentUser - Modify the requesters User account settings. Returns a User object on success. Fires a User Update Gateway event.
//
// All parameters to this endpoint are optional.
//
//goland:noinspection GoUnusedExportedFunction
func ModifyCurrentUser(username *string, avatar *dataurl.DataURL) *User {
	payload := struct {
		Username string `json:"username,omitempty"`
		Avatar   string `json:"avatar,omitempty"`
	}{}

	if username != nil {
		payload.Username = *username
	}
	if avatar != nil {
		payload.Avatar = avatar.String()
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(modifyCurrentUser, api), payload, nil)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var user *User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return user
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
