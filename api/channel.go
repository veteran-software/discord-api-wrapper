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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/veteran-software/discord-api-wrapper/logging"
	"github.com/veteran-software/discord-api-wrapper/routes"
	"github.com/veteran-software/discord-api-wrapper/utilities"
)

/* CHANNEL OBJECT */

// Channel - Represents a guild or DM channel within Discord.
type Channel struct {
	// ID - the id of this channel
	ID                         Snowflake      `json:"id"`
	Type                       ChannelType    `json:"type"`
	GuildID                    Snowflake      `json:"guild_id,omitempty"`
	Position                   int8           `json:"position,omitempty"`
	PermissionOverwrites       []Overwrite    `json:"permission_overwrites,omitempty"`
	Name                       string         `json:"name,omitempty"`
	Topic                      *string        `json:"topic,omitempty"`
	Nsfw                       bool           `json:"nsfw,omitempty"`
	LastMessageID              *Snowflake     `json:"last_message_id,omitempty"`
	Bitrate                    int64          `json:"bitrate,omitempty"`
	UserLimit                  int64          `json:"user_limit,omitempty"`
	RateLimitPerUser           int64          `json:"rate_limit_per_user,omitempty"`
	Recipients                 []User         `json:"recipients,omitempty"`
	Icon                       *string        `json:"icon,omitempty"`
	OwnerID                    Snowflake      `json:"owner_id,omitempty"`
	ApplicationID              Snowflake      `json:"application_id,omitempty"`
	ParentID                   *Snowflake     `json:"parent_id,omitempty"`
	LastPinTimestamp           *time.Time     `json:"last_pin_timestamp,omitempty"`
	RtcRegion                  *string        `json:"rtc_region,omitempty"`
	VideoQualityMode           int64          `json:"video_quality_mode,omitempty"`
	MessageCount               int64          `json:"message_count,omitempty"`
	MemberCount                int64          `json:"member_count,omitempty"`
	ThreadMetadata             ThreadMetadata `json:"thread_metadata,omitempty"`
	Member                     ThreadMember   `json:"member,omitempty"`
	DefaultAutoArchiveDuration int64          `json:"default_auto_archive_duration,omitempty"`
	// Only available from Interaction Webhooks
	Permissions string `json:"permissions"`
}

type ChannelType int

const (
	GuildText ChannelType = iota
	DM
	GuildVoice
	GroupDM
	GuildCategory
	GuildNews
	GuildStore
	GuildNewsThread ChannelType = iota + 3
	GuildPublicThread
	GuildPrivateThread
	GuildStageVoice
)

type VideoQualityMode int

//goland:noinspection GoUnusedConst
const (
	// Auto - Discord chooses the quality for optimal performance
	Auto VideoQualityMode = iota + 1
	// Full - 720p
	Full
)

/* MESSAGE OBJECT */

type Message struct {
	ID                Snowflake          `json:"id,omitempty"`
	ChannelID         Snowflake          `json:"channel_id,omitempty"`
	GuildID           Snowflake          `json:"guild_id,omitempty"`
	Author            User               `json:"author,omitempty"`
	Member            GuildMember        `json:"member,omitempty"`
	Content           string             `json:"content,omitempty"`
	Timestamp         time.Time          `json:"timestamp,omitempty"`
	EditedTimestamp   *time.Time         `json:"edited_timestamp,omitempty"`
	TTS               bool               `json:"tts,omitempty"`
	MentionEveryone   bool               `json:"mention_everyone,omitempty"`
	Mentions          []User             `json:"mentions,omitempty"`
	MentionRoles      []Snowflake        `json:"mention_roles,omitempty"`
	MentionChannels   []Channel          `json:"mention_channels,omitempty"`
	Attachments       []Attachment       `json:"attachments,omitempty"`
	Embeds            []Embed            `json:"embeds,omitempty"`
	Reactions         []ReactionObject   `json:"reactions,omitempty"`
	Nonce             interface{}        `json:"nonce,string,omitempty"`
	Pinned            bool               `json:"pinned,omitempty"`
	WebhookID         Snowflake          `json:"webhook_id,omitempty"`
	Type              MessageType        `json:"type,omitempty"`
	ApplicationID     Snowflake          `json:"application_id,omitempty"`
	MessageReference  MessageReference   `json:"message_reference,omitempty"`
	Flags             MessageFlags       `json:"flags,omitempty"`
	ReferencedMessage *Message           `json:"referenced_message,omitempty"`
	Interaction       MessageInteraction `json:"interaction,omitempty"`
	Thread            Channel            `json:"thread,omitempty"`
	Components        []Component        `json:"components,omitempty"`
}

type MessageType int

//goland:noinspection GoUnusedConst,SpellCheckingInspection
const (
	Default MessageType = iota
	RecipientAdd
	RecipientRemove
	Call
	ChannelNameChange
	ChannelIconChange
	ChannelPinnedMessage
	GuildMemberJoin
	UserPremiumGuildSubscription
	UserPremiumGuildSubscriptionTier1
	UserPremiumGuildSubscriptionTier2
	UserPremiumGuildSubscriptionTier3
	ChannelFollowAdd
	_
	GuildDiscoveryDisqualified
	GuildDiscoveryRequalified
	GuildDiscoveryGracePeriodInitialWarning
	GuildDiscoveryGracePeriodFinalWarning
	ThreadCreated
	Reply
	ChatInputCommand
	ThreadStarterMessage
	GuildInviteReminder
	ContextMenuCommand
)

type MessageFlags int

const (
	CrossPosted                      MessageFlags = 1 << 0 // this message has been published to subscribed channels (via Channel Following)
	IsCrossPost                      MessageFlags = 1 << 1 // this message originated from a message in another channel (via Channel Following)
	SuppressEmbeds                   MessageFlags = 1 << 2 // do not include any embeds when serializing this message
	SourceMessageDeleted             MessageFlags = 1 << 3 // the source message for this crosspost has been deleted (via Channel Following)
	Urgent                           MessageFlags = 1 << 4 // this message came from the urgent message system
	HasThread                        MessageFlags = 1 << 5 // this message has an associated thread, with the same id as the message
	Ephemeral                        MessageFlags = 1 << 6 // this message is only visible to the user who invoked the Interaction
	Loading                          MessageFlags = 1 << 7 // this message is an Interaction Response and the bot is "thinking"
	FailedToMentionSomeRolesInThread MessageFlags = 1 << 8 // this message failed to mention some roles and add their members to the thread
)

/* MESSAGE REFERENCE OBJECT */

type MessageReference struct {
	MessageID       Snowflake `json:"message_id,omitempty"`
	ChannelID       Snowflake `json:"channel_id,omitempty"`
	GuildID         Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"` // default true
}

/* FOLLOWED CHANNEL OBJECT */

// TODO

/* REACTION OBJECT */

type ReactionObject struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

/* OVERWRITE OBJECT */

type Overwrite struct {
	ID    string         `json:"id"`
	Type  PermissionType `json:"type"`
	Allow string         `json:"allow"`
	Deny  string         `json:"deny"`
}

type PermissionType int

const (
	PermissionRole PermissionType = iota
	PermissionMember
)

/* THREAD METADATA OBJECT */

type ThreadMetadata struct {
	Archived            bool      `json:"archived"`
	AutoArchiveDuration int64     `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time `json:"archive_timestamp"`
	Locked              bool      `json:"locked"`
	Invitable           bool      `json:"invitable,omitempty"`
}

/* THREAD MEMBER OBJECT */

type ThreadMember struct {
	ID            *Snowflake `json:"id"`
	UserID        *Snowflake `json:"user_id"`
	JoinTimestamp time.Time  `json:"join_timestamp"`
	Flags         int64      `json:"flags"`
}

/* EMBED OBJECT */

/*
Embed

Title: title of embed

Type: EmbedType (always RichEmbed for webhook embeds)

Description: description of embed

URL: url of embed

Timestamp: timestamp of embed content

Color: color code of the embed

Footer: footer information

Image: image information

Thumbnail: thumbnail information

Author: author information

Fields: fields information
*/
type Embed struct {
	Title       string     `json:"title,omitempty"`
	Type        EmbedType  `json:"type,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url,omitempty"`
	Timestamp   string     `json:"timestamp,omitempty"`
	Color       int64      `json:"color,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Fields      []*Field   `json:"fields,omitempty"`
}

/*
EmbedType

Embed types are "loosely defined" and, for the most part, are not used by our clients for rendering. Embed attributes power what is rendered. Embed types should be considered deprecated and might be removed in a future API version.

--------

RichEmbed: generic embed rendered from embed attributes

imageEmbed: image embed

videoEmbed: video embed

gifVEmbed: animated gif image embed rendered as a video embed

articleEmbed: article embed

linkEmbed: link embed
*/
type EmbedType string

const (
	RichEmbed    EmbedType = "rich"
	imageEmbed   EmbedType = "image"
	videoEmbed   EmbedType = "video"
	gifVEmbed    EmbedType = "gifv"
	articleEmbed EmbedType = "article"
	linkEmbed    EmbedType = "link"
)

type Thumbnail struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Video struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Image struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Provider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Author struct {
	Name    string  `json:"name,omitempty"`
	URL     string  `json:"url,omitempty"`
	IconURL *string `json:"icon_url,omitempty"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

/* ATTACHMENT OBJECT */

type Attachment struct {
	ID          Snowflake `json:"id"`
	Filename    string    `json:"filename"`
	Description string    `json:"description,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Size        int       `json:"size"`
	URL         string    `json:"url"`
	ProxyURL    string    `json:"proxy_url"`
	Height      int       `json:"height,omitempty"`
	Width       int       `json:"width,omitempty"`
	Ephemeral   bool      `json:"ephemeral,omitempty"`
}

/* CHANNEL MENTION OBJECT */

type ChannelMention struct {
	ID      Snowflake   `json:"id"`
	GuildID Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}

/* ALLOWED MENTIONS OBJECT */

type AllowedMentionType string

const (
	RoleMentions     AllowedMentionType = "roles"
	UserMentions     AllowedMentionType = "users"
	EveryoneMentions AllowedMentionType = "everyone"
)

type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []Snowflake          `json:"roles,omitempty"`
	Users       []Snowflake          `json:"users,omitempty"`
	RepliedUser bool                 `json:"replied_user,omitempty"`
}

/* HELPER FUNCTIONS */

// Embed limits
const (
	titleLimit       = 256
	descriptionLimit = 4096
	fieldCount       = 25
	fieldNameLimit   = 256
	fieldValueLimit  = 1024
	footerTextLimit  = 2048
	authorNameLimit  = 256
)

func (e *Embed) isValidLength() bool {
	if len(e.Title) <= titleLimit && len(e.Description) <= descriptionLimit && len(e.Fields) <= fieldCount && len(e.Footer.Text) <= footerTextLimit && len(e.Author.Name) <= authorNameLimit {
		for _, field := range e.Fields {
			if len(field.Name) > fieldNameLimit || len(field.Value) > fieldValueLimit {
				return false
			}
		}

		return true
	}

	return false
}

func (c *Channel) String() string {
	var chanType string

	switch c.Type {
	case GuildText:
		chanType = "GTC:"
	case DM:
		chanType = "DM:"
	case GroupDM:
		chanType = "GDM:"
	case GuildNews:
		chanType = "GNC:"
	case GuildNewsThread:
		chanType = "GNT:"
	case GuildPublicThread:
		chanType = "GPuT:"
	case GuildPrivateThread:
		chanType = "GPrT:"
	}

	return chanType + c.Name + "(" + c.ID.String() + ")"
}

/* API endpoints */

// GetChannel
//
// Get a channel by ID.
//
// Returns a channel object.
//
// If the channel is a thread, a thread member object is included in the returned result.
func (c *Channel) GetChannel() *Channel {
	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(routes.Channels_, api, c.ID), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channel Channel
	err = json.NewDecoder(resp.Body).Decode(&channel)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return &channel
}

/*ModifyChannel

Update a channel's settings.

Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.

All JSON parameters are optional.

This endpoint supports the "X-Audit-Log-Reason" header.
*/
func (c *Channel) ModifyChannel(dm *map[string]interface{}, guildChannel *map[string]interface{}, name *string, icon *base64.Encoding, reason *string) *Channel {
	var payload interface{}

	switch c.Type {
	case GroupDM:
		payload = struct {
			Name string `json:"name"` // 1-100 character channel name
			Icon string `json:"icon"` // base64 encoded icon
		}{
			Name: fmt.Sprintf("%v", (*dm)["name"]),
			Icon: fmt.Sprintf("%v", (*dm)["icon"]),
		}
	case GuildPublicThread:
		fallthrough
	case GuildPrivateThread:
		fallthrough
	case GuildNewsThread:
		archived, _ := strconv.ParseBool(fmt.Sprintf("%v", (*dm)["archived"]))

		payload = struct {
			Name                string `json:"name"`                  // 1-100 character channel name
			Archived            bool   `json:"archived"`              // whether the thread is archived
			AutoArchiveDuration int    `json:"auto_archive_duration"` // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
			Locked              bool   `json:"locked"`                // whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
			Invitable           bool   `json:"invitable"`             // whether non-moderators can add other non-moderators to a thread; only available on private threads
			RateLimitPerUser    *int   `json:"rate_limit_per_user"`   // amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages, manage_thread, or manage_channel, are unaffected
		}{
			Name:     fmt.Sprintf("%v", (*dm)["name"]),
			Archived: archived,
		}
	}

	if c.Type == GroupDM {
		payload = struct {
			Name string `json:"name"` // 1-100 character channel name
			Icon string `json:"icon"` // base64 encoded icon
		}{
			Name: fmt.Sprintf("%v", (*dm)["name"]),
			Icon: fmt.Sprintf("%v", (*dm)["icon"]),
		}

	} else if c.Type == GuildText || c.Type == GuildVoice || c.Type == GuildCategory || c.Type == GuildNews || c.Type == GuildStore || c.Type == GuildStageVoice {
		payload = struct {
			Name string          `json:"name"` // 1-100 character channel name
			Icon base64.Encoding `json:"icon"` // base64 encoded icon
		}{
			Name: fmt.Sprintf("%v", (*dm)["name"]),
			Icon: *icon,
		}
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(routes.Channels_, api, c.ID), &payload, reason)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channel Channel
	err = json.NewDecoder(resp.Body).Decode(&channel)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return &channel
}

/*
ModifyGuildChannelJSON

Requires the MANAGE_CHANNELS permission for the guild.

Fires a Channel Update Gateway event.

If modifying a category, individual Channel Update events will fire for each child channel that also changes.

If modifying permission overwrites, the MANAGE_ROLES permission is required.

Only permissions your bot has in the guild or channel can be allowed/denied (unless your bot has a MANAGE_ROLES overwrite in the channel).
*/
type ModifyGuildChannelJSON struct {
	// All
	Name                 string       `json:"name"`                  // 1-100 character channel name
	Position             *int         `json:"position,omitempty"`    // the position of the channel in the left-hand listing
	PermissionOverwrites *[]Overwrite `json:"permission_overwrites"` // channel or category-specific permissions

	// Text
	RateLimitPerUser *int `json:"rate_limit_per_user"` // amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages or manage_channel, are unaffected

	// Text, News
	Type                       ChannelType `json:"type,omitempty"`                // the type of channel; only conversion between text and news is supported and only in guilds with the "NEWS" feature
	Topic                      *string     `json:"topic"`                         // 0-1024 character channel topic
	DefaultAutoArchiveDuration *int        `json:"default_auto_archive_duration"` // the default duration that the clients use (not the API) for newly created threads in the channel, in minutes, to automatically archive the thread after recent activity

	// Text, News, Store
	Nsfw *bool `json:"nsfw"` // whether the channel is nsfw

	// Text, News, Store, Voice
	ParentID *Snowflake `json:"parent_id"` // id of the new parent category for a channel

	// Voice
	Bitrate          *int             `json:"bitrate"`            // the bitrate (in bits) of the voice channel; 8000 to 96000 (128000 for VIP servers)
	UserLimit        *int             `json:"user_limit"`         // the user limit of the voice channel; 0 refers to no limit, 1 to 99 refers to a user limit
	RtcRegion        *string          `json:"rtc_region"`         // channel voice region id, automatic when set to null
	VideoQualityMode VideoQualityMode `json:"video_quality_mode"` // the camera video quality mode of the voice channel
}

/*
ModifyThreadJSON

When setting archived to false, when locked is also false, only the SEND_MESSAGES permission is required.

Otherwise, requires the MANAGE_THREADS permission. Fires a Thread Update Gateway event. Requires the thread to have archived set to false or be set to false in the request.
*/
type ModifyThreadJSON struct {
	Name                string `json:"name"`                  // 1-100 character channel name
	Archived            bool   `json:"archived"`              // whether the thread is archived
	AutoArchiveDuration int    `json:"auto_archive_duration"` // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	Locked              bool   `json:"locked"`                // whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
	Invitable           bool   `json:"invitable"`             // whether non-moderators can add other non-moderators to a thread; only available on private threads
	RateLimitPerUser    *int   `json:"rate_limit_per_user"`   // amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages, manage_thread, or manage_channel, are unaffected
}

/*
DeleteChannel

Delete a channel, or close a private message.

Requires the MANAGE_CHANNELS permission for the guild, or MANAGE_THREADS if the channel is a thread.

Deleting a category does not delete its child channels; they will have their parent_id removed and a Channel Update Gateway event will fire for each of them.

Returns a channel object on success. Fires a Channel Delete Gateway event (or Thread Delete if the channel was a thread).

Deleting a guild channel cannot be undone. Use this with caution, as it is impossible to undo this action when performed on a guild channel. In contrast, when used with a private message, it is possible to undo the action by opening a private message with the recipient again.

For Community guilds, the Rules or Guidelines channel and the Community Updates channel cannot be deleted.

This endpoint supports the X-Audit-Log-Reason header.
*/
func (c *Channel) DeleteChannel() (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_, api, c.ID)
}

/*
GetChannelMessages

Returns the messages for a channel.

If operating on a guild channel, this endpoint requires the VIEW_CHANNEL permission to be present on the current user.

If the current user is missing the 'READ_MESSAGE_HISTORY' permission in the channel then this will return no messages (since they cannot read the message history).

Returns an array of message objects on success.

SUPPORTS: "around : Snowflake"; "before : Snowflake"; "after : Snowflake"; "limit : int" ; nil
*/
func (c *Channel) GetChannelMessages(opts *map[string]interface{}) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_MessagesQsp, api, c.ID.String(), *utilities.ParseQueryString(opts))
}

/*
GetChannelMessage

Returns a specific message in the channel.

If operating on a guild channel, this endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present on the current user.

Returns a message object on success.
*/
func (c *Channel) GetChannelMessage(messageID string) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_Messages_, api, c.ID.String(), messageID)
}

/*
CreateMessage

Post a message to a guild text or DM channel. Returns a message object.

Fires a Message Create Gateway event.

See message formatting for more information on how to properly format messages.

Limitations
   * When operating on a guild channel, the current user must have the SEND_MESSAGES permission.
   * When sending a message with tts (text-to-speech) set to true, the current user must have the SEND_TTS_MESSAGES permission.
   * When creating a message as a reply to another message, the current user must have the READ_MESSAGE_HISTORY permission.
       * The referenced message must exist and cannot be a system message.
   * The maximum request size when sending a message is 8 MB
   * For the embed object, you can set every field except type (it will be rich regardless of if you try to set it), provider, video, and any height, width, or proxy_url values for images.
   * Files can only be uploaded when using the multipart/form-data content type.

You may create a message as a reply to another message.

To do so, include a message_reference with a message_id.

The channel_id and guild_id in the message_reference are optional, but will be validated if provided.

Note that when sending a message, you must provide a value for at least one of content, embeds, or file.

For a file attachment, the Content-Disposition subpart header MUST contain a filename parameter.

This endpoint supports both application/json and multipart/form-data bodies.

When uploading files the multipart/form-data content type must be used.

Note that in multipart form data, the embeds and allowed_mentions fields cannot be used.

You can pass a stringified JSON body as a form value as payload_json instead.

If you supply a payload_json form value, all fields except for file fields will be ignored in the form data.
*/
func (c *Channel) CreateMessage() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Messages, api, c.ID)
}

type CreateMessageJSON struct {
	Content          string           `json:"content"`           // the message contents (up to 2000 characters)
	TTS              bool             `json:"tts"`               // true if this is a TTS message
	Embeds           []Embed          `json:"embeds"`            // embedded rich content (up to 6000 characters)
	AllowedMentions  AllowedMentions  `json:"allowed_mentions"`  // allowed mentions for the message
	MessageReference MessageReference `json:"message_reference"` // include to make your message a reply
	Components       []Component      `json:"components"`        // the components to include with the message
	StickerIDs       []Snowflake      `json:"sticker_ids"`       // IDs of up to 3 stickers in the server to send in the message
	PayloadJson      string           `json:"payload_json"`      // JSON encoded body of non-file params
	Attachments      []Attachment     `json:"attachments"`       // attachment objects with filename and description
	Flags            MessageFlags     `json:"flags"`             // message flags combined as a bitfield (only SUPPRESS_EMBEDS can be set)
}

/*
CrosspostMessage

Crosspost a message in a News Channel to following channels.

This endpoint requires the 'SEND_MESSAGES' permission, if the current user sent the message, or additionally the 'MANAGE_MESSAGES' permission, for all other messages, to be present for the current user.

Returns a message object.
*/
//goland:noinspection SpellCheckingInspection
func (c *Channel) CrosspostMessage(messageID string) (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Messages_Crosspost, api, c.ID.String(), messageID)
}

/*
CreateReaction

Create a reaction for the message.

This endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present on the current user.

Additionally, if nobody else has reacted to the message using this emoji, this endpoint requires the 'ADD_REACTIONS' permission to be present on the current user.

Returns a 204 empty response on success.

The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.

To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
*/
func (c *Channel) CreateReaction(messageID Snowflake, emoji string) (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_Messages_Reactions_Me, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji))
}

/*
DeleteOwnReaction

Delete a reaction the current user has made for the message.

Returns a 204 empty response on success.

The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.

To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
*/
func (c *Channel) DeleteOwnReaction(messageID Snowflake, emoji string) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Messages_Reactions_Me, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji))
}

/*
DeleteUserReaction

Deletes another user's reaction.

This endpoint requires the 'MANAGE_MESSAGES' permission to be present on the current user.

Returns a 204 empty response on success. The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.

To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
*/
func (c *Channel) DeleteUserReaction(messageID Snowflake, emoji string, userID Snowflake) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Messages_Reactions__, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji), userID.String())
}

/*
GetReactions

Get a list of users that reacted with this emoji.

Returns an array of user objects on success.

The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.

To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.

OPTS SUPPORTS: "after : Snowflake"; "limit : int", nil
*/
func (c *Channel) GetReactions(messageID Snowflake, emoji string, opts *map[string]interface{}) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_Messages_Reactions__, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji), *utilities.ParseQueryString(opts))
}

/*
DeleteAllReactions

Deletes all reactions on a message.

This endpoint requires the 'MANAGE_MESSAGES' permission to be present on the current user.

Fires a Message Reaction Remove All Gateway event.
*/
func (c *Channel) DeleteAllReactions(messageID Snowflake) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Messages_Reactions, api, c.ID.String(), messageID.String())
}

/*
DeleteAllReactionsForEmoji

Deletes all the reactions for a given emoji on a message.

This endpoint requires the MANAGE_MESSAGES permission to be present on the current user.

Fires a Message Reaction Remove Emoji Gateway event.

The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.

To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
*/
func (c *Channel) DeleteAllReactionsForEmoji(messageID Snowflake, emoji string) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Messages_Reactions_, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji))
}

/*
EditMessage

Edit a previously sent message.

The fields content, embeds, and flags can be edited by the original message author.

Other users can only edit flags and only if they have the MANAGE_MESSAGES permission in the corresponding channel.

When specifying flags, ensure to include all previously set flags/bits in addition to ones that you are modifying.

Only flags documented in the table below may be modified by users (unsupported flag changes are currently ignored without error).


When the content field is edited, the mentions array in the message object will be reconstructed from scratch based on the new content.

The allowed_mentions field of the edit request controls how this happens.

If there is no explicit allowed_mentions in the edit request, the content will be parsed with default allowances, that is, without regard to whether or not an allowed_mentions was present in the request that originally created the message.

Returns a message object.

Fires a Message Update Gateway event.
*/
func (c *Channel) EditMessage(messageID string) (method, route string) {
	return http.MethodPatch, fmt.Sprintf(routes.Channels_Messages_, api, c.ID.String(), messageID)
}

type EditMessageJSON struct {
	Content         string          `json:"content"`
	Embeds          []Embed         `json:"embeds"`
	Flags           int             `json:"flags"`
	AllowedMentions AllowedMentions `json:"allowed_mentions"`
	Components      []Component     `json:"components"`
	PayloadJson     string          `json:"payload_json"`
	Attachments     []Attachment    `json:"attachments"`
}

/*
DeleteMessage

Delete a message.

If operating on a guild channel and trying to delete a message that was not sent by the current user, this endpoint requires the MANAGE_MESSAGES permission.

Returns a 204 empty response on success.

Fires a Message Delete Gateway event.

This endpoint supports the "X-Audit-Log-Reason" header.
*/
func (c *Channel) DeleteMessage(messageID string) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Messages_, api, c.ID.String(), messageID)
}

/*
BulkDeleteMessages

Delete multiple messages in a single request.

This endpoint can only be used on guild channels and requires the MANAGE_MESSAGES permission.

Returns a 204 empty response on success.

Fires a Message Delete Bulk Gateway event.

Any message IDs given that do not exist or are invalid will count towards the minimum and maximum message count (currently 2 and 100 respectively).

This endpoint will not delete messages older than 2 weeks, and will fail with a 400 BAD REQUEST if any message provided is older than that or if any duplicate message IDs are provided.

This endpoint supports the "X-Audit-Log-Reason" header.
*/
func (c *Channel) BulkDeleteMessages() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_MessagesBulkDelete, api, c.ID.String())
}

type BulkDeleteJSON struct {
	Messages []Snowflake `json:"messages"`
}

/*
EditChannelPermissions

Edit the channel permission overwrites for a user or role in a channel.

Only usable for guild channels.

Requires the MANAGE_ROLES permission.

Only permissions your bot has in the guild or channel can be allowed/denied (unless your bot has a MANAGE_ROLES overwrite in the channel).

Returns a 204 empty response on success.

For more information about permissions, see permissions.

This endpoint supports the "X-Audit-Log-Reason" header.
*/
func (c *Channel) EditChannelPermissions(overwriteID Snowflake) (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_Permissions_, api, c.ID.String(), overwriteID.String())
}

type EditChannelPermissionsJSON struct {
	Allow *string        `json:"allow,omitempty"`
	Deny  *string        `json:"deny,omitempty"`
	Type  PermissionType `json:"type"`
}

/*
GetChannelInvites

Returns a list of invite objects (with invite metadata) for the channel.

Only usable for guild channels.

Requires the MANAGE_CHANNELS permission.
*/
func (c *Channel) GetChannelInvites() (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_Invites, api, c.ID.String())
}

/*
CreateChannelInvite

Create a new invite object for the channel.

Only usable for guild channels.

Requires the CREATE_INSTANT_INVITE permission.

All JSON parameters for this route are optional, however the request body is not.

If you are not sending any fields, you still have to send an empty JSON object ({}).

Returns an Invite object. Fires an Invite Create Gateway event.

This endpoint supports the X-Audit-Log-Reason header.
*/
func (c *Channel) CreateChannelInvite() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Invites, api, c.ID.String())
}

type CreateChannelJSON struct {
	MaxAge              uint64           `json:"max_age"`               // duration of invite in seconds before expiry, or 0 for never. between 0 and 604800 (7 days)
	MaxUses             int              `json:"max_uses"`              // max number of uses or 0 for unlimited. between 0 and 100
	Temporary           bool             `json:"temporary"`             // whether this invite only grants temporary membership
	Unique              bool             `json:"unique"`                // if true, don't try to reuse a similar invite (useful for creating many unique one time use invites)
	TargetType          InviteTargetType `json:"target_type"`           // the type of target for this voice channel invite
	TargetUserID        Snowflake        `json:"target_user_id"`        // the id of the user whose stream to display for this invite, required if target_type is 1, the user must be streaming in the channel
	TargetApplicationID Snowflake        `json:"target_application_id"` // the id of the embedded application to open for this invite, required if target_type is 2, the application must have the EMBEDDED flag
}

/*
DeleteChannelPermission

Delete a channel permission overwrite for a user or role in a channel.

Only usable for guild channels.

Requires the MANAGE_ROLES permission.

Returns a 204 empty response on success.

For more information about permissions, see permissions

This endpoint supports the "X-Audit-Log-Reason" header.
*/
func (c *Channel) DeleteChannelPermission(overwriteID Snowflake) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Permissions_, api, c.ID.String(), overwriteID.String())
}

/*
FollowNewsChannel

Follow a News Channel to send messages to a target channel.

Requires the MANAGE_WEBHOOKS permission in the target channel.

Returns a followed channel object.
*/
func (c *Channel) FollowNewsChannel() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Followers, api, c.ID.String())
}

type FollowNewsChannelJSON struct {
	WebhookChannelID Snowflake `json:"webhook_channel_id"`
}

/*
TriggerTypingIndicator

Post a typing indicator for the specified channel.

Generally bots should not implement this route.

However, if a bot is responding to a command and expects the computation to take a few seconds, this endpoint may be called to let the user know that the bot is processing their message.

Returns a 204 empty response on success.

Fires a Typing Start Gateway event.
*/
func (c *Channel) TriggerTypingIndicator() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Typing, api, c.ID.String())
}

/*
GetPinnedMessages

Returns all pinned messages in the channel as an array of message objects.
*/
func (c *Channel) GetPinnedMessages() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Pins, api, c.ID.String())
}

/*
PinMessage

Pin a message in a channel.

Requires the MANAGE_MESSAGES permission.

Returns a 204 empty response on success.

    The max pinned messages is 50.

    This endpoint supports the X-Audit-Log-Reason header.
*/
func (c *Channel) PinMessage(messageID Snowflake) (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_Pins_, api, c.ID.String(), messageID.String())
}

/*
UnpinMessage

Unpin a message in a channel.

Requires the MANAGE_MESSAGES permission.

Returns a 204 empty response on success.

    This endpoint supports the X-Audit-Log-Reason header.
*/
func (c *Channel) UnpinMessage(messageID Snowflake) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Pins_, api, c.ID.String(), messageID.String())
}

/*
GroupDmAddRecipient

Adds a recipient to a Group DM using their access token.

REQUIRES: gdm.join SCOPE
*/
func (c *Channel) GroupDmAddRecipient(userID Snowflake) (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_Recipients_, api, c.ID.String(), userID.String())
}

type GroupDmAddRecipientJSON struct {
	AccessToken string `json:"access_token"` // access token of a user that has granted your app the gdm.join scope
	Nick        string `json:"nick"`         // nickname of the user being added
}

/*
GroupDmRemoveRecipient

Removes a recipient from a Group DM.
*/
func (c *Channel) GroupDmRemoveRecipient(userID Snowflake) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_Recipients_, api, c.ID.String(), userID.String())
}

/*
StartThreadWithMessage

Creates a new thread from an existing message.

Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.

Fires a Thread Create Gateway event.

When called on a GuildText channel, creates a GuildPublicThread.

When called on a GuildNews channel, creates a GuildNewsThread.

The id of the created thread will be the same as the id of the message, and as such a message can only have a single thread created from it.

    This endpoint supports the X-Audit-Log-Reason header.
*/
func (c *Channel) StartThreadWithMessage(messageID Snowflake) (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Messages_Threads, api, c.ID.String(), messageID.String())
}

type StartThreadWithMessageJSON struct {
	Name                string  `json:"name"`                            // 1-100 character channel name
	AutoArchiveDuration uint64  `json:"auto_archive_duration,omitempty"` // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	RateLimitPerUser    *uint64 `json:"rate_limit_per_user,omitempty"`   // amount of seconds a user has to wait before sending another message (0-21600)
}

/*
StartThreadWithoutMessage

Creates a new thread that is not connected to an existing message.

The created thread defaults to a GUILD_PRIVATE_THREAD *.

Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.

Fires a Thread Create Gateway event.

    This endpoint supports the X-Audit-Log-Reason header.

* Creating a private thread requires the server to be boosted. The guild features will indicate if that is possible for the guild.


*/
func (c *Channel) StartThreadWithoutMessage() (method, route string) {
	return http.MethodPost, fmt.Sprintf(routes.Channels_Threads, api, c.ID.String())
}

type StartThreadWithoutMessageJSON struct {
	Name                string `json:"name"`                  // 1-100 character channel name
	AutoArchiveDuration uint64 `json:"auto_archive_duration"` // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080

	// In API v9, type defaults to GuildPrivateThread in order to match the behavior when thread documentation was first published.
	// In API v10 this will be changed to be a required field, with no default.
	Type             ChannelType `json:"type"`                          // the type of thread to create
	Invitable        bool        `json:"invitable"`                     // whether non-moderators can add other non-moderators to a thread; only available when creating a private thread
	RateLimitPerUser *uint64     `json:"rate_limit_per_user,omitempty"` // amount of seconds a user has to wait before sending another message (0-21600)
}

/*
JoinThread

Adds the current user to a thread.

Also requires the thread is not archived.

Returns a 204 empty response on success.

Fires a Thread Members Update Gateway event.
*/
func (c *Channel) JoinThread() (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_ThreadMembersMe, api, c.ID.String())
}

/*
AddThreadMember

Adds another member to a thread.

Requires the ability to send messages in the thread.

Also requires the thread is not archived.

Returns a 204 empty response if the member is successfully added or was already a member of the thread.

Fires a Thread Members Update Gateway event.
*/
func (c *Channel) AddThreadMember(userID Snowflake) (method, route string) {
	return http.MethodPut, fmt.Sprintf(routes.Channels_ThreadMembers_, api, c.ID.String(), userID.String())
}

/*
LeaveThread

Removes the current user from a thread.

Also requires the thread is not archived.

Returns a 204 empty response on success.

Fires a Thread Members Update Gateway event.
*/
func (c *Channel) LeaveThread() (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_ThreadMembersMe, api, c.ID.String())
}

/*
RemoveThreadMember

Removes another member from a thread.

Requires the MANAGE_THREADS permission, or the creator of the thread if it is a GuildPrivateThread.

Also requires the thread is not archived.

Returns a 204 empty response on success.

Fires a Thread Members Update Gateway event.
*/
func (c *Channel) RemoveThreadMember(userID Snowflake) (method, route string) {
	return http.MethodDelete, fmt.Sprintf(routes.Channels_ThreadMembers_, api, c.ID.String(), userID.String())
}

/*
GetThreadMember

Returns a thread member object for the specified user if they are a member of the thread, returns a 404 response otherwise.
*/
func (c *Channel) GetThreadMember(userID Snowflake) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_ThreadMembers_, api, c.ID.String(), userID.String())
}

/*
ListThreadMembers

Returns array of thread members objects that are members of the thread.

    This endpoint is restricted according to whether the GUILD_MEMBERS Privileged Intent is enabled for your application.
*/
func (c *Channel) ListThreadMembers() (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_ThreadMembers, api, c.ID.String())
}

/*
ListPublicArchivedThreads

Returns archived threads in the channel that are public.

When called on a GUILD_TEXT channel, returns threads of type GUILD_PUBLIC_THREAD.

When called on a GUILD_NEWS channel returns threads of type GUILD_NEWS_THREAD.

Threads are ordered by archive_timestamp, in descending order.

Requires the READ_MESSAGE_HISTORY permission.
*/
func (c *Channel) ListPublicArchivedThreads(opts *map[string]interface{}) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_ThreadsArchivedPublicQsp, api, c.ID.String(), *utilities.ParseQueryString(opts))
}

/*
ListPrivateArchivedThreads

Returns archived threads in the channel that are of type GUILD_PRIVATE_THREAD.

Threads are ordered by archive_timestamp, in descending order.

Requires both the READ_MESSAGE_HISTORY and MANAGE_THREADS permissions.
*/
func (c *Channel) ListPrivateArchivedThreads(opts *map[string]interface{}) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_ThreadsArchivedPrivateQsp, api, c.ID.String(), *utilities.ParseQueryString(opts))
}

/*
ListJoinedPrivateArchivedThreads

Returns archived threads in the channel that are of type GUILD_PRIVATE_THREAD, and the user has joined.

Threads are ordered by their id, in descending order.

Requires the READ_MESSAGE_HISTORY permission.
*/
func (c *Channel) ListJoinedPrivateArchivedThreads(opts *map[string]interface{}) (method, route string) {
	return http.MethodGet, fmt.Sprintf(routes.Channels_UsersMeThreadsArchivedPrivateQsp, api, c.ID.String(), *utilities.ParseQueryString(opts))
}

type ListArchivedThreadsResponse struct {
	Threads []Channel      `json:"threads"`  // the public, archived threads
	Members []ThreadMember `json:"members"`  // a thread member object for each returned thread the current user has joined
	HasMore bool           `json:"has_more"` // whether there are potentially additional threads that could be returned on a subsequent call
}
