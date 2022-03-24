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

// Application - an application which operates on Discord, commonly referred to as bots
//goland:noinspection SpellCheckingInspection
type Application struct {
	ID                  Snowflake `json:"id"`                             // the id of the app
	Name                string    `json:"name"`                           // the name of the app
	Icon                *string   `json:"icon"`                           // the icon hash of the app
	Description         string    `json:"description"`                    // the description of the app
	RpcOrigins          []string  `json:"rpc_origins"`                    // an array of rpc origin urls, if rpc is enabled
	BotPublic           bool      `json:"bot_public"`                     // when false only app owner can join the app's bot to guilds
	BotRequireCodeGrant bool      `json:"bot_require_code_grant"`         // when true the app's bot will only join upon completion of the full oauth2 code grant flow
	TermsOfServiceURL   string    `json:"terms_of_service_url,omitempty"` // the url of the app's terms of service
	PrivacyPolicyURL    string    `json:"privacy_policy_url,omitempty"`   // the url of the app's privacy policy
	Owner               User      `json:"owner,omitempty"`                // partial user object containing info on the owner of the application
	// Summary
	// Deprecated: application.summary now returns an empty string. This field will be removed in gateway v11
	Summary      string           `json:"summary"`               // deprecated: previously if this application was a game sold on Discord, this field would be the summary field for the store page of its primary SKU; now an empty string
	VerifyKey    string           `json:"verify_key"`            // the hex encoded key for verification in interactions and the GameSDK's GetTicket
	Team         *Team            `json:"team"`                  // if the application belongs to a team, this will be a list of the members of that team
	GuildID      Snowflake        `json:"guild_id,omitempty"`    // if this application is a game sold on Discord, this field will be the guild to which it has been linked
	PrimarySkuID Snowflake        `json:"primary_sku_id"`        // if this application is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	Slug         string           `json:"slug,omitempty"`        // if this application is a game sold on Discord, this field will be the URL slug that links to the store page
	CoverImage   string           `json:"cover_image,omitempty"` // the application's default rich presence invite cover image hash
	Flags        ApplicationFlags `json:"flags,omitempty"`       // the application's public ApplicationFlags
}

// ApplicationFlags - the application's public ApplicationFlags
type ApplicationFlags int64

//goland:noinspection GoUnusedConst
const (
	GatewayPresence               ApplicationFlags = 1 << 12 // Intent required for bots in 100 or more servers to receive presence_update events
	GatewayPresenceLimited        ApplicationFlags = 1 << 13 // Intent required for bots in under 100 servers to receive presence_update events, found in Bot Settings
	GatewayGuildMembers           ApplicationFlags = 1 << 14 // Intent required for bots in 100 or more servers to receive member-related events like guild_member_add. See list of member-related events under GUILD_MEMBERS
	GatewayGuildMembersLimited    ApplicationFlags = 1 << 15 // Intent required for bots in under 100 servers to receive member-related events like guild_member_add, found in Bot Settings. See list of member-related events under GUILD_MEMBERS
	VerificationPendingGuildLimit ApplicationFlags = 1 << 16 // Indicates unusual growth of an app that prevents verification
	Embedded                      ApplicationFlags = 1 << 17 // Indicates if an app is embedded within the Discord client (currently unavailable publicly)
	GatewayMessageContent         ApplicationFlags = 1 << 18 // Intent required for bots in 100 or more servers to receive message content
	GatewayMessageContentLimited  ApplicationFlags = 1 << 19 // Intent required for bots in under 100 servers to receive message content, found in Bot Settings
)
