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

/* APPLICATION COMMAND OBJECT */

/*
ApplicationCommand

https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure

A command, or each individual subcommand, can have a maximum of 25 options

An application command is the base "command" model that belongs to an application. This is what you are creating when you POST a new command.

Required options must be listed before optional options

--------

ID: unique id of the command

ApplicationID: unique id of the parent application

GuildID: guild id of the command, if not global

Name:  1-32 lowercase character name matching ^[\w-]{1,32}$

Description: 1-100 character description

Options: the parameters for the command

DefaultPermissions: whether the command is enabled by default when the app is added to a guild
*/
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

/*
ApplicationCommandType

https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
*/
type ApplicationCommandType int

const (
	CommandTypeChatInput ApplicationCommandType = iota + 1
	CommandTypeUser
	CommandTypeMessage
)

/*
ApplicationCommandOption

https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure

You can specify a maximum of 25 choices per option
*/
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

/*
ApplicationCommandOptionType

https:discord.com/developers/docs/interactions/application-command#application-command-object-application-command-option-type
*/
type ApplicationCommandOptionType int

const (
	OptionTypeSubCommand = iota + 1
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

/*
ApplicationCommandOptionChoice

https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure

If you specify choices for an option, they are the only valid values for a user to pick
*/
type ApplicationCommandOptionChoice struct {
	Name  string      `json:"name"`         // 1-100 character choice name
	Value interface{} `json:"value,string"` // value of the choice, up to 100 characters if string
}

/*
ApplicationCommandInteractionDataOption

https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure

All options have names, and an option can either be a parameter and input value--in which case value will be set--or it can denote a subcommand or group--in which case it will contain a top-level key and another array of options.

value and options are mutually exclusive.

--------

Name: the name of the parameter

Type: value of ApplicationCommandOptionType

Value: the value of the pair

Options: present if this option is a group or subcommand
*/
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`              // the name of the parameter
	Type    ApplicationCommandOptionType               `json:"type"`              // value of application command option type
	Value   interface{}                                `json:"value,omitempty"`   // the value of the pair
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"` // present if this option is a group or subcommand
	Focused bool                                       `json:"focused,omitempty"` // true if this option is the currently focused option for autocomplete
}

/* APPLICATION COMMAND PERMISSIONS OBJECT */

/*
GuildApplicationCommandPermissions

https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-guild-application-command-permissions-structure

Returned when fetching the permissions for a command in a guild.
*/
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                       `json:"id"`
	ApplicationID Snowflake                       `json:"application_id"`
	GuildID       Snowflake                       `json:"guild_id"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}

/*
ApplicationCommandPermissions

https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure

Application command permissions allow you to enable or disable command for specific users or roles within a guild.

ID: the id of the role or user

Type: role or user

Permission: true to allow, false, to disallow
*/
type ApplicationCommandPermissions struct {
	ID         Snowflake                        `json:"id"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

/*
ApplicationCommandPermissionType

https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permission-type
*/
type ApplicationCommandPermissionType int

const (
	_ ApplicationCommandPermissionType = iota
	PermissionTypeRole
	PermissionTypeUser
)

/* INTERACTION OBJECT */

/* ENDPOINTS */

/**
Interaction endpoints
*/

// GLOBAL COMMANDS

func GetGlobalApplicationCommands(applicationID string) (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/commands", api, applicationID)
}

func CreateGlobalApplicationCommand(applicationID string) (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/applications/%s/commands", api, applicationID)
}

type CreateApplicationCommandJSON struct {
	Name              string                      `json:"name"`
	Description       string                      `json:"description"`
	Options           *[]ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission bool                        `json:"default_permission,omitempty"` // default true
}

// GetGlobalApplicationCommand
// Fetch a global command for your application. Returns an application command object.
func (i *Interaction) GetGlobalApplicationCommand() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/commands/%s", api, i.ApplicationID, i.Data.ID)
}

func (i *Interaction) EditGlobalApplicationCommand() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf("%s/applications/%s/commands/%s", api, i.ApplicationID, i.Data.ID)
}

type EditApplicationCommandJSON struct {
	CreateApplicationCommandJSON
}

func (i *Interaction) BulkOverwriteGlobalApplicationCommands() (method string, route string) {
	return http.MethodPut, fmt.Sprintf("%s/applications/%s/commands", api, i.ApplicationID)
}

func DeleteGlobalApplicationCommand(applicationID string, commandID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/applications/%s/commands/%s", api, applicationID, commandID)
}

// GUILD COMMANDS

func (i *Interaction) GetGuildApplicationCommands() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/guilds/%s/commands", api, i.ApplicationID, i.GuildID)
}

func GetGuildApplicationCommands(applicationID string, guildID string) (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/guilds/%s/commands", api, applicationID, guildID)
}

func CreateGuildApplicationCommand(applicationID string, guildID string) (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/applications/%s/guilds/%s/commands", api, applicationID, guildID)
}

type CreateGuildApplicationCommandJSON struct {
	CreateApplicationCommandJSON
}

func (i *Interaction) GetGuildApplicationCommand() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/guilds/%s/commands/%s", api, i.ApplicationID, i.GuildID, i.Data.ID)
}

func (i *Interaction) EditGuildApplicationCommand() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf("%s/applications/%s/guilds/%s/commands/%s", api, i.ApplicationID, i.GuildID, i.Data.ID)
}

type EditGuildApplicationCommandJSON struct {
	CreateApplicationCommandJSON
}

func DeleteGuildApplicationCommand(applicationID string, guildID string, commandID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/applications/%s/guilds/%s/commands/%s", api, applicationID, guildID, commandID)
}

func (i *Interaction) BulkOverwriteGuildApplicationCommands() (method string, route string) {
	return http.MethodPut, fmt.Sprintf("%s/applications/%s/guilds/%s/commands", api, i.ApplicationID, i.GuildID)
}

// INTERACTION ENDPOINTS

// CreateInteractionResponse Create a response to an Interaction from the gateway.
func (i *Interaction) CreateInteractionResponse() (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/interactions/%s/%s/callback", api, i.ID.String(), i.Token)
}

// GetOriginalInteractionResponse Returns the initial Interaction response.
func (i *Interaction) GetOriginalInteractionResponse() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/webhooks/%s/%s/messages/@original", api, i.ApplicationID, i.Token)
}

// EditOriginalInteractionResponse Edits the initial Interaction response.
func (i *Interaction) EditOriginalInteractionResponse() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf("%s/webhooks/%s/%s/messages/@original", api, i.ApplicationID, i.Token)
}

// DeleteOriginalInteractionResponse Deletes the initial Interaction response. Returns 204 on success.
func (i *Interaction) DeleteOriginalInteractionResponse() (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/webhooks/%s/%s/messages/@original", api, i.ApplicationID, i.Token)
}

func (i *Interaction) CreateFollowupMessage() (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/webhooks/%s/%s", api, i.ApplicationID, i.Token)
}

func (i *Interaction) EditFollowupMessage() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf("%s/webhooks/%s/%s/messages/%s", api, i.ApplicationID, i.Token, i.Message.ID)
}

func (i *Interaction) DeleteFollowupMessage() (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/webhooks/%s/%s/messages/%s", api, i.ApplicationID, i.Token, i.Message.ID)
}

func (i *Interaction) GetGuildApplicationCommandPermissions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/guilds/%s/command/permissions", api, i.ApplicationID, i.GuildID)
}

func (i *Interaction) GetApplicationCommandPermissions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/applications/%s/guilds/%s/command/%s/permissions", api, i.ApplicationID, i.GuildID, i.Data.ID)
}

func (i *Interaction) EditApplicationCommandPermissions() (method string, route string) {
	return http.MethodPut, fmt.Sprintf("%s/applications/%s/guilds/%s/command/%s/permissions", api, i.ApplicationID, i.GuildID, i.Data.ID)
}

func (i *Interaction) BatchEditApplicationCommandPermissions() (method string, route string) {
	return http.MethodPut, fmt.Sprintf("%s/applications/%s/guilds/%s/command/permissions", api, i.ApplicationID, i.GuildID)
}
