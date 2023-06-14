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

package oauth2

import (
	"net/url"
)

// BaseAuthorizationURL - Base authorization URL
//
//goland:noinspection GoUnusedExportedFunction
func BaseAuthorizationURL() *url.URL {
	u, _ := url.Parse("https://discord.com/api/oauth2/authorize")

	return u
}

// TokenURL - Token URL
//
//goland:noinspection GoUnusedExportedFunction
func TokenURL() *url.URL {
	u, _ := url.Parse("https://discord.com/api/oauth2/token")

	return u
}

// TokenRevocationURL - Token Revocation URL
//
//goland:noinspection GoUnusedExportedFunction
func TokenRevocationURL() *url.URL {
	u, _ := url.Parse("https://discord.com/api/oauth2/token/revoke")

	return u
}
