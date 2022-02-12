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

package routes

//goland:noinspection GoSnakeCaseUsage
const (
	Applications_Commands                   = "%s/applications/%s/commands"
	Applications_Commands_                  = "%s/applications/%s/commands/%s"
	Applications_Guilds_Command_Permissions = "%s/applications/%s/guilds/%s/command/%s/permissions"
	Applications_Guilds_CommandPermissions  = "%s/applications/%s/guilds/%s/command/permissions"
	Applications_Guilds_Commands            = "%s/applications/%s/guilds/%s/commands"
	Applications_Guilds_Commands_           = "%s/applications/%s/guilds/%s/commands/%s"
	Interaction__Callback                   = "%s/interactions/%s/%s/callback"
	Webhooks__                              = "%s/webhooks/%s/%s"
	Webhooks__Messages_                     = "%s/webhooks/%s/%s/messages/%s"
	Webhooks__MessagesOriginal              = "%s/webhooks/%s/%s/messages/@original"
)
