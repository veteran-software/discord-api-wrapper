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
	"strconv"

	"github.com/veteran-software/discord-api-wrapper/logging"
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
