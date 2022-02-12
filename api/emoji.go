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

	"github.com/veteran-software/discord-api-wrapper/routes"
)

/*
Emoji

ID: emoji id

Name: emoji name

Roles: roles allowed to use this emoji

User: user that created this emoji

RequireColons: whether this emoji must be wrapped in colons

Managed: whether this emoji is managed

Animated: whether this emoji is animated

Available: whether this emoji can be used, may be false due to loss of Server Boosts
*/
type Emoji struct {
	ID            *Snowflake `json:"id"`
	Name          string     `json:"name"`
	Roles         []Role     `json:"roles,omitempty"`
	User          *User      `json:"user,omitempty"`
	RequireColons bool       `json:"require_colons,omitempty"`
	Managed       bool       `json:"managed,omitempty"`
	Animated      bool       `json:"animated,omitempty"`
	Available     bool       `json:"available,omitempty"`
}

// ListGuildEmojis
// Returns a list of emoji objects for the given guild.
func (g *Guild) ListGuildEmojis() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Guilds_Emojis, api, g.ID.String())
}

// GetGuildEmoji
// Returns an emoji object for the given guild and emoji IDs.
func (g *Guild) GetGuildEmoji(emoji Emoji) (method string, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Guilds_Emojis_, api, g.ID.String(), emoji.ID.String())
}

func (g *Guild) CreateGuildEmoji() (method string, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Guilds_Emojis, api, g.ID.String())
}
