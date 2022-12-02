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

/*
Webhooks are a low-effort way to post messages to channels in Discord.

They do not require a bot user or authentication to use.
*/

// Webhook - Used to represent a webhook.
type Webhook struct {
	ID            Snowflake   `json:"id"`                       // the id of the webhook
	Type          WebhookType `json:"type"`                     // the type of the webhook
	GuildID       *Snowflake  `json:"guild_id,omitempty"`       // the guild id this webhook is for, if any
	ChannelID     *Snowflake  `json:"channel_id"`               // the channel id this webhook is for, if any
	User          User        `json:"user,omitempty"`           // the user this webhook was created by (not returned when getting a webhook with its token)
	Name          *string     `json:"name"`                     // the default name of the webhook
	Avatar        *string     `json:"avatar"`                   // the default user avatar hash of the webhook
	Token         string      `json:"token,omitempty"`          // the secure token of the webhook (returned for Incoming Webhooks)
	ApplicationID *Snowflake  `json:"application_id"`           // the bot/OAuth2 application that created this webhook
	SourceGuild   Guild       `json:"source_guild,omitempty"`   // the guild of the channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceChannel Channel     `json:"source_channel,omitempty"` // the channel that this webhook is following (returned for Channel Follower Webhooks)
	URL           string      `json:"url,omitempty"`            // the url used for executing the webhook (returned by the webhooks OAuth2 flow)
}

// WebhookType - the type of the webhook
type WebhookType int

//goland:noinspection GoUnusedConst
const (
	WebhookTypeIncoming        WebhookType = iota + 1 // Incoming Webhooks can post messages to channels with a generated token
	WebhookTypeChannelFollower                        // Channel Follower Webhooks are internal webhooks used with Channel Following to post new messages into channels
	WebhookTypeApplication                            // Application webhooks are webhooks used with Interactions
)
