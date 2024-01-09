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

package integrations

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* INTEGRATIONS */

type (
	// IntegrationCreate - Sent when an integration is created. The inner payload is an integration object with an additional guild_id key
	IntegrationCreate struct {
		api.Integration
		GuildID api.Snowflake `json:"guild_id"`
	}

	// IntegrationUpdate - Sent when an integration is updated. The inner payload is an integration object with an additional guild_id key
	IntegrationUpdate struct {
		api.Integration
		GuildID api.Snowflake `json:"guild_id"`
	}

	// IntegrationDelete - Sent when an integration is deleted.
	IntegrationDelete struct {
		ID            api.Snowflake `json:"id"`
		GuildID       api.Snowflake `json:"guild_id"`
		ApplicationID api.Snowflake `json:"application_id"`
	}
)
