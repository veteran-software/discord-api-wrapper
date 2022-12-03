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
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// GetGuildAuditLog - Returns an audit log object for the guild.
//
// Requires the ViewAuditLog permission.
func (g *Guild) GetGuildAuditLog(userID *Snowflake, actionType *uint64, before *Snowflake, limit *uint64) (*AuditLog, error) {
	u := parseRoute(fmt.Sprintf(getGuildAuditLog, api, g.ID.String()))

	// Set the optional qsp
	q := u.Query()
	if userID != nil {
		q.Set("user_id", userID.String())
	}
	if actionType != nil {
		q.Set("action_type", strconv.FormatUint(*actionType, 10))
	}
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		if *limit >= 1 && *limit <= 100 {
			q.Set("limit", strconv.FormatUint(*limit, 10))
		} else {
			return nil, errors.New("the limit filter must be >= 1 && <= 100")
		}
	}
	// If there's any of the optional qsp present, encode and add to the URL
	if len(q) != 0 {
		u.RawQuery = q.Encode()
	}

	var log AuditLog
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &log)

	return &log, err
}
