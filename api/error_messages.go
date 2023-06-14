/*
 * Copyright (c) 2023. Veteran Software
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

type ArrayError struct {
	Code   int `json:"code"`
	Errors struct {
		Activities map[string]struct {
			Platform struct {
				Errors []struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				} `json:"_errors"`
			} `json:"platform"`
			Type struct {
				Errors []struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				} `json:"_errors"`
			} `json:"type"`
		} `json:"activities"`
	} `json:"errors"`
	Message string `json:"message"`
}

type ObjectError struct {
	Code   int `json:"code"`
	Errors struct {
		AccessToken struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"access_token"`
	} `json:"errors"`
	Message string `json:"message"`
}

type RequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  struct {
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"_errors"`
	} `json:"errors,omitempty"`
}
