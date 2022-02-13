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
	"fmt"
	"net/http"
)

// ApplicationCommand - A command, or each individual subcommand, can have a maximum of 25 options
//
// An application command is the base "command" model that belongs to an application. This is what you are creating when you POST a new command.
//
// Required options must be listed before optional options
type ApplicationCommand struct {
	ID                 Snowflake                  `json:"id,omitempty"`                  // unique id of the command
	Type               ApplicationCommandType     `json:"type,omitempty"`                // the type of command, defaults 1 if not set
	ApplicationID      Snowflake                  `json:"application_id"`                // unique id of the parent application
	GuildID            Snowflake                  `json:"guild_id,omitempty"`            // guild id of the command, if not global
	Name               string                     `json:"name"`                          // max 32 chars, must follow ^[\w-]{1,32}$ regex
	Description        string                     `json:"description"`                   // 1-100 character description for CHAT_INPUT command, empty string for USER and MESSAGE command
	Options            []ApplicationCommandOption `json:"options,omitempty"`             // the parameters for the command, max 25; CHAT_INPUT
	DefaultPermissions bool                       `json:"default_permissions,omitempty"` // default true; whether the command is enabled by default when added to a guild
	Version            Snowflake                  `json:"version"`                       // autoincrementing version identifier updated during substantial record changes
}

// ApplicationCommandType - The type of application command
type ApplicationCommandType int

//goland:noinspection GoUnusedConst
const (
	CommandTypeChatInput ApplicationCommandType = iota + 1
	CommandTypeUser
	CommandTypeMessage
)

// ApplicationCommandOption - You can specify a maximum of 25 choices per option
type ApplicationCommandOption struct {
	Type         ApplicationCommandOptionType     `json:"type"`                    // the type of option
	Name         string                           `json:"name"`                    // 1-32 character name
	Description  string                           `json:"description"`             // 1-100 character description
	Required     bool                             `json:"required,omitempty"`      // if the parameter is required or optional--default `false`
	Choices      []ApplicationCommandOptionChoice `json:"choices,omitempty"`       // choices for STRING, INTEGER, and NUMBER types for the user to pick from, max 25
	Options      []ApplicationCommandOption       `json:"options,omitempty"`       // if the option is a subcommand or subcommand group type, these nested options will be the parameters
	ChannelTypes []ChannelType                    `json:"channel_types,omitempty"` // if the option is a channel type, the channels shown will be restricted to these types
	MinValue     interface{}                      `json:"min_value,omitempty"`     // if the option is an INTEGER or NUMBER type, the minimum value permitted; integer for INTEGER options, double for NUMBER options
	MaxValue     interface{}                      `json:"max_value,omitempty"`     // if the option is an INTEGER or NUMBER type, the maximum value permitted; integer for INTEGER options, double for NUMBER options
	Autocomplete bool                             `json:"autocomplete,omitempty"`  // enable autocomplete interactions for this option
}

// ApplicationCommandOptionType - The option type of the command
type ApplicationCommandOptionType int

//goland:noinspection GoUnusedConst
const (
	OptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	OptionTypeSubCommandGroup
	OptionTypeString
	OptionTypeInteger // Any integer between -2^53 and 2^53
	OptionTypeBoolean
	OptionTypeUser
	OptionTypeChannel // Includes all channel types + categories
	OptionTypeRole
	OptionTypeMentionable // Includes users and roles
	OptionTypeNumber      // Any double between -2^53 and 2^53
)

// ApplicationCommandOptionChoice - If you specify choices for an option, they are the only valid values for a user to pick
type ApplicationCommandOptionChoice struct {
	Name  string      `json:"name"`         // 1-100 character choice name
	Value interface{} `json:"value,string"` // value of the choice, up to 100 characters if string
}

// ApplicationCommandInteractionDataOption - All options have names, and an option can either be a parameter and input value--in which case value will be set--or it can denote a subcommand or group--in which case it will contain a top-level key and another array of options.
//
// value and options are mutually exclusive.
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`              // the name of the parameter
	Type    ApplicationCommandOptionType               `json:"type"`              // value of application command option type
	Value   interface{}                                `json:"value,omitempty"`   // the value of the pair
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"` // present if this option is a group or subcommand
	Focused bool                                       `json:"focused,omitempty"` // true if this option is the currently focused option for autocomplete
}

// GuildApplicationCommandPermissions - Returned when fetching the permissions for a command in a guild.
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                       `json:"id"`
	ApplicationID Snowflake                       `json:"application_id"`
	GuildID       Snowflake                       `json:"guild_id"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}

// ApplicationCommandPermissions - Application command permissions allow you to enable or disable command for specific users or roles within a guild.
type ApplicationCommandPermissions struct {
	ID         Snowflake                        `json:"id"`         // ID: the id of the role or user
	Type       ApplicationCommandPermissionType `json:"type"`       // Type: role or user
	Permission bool                             `json:"permission"` // Permission: true to allow, false, to disallow
}

// ApplicationCommandPermissionType - The permission type for the command
type ApplicationCommandPermissionType int

//goland:noinspection GoUnusedConst
const (
	PermissionTypeRole ApplicationCommandPermissionType = iota + 1
	PermissionTypeUser
)

// GetGlobalApplicationCommands - Fetch all the global commands for your application.
//
// Returns an array of application command objects.
func GetGlobalApplicationCommands(applicationID string) (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getGlobalApplicationCommands, api, applicationID)
}

// CreateGlobalApplicationCommand - Create a new global command.
//
// New global commands will be available in all guilds after 1 hour. Returns 201 and an application command object.
//
//    Creating a command with the same name as an existing command for your application will overwrite the old command.
func CreateGlobalApplicationCommand(applicationID string) (method string, route string) {
	return http.MethodPost, fmt.Sprintf(createGlobalApplicationCommand, api, applicationID)
}

type CreateApplicationCommandJSON struct {
	Name              string                      `json:"name"`
	Description       string                      `json:"description"`
	Options           *[]ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission bool                        `json:"default_permission,omitempty"` // default true
}

// GetGlobalApplicationCommand - Fetch a global command for your application.
//
// Returns an application command object.
func (i *Interaction) GetGlobalApplicationCommand() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getGlobalApplicationCommand, api, i.ApplicationID, i.Data.ID)
}

// EditGlobalApplicationCommand - Edit a global command. Updates will be available in all guilds after 1 hour.
//
// Returns 200 and an application command object.
//
//    All JSON parameters for this endpoint are optional.
func (i *Interaction) EditGlobalApplicationCommand() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf(editGlobalApplicationCommand, api, i.ApplicationID, i.Data.ID)
}

type EditApplicationCommandJSON struct {
	CreateApplicationCommandJSON
}

// DeleteGlobalApplicationCommand - Deletes a global command. Returns 204 No Content on success.
func DeleteGlobalApplicationCommand(applicationID string, commandID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf(deleteGlobalApplicationCommand, api, applicationID, commandID)
}

// BulkOverwriteGlobalApplicationCommands - Takes a list of application commands, overwriting the existing global command list for this application.
//
// Updates will be available in all guilds after 1 hour.
//
// Returns 200 and a list of application command objects.
//
// Commands that do not already exist will count toward daily application command create limits.
func (i *Interaction) BulkOverwriteGlobalApplicationCommands() (method string, route string) {
	return http.MethodPut, fmt.Sprintf(bulkOverwriteGlobalApplicationCommands, api, i.ApplicationID)
}

// GetGuildApplicationCommands - Fetch all the guild commands for your application for a specific guild.
//
// Returns an array of application command objects.
func (i *Interaction) GetGuildApplicationCommands() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getGuildApplicationCommands, api, i.ApplicationID, i.GuildID)
}

// GetGuildApplicationCommands - Fetch all the guild commands for your application for a specific guild.
//
// Returns an array of application command objects.
func GetGuildApplicationCommands(applicationID string, guildID string) (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getGuildApplicationCommands, api, applicationID, guildID)
}

// CreateGuildApplicationCommand - Create a new guild command.
//
// New guild commands will be available in the guild immediately.
//
// Returns 201 and an application command object.
//
// If the command did not already exist, it will count toward daily application command create limits.
func CreateGuildApplicationCommand(applicationID string, guildID string) (method string, route string) {
	return http.MethodPost, fmt.Sprintf(createGuildApplicationCommand, api, applicationID, guildID)
}

type CreateGuildApplicationCommandJSON struct {
	CreateApplicationCommandJSON
}

// GetGuildApplicationCommand - Fetch a guild command for your application. Returns an application command object.
func (i *Interaction) GetGuildApplicationCommand() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getGuildApplicationCommand, api, i.ApplicationID, i.GuildID, i.Data.ID)
}

// EditGuildApplicationCommand - Edit a guild command. Updates for guild commands will be available immediately.
//
// Returns 200 and an application command object.
//
//    All parameters for this endpoint are optional.
func (i *Interaction) EditGuildApplicationCommand() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf(editGuildApplicationCommand, api, i.ApplicationID, i.GuildID, i.Data.ID)
}

type EditGuildApplicationCommandJSON struct {
	CreateApplicationCommandJSON
}

// DeleteGuildApplicationCommand - Delete a guild command. Returns 204 No Content on success.
func DeleteGuildApplicationCommand(applicationID string, guildID string, commandID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf(deleteGuildApplicationCommand, api, applicationID, guildID, commandID)
}

// BulkOverwriteGuildApplicationCommands - Takes a list of application commands, overwriting the existing command list for this application for the targeted guild.
//
// Returns 200 and a list of application command objects.
func (i *Interaction) BulkOverwriteGuildApplicationCommands() (method string, route string) {
	return http.MethodPut, fmt.Sprintf(bulkOverwriteGuildApplicationCommands, api, i.ApplicationID, i.GuildID.String())
}

// GetGuildApplicationCommandPermissions - Fetches command permissions for all commands for your application in a guild.
//
// Returns an array of guild application command permissions objects.
func (i *Interaction) GetGuildApplicationCommandPermissions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getGuildApplicationCommandPermissions, api, i.ApplicationID, i.GuildID)
}

// GetApplicationCommandPermissions - Fetches command permissions for a specific command for your application in a guild.
//
// Returns a guild application command permissions object.
func (i *Interaction) GetApplicationCommandPermissions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf(getApplicationCommandPermissions, api, i.ApplicationID, i.GuildID, i.Data.ID)
}

// EditApplicationCommandPermissions - Edits command permissions for a specific command for your application in a guild.
//
// You can only add up to 10 permission overwrites for a command.
//
// Returns a GuildApplicationCommandPermissions object.
//
//   This endpoint will overwrite existing permissions for the command in that guild
//   Deleting or renaming a command will permanently delete all permissions for that command
func (i *Interaction) EditApplicationCommandPermissions() (method string, route string) {
	return http.MethodPut, fmt.Sprintf(editApplicationCommandPermissions, api, i.ApplicationID, i.GuildID, i.Data.ID)
}

type EditApplicationCommandPermissionsJSON struct {
	Permissions []ApplicationCommandPermissions `json:"permissions"` // the permissions for the command in the guild
}

// BatchEditApplicationCommandPermissions - Batch edits permissions for all commands in a guild.
//
// Takes an array of partial guild application command permissions objects including id and permissions.
//
// You can only add up to 10 permission overwrites for a command.
//
// Returns an array of GuildApplicationCommandPermissions objects.
//
//   This endpoint will overwrite all existing permissions for all commands in a guild, including slash commands, user commands, and message commands.
func (i *Interaction) BatchEditApplicationCommandPermissions() (method string, route string) {
	return http.MethodPut, fmt.Sprintf(batchEditApplicationCommandPermissions, api, i.ApplicationID, i.GuildID)
}
