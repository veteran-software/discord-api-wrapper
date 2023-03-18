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

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// GetCurrentUser - Returns the user object of the requesters account.
//
// For OAuth2, this requires the `identify` scope, which will return the object without an email, and optionally the `email` scope, which returns the object with an email.
//
//goland:noinspection GoUnusedExportedFunction
func GetCurrentUser() (*User, error) {
	u := parseRoute(fmt.Sprintf(getCurrentUser, api))

	var user *User
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &user)

	return user, err
}

// GetUser - Returns a User object for a given user ID.
//
//goland:noinspection GoUnusedExportedFunction
func (u *User) GetUser() (*User, error) {
	route := parseRoute(fmt.Sprintf(getUser, api, u.ID.String()))

	var user *User
	err := json.Unmarshal(fireGetRequest(route, nil, nil), &user)

	return user, err
}

// ModifyCurrentUser - Modify the requesters user account settings. Returns a User object on success.
//
//	All parameters to this endpoint are optional.
//
//goland:noinspection GoUnusedExportedFunction
func ModifyCurrentUser(payload *ModifyCurrentUserJSON) (*User, error) {
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
//
//goland:noinspection GoUnusedExportedFunction
func GetCurrentUserGuilds(before *Snowflake, after *Snowflake, limit *uint64) ([]*Guild, error) {
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

	var guilds []*Guild
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
//	You should not use this endpoint to DM everyone in a server about something.
//
//	DMs should generally be initiated by a user action.
//
//	If you open a significant amount of DMs too quickly, your bot may be rate limited or blocked from opening new ones.
//
//goland:noinspection GoUnusedExportedFunction
func CreateDM(payload *CreateDmJSON) (*Channel, error) {
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
//	This endpoint is limited to 10 active group DMs.
//
//goland:noinspection GoUnusedExportedFunction
func CreateGroupDM(payload *CreateDmJSON) (*Channel, error) {
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
//
//goland:noinspection GoUnusedExportedFunction
func GetUserConnections() ([]*Connection, error) {
	u := parseRoute(fmt.Sprintf(getUserConnections, api))

	var connections []*Connection
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

// GetUserApplicationRoleConnection - Returns the application role connection for the user.
//
// Requires an OAuth2 access token with role_connections.write scope for the application specified in the path.
//
//goland:noinspection GoUnusedExportedFunction
func GetUserApplicationRoleConnection(applicationID *Snowflake) (*ApplicationRoleConnection, error) {
	u := parseRoute(fmt.Sprintf(getUserApplicationRoleConnection, api, applicationID.String()))

	var connection *ApplicationRoleConnection
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &connection)

	return connection, err
}

// UpdateUserApplicationRoleConnection - Updates and returns the application role connection for the user.
//
// Requires an OAuth2 access token with role_connections.write scope for the application specified in the path.
//
//goland:noinspection GoUnusedExportedFunction
func UpdateUserApplicationRoleConnection(applicationID *Snowflake,
	payload *ApplicationRoleConnection) (*ApplicationRoleConnection, error) {
	u := parseRoute(fmt.Sprintf(modifyUserApplicationRoleConnection, api, applicationID.String()))

	var connection *ApplicationRoleConnection
	err := json.Unmarshal(firePutRequest(u, payload, nil), &connection)

	return connection, err
}
