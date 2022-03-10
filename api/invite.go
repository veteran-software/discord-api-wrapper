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
)

//Invite - Represents a code that when used, adds a user to a guild or group DM channel.
type Invite struct {
	Code                     string              `json:"code"`                                 // the invite code (unique ID)
	Guild                    Guild               `json:"guild,omitempty"`                      // the guild this invite is for
	Channel                  *Channel            `json:"channel"`                              // the channel this invite is for
	Inviter                  User                `json:"inviter,omitempty"`                    // the user who created the invite
	TargetType               InviteTargetType    `json:"target_type,omitempty"`                // the type of target for this voice channel invite
	TargetUser               User                `json:"target_user,omitempty"`                // the user whose stream to display for this voice channel stream invite
	TargetApplication        Application         `json:"target_application,omitempty"`         // the embedded application to open for this voice channel embedded application invite
	ApproximatePresenceCount uint64              `json:"approximate_presence_count,omitempty"` // approximate count of online members, returned from the GET /invites/<code> endpoint when with_counts is true
	ApproximateMemberCount   uint64              `json:"approximate_member_count,omitempty"`   // approximate count of total members, returned from the GET /invites/<code> endpoint when with_counts is true
	ExpiresAt                *time.Time          `json:"expires_at,omitempty"`                 // the expiration date of this invite, returned from the GET /invites/<code> endpoint when with_expiration is true
	StageInstance            InviteStageInstance `json:"stage_instance,omitempty"`             // stage instance data if there is a public Stage instance in the Stage channel this invite is for
	GuildScheduledEvent      GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`      // guild scheduled event data, only included if guild_scheduled_event_id contains a valid guild scheduled event id
}

// InviteTargetType - the type of target for this voice channel invite
type InviteTargetType int

//goland:noinspection GoUnusedConst
const (
	TargetTypeStream              InviteTargetType = iota + 1 // STREAM
	TargetTypeEmbeddedApplication                             // EMBEDDED_APPLICATION
)

// InviteMetadata - Extra information about an invite, will extend the invite object.
type InviteMetadata struct {
	Uses      uint64    `json:"uses"`       // number of times this invite has been used
	MaxUses   uint64    `json:"max_uses"`   // max number of times this invite can be used
	MaxAge    uint64    `json:"max_age"`    // duration (in seconds) after which the invite expires
	Temporary bool      `json:"temporary"`  // whether this invite only grants temporary membership
	CreatedAt time.Time `json:"created_at"` // when this invite was created
}

// InviteStageInstance - stage instance data if there is a public Stage instance in the Stage channel this invite is for
type InviteStageInstance struct {
	Members          []GuildMember `json:"members"`
	ParticipantCount uint64        `json:"participant_count"`
	SpeakerCount     uint64        `json:"speaker_count"`
	Topic            string        `json:"topic"`
}

// GetInvite - Returns an Invite object for the given code.
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
	return http.MethodGet, fmt.Sprintf(getInvite, api, i.Code, q)
}

// DeleteInvite - Delete an Invite.
//
// Requires the MANAGE_CHANNELS permission on the channel this invite belongs to, or MANAGE_GUILD to remove any invite across the guild.
//
// Returns an Invite object on success.
//
// Fires an Invite Delete Gateway event.
func (i *Invite) DeleteInvite() (method, route string) {
	return http.MethodDelete, fmt.Sprintf(deleteInvite, api, i.Code)
}
