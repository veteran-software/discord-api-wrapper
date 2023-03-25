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

package api

import (
	"encoding/json"
	"fmt"

	log "github.com/veteran-software/nowlive-logging"
)

// GetGuildTemplate - Returns a GuildTemplate object for the given code.
//
//goland:noinspection GoUnusedExportedFunction
func GetGuildTemplate(templateCode string) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(getGuildTemplate, api, templateCode))

	var guildTemplate *GuildTemplate
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildTemplate)

	return guildTemplate, err
}

// CreateGuildFromGuildTemplate - Create a new guild based on a template. Returns a guild object on success. Fires a GuildCreate Gateway event.
//
//	This endpoint can be used only by bots in less than 10 guilds.
//
//goland:noinspection GoUnusedExportedFunction
func CreateGuildFromGuildTemplate(templateCode string, payload *CreateGuildFromGuildTemplateJSON) (*Guild, error) {
	u := parseRoute(fmt.Sprintf(createGuildFromGuildTemplate, api, templateCode))

	var guild *Guild
	responseBytes, err := firePostRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guild)

	return guild, err
}

// CreateGuildFromGuildTemplateJSON - JSON payload
type CreateGuildFromGuildTemplateJSON struct {
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"`
}

// GetGuildTemplates - Returns an array of GuildTemplate objects. Requires the ManageGuild permission.
func (g *Guild) GetGuildTemplates() ([]*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(getGuildTemplates, api, g.ID.String()))

	var guildTemplates []*GuildTemplate
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildTemplates)

	return guildTemplates, err
}

func (g *Guild) CreateGuildTemplate(payload *CreateGuildTemplateJSON) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(createGuildTemplate, api, g.ID.String()))

	var guildTemplate *GuildTemplate
	responseBytes, err := firePostRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildTemplate)

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
	responseBytes, err := firePutRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildTemplate)

	return guildTemplate, err
}

// ModifyGuildTemplate - Modifies the template's metadata.
//
// Requires the ManageGuild permission.
//
// Returns the GuildTemplate object on success.
func (g *Guild) ModifyGuildTemplate(templateCode string, payload *ModifyGuildTemplateJSON) (*GuildTemplate, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildTemplate, api, g.ID.String(), templateCode))

	var guildTemplate *GuildTemplate
	responseBytes, err := firePatchRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildTemplate)

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
