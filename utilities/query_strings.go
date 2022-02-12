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

package utilities

import (
	"fmt"
	"strings"
)

func ParseQueryString(opts *map[string]interface{}) *string {
	var qsp []string
	for k, v := range *opts {
		if strings.ToLower(k) == "around" ||
			strings.ToLower(k) == "before" ||
			strings.ToLower(k) == "after" ||
			strings.ToLower(k) == "limit" ||
			strings.ToLower(k) == "with_counts" ||
			strings.ToLower(k) == "with_expiration" ||
			strings.ToLower(k) == "guild_scheduled_event_id" {
			qsp = append(qsp, fmt.Sprintf("%s=%v", strings.ToLower(k), v))
		}
	}

	var q string
	if len(qsp) > 0 {
		q = "?" + strings.Join(qsp, "&")
		return &q
	}

	return nil
}
