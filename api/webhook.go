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
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
Webhooks are a low-effort way to post messages to channels in Discord.

They do not require a bot user or authentication to use.
*/

const thrID = "thread_id="

// Webhook - Used to represent a webhook.
type Webhook struct {
	ID            Snowflake   `json:"id"`                       // ID - the id of the webhook
	Type          WebhookType `json:"type"`                     // Type - the type of the webhook
	GuildID       *Snowflake  `json:"guild_id,omitempty"`       // GuildID - the guild id this webhook is for, if any
	ChannelID     *Snowflake  `json:"channel_id"`               // ChannelID - the channel id this webhook is for, if any
	User          User        `json:"user,omitempty"`           // User - the user this webhook was created by (not returned when getting a webhook with its token)
	Name          *string     `json:"name"`                     // Name - the default name of the webhook
	Avatar        *string     `json:"avatar"`                   // Avatar - the default user avatar hash of the webhook
	Token         string      `json:"token,omitempty"`          // Token - the secure token of the webhook (returned for Incoming Webhooks)
	ApplicationID *Snowflake  `json:"application_id"`           // ApplicationID - the bot/OAuth2 application that created this webhook
	SourceGuild   Guild       `json:"source_guild,omitempty"`   // SourceGuild - the guild of the channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceChannel Channel     `json:"source_channel,omitempty"` // SourceChannel - the channel that this webhook is following (returned for Channel Follower Webhooks)
	URL           string      `json:"url,omitempty"`            // URL - the url used for executing the webhook (returned by the webhooks OAuth2 flow)
}

type WebhookType int

const (
	WebhookTypeIncoming        WebhookType = iota + 1 // WebhookTypeIncoming - Incoming Webhooks can post messages to channels with a generated token
	WebhookTypeChannelFollower                        // WebhookTypeChannelFollower - Channel Follower Webhooks are internal webhooks used with Channel Following to post new messages into channels
	WebhookTypeApplication                            // WebhookTypeApplication - Application webhooks are webhooks used with Interactions
)

// CreateWebhook
//
// Create a new webhook.
//
// Requires the ManageWebhooks permission.
//
// Returns a Webhook object on success.
//
// Webhook names follow our naming restrictions that can be found in our Usernames and Nicknames documentation, with the following additional stipulations:
//
// 	* Webhook names cannot be: 'clyde'
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (c *Channel) CreateWebhook() (string, string) {
	return http.MethodPost, fmt.Sprintf(createWebhook, api, c.ID.String())
}

type CreateWebhookJSON struct {
	Name   string           `json:"name"`             // Name - name of the webhook (1-80 characters)
	Avatar *base64.Encoding `json:"avatar,omitempty"` // Avatar - image for the default webhook avatar
}

// GetChannelWebhooks - Returns a list of channel webhook objects. Requires the ManageWebhooks permission.
func (c *Channel) GetChannelWebhooks() (string, string) {
	return http.MethodGet, fmt.Sprintf(getChannelWebhooks, api, c.ID.String())
}

// GetGuildWebhooks - Returns a list of guild webhook objects. Requires the ManageWebhooks permission.
func (g *Guild) GetGuildWebhooks() (string, string) {
	return http.MethodGet, fmt.Sprintf(getGuildWebhooks, api, g.ID.String())
}

// GetWebhook - Returns the new webhook object for the given id.
func (w *Webhook) GetWebhook() (string, string) {
	return http.MethodGet, fmt.Sprintf(getWebhook, api, w.ID.String())
}

// GetWebhookWithToken - Same as above, except this call does not require authentication and returns no user in the webhook object.
func (w *Webhook) GetWebhookWithToken() (string, string) {
	return http.MethodGet, fmt.Sprintf("%s/webhooks/%s/%s", api, w.ID.String(), w.Token)
}

// ModifyWebhook - Modify a webhook. Requires the ManageWebhooks permission. Returns the updated webhook object on success.
//
// All parameters to this endpoint are optional
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (w *Webhook) ModifyWebhook() (string, string) {
	return http.MethodPatch, fmt.Sprintf(modifyWebhook, api, w.ID.String())
}

type ModifyWebhookJSON struct {
	Name      string           `json:"name,omitempty"`       // Name - the default name of the webhook
	Avatar    *base64.Encoding `json:"avatar,omitempty"`     // Avatar - image for the default webhook avatar
	ChannelID Snowflake        `json:"channel_id,omitempty"` // ChannelID - the new channel id this webhook should be moved to
}

// ModifyWebhookWithToken - Same as above, except this call does not require authentication, does not accept a channel_id parameter in the body, and does not return a user in the webhook object.
func (w *Webhook) ModifyWebhookWithToken() (string, string) {
	return http.MethodPatch, fmt.Sprintf(modifyWebhookWithToken, api, w.ID.String(), w.Token)
}

// DeleteWebhook - Delete a webhook permanently. Requires the ManageWebhooks permission. Returns a 204 No Content response on success.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (w *Webhook) DeleteWebhook() (string, string) {
	return http.MethodDelete, fmt.Sprintf(deleteWebhook, api, w.ID.String())
}

// DeleteWebhookWithToken - Same as above, except this call does not require authentication.
func (w *Webhook) DeleteWebhookWithToken() (string, string) {
	return http.MethodDelete, fmt.Sprintf(deleteWebhookWithToken, api, w.ID.String(), w.Token)
}

// ExecuteWebhook - Refer to Uploading Files for details on attachments and multipart/form-data requests.
//
// Note that when sending a message, you must provide a value for at least one of content, embeds, or file.
//
// wait and threadID are optional; pass nil if not needed
func (w *Webhook) ExecuteWebhook(wait *bool, threadID *Snowflake) (string, string) {
	var qsp []string
	if wait != nil {
		qsp = append(qsp, "wait="+strconv.FormatBool(*wait))
	}
	if threadID != nil {
		qsp = append(qsp, thrID+threadID.String())
	}
	var query string
	if len(qsp) > 0 {
		query = "?" + strings.Join(qsp, "&")
	}

	return http.MethodPost, fmt.Sprintf(executeWebhook, api, w.ID.String(), w.Token, query)
}

type ExecuteWebhookJSON struct {
	Content         string          `json:"content"`                    // Content - the message contents (up to 2000 characters); Required - one of content, file, embeds
	Username        string          `json:"username,omitempty"`         // Username - override the default username of the webhook; Required - false
	AvatarURL       string          `json:"avatar_url,omitempty"`       // AvatarURL - override the default avatar of the webhook; Required - false
	Tts             bool            `json:"tts,omitempty"`              // Tts - true if this is a TTS message; Required - false
	Embeds          []Embed         `json:"embeds"`                     // Embeds - embedded rich content; Required - one of content, file, embeds
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitempty"` // AllowedMentions - allowed mentions for the message; Required - false
	Components      []Component     `json:"components,omitempty"`       // Components - the components to include with the message - Required - false
	PayloadJson     string          `json:"payload_json"`               // PayloadJson - JSON encoded body of non-file params; Required - "multipart/form-data" only
	Attachments     []Attachment    `json:"attachments,omitempty"`      // Attachments - Attachment objects with filename and description; Required - false
	Flags           MessageFlags    `json:"flags,omitempty"`            // Flags - MessageFlags combined as a bitfield (only SuppressEmbeds can be set)
}

// ExecuteSlackCompatibleWebhook - Refer to Slack's documentation for more information. We do not support Slack's channel, icon_emoji, mrkdwn, or mrkdwn_in properties.
//
// wait and threadID are optional; pass nil if not needed
func (w *Webhook) ExecuteSlackCompatibleWebhook(wait *bool, threadID *Snowflake) (string, string) {
	var qsp []string
	if wait != nil {
		qsp = append(qsp, "wait="+strconv.FormatBool(*wait))
	}
	if threadID != nil {
		qsp = append(qsp, thrID+threadID.String())
	}
	var query string
	if len(qsp) > 0 {
		query = "?" + strings.Join(qsp, "&")
	}

	return http.MethodPost, fmt.Sprintf(executeSlackCompatibleWebhook, api, w.ID.String(), w.Token, query)
}

// ExecuteGitHubCompatibleWebhook
//
// Add a new webhook to your GitHub repo (in the repo's settings), and use this endpoint as the "Payload URL."
//
// You can choose what events your Discord channel receives by choosing the "Let me select individual events" option and selecting individual events for the new webhook you're configuring.
//
// wait and threadID are optional; pass nil if not needed
func (w *Webhook) ExecuteGitHubCompatibleWebhook(wait *bool, threadID *Snowflake) (string, string) {
	var qsp []string
	if wait != nil {
		qsp = append(qsp, "wait="+strconv.FormatBool(*wait))
	}
	if threadID != nil {
		qsp = append(qsp, thrID+threadID.String())
	}
	var query string
	if len(qsp) > 0 {
		query = "?" + strings.Join(qsp, "&")
	}

	return http.MethodPost, fmt.Sprintf(executeGitHubCompatibleWebhook, api, w.ID.String(), w.Token, query)
}

// GetWebhookMessage - Returns a previously-sent webhook message from the same token. Returns a message object on success.
//
// threadID is optional; pass nil if not needed
func (w *Webhook) GetWebhookMessage(msgID Snowflake, threadID *Snowflake) (string, string) {
	var query string
	if threadID != nil {
		query = "?" + thrID + threadID.String()
	}

	return http.MethodGet, fmt.Sprintf(getWebhookMessage, api, w.ID.String(), w.Token, msgID.String(), query)
}

// EditWebhookMessage
//
// Edits a previously-sent webhook message from the same token. Returns a message object on success.
//
// When the content field is edited, the mentions array in the message object will be reconstructed from scratch based on the new content.
// The allowed_mentions field of the edit request controls how this happens.
// If there is no explicit allowed_mentions in the edit request, the content will be parsed with default allowances, that is, without regard to whether or not an allowed_mentions was present in the request that originally created the message.
//
// Refer to Uploading Files for details on attachments and multipart/form-data requests.
// Any provided files will be appended to the message.
// To remove or replace files you will have to supply the "attachments" field which specifies the files to retain on the message after edit.
//
// Starting with API v10, the attachments array must contain all attachments that should be present after edit, including retained and new attachments provided in the request body.
//
// All JSON parameters to this endpoint are optional and nullable.
//
// threadID is optional; pass nil if not needed
func (w *Webhook) EditWebhookMessage(msgID Snowflake, threadID *Snowflake) (string, string) {
	var query string
	if threadID != nil {
		query = "?" + thrID + threadID.String()
	}

	return http.MethodPatch, fmt.Sprintf(editWebhookMessage, api, w.ID.String(), w.Token, msgID.String(), query)
}

// EditWebhookMessageJSON - All parameters to this endpoint are optional and nullable.
type EditWebhookMessageJSON struct {
	Content         *string          `json:"content,omitempty"`          // Content - the message contents (up to 2000 characters)
	Embeds          *[]Embed         `json:"embeds,omitempty"`           // Embeds - embedded rich content
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"` // AllowedMentions - allowed mentions for the message
	Components      *[]Component     `json:"components,omitempty"`       // Components - the components to include with the message
	PayloadJson     *string          `json:"payload_json,omitempty"`     // PayloadJson - JSON encoded body of non-file params (multipart/form-data only)
	Attachments     *[]Attachment    `json:"attachments,omitempty"`      // Attachments - attached files to keep and possible descriptions for new files
}

// DeleteWebhookMessage - Deletes a message that was created by the webhook. Returns a 204 No Content response on success.
//
// threadID is optional; pass nil if not needed
func (w *Webhook) DeleteWebhookMessage(msgID Snowflake, threadID *Snowflake) (string, string) {
	var query string
	if threadID != nil {
		query = "?" + thrID + threadID.String()
	}

	return http.MethodDelete, fmt.Sprintf(deleteWebhookMessage, api, w.ID.String(), w.Token, msgID.String(), query)
}
