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
	"time"
)

// GuildTemplate - Represents a code that when used, creates a guild based on a snapshot of an existing guild.
type GuildTemplate struct {
	Code                  string    `json:"code"`
	Name                  string    `json:"name"`
	Description           *string   `json:"description"`
	UsageCount            uint64    `json:"usage_count"`
	CreatorID             Snowflake `json:"creator_id"`
	Creator               User      `json:"creator"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	SourceGuildID         Snowflake `json:"source_guild_id"`
	SerializedSourceGuild Guild     `json:"serialized_source_guild"`
	IsDirty               *bool     `json:"is_dirty"`
}

// GetGuildTemplate - Returns a GuildTemplate object for the given code.
//goland:noinspection GoUnusedExportedFunction
func GetGuildTemplate(templateCode string) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(getGuildTemplate, api, templateCode))

	var guildTemplate *GuildTemplate
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guildTemplate)

	return guildTemplate, err
}

// CreateGuildFromGuildTemplate - Create a new guild based on a template. Returns a guild object on success. Fires a GuildCreate Gateway event.
//
//    This endpoint can be used only by bots in less than 10 guilds.
//goland:noinspection GoUnusedExportedFunction
func CreateGuildFromGuildTemplate(templateCode string, payload CreateGuildFromGuildTemplateJSON) (*Guild, error) {
	u := parseRoute(fmt.Sprintf(createGuildFromGuildTemplate, api, templateCode))

	var guild *Guild
	err := json.Unmarshal(firePostRequest(u, payload, nil), &guild)

	return guild, err
}

// CreateGuildFromGuildTemplateJSON - JSON payload
type CreateGuildFromGuildTemplateJSON struct {
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"`
}

// GetGuildTemplates - Returns an array of GuildTemplate objects. Requires the ManageGuild permission.
func (g *Guild) GetGuildTemplates() ([]GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(getGuildTemplates, api, g.ID.String()))

	var guildTemplates []GuildTemplate
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &guildTemplates)

	return guildTemplates, err
}

func (g *Guild) CreateGuildTemplate(payload CreateGuildTemplateJSON) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(createGuildTemplate, api, g.ID.String()))

	var guildTemplate *GuildTemplate
	err := json.Unmarshal(firePostRequest(u, payload, nil), &guildTemplate)

	return guildTemplate, err
}

// CreateGuildTemplateJSON - JSON payload
type CreateGuildTemplateJSON struct {
	Name        string  `json:"name,omitempty"`        // name of the template (1-100 characters)
	Description *string `json:"description,omitempty"` // description for the template (0-120 characters)
}

// SyncGuildTemplate - Syncs the template to the guild's current state.
//
// Requires the ManageGuild permission. Returns the GuildTemplate object on success.
func (g *Guild) SyncGuildTemplate(templateCode string) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(syncGuildTemplate, api, g.ID.String(), templateCode))

	var guildTemplate *GuildTemplate
	err := json.Unmarshal(firePutRequest(u, nil, nil), &guildTemplate)

	return guildTemplate, err
}

// ModifyGuildTemplate - Modifies the template's metadata.
//
// Requires the ManageGuild permission.
//
// Returns the GuildTemplate object on success.
func (g *Guild) ModifyGuildTemplate(templateCode string, payload ModifyGuildTemplateJSON) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildTemplate, api, g.ID.String(), templateCode))

	var guildTemplate *GuildTemplate
	err := json.Unmarshal(firePatchRequest(u, payload, nil), &guildTemplate)

	return guildTemplate, err
}

type ModifyGuildTemplateJSON struct {
	CreateGuildTemplateJSON
}

// DeleteGuildTemplate - Deletes the template.
//
// Requires the ManageGuild permission. Returns the deleted GuildTemplate object on success.
// TODO: This DELETE endpoint returns a payload; update the delete request and all other methods that use it accordingly
func (g *Guild) DeleteGuildTemplate(templateCode string) error {
	u := parseRoute(fmt.Sprintf(deleteGuildTemplate, api, g.ID.String(), templateCode))

	return fireDeleteRequest(u, nil)
}
