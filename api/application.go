/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import "github.com/veteran-software/discord-api-wrapper/v10/oauth2"

// Application - an application which operates on Discord, commonly referred to as bots
//
//goland:noinspection SpellCheckingInspection
type Application struct {
	ID                             Snowflake        `json:"id"`                                          // ID of the app
	Name                           string           `json:"name"`                                        // Name of the app
	Icon                           *string          `json:"icon"`                                        // Icon hash of the app
	Description                    string           `json:"description"`                                 // Description of the app
	RpcOrigins                     []string         `json:"rpc_origins,omitempty"`                       // List of RPC origin URLs, if RPC is enabled
	BotPublic                      bool             `json:"bot_public"`                                  // When false, only the app owner can add the app to guilds
	BotRequireCodeGrant            bool             `json:"bot_require_code_grant"`                      // When true, the app's bot will only join upon completion of the full OAuth2 code grant flow
	Bot                            User             `json:"bot,omitempty"`                               // Partial user object for the bot user associated with the app
	TermsOfServiceURL              string           `json:"terms_of_service_url,omitempty"`              // URL of the app's Terms of Service
	PrivacyPolicyURL               string           `json:"privacy_policy_url,omitempty"`                // URL of the app's Privacy Policy
	Owner                          User             `json:"owner,omitempty"`                             // Partial user object for the owner of the app
	VerifyKey                      string           `json:"verify_key"`                                  // Hex encoded key for verification in interactions and the GameSDK's GetTicket
	Team                           *Team            `json:"team"`                                        // If the app belongs to a team, this will be a list of the members of that team
	GuildID                        Snowflake        `json:"guild_id,omitempty"`                          // Guild associated with the app. For example, a developer support server.
	Guild                          Guild            `json:"guild,omitempty"`                             // Partial object of the associated guild
	PrimarySkuID                   Snowflake        `json:"primary_sku_id,omitempty"`                    // If this app is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	Slug                           string           `json:"slug,omitempty"`                              // If this app is a game sold on Discord, this field will be the URL slug that links to the store page
	CoverImage                     string           `json:"cover_image,omitempty"`                       // App's default rich presence invite cover image hash
	Flags                          ApplicationFlags `json:"flags,omitempty"`                             // App's public flags
	ApproximateGuildCount          int64            `json:"approximate_guild_count,omitempty"`           // Approximate count of guilds the app has been added to
	RedirectUris                   []string         `json:"redirect_uris,omitempty"`                     // Array of redirect URIs for the app
	InteractionsEndpointUrl        string           `json:"interactions_endpoint_url,omitempty"`         // Interactions endpoint URL for the app
	RoleConnectionsVerificationURL string           `json:"role_connections_verification_url,omitempty"` // Role connection verification URL for the app
	Tags                           []string         `json:"tags,omitempty"`                              // List of tags describing the content and functionality of the app. Max of 5 tags.
	InstallParams                  InstallParams    `json:"install_params,omitempty"`                    // Settings for the app's default in-app authorization link, if enabled
	CustomInstallURL               string           `json:"custom_install_url,omitempty"`                // Default custom authorization URL for the app, if enabled
}

// ApplicationFlags - the application's public ApplicationFlags
type ApplicationFlags int64

//goland:noinspection GoUnusedConst
const (
	ApplicationAutoModerationRuleCreateBadge ApplicationFlags = 1 << 6  // Indicates if an app uses the Auto Moderation API
	GatewayPresence                          ApplicationFlags = 1 << 12 // Intent required for bots in 100 or more servers to receive presence_update events
	GatewayPresenceLimited                   ApplicationFlags = 1 << 13 // Intent required for bots in under 100 servers to receive presence_update events, found in Bot Settings
	GatewayGuildMembers                      ApplicationFlags = 1 << 14 // Intent required for bots in 100 or more servers to receive member-related events like guild_member_add. See list of member-related events under GUILD_MEMBERS
	GatewayGuildMembersLimited               ApplicationFlags = 1 << 15 // Intent required for bots in under 100 servers to receive member-related events like guild_member_add, found in Bot Settings. See list of member-related events under GUILD_MEMBERS
	VerificationPendingGuildLimit            ApplicationFlags = 1 << 16 // Indicates unusual growth of an app that prevents verification
	Embedded                                 ApplicationFlags = 1 << 17 // Indicates if an app is embedded within the Discord client (currently unavailable publicly)
	GatewayMessageContent                    ApplicationFlags = 1 << 18 // Intent required for bots in 100 or more servers to receive message content
	GatewayMessageContentLimited             ApplicationFlags = 1 << 19 // Intent required for bots in under 100 servers to receive message content, found in Bot Settings
	ApplicationCommandBadge                  ApplicationFlags = 1 << 23 // Indicates if an app has registered global application commands
)

// InstallParams - settings for the application's default in-app authorization link, if enabled
type InstallParams struct {
	Scopes      []*oauth2.Scopes `json:"scopes"`      // Scopes to add the application to the server with
	Permissions string           `json:"permissions"` // Permissions to request for the bot role
}
