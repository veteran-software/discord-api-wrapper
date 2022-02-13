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
	"strings"
	"time"

	"github.com/veteran-software/discord-api-wrapper/routes"
)

/*
Invite Object

Represents a code that when used, adds a user to a guild or group DM channel.
*/

type Invite struct {
	Code                     string              `json:"code"`
	Guild                    Guild               `json:"guild,omitempty"`
	Channel                  *Channel            `json:"channel"`
	Inviter                  User                `json:"inviter,omitempty"`
	TargetType               InviteTargetType    `json:"target_type,omitempty"`
	TargetUser               User                `json:"target_user,omitempty"`
	TargetApplication        Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount uint64              `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   uint64              `json:"approximate_member_count,omitempty"`
	ExpiresAt                time.Time           `json:"expires_at,omitempty"`
	StageInstance            InviteStageInstance `json:"stage_instance,omitempty"`
	GuildScheduledEvent      GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

type InviteTargetType int

const (
	TargetTypeStream InviteTargetType = iota + 1
	TargetTypeEmbeddedApplication
)

/* INVITE METADATA OBJECT */

type InviteMetadata struct {
	Uses      uint64    `json:"uses"`
	MaxUses   uint64    `json:"max_uses"`
	MaxAge    uint64    `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt time.Time `json:"created_at"`
}

/* INVITE STAGE INSTANCE OBJECT */

type InviteStageInstance struct {
	Members          []GuildMember `json:"members"`
	ParticipantCount uint64        `json:"participant_count"`
	SpeakerCount     uint64        `json:"speaker_count"`
	Topic            string        `json:"topic"`
}

/* ENDPOINTS */

/*
GetInvite

Returns an Invite object for the given code.
*/
func (i *Invite) GetInvite(withCounts *bool, withExpiration *bool, guildScheduleEventID *Snowflake) (method, route string) {
	var qsp []string
	if withCounts != nil {
		qsp = append(qsp, "with_counts="+strconv.FormatBool(*withCounts))
	}
	if withExpiration != nil {
		qsp = append(qsp, "with_expiration="+strconv.FormatBool(*withExpiration))
	}
	if guildScheduleEventID != nil {
		qsp = append(qsp, "guild_scheduled_event_id="+(*guildScheduleEventID).String())
	}
	var q string
	if len(qsp) > 0 {
		q = "?" + strings.Join(qsp, "&")
	}
	return http.MethodGet, fmt.Sprintf(routes.Invites_Qsp, api, i.Code, q)
}

/*
DeleteInvite

Delete an Invite.

Requires the MANAGE_CHANNELS permission on the channel this invite belongs to, or MANAGE_GUILD to remove any invite across the guild.

Returns an Invite object on success.

Fires an Invite Delete Gateway event.
*/
func (i *Invite) DeleteInvite() (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Invites_, api, i.Code)
}
