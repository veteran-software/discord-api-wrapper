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

	log "github.com/veteran-software/nowlive-logging"
)

// GetInvite - Returns an Invite object for the given code.
func (i *Invite) GetInvite(withCounts *bool, withExpiration *bool, guildScheduledEventID *Snowflake) (*Invite,
	error) {
	u := parseRoute(fmt.Sprintf(getInvite, api, *i.Code))

	q := u.Query()
	if withCounts != nil {
		q.Set("with_counts", strconv.FormatBool(*withCounts))
	}
	if withExpiration != nil {
		q.Set("with_expiration", strconv.FormatBool(*withExpiration))
	}
	if guildScheduledEventID != nil {
		q.Set("guild_scheduled_event_id", guildScheduledEventID.String())
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var invite *Invite
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &invite)

	return invite, err
}

// DeleteInvite - Delete an Invite.
//
// Requires the ManageChannels permission on the channel this invite belongs to, or ManageGuild to remove any invite across the guild.
//
// Returns an Invite object on success.
//
// Fires an InviteDelete Gateway event.
//
//	This endpoint supports the `X-Audit-Log-Reason` header.
func (i *Invite) DeleteInvite(reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteInvite, api, *i.Code))

	return fireDeleteRequest(u, reason)
}
