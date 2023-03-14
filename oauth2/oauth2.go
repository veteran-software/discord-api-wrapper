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

package oauth2

import (
	"net/url"

	log "github.com/veteran-software/nowlive-logging"
)

// BaseAuthorizationURL - Base authorization URL
//
//goland:noinspection GoUnusedExportedFunction
func BaseAuthorizationURL() *url.URL {
	u, err := url.Parse("https://discord.com/api/oauth2/authorize")
	if err != nil {
		log.Errorln(log.FuncName(), err)
	}

	return u
}

// TokenURL - Token URL
//
//goland:noinspection GoUnusedExportedFunction
func TokenURL() *url.URL {
	u, err := url.Parse("https://discord.com/api/oauth2/token")
	if err != nil {
		log.Errorln(log.FuncName(), err)
	}

	return u
}

// TokenRevocationURL - Token Revocation URL
//
//goland:noinspection GoUnusedExportedFunction
func TokenRevocationURL() *url.URL {
	u, err := url.Parse("https://discord.com/api/oauth2/token/revoke")
	if err != nil {
		log.Errorln(log.FuncName(), err)
	}

	return u
}

// Scopes
//
// These are a list of all the OAuth2 scopes that Discord supports.
//
// Some scopes require approval from Discord to use.
//
// Requesting them from a user without approval from Discord may cause errors or undocumented behavior in the OAuth2 flow.
type Scopes string

// `guilds.join` and `bot` require you to have a bot account linked to your application.
// Also, in order to add a user to a guild, your bot has to already belong to that guild.
//
//goland:noinspection GoUnusedConst
const (
	ActivitiesRead                        Scopes = "activities.read"                          // allows your app to fetch data from a user's "Now Playing/Recently Played" list - requires Discord approval
	ActivitiesWrite                       Scopes = "activities.write"                         // allows your app to update a user's activity - requires Discord approval (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER)
	ApplicationsBuildsRead                Scopes = "applications.builds.read"                 // allows your app to read build data for a user's applications
	ApplicationsBuildsUpload              Scopes = "applications.builds.upload"               // allows your app to upload/update builds for a user's applications - requires Discord approval
	ApplicationsCommands                  Scopes = "applications.commands"                    // allows your app to use commands in a guild
	ApplicationsCommandsUpdate            Scopes = "applications.commands.update"             // allows your app to update its commands using a Bearer token - client credentials grant only
	ApplicationsCommandsPermissionsUpdate Scopes = "applications.commands.permissions.update" // allows your app to update permissions for its commands in a guild a user has permissions to
	ApplicationsEntitlements              Scopes = "applications.entitlements"                // allows your app to read entitlements for a user's applications
	ApplicationsStoreUpdate               Scopes = "applications.store.update"                // allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications
	Bot                                   Scopes = "bot"                                      // for oauth2 bots, this puts the bot in the user's selected guild by default
	Connections                           Scopes = "connections"                              // allows `/users/@me/connections` to return linked third-party accounts
	DmChannelsRead                        Scopes = "dm_channels.read"                         // allows your app to see information about the user's DMs and group DMs - requires Discord approval
	Email                                 Scopes = "email"                                    // enables `/users/@me` to return an email
	GdmJoin                               Scopes = "gdm.join"                                 // allows your app to join users to a group dm
	Guilds                                Scopes = "guilds"                                   // allows `/users/@me/guilds` to return basic information about all of a user's guilds
	GuildsJoin                            Scopes = "guilds.join"                              // allows `/guilds/{guild.id}/members/{user.id}` to be used for joining users to a guild
	GuildsMembersRead                     Scopes = "guilds.members.read"                      // allows `/users/@me/guilds/{guild.id}/member` to return a user's member information in a guild
	Identify                              Scopes = "identify"                                 // allows `/users/@me` without email
	MessagesRead                          Scopes = "messages.read"                            // for local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates)
	RelationshipsRead                     Scopes = "relationships.read"                       // allows your app to know a user's friends and implicit relationships - requires Discord approval
	Rpc                                   Scopes = "rpc"                                      // for local rpc server access, this allows you to control a user's local Discord client - requires Discord approval
	RpcActivitiesWrite                    Scopes = "rpc.activities.write"                     // for local rpc server access, this allows you to update a user's activity - requires Discord approval
	RpcNotificationsRead                  Scopes = "rpc.notifications.read"                   // for local rpc server access, this allows you to receive notifications pushed out to the user - requires Discord approval
	RpcVoiceRead                          Scopes = "rpc.voice.read"                           // for local rpc server access, this allows you to read a user's voice settings and listen for voice events - requires Discord approval
	RpcVoiceWrite                         Scopes = "rpc.voice.write"                          // for local rpc server access, this allows you to update a user's voice settings - requires Discord approval
	Voice                                 Scopes = "voice"                                    // allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval                                   //
	WebhookIncoming                       Scopes = "webhook.incoming"                         // this generates a webhook that is returned in the oauth token response for authorization code grants
)
