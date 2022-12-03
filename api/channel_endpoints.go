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
	"fmt"
	"strconv"
	"time"
)

// ListPrivateArchivedThreads - Returns archived threads in the channel that are of type GuildPrivateThread.
//
// Threads are ordered by archive_timestamp, in descending order.
//
// Requires both the READ_MESSAGE_HISTORY and MANAGE_THREADS permissions.
func (c *Channel) ListPrivateArchivedThreads(before *time.Time, limit *int) (*ThreadListResponse, error) {
	u := parseRoute(fmt.Sprintf(listPrivateArchivedThreads, api, c.ID.String()))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var threadListResponse *ThreadListResponse
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadListResponse)

	return threadListResponse, err
}

// ListJoinedPrivateArchivedThreads - Returns archived threads in the channel that are of type GuildPrivateThread, and the user has joined.
//
// Threads are ordered by their id, in descending order.
//
// Requires the READ_MESSAGE_HISTORY permission.
func (c *Channel) ListJoinedPrivateArchivedThreads(before *Snowflake, limit *int) (*ThreadListResponse, error) {
	u := parseRoute(fmt.Sprintf(listJoinedPrivateArchivedThreads, api, c.ID.String()))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var threadListResponse *ThreadListResponse
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadListResponse)

	return threadListResponse, err
}
