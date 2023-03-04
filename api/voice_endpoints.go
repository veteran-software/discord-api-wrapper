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
)

// ListVoiceRegions - Returns an array of voice region objects that can be used when setting a voice or stage channel's `rtc_region`.
//
//goland:noinspection GoUnusedExportedFunction
func ListVoiceRegions() ([]*VoiceRegion, error) {
	u := parseRoute(fmt.Sprintf(listVoiceRegions, api))

	var voiceRegions []*VoiceRegion
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &voiceRegions)

	return voiceRegions, err
}
