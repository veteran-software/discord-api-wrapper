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

//goland:noinspection SpellCheckingInspection
type Application struct {
	ID                  Snowflake        `json:"id"`                             // ID - the id of the app
	Name                string           `json:"name"`                           // Name - the name of the app
	Icon                *string          `json:"icon"`                           // Icon - the icon hash of the app
	Description         string           `json:"description"`                    // Description - the description of the app
	RpcOrigins          []string         `json:"rpc_origins"`                    // RpcOrigins - an array of rpc origin urls, if rpc is enabled
	BotPublic           bool             `json:"bot_public"`                     // BotPublic - when false only app owner can join the app's bot to guilds
	BotRequireCodeGrant bool             `json:"bot_require_code_grant"`         // BotRequireCodeGrant - when true the app's bot will only join upon completion of the full oauth2 code grant flow
	TermsOfServiceURL   string           `json:"terms_of_service_url,omitempty"` // TermsOfServiceURL - the url of the app's terms of service
	PrivacyPolicyURL    string           `json:"privacy_policy_url,omitempty"`   // PrivacyPolicyURL - the url of the app's privacy policy
	Owner               User             `json:"owner,omitempty"`                // Owner - partial user object containing info on the owner of the application
	Summary             string           `json:"summary"`                        // Summary - if this application is a game sold on Discord, this field will be the summary field for the store page of its primary sku
	VerifyKey           string           `json:"verify_key"`                     // VerifyKey - the hex encoded key for verification in interactions and the GameSDK's GetTicket
	Team                *Team            `json:"team"`                           // Team - if the application belongs to a team, this will be a list of the members of that team
	GuildID             Snowflake        `json:"guild_id,omitempty"`             // GuildID - if this application is a game sold on Discord, this field will be the guild to which it has been linked
	PrimarySkuID        Snowflake        `json:"primary_sku_id"`                 // PrimarySkuID - if this application is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	Slug                string           `json:"slug,omitempty"`                 // Slug - if this application is a game sold on Discord, this field will be the URL slug that links to the store page
	CoverImage          string           `json:"cover_image,omitempty"`          // CoverImage - the application's default rich presence invite cover image hash
	Flags               ApplicationFlags `json:"flags,omitempty"`                // Flags - the application's public ApplicationFlags
}

type ApplicationFlags int64

const (
	GatewayPresence               ApplicationFlags = 1 << 12
	GatewayPresenceLimited        ApplicationFlags = 1 << 13
	GatewayGuildMembers           ApplicationFlags = 1 << 14
	GatewayGuildMembersLimited    ApplicationFlags = 1 << 15
	VerificationPendingGuildLimit ApplicationFlags = 1 << 16
	Embedded                      ApplicationFlags = 1 << 17
	GatewayMessageContent         ApplicationFlags = 1 << 18
	GatewayMessageContentLimited  ApplicationFlags = 1 << 19
)
