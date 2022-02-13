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
	Channels_Webhooks      = "%s/channels/%s/webhooks"
	Guilds_Webhooks        = "%s/guilds/%s/webhooks"
	Webhooks_              = "%s/webhooks/%s"
	Webhooks__Qsp          = "%s/webhooks/%s/%s%s"
	Webhooks__GitHubQsp    = "%s/webhooks/%s/%s/github%s"
	Webhooks__Messages_Qsp = "%s/webhooks/%s/%s/messages/%s%s"
	Webhooks__SlackQsp     = "%s/webhooks/%s/%s/slack%s"
)