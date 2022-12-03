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
)

// GetGlobalApplicationCommands - Fetch all the global commands for your application.
//
// Returns an array of application command objects.
//
//goland:noinspection GoUnusedExportedFunction
func GetGlobalApplicationCommands(applicationID Snowflake, withLocalizations bool) ([]ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(getGlobalApplicationCommands, api, applicationID.String()))

	q := u.Query()
	q.Set("with_localizations", strconv.FormatBool(withLocalizations))
	u.RawQuery = q.Encode()

	var commands []ApplicationCommand
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &commands)

	return commands, err
}

// CreateGlobalApplicationCommand - Create a new global command.
//
// New global commands will be available in all guilds after 1 hour. Returns 201 and an application command object.
//
//	Creating a command with the same name as an existing command for your application will overwrite the old command.
//
//goland:noinspection GoUnusedExportedFunction
func CreateGlobalApplicationCommand(applicationID Snowflake, payload CreateApplicationCommandJSON) (*ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(createGlobalApplicationCommand, api, applicationID.String()))

	var command *ApplicationCommand
	err := json.Unmarshal(firePostRequest(u, payload, nil), &command)

	return command, err
}

// CreateApplicationCommandJSON - JSON payload structure
type CreateApplicationCommandJSON struct {
	Name                     string                      `json:"name"`                                 // 1-32 character name
	NameLocalizations        *LocalizationDict           `json:"name_localizations,omitempty"`         // Localization dictionary for the name field. Values follow the same restrictions as name
	Description              string                      `json:"description"`                          // 1-100 character description
	DescriptionLocalizations *LocalizationDict           `json:"description_localizations,omitempty"`  // Localization dictionary for the description field. Values follow the same restrictions as description
	Options                  *[]ApplicationCommandOption `json:"options,omitempty"`                    // the parameters for the command
	DefaultMemberPermissions *string                     `json:"default_member_permissions,omitempty"` // Set of permissions represented as a bit set
	DmPermission             *bool                       `json:"dm_permission,omitempty"`              // Indicates whether the command is available in DMs with the app, only for globally-scoped commands. By default, commands are visible.
	DefaultPermission        bool                        `json:"default_permission,omitempty"`         // whether the command is enabled by default when the app is added to a guild; default true
	Type                     ApplicationCommandType      `json:"type,omitempty"`                       // the type of command, defaults 1 if not set
}

// GetGlobalApplicationCommand - Fetch a global command for your application.
//
// Returns an application command object.
//
// Includes localizations by default
func (i *Interaction) GetGlobalApplicationCommand() (*ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(getGlobalApplicationCommand, api, i.ApplicationID.String(), i.Data.ID.String()))

	var commands *ApplicationCommand
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &commands)

	return commands, err
}

// EditGlobalApplicationCommand - Edit a global command. Updates will be available in all guilds after 1 hour.
//
// Returns 200 and an application command object.
//
//	All JSON parameters for this endpoint are optional.
func (i *Interaction) EditGlobalApplicationCommand(payload EditApplicationCommandJSON) (*ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(editGlobalApplicationCommand, api, i.ApplicationID.String(), i.Data.ID.String()))

	var commands *ApplicationCommand
	err := json.Unmarshal(firePatchRequest(u, payload, nil), &commands)

	return commands, err
}

// EditApplicationCommandJSON - JSON payload structure
type EditApplicationCommandJSON struct {
	Name                     string                      `json:"name"`                                 // 1-32 character name
	NameLocalizations        *LocalizationDict           `json:"name_localizations,omitempty"`         // Localization dictionary for the name field. Values follow the same restrictions as name
	Description              string                      `json:"description"`                          // 1-100 character description
	DescriptionLocalizations *LocalizationDict           `json:"description_localizations,omitempty"`  // Localization dictionary for the description field. Values follow the same restrictions as description
	Options                  *[]ApplicationCommandOption `json:"options,omitempty"`                    // the parameters for the command
	DefaultMemberPermissions *string                     `json:"default_member_permissions,omitempty"` // Set of permissions represented as a bit set
	DmPermission             *bool                       `json:"dm_permission,omitempty"`              // Indicates whether the command is available in DMs with the app, only for globally-scoped commands. By default, commands are visible.
	DefaultPermission        bool                        `json:"default_permission,omitempty"`         // whether the command is enabled by default when the app is added to a guild; default true
}

// DeleteGlobalApplicationCommand - Deletes a global command. Returns 204 No Content on success.
//
//goland:noinspection GoUnusedExportedFunction
func DeleteGlobalApplicationCommand(applicationID Snowflake, commandID string) error {
	u := parseRoute(fmt.Sprintf(deleteGlobalApplicationCommand, api, applicationID.String(), commandID))

	return fireDeleteRequest(u, nil)
}

// BulkOverwriteGlobalApplicationCommands - Takes a list of application commands, overwriting the existing global command list for this application.
//
// Updates will be available in all guilds after 1 hour.
//
// Returns 200 and a list of application command objects.
//
// Commands that do not already exist will count toward daily application command create limits.
//
//goland:noinspection GoUnusedExportedFunction
func BulkOverwriteGlobalApplicationCommands(applicationID Snowflake, payload []ApplicationCommand) ([]ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(bulkOverwriteGlobalApplicationCommands, api, applicationID.String()))

	var commands []ApplicationCommand
	err := json.Unmarshal(firePutRequest(u, payload, nil), &commands)

	return commands, err
}

// GetGuildApplicationCommands - Fetch all the guild commands for your application for a specific guild.
//
// Returns an array of application command objects.
func (i *Interaction) GetGuildApplicationCommands(withLocalizations bool) ([]ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(getGuildApplicationCommands, api, i.ApplicationID.String(), i.GuildID.String()))

	q := u.Query()
	q.Set("with_localizations", strconv.FormatBool(withLocalizations))
	u.RawQuery = q.Encode()

	var commands []ApplicationCommand
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &commands)

	return commands, err
}

// GetGuildApplicationCommands - Fetch all the guild commands for your application for a specific guild.
//
// Returns an array of application command objects.
//
//goland:noinspection GoUnusedExportedFunction
func GetGuildApplicationCommands(applicationID Snowflake, guildID Snowflake, withLocalizations bool) ([]ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(getGuildApplicationCommands, api, applicationID.String(), guildID.String()))

	q := u.Query()
	q.Set("with_localizations", strconv.FormatBool(withLocalizations))
	u.RawQuery = q.Encode()

	var commands []ApplicationCommand
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &commands)

	return commands, err
}

// CreateGuildApplicationCommand - Create a new guild command.
//
// New guild commands will be available in the guild immediately.
//
// Returns 201 and an application command object.
//
// If the command did not already exist, it will count toward daily application command create limits.
//
//goland:noinspection GoUnusedExportedFunction
func CreateGuildApplicationCommand(applicationID Snowflake, guildID Snowflake, payload CreateApplicationCommandJSON) (*ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(createGuildApplicationCommand, api, applicationID.String(), guildID.String()))

	var command *ApplicationCommand
	err := json.Unmarshal(firePostRequest(u, payload, nil), &command)

	return command, err
}

// GetGuildApplicationCommand - Fetch a guild command for your application. Returns an application command object.
func (i *Interaction) GetGuildApplicationCommand() (*ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(getGuildApplicationCommand, api, i.ApplicationID.String(), i.GuildID.String(), i.Data.ID.String()))

	var command *ApplicationCommand
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &command)

	return command, err
}

// EditGuildApplicationCommand - Edit a guild command. Updates for guild commands will be available immediately.
//
// Returns 200 and an application command object.
//
//	All parameters for this endpoint are optional.
func (i *Interaction) EditGuildApplicationCommand(payload EditApplicationCommandJSON) (*ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(editGuildApplicationCommand, api, i.ApplicationID.String(), i.GuildID.String(), i.Data.ID.String()))

	var command *ApplicationCommand
	err := json.Unmarshal(firePatchRequest(u, payload, nil), &command)

	return command, err
}

// DeleteGuildApplicationCommand - Delete a guild command. Returns 204 No Content on success.
//
//goland:noinspection GoUnusedExportedFunction
func DeleteGuildApplicationCommand(applicationID Snowflake, guildID Snowflake, commandID string) error {
	u := parseRoute(fmt.Sprintf(deleteGuildApplicationCommand, api, applicationID.String(), guildID.String(), commandID))

	return fireDeleteRequest(u, nil)
}

// BulkOverwriteGuildApplicationCommands - Takes a list of application commands, overwriting the existing command list for this application for the targeted guild.
//
// Returns 200 and a list of application command objects.
func (i *Interaction) BulkOverwriteGuildApplicationCommands(payload []ApplicationCommand) ([]ApplicationCommand, error) {
	u := parseRoute(fmt.Sprintf(bulkOverwriteGuildApplicationCommands, api, i.ApplicationID.String(), i.GuildID.String()))

	var commands []ApplicationCommand
	err := json.Unmarshal(firePutRequest(u, payload, nil), &commands)

	return commands, err
}

// GetGuildApplicationCommandPermissions - Fetches command permissions for all commands for your application in a guild.
//
// Returns an array of guild application command permissions objects.
func (i *Interaction) GetGuildApplicationCommandPermissions() ([]GuildApplicationCommandPermissions, error) {
	u := parseRoute(fmt.Sprintf(getGuildApplicationCommandPermissions, api, i.ApplicationID.String(), i.GuildID.String()))

	var commandPerms []GuildApplicationCommandPermissions
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &commandPerms)

	return commandPerms, err
}

// GetApplicationCommandPermissions - Fetches command permissions for a specific command for your application in a guild.
//
// Returns a guild application command permissions object.
func (i *Interaction) GetApplicationCommandPermissions() (*GuildApplicationCommandPermissions, error) {
	u := parseRoute(fmt.Sprintf(getApplicationCommandPermissions, api, i.ApplicationID.String(), i.GuildID.String(), i.Data.ID.String()))

	var commandPerms *GuildApplicationCommandPermissions
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &commandPerms)

	return commandPerms, err
}

// EditApplicationCommandPermissions
//
//	This endpoint will overwrite existing permissions for the command in that guild
//
// Edits command permissions for a specific command for your application in a guild and returns a GuildApplicationCommandPermissions object.
//
// You can add up to 100 permission overwrites for a command.
//
//	This endpoint requires authentication with a `Bearer` token that has permission to manage the guild and its roles. For more information, read above about application command permissions.
//
//	Deleting or renaming a command will permanently delete all permissions for the command
//
// TODO: Find the best way to handle the requirement for needing a Bearer Token to use this endpoint
func (i *Interaction) EditApplicationCommandPermissions(payload EditApplicationCommandPermissionsJSON) (*GuildApplicationCommandPermissions, error) {
	u := parseRoute(fmt.Sprintf(editApplicationCommandPermissions, api, i.ApplicationID.String(), i.GuildID.String(), i.Data.ID.String()))

	var commandPerms *GuildApplicationCommandPermissions
	err := json.Unmarshal(firePutRequest(u, payload, nil), &commandPerms)

	return commandPerms, err
}

// EditApplicationCommandPermissionsJSON - JSON payload structure
type EditApplicationCommandPermissionsJSON struct {
	Permissions []ApplicationCommandPermissions `json:"permissions"` // the permissions for the command in the guild
}
