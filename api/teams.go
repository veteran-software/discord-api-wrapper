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

// Team - Teams are groups of developers on Discord who want to collaborate on apps.
//
// On other platforms, these may be referred to as "organizations", "companies", or "teams".
type Team struct {
	Icon        *string      `json:"icon"`          // a hash of the image of the team's icon
	ID          Snowflake    `json:"id"`            // the unique id of the team
	Members     []TeamMember `json:"members"`       // the members of the team
	Name        string       `json:"name"`          // the name of the team
	OwnerUserID Snowflake    `json:"owner_user_id"` // the user id of the current team owner
}

// TeamMember - representation of a team member
//goland:noinspection GrazieInspection
type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"` // the user's membership state on the team
	Permissions     []string        `json:"permissions"`      // will always be ["*"]
	TeamID          Snowflake       `json:"team_id"`          // the id of the parent team of which they are a member
	User            User            `json:"user"`             // the avatar, discriminator, id, and username of the user
}

// MembershipState - Current state of a team member
type MembershipState int

//goland:noinspection GoUnusedConst
const (
	Invited  MembershipState = iota + 1 // INVITED
	Accepted                            // ACCEPTED
)
