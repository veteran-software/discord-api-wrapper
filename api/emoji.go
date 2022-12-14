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

// Emoji - Routes for controlling emojis do not follow the normal rate limit conventions.
//
// These routes are specifically limited on a per-guild basis to prevent abuse.
//
// This means that the quota returned by our APIs may be inaccurate, and you may encounter 429s.
type Emoji struct {
	ID            *Snowflake `json:"id"`                       // ID - emoji id
	Name          string     `json:"name"`                     // Name - emoji name
	Roles         []*Role    `json:"roles,omitempty"`          // Roles - roles allowed to use this emoji
	User          *User      `json:"user,omitempty"`           // User - user that created this emoji
	RequireColons bool       `json:"require_colons,omitempty"` // RequireColons - whether this emoji must be wrapped in colons
	Managed       bool       `json:"managed,omitempty"`        // Managed - whether this emoji is managed
	Animated      bool       `json:"animated,omitempty"`       // Animated - whether this emoji is animated
	Available     bool       `json:"available,omitempty"`      // Available - whether this emoji can be used, may be false due to loss of Server Boosts
}
