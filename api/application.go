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
	// Deprecated: application.summary now returns an empty string. This field will be removed in gateway v11
	Summary      string           `json:"summary"`               // if this application is a game sold on Discord, this field will be the summary field for the store page of its primary sku
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
	GatewayPresence               ApplicationFlags = 1 << 12 // GATEWAY_PRESENCE
	GatewayPresenceLimited        ApplicationFlags = 1 << 13 // GATEWAY_PRESENCE_LIMITED
	GatewayGuildMembers           ApplicationFlags = 1 << 14 // GATEWAY_GUILD_MEMBERS
	GatewayGuildMembersLimited    ApplicationFlags = 1 << 15 // GATEWAY_GUILD_MEMBERS_LIMITED
	VerificationPendingGuildLimit ApplicationFlags = 1 << 16 // VERIFICATION_PENDING_GUILD_LIMIT
	Embedded                      ApplicationFlags = 1 << 17 // EMBEDDED
	GatewayMessageContent         ApplicationFlags = 1 << 18 // GATEWAY_MESSAGE_CONTENT
	GatewayMessageContentLimited  ApplicationFlags = 1 << 19 // GATEWAY_MESSAGE_CONTENT_LIMITED
)
