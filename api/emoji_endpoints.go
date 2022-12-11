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
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// ListGuildEmojis - Returns a list of emoji objects for the given guild.
func (g *Guild) ListGuildEmojis() ([]*Emoji, error) {
	u := parseRoute(fmt.Sprintf(listGuildEmojis, api, g.ID.String()))

	var emojis []*Emoji
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &emojis)

	return emojis, err
}

// GetGuildEmoji - Returns an emoji object for the given guild and emoji IDs.
func (g *Guild) GetGuildEmoji(emoji *Emoji) (*Emoji, error) {
	u := parseRoute(fmt.Sprintf(getGuildEmoji, api, g.ID.String(), emoji.ID.String()))

	var e *Emoji
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &e)

	return e, err
}

// CreateGuildEmoji - Create a new emoji for the guild.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns the new Emoji object on success.
//
// Fires a Guild Emojis Update Gateway event.
//
// Emojis and animated emojis have a maximum file size of 256kb.
//
// Attempting to upload an emoji larger than this limit will fail and return "400 Bad Request" and an error message, but not a JSON status code.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) CreateGuildEmoji(payload *CreateEmojiJSON, reason *string) (*Emoji, error) {
	u := parseRoute(fmt.Sprintf(createGuildEmoji, api, g.ID.String()))

	var emoji *Emoji
	err := json.Unmarshal(firePostRequest(u, payload, reason), &emoji)

	return emoji, err
}

// CreateEmojiJSON - Parameters to pass in the JSON payload
//
// TODO: Validate the base64.Encoding
type CreateEmojiJSON struct {
	Name  string          `json:"name"`  // Name - name of the emoji
	Image base64.Encoding `json:"image"` // Image - the 128x128 emoji image
	Roles []Snowflake     `json:"roles"` // Roles - roles allowed to use this emoji
}

// ModifyGuildEmoji - Modify the given emoji.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns the updated emoji object on success.
//
// Fires a Guild Emojis Update Gateway event.
//
// All JSON parameters to this endpoint are optional.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildEmoji(emoji *Emoji, payload *ModifyGuildEmojiJSON, reason *string) (*Emoji, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildEmoji, api, g.ID.String(), emoji.ID.String()))

	var e *Emoji
	err := json.Unmarshal(firePatchRequest(u, payload, reason), &e)

	return e, err
}

// ModifyGuildEmojiJSON - Parameters to pass in the JSON payload
type ModifyGuildEmojiJSON struct {
	Name  string       `json:"name,omitempty"`  // Name - name of the emoji
	Roles []*Snowflake `json:"roles,omitempty"` // Roles - roles allowed to use this emoji
}

// DeleteGuildEmoji - Delete the given emoji.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns "204 No Content" on success.
//
// Fires a GuildEmojisUpdate Gateway event.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) DeleteGuildEmoji(emoji *Emoji, reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteGuildEmoji, api, g.ID.String(), emoji.ID.String()))

	return fireDeleteRequest(u, reason)
}
