package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

/* CHANNEL OBJECT */

type Channel struct {
	ID                         string         `json:"id"`
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
	_
	_
	_
	GuildNewsThread
	GuildPublicThread
	GuildPrivateThread
	GuildStageVoice
)

type VideoQualityMode int

const (
	Auto VideoQualityMode = iota + 1
	Full
)

/* MESSAGE OBJECT */

type Message struct {
	ID                Snowflake              `json:"id,omitempty"`
	ChannelID         Snowflake              `json:"channel_id,omitempty"`
	GuildID           Snowflake              `json:"guild_id,omitempty"`
	Author            User                   `json:"author,omitempty"`
	Member            GuildMember            `json:"member,omitempty"`
	Content           string                 `json:"content,omitempty"`
	Timestamp         time.Time              `json:"timestamp,omitempty"`
	EditedTimestamp   *time.Time             `json:"edited_timestamp,omitempty"`
	TTS               bool                   `json:"tts,omitempty"`
	MentionEveryone   bool                   `json:"mention_everyone,omitempty"`
	Mentions          []User                 `json:"mentions,omitempty"`
	MentionRoles      []Snowflake            `json:"mention_roles,omitempty"`
	MentionChannels   []Channel              `json:"mention_channels,omitempty"`
	Attachments       []Attachment           `json:"attachments,omitempty"`
	Embeds            []Embed                `json:"embeds,omitempty"`
	Reactions         []ReactionObject       `json:"reactions,omitempty"`
	Nonce             interface{}            `json:"nonce,string,omitempty"`
	Pinned            bool                   `json:"pinned,omitempty"`
	WebhookID         Snowflake              `json:"webhook_id,omitempty"`
	Type              MessageType            `json:"type,omitempty"`
	ApplicationID     Snowflake              `json:"application_id,omitempty"`
	MessageReference  MessageReferenceObject `json:"message_reference,omitempty"`
	Flags             MessageFlags           `json:"flags,omitempty"`
	ReferencedMessage *Message               `json:"referenced_message,omitempty"`
	Interaction       MessageInteraction     `json:"interaction,omitempty"`
	Thread            Channel                `json:"thread,omitempty"`
	Components        []Component            `json:"components,omitempty"`
}

type MessageType int

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
	Reply                // only in API v8+
	ChatInputCommand     // only in API v8+
	ThreadStarterMessage // only in API v9+
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

type MessageReferenceObject struct {
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
	title       = 256
	description = 4096
	fieldCount  = 25
	fieldName   = 256
	fieldValue  = 1024
	footerText  = 2048
	authorName  = 256
)

/* API endpoints */

// GetChannel Get a channel by ID. Returns a channel object. If the channel is a thread, a thread member object is included in the returned result.
func (c *Channel) GetChannel() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/channels/%s", api, c.ID)
}

// ModifyChannel Update a channel's settings. Returns a channel on success, and a 400 BAD REQUEST on invalid parameters. All JSON parameters are optional.
// This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) ModifyChannel() (method string, route string) {
	return http.MethodPatch, fmt.Sprintf("%s/channels/%s", api, c.ID)
}

// ModifyChannelGroupDMJSON Fires a Channel Update Gateway event.
type ModifyChannelGroupDMJSON struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type ModifyChannelGuildChannelJSON struct {
	// All
	Name                 string       `json:"name"`
	Position             *int         `json:"position,omitempty"`
	PermissionOverwrites *[]Overwrite `json:"permission_overwrites,omitempty"`

	// Text
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`

	// Text, News
	Type                       int     `json:"type,omitempty"`
	Topic                      *string `json:"topic,omitempty"`
	DefaultAutoArchiveDuration *int    `json:"default_auto_archive_duration,omitempty"`

	// Text, News, Store
	Nsfw *bool `json:"nsfw,omitempty"`

	// Text, News, Store, Voice
	ParentID *Snowflake `json:"parent_id,omitempty"`

	// Voice
	Bitrate          *int             `json:"bitrate,omitempty"`
	UserLimit        *int             `json:"user_limit,omitempty"`
	RtcRegion        *string          `json:"rtc_region,omitempty"`
	VideoQualityMode VideoQualityMode `json:"video_quality_mode,omitempty"`
}

type ModifyChannelThreadJSON struct {
	Name                string `json:"name"`
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	Locked              bool   `json:"locked"`
	Invitable           bool   `json:"invitable"`
	RateLimitPerUser    *int   `json:"rate_limit_per_user,omitempty"`
}

// DeleteChannel
// Delete a channel, or close a private message.
// Requires the MANAGE_CHANNELS permission for the guild, or MANAGE_THREADS if the channel is a thread.
// Deleting a category does not delete its child channels; they will have their parent_id removed and a Channel Update Gateway event will fire for each of them.
// Returns a channel object on success. Fires a Channel Delete Gateway event (or Thread Delete if the channel was a thread).
// Deleting a guild channel cannot be undone. Use this with caution, as it is impossible to undo this action when performed on a guild channel. In contrast, when used with a private message, it is possible to undo the action by opening a private message with the recipient again.
// For Community guilds, the Rules or Guidelines channel and the Community Updates channel cannot be deleted.
// This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) DeleteChannel() (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/channels/%s", api, c.ID)
}

// GetChannelMessages
// Returns the messages for a channel.
// If operating on a guild channel, this endpoint requires the VIEW_CHANNEL permission to be present on the current user.
// If the current user is missing the 'READ_MESSAGE_HISTORY' permission in the channel then this will return no messages (since they cannot read the message history).
// Returns an array of message objects on success.
func (c *Channel) GetChannelMessages() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/channels/%s/messages", api, c.ID)
}

// GetChannelMessage
// Returns a specific message in the channel.
// If operating on a guild channel, this endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present on the current user.
// Returns a message object on success.
func (c *Channel) GetChannelMessage(messageID string) (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/channels/%s/messages/%s", api, c.ID, messageID)
}

// CreateMessage
// Post a message to a guild text or DM channel. Returns a message object.
// Fires a Message Create Gateway event.
// See message formatting for more information on how to properly format messages.
//
// Limitations
//    * When operating on a guild channel, the current user must have the SEND_MESSAGES permission.
//    * When sending a message with tts (text-to-speech) set to true, the current user must have the SEND_TTS_MESSAGES permission.
//    * When creating a message as a reply to another message, the current user must have the READ_MESSAGE_HISTORY permission.
//        * The referenced message must exist and cannot be a system message.
//    * The maximum request size when sending a message is 8 MB
//    * For the embed object, you can set every field except type (it will be rich regardless of if you try to set it), provider, video, and any height, width, or proxy_url values for images.
//    * Files can only be uploaded when using the multipart/form-data content type.
//
// You may create a message as a reply to another message.
// To do so, include a message_reference with a message_id.
// The channel_id and guild_id in the message_reference are optional, but will be validated if provided.
//
// Note that when sending a message, you must provide a value for at least one of content, embeds, or file.
//
// For a file attachment, the Content-Disposition subpart header MUST contain a filename parameter.
//
// This endpoint supports both application/json and multipart/form-data bodies.
// When uploading files the multipart/form-data content type must be used.
// Note that in multipart form data, the embeds and allowed_mentions fields cannot be used.
// You can pass a stringified JSON body as a form value as payload_json instead.
// If you supply a payload_json form value, all fields except for file fields will be ignored in the form data.
func (c *Channel) CreateMessage() (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/channels/%s/messages", api, c.ID)
}

type CreateMessageJSON struct {
	Content          *string                 `json:"content,omitempty"`
	TTS              *bool                   `json:"tts,omitempty"`
	Embeds           []*Embed                `json:"embeds,omitempty"`
	AllowedMentions  *AllowedMentions        `json:"allowed_mentions,omitempty"`
	MessageReference *MessageReferenceObject `json:"message_reference,omitempty"`
	Components       []*Component            `json:"components,omitempty"`
	StickerIDs       []*Snowflake            `json:"sticker_ids,omitempty"`
	PayloadJson      *string                 `json:"payload_json,omitempty"`
	Attachments      []Attachment            `json:"attachments"`
}

// CrosspostMessage
// Crosspost a message in a News Channel to following channels.
// This endpoint requires the 'SEND_MESSAGES' permission, if the current user sent the message, or additionally the 'MANAGE_MESSAGES' permission, for all other messages, to be present for the current user.
//
// Returns a message object.
func (c *Channel) CrosspostMessage(messageID string) (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/channels/%s/messages/%s/crosspost", api, c.ID, messageID)
}

// CreateReaction
// Create a reaction for the message.
// This endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present on the current user.
// Additionally, if nobody else has reacted to the message using this emoji, this endpoint requires the 'ADD_REACTIONS' permission to be present on the current user.
// Returns a 204 empty response on success.
// The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) CreateReaction(messageID string, emoji string) (method string, route string) {
	return http.MethodPut, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", api, c.ID, messageID, url.QueryEscape(emoji))
}

// DeleteOwnReaction
// Delete a reaction the current user has made for the message.
// Returns a 204 empty response on success.
// The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) DeleteOwnReaction(messageID string, emoji string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/@me", api, c.ID, messageID, url.QueryEscape(emoji))
}

// DeleteUserReaction
// Deletes another user's reaction.
// This endpoint requires the 'MANAGE_MESSAGES' permission to be present on the current user.
// Returns a 204 empty response on success. The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) DeleteUserReaction(messageID string, emoji string, userID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions/%s/%s", api, c.ID, messageID, url.QueryEscape(emoji), userID)
}

// TODO: GetReactions endpoint omitted for now as we won't immediately need it

// DeleteAllReactions
// Deletes all reactions on a message.
// This endpoint requires the 'MANAGE_MESSAGES' permission to be present on the current user.
// Fires a Message Reaction Remove All Gateway event.
func (c *Channel) DeleteAllReactions(messageID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s/reactions", api, c.ID, messageID)
}

// TODO: DeleteAllReactionsForEmoji endpoint omitted for now as we won't immediately need it

// EditMessage
// Edit a previously sent message.
// The fields content, embeds, and flags can be edited by the original message author.
// Other users can only edit flags and only if they have the MANAGE_MESSAGES permission in the corresponding channel.
// When specifying flags, ensure to include all previously set flags/bits in addition to ones that you are modifying.
// Only flags documented in the table below may be modified by users (unsupported flag changes are currently ignored without error).
//
// When the content field is edited, the mentions array in the message object will be reconstructed from scratch based on the new content.
// The allowed_mentions field of the edit request controls how this happens.
// If there is no explicit allowed_mentions in the edit request, the content will be parsed with default allowances, that is, without regard to whether or not an allowed_mentions was present in the request that originally created the message.
//
// Returns a message object.
// Fires a Message Update Gateway event.
func (c *Channel) EditMessage(messageID string) (method string, route string) {
	return http.MethodPatch, fmt.Sprintf("%s/channels/%s/messages/%s", api, c.ID, messageID)
}

type EditMessageJSON struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	Flags           *int             `json:"flags,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []*Component     `json:"components,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
	PayloadJson     *string          `json:"payload_json,omitempty"`
}

// DeleteMessage
// Delete a message.
// If operating on a guild channel and trying to delete a message that was not sent by the current user, this endpoint requires the MANAGE_MESSAGES permission.
// Returns a 204 empty response on success.
// Fires a Message Delete Gateway event.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) DeleteMessage(messageID string) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/channels/%s/messages/%s", api, c.ID, messageID)
}

// TODO: BulkDeleteMessages endpoint omitted for the moment

// TODO: Not implementing any Thread endpoints right now because we just don't need them
