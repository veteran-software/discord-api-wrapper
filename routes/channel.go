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

//goland:noinspection GoSnakeCaseUsage,SpellCheckingInspection
const (
	Channels_                                 = "%s/channels/%s"
	Channels_Followers                        = "%s/channels/%s/followers"
	Channels_Invites                          = "%s/channels/%s/invites"
	Channels_Messages                         = "%s/channels/%s/messages"
	Channels_MessagesQsp                      = "%s/channels/%s/messages%s"
	Channels_Messages_                        = "%s/channels/%s/messages/%s"
	Channels_Messages_Crosspost               = "%s/channels/%s/messages/%s/crosspost"
	Channels_Messages_Reactions               = "%s/channels/%s/messages/%s/reactions"
	Channels_Messages_Reactions_              = "%s/channels/%s/messages/%s/reactions/%s"
	Channels_Messages_Reactions__             = "%s/channels/%s/messages/%s/reactions/%s/%s"
	Channels_Messages_Reactions_Me            = "%s/channels/%s/messages/%s/reactions/%s/@me"
	Channels_Messages_Threads                 = "%s/channels/%s/messages/%s/threads"
	Channels_MessagesBulkDelete               = "%s/channels/%s/messages/bulk-delete"
	Channels_Permissions_                     = "%s/channels/%s/permissions/%s"
	Channels_Pins                             = "%s/channels/%s/pins"
	Channels_Pins_                            = "%s/channels/%s/pins/%s"
	Channels_Recipients_                      = "%s/channels/%s/recipients/%s"
	Channels_ThreadMembers                    = "%s/channels/%s/thread-members"
	Channels_ThreadMembers_                   = "%s/channels/%s/thread-members/%s"
	Channels_ThreadMembersMe                  = "%s/channels/%s/thread-members/@me"
	Channels_Threads                          = "%s/channels/%s/threads"
	Channels_ThreadsArchivedPrivateQsp        = "%s/channels/%s/threads/archived/public%s"
	Channels_ThreadsArchivedPublicQsp         = "%s/channels/%s/threads/archived/public%s"
	Channels_Typing                           = "%s/channels/%s/typing"
	Channels_UsersMeThreadsArchivedPrivateQsp = "%s/channels/%s/users/@me/threads/archived/private%s"
)
