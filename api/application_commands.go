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
	"net/http"
	"strconv"

	"github.com/veteran-software/discord-api-wrapper/v10/logging"
)

// ApplicationCommand - A command, or each individual subcommand, can have a maximum of 25 options
//
// An application command is the base "command" model that belongs to an application. This is what you are creating when you POST a new command.
//
// Required options must be listed before optional options
//
//goland:noinspection SpellCheckingInspection
type ApplicationCommand struct {
	ID                       Snowflake                  `json:"id,omitempty"`                         // unique id of the command
	Type                     ApplicationCommandType     `json:"type,omitempty"`                       // the type of command, defaults 1 if not set
	ApplicationID            Snowflake                  `json:"application_id"`                       // unique id of the parent application
	GuildID                  Snowflake                  `json:"guild_id,omitempty"`                   // guild id of the command, if not global
	Name                     string                     `json:"name"`                                 // max 32 chars, must follow ^[\w-]{1,32}$ regex
	NameLocalizations        *LocalizationDict          `json:"name_localizations,omitempty"`         // Localization dictionary for the name field. Values follow the same restrictions as name
	Description              string                     `json:"description"`                          // 1-100 character description for CHAT_INPUT command, empty string for USER and MESSAGE command
	DescriptionLocalizations *LocalizationDict          `json:"description_localizations,omitempty"`  // Localization dictionary for the description field. Values follow the same restrictions as description
	Options                  []ApplicationCommandOption `json:"options,omitempty"`                    // the parameters for the command, max 25; CHAT_INPUT
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"` // Set of permissions represented as a bit set
	DmPermission             *bool                      `json:"dm_permission,omitempty"`              // Indicates whether the command is available in DMs with the app, only for globally-scoped commands. By default, commands are visible.
	Version                  Snowflake                  `json:"version"`                              // autoincrementing version identifier updated during substantial record changes
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
	Type                     ApplicationCommandOptionType     `json:"type"`                                // the type of option
	Name                     string                           `json:"name"`                                // 1-32 character name
	NameLocalizations        *LocalizationDict                `json:"name_localizations,omitempty"`        // Localization dictionary for the name field. Values follow the same restrictions as name
	Description              string                           `json:"description"`                         // 1-100 character description
	DescriptionLocalizations *LocalizationDict                `json:"description_localizations,omitempty"` // Localization dictionary for the description field. Values follow the same restrictions as description
	Required                 bool                             `json:"required,omitempty"`                  // if the parameter is required or optional--default `false`
	Choices                  []ApplicationCommandOptionChoice `json:"choices,omitempty"`                   // choices for STRING, INTEGER, and NUMBER types for the user to pick from, max 25
	Options                  []ApplicationCommandOption       `json:"options,omitempty"`                   // if the option is a subcommand or subcommand group type, these nested options will be the parameters
	ChannelTypes             []ChannelType                    `json:"channel_types,omitempty"`             // if the option is a channel type, the channels shown will be restricted to these types
	MinValue                 any                              `json:"min_value,omitempty"`                 // if the option is an INTEGER or NUMBER type, the minimum value permitted; integer for INTEGER options, double for NUMBER options
	MaxValue                 any                              `json:"max_value,omitempty"`                 // if the option is an INTEGER or NUMBER type, the maximum value permitted; integer for INTEGER options, double for NUMBER options
	MinLength                int                              `json:"min_length,omitempty"`                // For option type STRING, the minimum allowed length (minimum of 0, maximum of 6000)
	MaxLength                int                              `json:"max_length,omitempty"`                // For option type STRING, the maximum allowed length (minimum of 1, maximum of 6000)
	Autocomplete             bool                             `json:"autocomplete,omitempty"`              // If autocomplete interactions are enabled for this STRING, INTEGER, or NUMBER type option
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
	OptionTypeAttachment  // attachment object
)

// ApplicationCommandOptionChoice - If you specify choices for an option, they are the only valid values for a user to pick
type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`                         // 1-100 character choice name
	NameLocalizations *LocalizationDict `json:"name_localizations,omitempty"` // Localization dictionary for the name field. Values follow the same restrictions as name
	Value             any               `json:"value"`                        // value of the choice, up to 100 characters if string
}

// ApplicationCommandInteractionDataOption - All options have names, and an option can either be a parameter and input value--in which case value will be set--or it can denote a subcommand or group--in which case it will contain a top-level key and another array of options.
//
//	value and options are mutually exclusive.
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`              // the name of the parameter
	Type    ApplicationCommandOptionType               `json:"type"`              // value of application command option type
	Value   any                                        `json:"value,omitempty"`   // the value of the pair
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"` // present if this option is a group or subcommand
	Focused bool                                       `json:"focused,omitempty"` // true if this option is the currently focused option for autocomplete
}

// GuildApplicationCommandPermissions - Returned when fetching the permissions for a command in a guild.
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                       `json:"id"`             // the id of the command
	ApplicationID Snowflake                       `json:"application_id"` // the id of the application the command belongs to
	GuildID       Snowflake                       `json:"guild_id"`       // the id of the guild
	Permissions   []ApplicationCommandPermissions `json:"permissions"`    // the permissions for the command in the guild
}

// ApplicationCommandPermissions - Application command permissions allow you to enable or disable command for specific users or roles within a guild.
type ApplicationCommandPermissions struct {
	ID         Snowflake                        `json:"id"`         // the id of the role or user
	Type       ApplicationCommandPermissionType `json:"type"`       // role or user
	Permission bool                             `json:"permission"` // true to allow, false, to disallow
}

// ApplicationCommandPermissionType - The permission type for the command
type ApplicationCommandPermissionType int

//goland:noinspection GoUnusedConst
const (
	PermissionTypeRole    ApplicationCommandPermissionType = iota + 1 // ROLE
	PermissionTypeUser                                                // USER
	PermissionTypeChannel                                             // CHANNEL
)

// PermissionConstantsEveryone
// All members in a guild
//
//goland:noinspection GoUnusedExportedFunction
func PermissionConstantsEveryone(guildID Snowflake) *Snowflake {
	return &guildID
}

// PermissionsConstantsAllChannels
// All channels in a guild
//
//goland:noinspection GoUnusedExportedFunction
func PermissionsConstantsAllChannels(guildID Snowflake) *Snowflake {
	snowflakeInt, err := strconv.ParseUint(string(guildID), 10, 64)
	if err != nil {
		logging.Errorln(err)
		return new(Snowflake)
	}

	updatedSnowflake := snowflakeInt - 1

	return StringToSnowflake(strconv.FormatUint(updatedSnowflake, 10))
}

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

// BatchEditApplicationCommandPermissions - Batch edits permissions for all commands in a guild.
//
// Takes an array of partial guild application command permissions objects including id and permissions.
//
// You can only add up to 10 permission overwrites for a command.
//
// Returns an array of GuildApplicationCommandPermissions objects.
//
//	This endpoint will overwrite all existing permissions for all commands in a guild, including slash commands, user commands, and message commands.
func (i *Interaction) BatchEditApplicationCommandPermissions() (method string, route string) {
	return http.MethodPut, fmt.Sprintf(batchEditApplicationCommandPermissions, api, i.ApplicationID.String(), i.GuildID.String())
}
