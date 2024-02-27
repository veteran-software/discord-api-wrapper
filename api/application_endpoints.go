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
	"errors"
	"fmt"
	"strconv"
	"strings"

	log "github.com/veteran-software/nowlive-logging"
)

//goland:noinspection GoUnusedExportedFunction
func GetCurrentApplication() (*Application, error) {
	u := parseRoute(fmt.Sprintf(getCurrentApplication, api))

	var application *Application
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &application)

	return application, err
}

//goland:noinspection GoUnusedExportedFunction
func EditCurrentApplication(customInstallUrl, description, roleConnectionsVerificationUrl, interactionsEndpointUrl *string, flags *int, tags *[]string) (*Application, error) {
	u := parseRoute(fmt.Sprintf(editCurrentApplication, api))

	// Set the optional qsp
	q := u.Query()
	if customInstallUrl != nil {
		q.Set("custom_install_url", *customInstallUrl)
	}
	if description != nil {
		q.Set("description", *description)
	}
	if roleConnectionsVerificationUrl != nil {
		q.Set("role_connections_verification_url", *roleConnectionsVerificationUrl)
	}
	if interactionsEndpointUrl != nil {
		q.Set("interactions_endpoint_url", *interactionsEndpointUrl)
	}
	if flags != nil {
		q.Set("flags", strconv.Itoa(*flags))
	}
	if tags != nil {
		if len(*tags) > 5 {
			return nil, errors.New("you cannot have more than 5 tags")
		}
		for _, tag := range *tags {
			if len(tag) > 20 {
				return nil, errors.New("tag cannot be longer than 20 characters long")
			}
		}

		q.Set("flags", fmt.Sprintf("[%v]", strings.Join(*tags, ",")))
	}
	// If there's any of the optional qsp present, encode and add to the URL
	if len(q) != 0 {
		u.RawQuery = q.Encode()
	}

	var application *Application
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &application)

	return application, err
}
