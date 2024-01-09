/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"encoding/json"

	log "github.com/veteran-software/nowlive-logging"
)

const (
	gatewayVersion = 10
)

type GetGatewayResponse struct {
	Url string `json:"url"`
}

//goland:noinspection GoUnusedExportedFunction
func GetGateway() (*GetGatewayResponse, error) {
	u := parseRoute(api + "/gateway")

	var gatewayResponseBytes *GetGatewayResponse
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &gatewayResponseBytes)

	return gatewayResponseBytes, nil
}

type GetGatewayBotResponse struct {
	Url               string `json:"url"`
	Shards            int    `json:"shards"`
	SessionStartLimit struct {
		Total          int `json:"total"`
		Remaining      int `json:"remaining"`
		ResetAfter     int `json:"reset_after"`
		MaxConcurrency int `json:"max_concurrency"`
	} `json:"session_start_limit"`
}

//goland:noinspection GoUnusedExportedFunction
func GetGatewayBot() (*GetGatewayBotResponse, error) {
	u := parseRoute(api + "/gateway/bot")

	var gatewayResponseBytes *GetGatewayBotResponse
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &gatewayResponseBytes)

	return gatewayResponseBytes, nil
}
