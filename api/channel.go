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
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Channel - Represents a guild or DM channel within Discord.
type Channel struct {
	ID                         Snowflake      `json:"id"`                                      // the id of this channel
	Type                       ChannelType    `json:"type"`                                    // the ChannelType
	GuildID                    Snowflake      `json:"guild_id,omitempty"`                      // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	Position                   int            `json:"position,omitempty"`                      // sorting position of the channel
	PermissionOverwrites       []Overwrite    `json:"permission_overwrites,omitempty"`         // explicit permission overwrites for members and roles
	Name                       string         `json:"name,omitempty"`                          // the name of the channel (1-100 characters)
	Topic                      *string        `json:"topic,omitempty"`                         // the channel topic (0-1024 characters)
	Nsfw                       bool           `json:"nsfw,omitempty"`                          // whether the channel is nsfw
	LastMessageID              *Snowflake     `json:"last_message_id,omitempty"`               // the id of the last message sent in this channel (may not point to an existing or valid message)
	Bitrate                    int64          `json:"bitrate,omitempty"`                       // the bitrate (in bits) of the voice channel
	UserLimit                  int64          `json:"user_limit,omitempty"`                    // the user limit of the voice channel
	RateLimitPerUser           int64          `json:"rate_limit_per_user,omitempty"`           // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Recipients                 []User         `json:"recipients,omitempty"`                    // the recipients of the DM
	Icon                       *string        `json:"icon,omitempty"`                          // icon hash of the group DM
	OwnerID                    Snowflake      `json:"owner_id,omitempty"`                      // id of the creator of the group DM or thread
	ApplicationID              Snowflake      `json:"application_id,omitempty"`                // application id of the group DM creator if it is bot-created
	ParentID                   *Snowflake     `json:"parent_id,omitempty"`                     // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	LastPinTimestamp           *time.Time     `json:"last_pin_timestamp,omitempty"`            // when the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned.
	RtcRegion                  *string        `json:"rtc_region,omitempty"`                    // voice region id for the voice channel, automatic when set to null
	VideoQualityMode           int64          `json:"video_quality_mode,omitempty"`            // the camera video quality mode of the voice channel, 1 when not present
	MessageCount               int64          `json:"message_count,omitempty"`                 // an approximate count of messages in a thread, stops counting at 50
	MemberCount                int64          `json:"member_count,omitempty"`                  // an approximate count of users in a thread, stops counting at 50
	ThreadMetadata             ThreadMetadata `json:"thread_metadata,omitempty"`               // thread-specific fields not needed by other channels
	Member                     ThreadMember   `json:"member,omitempty"`                        // ThreadMember for the current User, if they have joined the thread, only included on certain API endpoints
	DefaultAutoArchiveDuration int            `json:"default_auto_archive_duration,omitempty"` // default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	Permissions                string         `json:"permissions"`                             // computed permissions for the invoking user in the channel, including overwrites, only included when part of the resolved data received on a slash command interaction
	Flags                      ChannelFlag    `json:"flags,omitempty"`                         // channel flags combined as a bitfield
}

// ChannelType - the type of channel
type ChannelType int

//goland:noinspection SpellCheckingInspection,GoUnusedConst
const (
	GuildText          ChannelType = iota     // a text channel within a server
	DM                                        // a direct message between users
	GuildVoice                                // a voice channel within a server
	GroupDM                                   // a direct message between multiple users
	GuildCategory                             // an organizational category that contains up to 50 channels
	GuildNews                                 // a channel that users can follow and crosspost into their own server
	GuildStore                                // a channel in which game developers can sell their game on Discord
	GuildNewsThread    ChannelType = iota + 3 // a temporary sub-channel within a GuildNews channel
	GuildPublicThread                         // a temporary sub-channel within a GuildText channel
	GuildPrivateThread                        // a temporary sub-channel within a GuildText channel that is only viewable by those invited and those with the ManageThreads permission
	GuildStageVoice                           // a voice channel for hosting events with an audience
	GuildDirectory                            // the channel in a hub containing the listed servers
	GuildForum                                // (still in development) a channel that can only contain threads
)

// VideoQualityMode - the camera video quality mode of the voice channel, 1 when not present
type VideoQualityMode int

//goland:noinspection GoUnusedConst
const (
	Auto VideoQualityMode = iota + 1 // Discord chooses the quality for optimal performance
	Full                             // 720p
)

type ChannelFlag uint64

//goland:noinspection GoUnusedConst
const (
	Pinned ChannelFlag = 1 << 1
)

// Message - Represents a message sent in a channel within Discord.
//goland:noinspection SpellCheckingInspection
type Message struct {
	ID                Snowflake          `json:"id,omitempty"`                 // id of the message
	ChannelID         Snowflake          `json:"channel_id,omitempty"`         // id of the Channel the message was sent in
	GuildID           Snowflake          `json:"guild_id,omitempty"`           // id of the Guild the message was sent in
	Author            User               `json:"author,omitempty"`             // the author of this message (not guaranteed to be a valid user)
	Member            GuildMember        `json:"member,omitempty"`             // member properties for this message's author
	Content           string             `json:"content,omitempty"`            // contents of the message
	Timestamp         time.Time          `json:"timestamp,omitempty"`          // when this message was sent
	EditedTimestamp   *time.Time         `json:"edited_timestamp,omitempty"`   // when this message was edited (or null if never)
	TTS               bool               `json:"tts,omitempty"`                // whether this was a TTS message
	MentionEveryone   bool               `json:"mention_everyone,omitempty"`   // whether this message mentions everyone
	Mentions          []User             `json:"mentions,omitempty"`           // users specifically mentioned in the message
	MentionRoles      []Snowflake        `json:"mention_roles,omitempty"`      // roles specifically mentioned in this message
	MentionChannels   []Channel          `json:"mention_channels,omitempty"`   // channels specifically mentioned in this message
	Attachments       []Attachment       `json:"attachments,omitempty"`        // any attached files
	Embeds            []Embed            `json:"embeds,omitempty"`             // any embedded content
	Reactions         []Reaction         `json:"reactions,omitempty"`          // reactions to the message
	Nonce             any                `json:"nonce,omitempty"`              // used for validating a message was sent
	Pinned            bool               `json:"pinned,omitempty"`             // whether this message is pinned
	WebhookID         Snowflake          `json:"webhook_id,omitempty"`         // if the message is generated by a Webhook, this is the webhook's id
	Type              MessageType        `json:"type,omitempty"`               // the MessageType
	Activity          MessageActivity    `json:"activity,omitempty"`           // sent with Rich Presence-related chat embeds
	Application       Application        `json:"application,omitempty"`        // sent with Rich Presence-related chat embeds
	ApplicationID     Snowflake          `json:"application_id,omitempty"`     // if the message is an Interaction or application-owned webhook, this is the id of the application
	MessageReference  MessageReference   `json:"message_reference,omitempty"`  // data showing the source of a crosspost, channel follow add, pin, or reply message
	Flags             MessageFlags       `json:"flags,omitempty"`              // MessageFlags combined as a bitfield
	ReferencedMessage *Message           `json:"referenced_message,omitempty"` // the message associated with the MessageReference
	Interaction       MessageInteraction `json:"interaction,omitempty"`        // sent if the message is a response to an Interaction
	Thread            Channel            `json:"thread,omitempty"`             // the thread that was started from this message, includes ThreadMember object
	Components        []Component        `json:"components,omitempty"`         // sent if the message contains components like buttons, action rows, or other interactive components
	StickerItems      []string           `json:"sticker_items,omitempty"`      // sent if the message contains stickers
	Stickers          []string           `json:"stickers,omitempty"`           // Deprecated: the stickers sent with the message
}

// MessageType - type of message
type MessageType int

//goland:noinspection GoUnusedConst,SpellCheckingInspection
const (
	Default                                 MessageType = iota     // DEFAULT
	RecipientAdd                                                   // RECIPIENT_ADD
	RecipientRemove                                                // RECIPIENT_REMOVE
	Call                                                           // CALL
	ChannelNameChange                                              // CHANNEL_NAME_CHANGE
	ChannelIconChange                                              // CHANNEL_ICON_CHANGE
	ChannelPinnedMessage                                           // CHANNEL_PINNED_MESSAGE
	GuildMemberJoin                                                // GUILD_MEMBER_JOIN
	UserPremiumGuildSubscription                                   // USER_PREMIUM_GUILD_SUBSCRIPTION
	UserPremiumGuildSubscriptionTier1                              // USER_PREMIUM_GUILD_SUBSCRIPTION_TIER_1
	UserPremiumGuildSubscriptionTier2                              // USER_PREMIUM_GUILD_SUBSCRIPTION_TIER_2
	UserPremiumGuildSubscriptionTier3                              // USER_PREMIUM_GUILD_SUBSCRIPTION_TIER_3
	ChannelFollowAdd                                               // CHANNEL_FOLLOW_ADD
	GuildDiscoveryDisqualified              MessageType = iota + 1 // GUILD_DISCOVERY_DISQUALIFIED
	GuildDiscoveryRequalified                                      // GUILD_DISCOVERY_REQUALIFIED
	GuildDiscoveryGracePeriodInitialWarning                        // GUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING
	GuildDiscoveryGracePeriodFinalWarning                          // GUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING
	ThreadCreated                                                  // THREAD_CREATED
	Reply                                                          // REPLY
	ChatInputCommand                                               // CHAT_INPUT_COMMAND
	ThreadStarterMessage                                           // THREAD_STARTER_MESSAGE
	GuildInviteReminder                                            // GUILD_INVITE_REMINDER
	ContextMenuCommand                                             // CONTEXT_MENU_COMMAND
)

// MessageActivity - sent with Rich Presence-related chat embeds
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`     // type of message activity
	PartyID string              `json:"party_id"` // party_id from a Rich Presence event
}

// MessageActivityType - type of message activity
type MessageActivityType int

//goland:noinspection GoUnusedConst
const (
	MessageActivityTypeJoin        MessageActivityType = iota + 1 // JOIN
	MessageActivityTypeSpectate                                   // SPECTATE
	MessageActivityTypeListen                                     // LISTEN
	MessageActivityTypeJoinRequest MessageActivityType = iota + 2 // JOIN_REQUEST
)

// MessageFlags - MessageFlags combined as a bitfield
type MessageFlags int

//goland:noinspection GoUnusedConst,SpellCheckingInspection
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

// MessageReference - ChannelID is optional when creating a reply, but will always be present when receiving an event/response that includes this data model.
type MessageReference struct {
	MessageID       Snowflake `json:"message_id,omitempty"`         // id of the originating message
	ChannelID       Snowflake `json:"channel_id,omitempty"`         // id of the originating message's channel
	GuildID         Snowflake `json:"guild_id,omitempty"`           // id of the originating message's guild
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"` // when sending, whether to error if the referenced message doesn't exist instead of sending as a normal (non-reply) message, default true
}

// FollowedChannel - representation of a followed News Channel
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id"` // source Channel id
	WebhookID Snowflake `json:"webhook_id"` // created target Webhook id
}

// Reaction - representation of a message reaction
type Reaction struct {
	Count int   `json:"count"` // times this Emoji has been used to react
	Me    bool  `json:"me"`    // whether the current User reacted using this Emoji
	Emoji Emoji `json:"emoji"` // Emoji information
}

// Overwrite - representation of a permissions overwrite
type Overwrite struct {
	ID    Snowflake      `json:"id"`    // Role or User id
	Type  PermissionType `json:"type"`  // either PermissionRole or PermissionMember
	Allow string         `json:"allow"` // permission bit set
	Deny  string         `json:"deny"`  // permission bit set
}

// PermissionType - either PermissionRole or PermissionMember
type PermissionType int

//goland:noinspection GoUnusedConst
const (
	PermissionRole   PermissionType = iota // 0 (role)
	PermissionMember                       // 1 (member)
)

// ThreadMetadata - The thread metadata object contains a number of thread-specific channel fields that are not needed by other channel types.
//goland:noinspection SpellCheckingInspection
type ThreadMetadata struct {
	Archived            bool       `json:"archived"`                   // whether the thread is archived
	AutoArchiveDuration int        `json:"auto_archive_duration"`      // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	ArchiveTimestamp    time.Time  `json:"archive_timestamp"`          // timestamp when the thread's archive status was last changed, used for calculating recent activity
	Locked              bool       `json:"locked"`                     // whether the thread is locked; when a thread is locked, only users with ManageThreads can unarchive it
	Invitable           bool       `json:"invitable,omitempty"`        // whether non-moderators can add other non-moderators to a thread; only available on private threads
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"` // timestamp when the thread was created; only populated for threads created after 2022-01-09
}

// ThreadMember - A thread member is used to indicate whether a user has joined a thread or not.
type ThreadMember struct {
	ID            Snowflake `json:"id,omitempty"`      // the id of the thread
	UserID        Snowflake `json:"user_id,omitempty"` // the id of the User
	JoinTimestamp time.Time `json:"join_timestamp"`    // the time the current user last joined the thread
	Flags         int64     `json:"flags"`             // any user-thread settings, currently only used for notifications
}

// Embed - contains rich content
type Embed struct {
	Title       string     `json:"title,omitempty"`       // title of embed
	Type        EmbedType  `json:"type,omitempty"`        // EmbedType (always RichEmbed for webhook embeds)
	Description string     `json:"description,omitempty"` // description of embed
	URL         string     `json:"url,omitempty"`         // url of embed
	Timestamp   string     `json:"timestamp,omitempty"`   // timestamp of embed content
	Color       int64      `json:"color,omitempty"`       // color code of the embed
	Footer      *Footer    `json:"footer,omitempty"`      // footer information
	Image       *Image     `json:"image,omitempty"`       // image information
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`   // thumbnail information
	Video       *video     `json:"video,omitempty"`       // video information; cannot set this when sending an Embed
	Provider    *provider  `json:"provider,omitempty"`    // provider information; cannot set this when sending an Embed
	Author      *Author    `json:"author,omitempty"`      // author information
	Fields      []*Field   `json:"fields,omitempty"`      // fields information
}

// EmbedType - Embed types are "loosely defined" and, for the most part, are not used by our clients for rendering.
//
// Embed attributes power what is rendered.
//
// Embed types should be considered deprecated and might be removed in a future API version.
type EmbedType string

//goland:noinspection GoUnusedConst,SpellCheckingInspection
const (
	RichEmbed    EmbedType = "rich"    // generic embed rendered from embed attributes
	imageEmbed   EmbedType = "image"   // image embed
	videoEmbed   EmbedType = "video"   // video embed
	gifVEmbed    EmbedType = "gifv"    // animated gif image embed rendered as a video embed
	articleEmbed EmbedType = "article" // article embed
	linkEmbed    EmbedType = "link"    // link embed
)

// Thumbnail - thumbnail information
type Thumbnail struct {
	URL    string `json:"url,omitempty"`    // source url of thumbnail (only supports http(s) and attachments)
	Height int    `json:"height,omitempty"` // height of thumbnail
	Width  int    `json:"width,omitempty"`  // width of thumbnail
}

// video - video information
type video struct {
	URL    string `json:"url,omitempty"`    // source url of video
	Height int    `json:"height,omitempty"` // height of video
	Width  int    `json:"width,omitempty"`  // width of video
}

// Image - image information
type Image struct {
	URL    string `json:"url,omitempty"`    // source url of image (only supports http(s) and attachments)
	Height int    `json:"height,omitempty"` // height of image
	Width  int    `json:"width,omitempty"`  // width of image
}

// provider - provider information
type provider struct {
	Name string `json:"name,omitempty"` // name of provider
	URL  string `json:"url,omitempty"`  // url of provider
}

// Author - author information
type Author struct {
	Name    string  `json:"name,omitempty"`     // name of author
	URL     string  `json:"url,omitempty"`      // url of author
	IconURL *string `json:"icon_url,omitempty"` // url of author icon (only supports http(s) and attachments)
}

// Footer - footer information
type Footer struct {
	Text    string `json:"text"`               // footer text
	IconURL string `json:"icon_url,omitempty"` // url of footer icon (only supports http(s) and attachments)
}

// Field - fields information
type Field struct {
	Name   string `json:"name"`   // name of the field
	Value  string `json:"value"`  // value of the field
	Inline bool   `json:"inline"` // whether this field should display inline
}

// Attachment - For the attachments array in Message Create/Edit requests, only the id is required.
//goland:noinspection SpellCheckingInspection
type Attachment struct {
	ID          Snowflake `json:"id"`                     // attachment id
	Filename    string    `json:"filename"`               // name of file attached
	Description string    `json:"description,omitempty"`  // description for the file
	ContentType string    `json:"content_type,omitempty"` // the attachment's media type
	Size        uint64    `json:"size"`                   // size of file in bytes
	URL         string    `json:"url"`                    // source url of file
	ProxyURL    string    `json:"proxy_url"`              // a proxied url of file
	Height      *int64    `json:"height,omitempty"`       // height of file (if image)
	Width       *int64    `json:"width,omitempty"`        // width of file (if image)
	Ephemeral   bool      `json:"ephemeral,omitempty"`    // whether this attachment is ephemeral
}

// ChannelMention - representation of a Channel mention
type ChannelMention struct {
	ID      Snowflake   `json:"id"`       // id of the channel
	GuildID Snowflake   `json:"guild_id"` // id of the guild containing the channel
	Type    ChannelType `json:"type"`     // the ChannelType
	Name    string      `json:"name"`     // the name of the channel
}

// AllowedMentions - The allowed mention field allows for more granular control over mentions without various hacks to the message content.
//
// This will always validate against message content to avoid phantom pings (e.g. to ping everyone, you must still have @everyone in the message content), and check against user/bot permissions.
type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`                  // An array of AllowedMentionType to parse from the content.
	Roles       []Snowflake          `json:"roles,omitempty"`        // Array of role_ids to mention (Max size of 100)
	Users       []Snowflake          `json:"users,omitempty"`        // Array of user_ids to mention (Max size of 100)
	RepliedUser bool                 `json:"replied_user,omitempty"` // For replies, whether to mention the author of the message being replied to (default false)
}

// AllowedMentionType - the type of mention allowed
type AllowedMentionType string

//goland:noinspection GoUnusedConst
const (
	RoleMentions     AllowedMentionType = "roles"    // Controls role mentions
	UserMentions     AllowedMentionType = "users"    // Controls user mentions
	EveryoneMentions AllowedMentionType = "everyone" // Controls @everyone and @here mentions
)

// Additionally, the combined sum of characters in all title, description, field.name, field.value, footer.text, and author.name fields across all embeds attached to a message must not exceed 6000 characters.
//
// Violating any of these constraints will result in a Bad Request response.
const (
	titleLimit       = 256  // 256 characters
	descriptionLimit = 4096 // 4096 characters
	fieldCount       = 25   // Up to 25 field objects
	fieldNameLimit   = 256  // 256 characters
	fieldValueLimit  = 1024 // 1024 characters
	footerTextLimit  = 2048 // 2048 characters
	authorNameLimit  = 256  // 256 characters
)

// GetChannel - Get a channel by ID.
//
// Returns a channel object.
//
// If the channel is a thread, a thread member object is included in the returned result.
func (c *Channel) GetChannel() (*Channel, error) {
	u := parseRoute(fmt.Sprintf(getChannel, api, c.ID.String()))

	var channel Channel
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &channel)

	return &channel, err
}

// ModifyGroupDm - Fires a ChannelUpdate Gateway event.
func (c *Channel) ModifyGroupDm(payload ModifyGroupDmJSON, reason *string) (*Channel, error) {
	return c.modifyChannel(payload, reason)
}

type ModifyGroupDmJSON struct {
	Name string `json:"name"` // 1-100 character channel name
	Icon string `json:"icon"` // base64 encoded icon
}

func (c *Channel) ModifyGuildTextChannel(payload ModifyTextChannelJSON, reason *string) (*Channel, error) {
	return c.modifyChannel(payload, reason)
}

func (c *Channel) ModifyGuildNewsChannel(payload ModifyNewsChannelJSON, reason *string) (*Channel, error) {
	return c.modifyChannel(payload, reason)
}

func (c *Channel) ModifyThread(payload ModifyThreadJSON, reason *string) (*Channel, error) {
	return c.modifyChannel(payload, reason)
}

func (c *Channel) ModifyGuildVoiceChannel(payload ModifyGuildVoiceChannelJSON, reason *string) (*Channel, error) {
	return c.modifyChannel(payload, reason)
}

type ModifyAllChannelJSON struct {
	Name                 string      `json:"name"`                  // 1-100 character channel name
	Position             *int        `json:"position"`              // the position of the channel in the left-hand listing
	PermissionOverwrites []Overwrite `json:"permission_overwrites"` // channel or category-specific permissions
}

type ModifyNewsChannelJSON struct {
	ModifyAllChannelJSON

	Type                       ChannelType `json:"type"`                          // the type of channel; only conversion between text and news is supported and only in guilds with the "NEWS" feature
	Topic                      *string     `json:"topic"`                         // 0-1024 character channel topic
	Nsfw                       *bool       `json:"nsfw"`                          // whether the channel is nsfw
	ParentID                   *Snowflake  `json:"parent_id"`                     // id of the new parent category for a channel
	DefaultAutoArchiveDuration *uint64     `json:"default_auto_archive_duration"` // the default duration that the clients use (not the API) for newly created threads in the channel, in minutes, to automatically archive the thread after recent activity
}

type ModifyTextChannelJSON struct {
	ModifyNewsChannelJSON

	RateLimitPerUser *uint64 `json:"rate_limit_per_user"` // amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission ManageMessages, or ManageChannels, are unaffected
}

type ModifyGuildVoiceChannelJSON struct {
	ModifyAllChannelJSON

	Bitrate          *uint64          `json:"bitrate"`            // the bitrate (in bits) of the voice channel; 8000 to 96000 (128000 for VIP servers)
	UserLimit        *uint            `json:"user_limit"`         // the user limit of the voice channel; 0 refers to no limit, 1 to 99 refers to a user limit
	ParentID         *Snowflake       `json:"parent_id"`          // id of the new parent category for a channel
	RtcRegion        *string          `json:"rtc_region"`         // channel voice region id, automatic when set to null
	VideoQualityMode VideoQualityMode `json:"video_quality_mode"` // the camera video quality mode of the voice channel
}

// modifyChannel - Update a channel's settings. Returns a channel on success, and a 400 BAD REQUEST on invalid parameters. All JSON parameters are optional.
func (c *Channel) modifyChannel(payload any, reason *string) (*Channel, error) {
	// TODO: verify types on payload
	u := parseRoute(fmt.Sprintf(modifyChannel, api, c.ID.String()))

	var channel Channel
	err := json.Unmarshal(fireGetRequest(u, &payload, reason), &channel)

	return &channel, err
}

// ModifyThreadJSON - When setting archived to false, when locked is also false, only the SEND_MESSAGES permission is required.
//
// Otherwise, requires the MANAGE_THREADS permission. Fires a Thread Update Gateway event.
// Requires the thread to have archived set to false or be set to false in the request.
//goland:noinspection SpellCheckingInspection
type ModifyThreadJSON struct {
	Name                string `json:"name"`                  // 1-100 character channel name
	Archived            bool   `json:"archived"`              // whether the thread is archived
	AutoArchiveDuration int    `json:"auto_archive_duration"` // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	Locked              bool   `json:"locked"`                // whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
	Invitable           bool   `json:"invitable"`             // whether non-moderators can add other non-moderators to a thread; only available on private threads
	RateLimitPerUser    *int   `json:"rate_limit_per_user"`   // amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages, manage_thread, or manage_channel, are unaffected
}

// DeleteChannel - Delete a channel, or close a private message.
//
// Requires the ManageChannels permission for the guild, or ManageThreads if the channel is a thread.
//
// Deleting a category does not delete its child channels; they will have their parent_id removed and a ChannelUpdate Gateway event will fire for each of them.
//
// Returns a channel object on success. Fires a ChannelDelete Gateway event (or ThreadDelete if the channel was a thread).
//
//     Deleting a guild channel cannot be undone. Use this with caution, as it is impossible to undo this action when performed on a guild channel. In contrast, when used with a private message, it is possible to undo the action by opening a private message with the recipient again.
//
//     For Community guilds, the Rules or Guidelines channel and the Community Updates channel cannot be deleted.
//
//     This endpoint supports the `X-Audit-Log-Reason` header.
func (c *Channel) DeleteChannel(reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteChannel, api, c.ID.String()))

	return fireDeleteRequest(u, reason)
}

// GetChannelMessages - Returns the messages for a channel.
//
// If operating on a guild channel, this endpoint requires the ViewChannel permission to be present on the current user.
//
// If the current user is missing the ReadMessageHistory permission in the channel then this will return no messages (since they cannot read the message history).
//
// Returns an array of message objects on success.
//
// SUPPORTS: "around : Snowflake"; "before : Snowflake"; "after : Snowflake"; "limit : int" ; nil
//
//      The before, after, and around keys are mutually exclusive, only one may be passed at a time.
//
// TODO: Check permissions; required ViewChannel and ReadMessageHistory
func (c *Channel) GetChannelMessages(around *Snowflake, before *Snowflake, after *Snowflake, limit *int) ([]Message, error) {
	u := parseRoute(fmt.Sprintf(getChannelMessages, api, c.ID.String()))

	q := u.Query()
	if around != nil {
		q.Set("around", around.String())
	}
	if before != nil {
		q.Set("before", before.String())
	}
	if after != nil {
		q.Set("after", after.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var messages []Message
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &messages)

	return messages, err
}

// GetChannelMessage - Returns a specific message in the channel.
//
// If operating on a guild channel, this endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present on the current user.
//
// Returns a message object on success
func (c *Channel) GetChannelMessage(messageID string) (*Message, error) {
	u := parseRoute(fmt.Sprintf(getChannelMessage, api, c.ID.String(), messageID))

	var message *Message
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &message)

	return message, err
}

// CreateMessage - Post a message to a guild text or DM channel. Returns a message object.
//
//      Discord may strip certain characters from message content, like invalid unicode characters or characters which cause unexpected message formatting. If you are passing user-generated strings into message content, consider sanitizing the data to prevent unexpected behavior and utilizing allowed_mentions to prevent unexpected mentions.
//
// Fires a Message Create Gateway event.
//
// See message formatting for more information on how to properly format messages.
//
// Limitations
//   * When operating on a guild channel, the current user must have the SendMessages permission.
//   * When sending a message with tts (text-to-speech) set to true, the current user must have the SendTtsMessages permission.
//   * When creating a message as a reply to another message, the current user must have the ReadMessageHistory permission.
//       * The referenced message must exist and cannot be a system message.
//   * The maximum request size when sending a message is 8 MB
//   * For the embed object, you can set every field except type (it will be rich regardless of if you try to set it), provider, video, and any height, width, or proxy_url values for images.
//   * Files can only be uploaded when using the multipart/form-data content type.
//
// You may create a message as a reply to another message. To do so, include a `message_reference` with a `message_id`. The `channel_id` and `guild_id` in the `message_reference` are optional, but will be validated if provided.
//
//     Note that when sending a message, you must provide a value for at least one of content, embeds, or file.
//
// For a file attachment, the Content-Disposition subpart header MUST contain a filename parameter.
//
// This endpoint supports both application/json and multipart/form-data bodies.
//
// When uploading files the multipart/form-data content type must be used.
//
// Note that in multipart form data, the embeds and allowed_mentions fields cannot be used.
//
// You can pass a stringified JSON body as a form value as payload_json instead.
//
// If you supply a payload_json form value, all fields except for file fields will be ignored in the form data.
func (c *Channel) CreateMessage(payload CreateMessageJSON) (*Message, error) {
	u := parseRoute(fmt.Sprintf(createMessage, api, c.ID.String()))

	var message *Message
	err := json.Unmarshal(firePostRequest(u, payload, nil), &message)

	return message, err
}

// CreateMessageJSON - JSON payload structure
// TODO: files[n]
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

// CrosspostMessage - Crosspost a message in a News Channel to following channels.
//
// This endpoint requires the 'SEND_MESSAGES' permission, if the current user sent the message, or additionally the 'MANAGE_MESSAGES' permission, for all other messages, to be present for the current user.
//
// Returns a message object.
//goland:noinspection SpellCheckingInspection
func (c *Channel) CrosspostMessage(messageID string) (*Message, error) {
	u := parseRoute(fmt.Sprintf(crosspostMessage, api, c.ID.String(), messageID))

	var message *Message
	err := json.Unmarshal(firePostRequest(u, nil, nil), &message)

	return message, err
}

// CreateReaction - Create a reaction for the message.
//
// This endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present on the current user.
//
// Additionally, if nobody else has reacted to the message using this emoji, this endpoint requires the 'ADD_REACTIONS' permission to be present on the current user.
//
// Returns a 204 empty response on success.
//
// The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
//
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) CreateReaction(messageID Snowflake, emoji string) {
	u := parseRoute(fmt.Sprintf(createReaction, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji)))

	_ = firePutRequest(u, nil, nil)
}

// DeleteOwnReaction - Delete a reaction the current user has made for the message.
//
// Returns a 204 empty response on success.
//
// The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
//
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) DeleteOwnReaction(messageID Snowflake, emoji string) error {
	u := parseRoute(fmt.Sprintf(deleteOwnReaction, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji)))

	return fireDeleteRequest(u, nil)
}

// DeleteUserReaction - Deletes another user's reaction.
//
// This endpoint requires the 'MANAGE_MESSAGES' permission to be present on the current user.
//
// Returns a 204 empty response on success. The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
//
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) DeleteUserReaction(messageID Snowflake, emoji string, userID Snowflake) error {
	u := parseRoute(fmt.Sprintf(deleteUserReaction, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji), userID.String()))

	return fireDeleteRequest(u, nil)
}

// GetReactions - Get a list of users that reacted with this emoji.
//
// Returns an array of user objects on success.
//
// The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
//
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
//
// OPTS SUPPORTS: "after : Snowflake"; "limit : int", nil
func (c *Channel) GetReactions(messageID Snowflake, emoji string, after *Snowflake, limit *int) ([]User, error) {
	u := parseRoute(fmt.Sprintf(getReactions, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji)))

	q := u.Query()
	if after != nil {
		q.Set("after", after.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var users []User
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &users)

	return users, err
}

// DeleteAllReactions - Deletes all reactions on a message.
//
// This endpoint requires the 'MANAGE_MESSAGES' permission to be present on the current user.
//
// Fires a Message Reaction Remove All Gateway event.
func (c *Channel) DeleteAllReactions(messageID Snowflake) error {
	u := parseRoute(fmt.Sprintf(deleteAllReactions, api, c.ID.String(), messageID.String()))

	return fireDeleteRequest(u, nil)
}

// DeleteAllReactionsForEmoji - Deletes all the reactions for a given emoji on a message.
//
// This endpoint requires the MANAGE_MESSAGES permission to be present on the current user.
//
// Fires a Message Reaction Remove Emoji Gateway event.
//
// The emoji must be URL Encoded or the request will fail with 10014: Unknown Emoji.
//
// To use custom emoji, you must encode it in the format name:id with the emoji name and emoji id.
func (c *Channel) DeleteAllReactionsForEmoji(messageID Snowflake, emoji string) error {
	u := parseRoute(fmt.Sprintf(deleteAllReactionsForEmoji, api, c.ID.String(), messageID.String(), url.QueryEscape(emoji)))

	return fireDeleteRequest(u, nil)
}

// EditMessage - Edit a previously sent message.
//
// The fields content, embeds, and flags can be edited by the original message author.
// Other users can only edit flags and only if they have the ManageMessages permission in the corresponding channel.
// When specifying flags, ensure to include all previously set flags/bits in addition to ones that you are modifying.
// Only flags documented in the table below may be modified by users (unsupported flag changes are currently ignored without error).
//
// When the content field is edited, the mentions array in the message object will be reconstructed from scratch based on the new content.
// The allowed_mentions field of the edit request controls how this happens.
// If there is no explicit allowed_mentions in the edit request, the content will be parsed with default allowances, that is, without regard to whether or not an allowed_mentions was present in the request that originally created the message.
//
// Returns a message object.
//
// Fires a Message Update Gateway event.
func (c *Channel) EditMessage(messageID string, payload EditMessageJSON) (*Message, error) {
	u := parseRoute(fmt.Sprintf(editMessage, api, c.ID.String(), messageID))

	var message *Message
	err := json.Unmarshal(firePatchRequest(u, payload, nil), &message)

	return message, err
}

// EditMessageJSON - JSON payload structure
//
// All parameters are optional and nullable.
//
// TODO: files[n]
type EditMessageJSON struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	Flags           *int             `json:"flags,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []*Component     `json:"components,omitempty"`
	PayloadJson     *string          `json:"payload_json,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
}

// DeleteMessage - Delete a message.
//
// If operating on a guild channel and trying to delete a message that was not sent by the current user, this endpoint requires the MANAGE_MESSAGES permission.
//
// Returns a 204 empty response on success.
//
// Fires a MessageDelete Gateway event.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (c *Channel) DeleteMessage(messageID string, reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteMessage, api, c.ID.String(), messageID))

	return fireDeleteRequest(u, reason)
}

// BulkDeleteMessages - Delete multiple messages in a single request.
//
// This endpoint can only be used on guild channels and requires the MANAGE_MESSAGES permission.
//
// Returns a 204 empty response on success.
//
// Fires a Message Delete Bulk Gateway event.
//
// Any message IDs given that do not exist or are invalid will count towards the minimum and maximum message count (currently 2 and 100 respectively).
// This endpoint will not delete messages older than 2 weeks, and will fail with a 400 BAD REQUEST if any message provided is older than that or if any duplicate message IDs are provided.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (c *Channel) BulkDeleteMessages(payload BulkDeleteJSON, reason *string) error {
	if len(payload.Messages) < 2 || len(payload.Messages) > 100 {
		return errors.New("you can only bulk delete >= 2 && <= 100 messages at a time")
	}

	for _, message := range payload.Messages {
		if time.Since(time.Unix(message.ParseSnowflake().Timestamp, 0)).Hours() > float64(14*24) {
			return errors.New("cannot bulk delete message older than 2 weeks")
		}
	}
	u := parseRoute(fmt.Sprintf(bulkDeleteMessages, api, c.ID.String()))

	_ = firePostRequest(u, payload, reason)

	return nil
}

// BulkDeleteJSON - JSON payload structure
type BulkDeleteJSON struct {
	Messages []Snowflake `json:"messages"`
}

// EditChannelPermissions - Edit the channel permission overwrites for a user or role in a channel.
//
// Only usable for guild channels.
//
// Requires the ManageRoles permission.
//
// Only permissions your bot has in the guild or channel can be allowed/denied (unless your bot has a ManageRoles overwrite in the channel).
//
// Returns a 204 empty response on success.
//
// For more information about permissions, see permissions.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (c *Channel) EditChannelPermissions(overwriteID Snowflake, payload EditChannelPermissionsJSON, reason *string) {
	u := parseRoute(fmt.Sprintf(editChannelPermissions, api, c.ID.String(), overwriteID.String()))

	_ = firePutRequest(u, payload, reason)
}

// EditChannelPermissionsJSON - JSON payload structure
type EditChannelPermissionsJSON struct {
	Allow *string        `json:"allow,omitempty"` // the bitwise value of all allowed permissions (default "0")
	Deny  *string        `json:"deny,omitempty"`  // the bitwise value of all disallowed permissions (default "0")
	Type  PermissionType `json:"type"`            // 0 for a role or 1 for a member
}

// GetChannelInvites - Returns a list of invite objects (with invite metadata) for the channel.
//
// Only usable for guild channels.
//
// Requires the ManageChannels permission.
func (c *Channel) GetChannelInvites() ([]Invite, error) {
	u := parseRoute(fmt.Sprintf(getChannelInvites, api, c.ID.String()))

	var invites []Invite
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &invites)

	return invites, err
}

// CreateChannelInvite - Create a new invite object for the channel.
//
// Only usable for guild channels.
//
// Requires the CreateInstantInvite permission.
//
// All JSON parameters for this route are optional, however the request body is not.
//
// If you are not sending any fields, you still have to send an empty JSON object ({}).
//
// Returns an Invite object. Fires an Invite Create Gateway event.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) CreateChannelInvite(payload CreateChannelInviteJSON, reason *string) (*Invite, error) {
	u := parseRoute(fmt.Sprintf(getChannelInvites, api, c.ID.String()))

	var invite *Invite
	err := json.Unmarshal(firePostRequest(u, payload, reason), &invite)

	return invite, err
}

// CreateChannelInviteJSON - JSON payload structure
type CreateChannelInviteJSON struct {
	MaxAge              uint64           `json:"max_age"`               // duration of invite in seconds before expiry, or 0 for never. between 0 and 604800 (7 days)
	MaxUses             int              `json:"max_uses"`              // max number of uses or 0 for unlimited. between 0 and 100
	Temporary           bool             `json:"temporary"`             // whether this invite only grants temporary membership
	Unique              bool             `json:"unique"`                // if true, don't try to reuse a similar invite (useful for creating many unique one time use invites)
	TargetType          InviteTargetType `json:"target_type"`           // the type of target for this voice channel invite
	TargetUserID        Snowflake        `json:"target_user_id"`        // the id of the user whose stream to display for this invite, required if target_type is 1, the user must be streaming in the channel
	TargetApplicationID Snowflake        `json:"target_application_id"` // the id of the embedded application to open for this invite, required if target_type is 2, the application must have the Embedded flag
}

// DeleteChannelPermission - Delete a channel permission overwrite for a user or role in a channel.
//
// Only usable for guild channels.
//
// Requires the ManageRoles permission.
//
// Returns a 204 empty response on success.
//
// For more information about permissions, see permissions
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (c *Channel) DeleteChannelPermission(overwriteID Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteChannelPermission, api, c.ID.String(), overwriteID.String()))

	return fireDeleteRequest(u, reason)
}

// FollowNewsChannel - Follow a News Channel to send messages to a target channel.
//
// Requires the ManageWebhooks permission in the target channel.
//
// Returns a followed channel object.
func (c *Channel) FollowNewsChannel(payload FollowNewsChannelJSON) (*FollowedChannel, error) {
	u := parseRoute(fmt.Sprintf(followNewsChannel, api, c.ID.String()))

	var followedChannel *FollowedChannel
	err := json.Unmarshal(firePostRequest(u, payload, nil), &followedChannel)

	return followedChannel, err
}

// FollowNewsChannelJSON - JSON payload structure
type FollowNewsChannelJSON struct {
	WebhookChannelID Snowflake `json:"webhook_channel_id"`
}

// TriggerTypingIndicator - Post a typing indicator for the specified channel.
//
// Generally bots should not implement this route.
// However, if a bot is responding to a command and expects the computation to take a few seconds, this endpoint may be called to let the user know that the bot is processing their message.
//
// Returns a 204 empty response on success.
//
// Fires a Typing Start Gateway event.
func (c *Channel) TriggerTypingIndicator() {
	u := parseRoute(fmt.Sprintf(triggerTypingIndicator, api, c.ID.String()))

	_ = firePostRequest(u, nil, nil)
}

// GetPinnedMessages - Returns all pinned messages in the channel as an array of message objects.
func (c *Channel) GetPinnedMessages() ([]Message, error) {
	u := parseRoute(fmt.Sprintf(getPinnedMessages, api, c.ID.String()))

	var messages []Message
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &messages)

	return messages, err
}

// PinMessage - Pin a message in a channel.
//
// Requires the ManageMessages permission.
//
// Returns a 204 empty response on success.
//
//    The max pinned messages is 50.
//
//    This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) PinMessage(messageID Snowflake, reason *string) {
	u := parseRoute(fmt.Sprintf(pinMessage, api, c.ID.String(), messageID.String()))

	_ = firePutRequest(u, nil, reason)
}

// UnpinMessage - Unpin a message in a channel.
//
// Requires the ManageMessages permission.
//
// Returns a 204 empty response on success.
//
//    This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) UnpinMessage(messageID Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(unpinMessage, api, c.ID.String(), messageID.String()))

	return fireDeleteRequest(u, reason)
}

// GroupDmAddRecipient - Adds a recipient to a Group DM using their access token.
//
// REQUIRES: gdm.join SCOPE
func (c *Channel) GroupDmAddRecipient(userID Snowflake, payload GroupDmAddRecipientJSON) {
	u := parseRoute(fmt.Sprintf(groupDmAddRecipient, api, c.ID.String(), userID.String()))

	_ = firePutRequest(u, payload, nil)
}

// GroupDmAddRecipientJSON - JSON payload structure
//
// IMPORTANT: requires a Bearer token for the user
type GroupDmAddRecipientJSON struct {
	AccessToken string `json:"access_token"` // access token of a user that has granted your app the gdm.join scope
	Nick        string `json:"nick"`         // nickname of the user being added
}

// GroupDmRemoveRecipient - Removes a recipient from a Group DM.
func (c *Channel) GroupDmRemoveRecipient(userID Snowflake) error {
	u := parseRoute(fmt.Sprintf(groupDmRemoveRecipient, api, c.ID.String(), userID.String()))

	return fireDeleteRequest(u, nil)
}

// StartThreadWithMessage - Creates a new thread from an existing message.
//
// Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.
//
// Fires a ThreadCreate Gateway event.
//
// When called on a GuildText channel, creates a GuildPublicThread. When called on a GuildNews channel, creates a GuildNewsThread.
//
// Does not work on a GuildForum channel.
//
// The id of the created thread will be the same as the id of the source message, and as such a message can only have a single thread created from it.
//
//    This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) StartThreadWithMessage(messageID Snowflake, payload StartThreadWithMessageJSON, reason *string) (*Channel, error) {
	u := parseRoute(fmt.Sprintf(startThreadWithMessage, api, c.ID.String(), messageID.String()))

	var channel Channel
	err := json.Unmarshal(firePostRequest(u, payload, reason), &channel)

	return &channel, err
}

// StartThreadWithMessageJSON - JSON payload structure
type StartThreadWithMessageJSON struct {
	Name                string  `json:"name"`                            // 1-100 character channel name
	AutoArchiveDuration uint64  `json:"auto_archive_duration,omitempty"` // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	RateLimitPerUser    *uint64 `json:"rate_limit_per_user,omitempty"`   // amount of seconds a user has to wait before sending another message (0-21600)
}

// StartThreadWithoutMessage - Creates a new thread that is not connected to an existing message.
//
// Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.
//
// Fires a ThreadCreate Gateway event.
//
//    This endpoint supports the X-Audit-Log-Reason header.
//
// * Creating a GuildPrivateThread requires the server to be boosted. The GuildFeatures will indicate if that is possible for the guild.
func (c *Channel) StartThreadWithoutMessage(payload StartThreadWithoutMessageJSON, reason *string) (*Channel, error) {
	u := parseRoute(fmt.Sprintf(startThreadWithoutMessage, api, c.ID.String()))

	var channel Channel
	err := json.Unmarshal(firePostRequest(u, payload, reason), &channel)

	return &channel, err
}

// StartThreadWithoutMessageJSON - JSON payload structure
type StartThreadWithoutMessageJSON struct {
	Name                string      `json:"name"`                          // 1-100 character channel name
	AutoArchiveDuration uint64      `json:"auto_archive_duration"`         // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	Type                ChannelType `json:"type"`                          // the type of thread to create
	Invitable           bool        `json:"invitable"`                     // whether non-moderators can add other non-moderators to a thread; only available when creating a private thread
	RateLimitPerUser    *uint64     `json:"rate_limit_per_user,omitempty"` // amount of seconds a user has to wait before sending another message (0-21600)
}

// StartThreadInForumChannel
//
// Creates a new thread in a forum channel, and sends a message within the created thread. Returns a Channel, with a nested Message object, on success, and a 400 BAD REQUEST on invalid parameters. Fires a ThreadCreate and Message Create Gateway event.
//
//     The type of the created thread is GuildPublicThread.
//     See message formatting for more information on how to properly format messages.
//     The current user must have the SendMessages permission (CreatePublicThreads is ignored).
//     The maximum request size when sending a message is 8MiB.
//     For the embed object, you can set every field except type (it will be rich regardless of if you try to set it), provider, video, and any height, width, or proxy_url values for images.
//     Examples for file uploads are available in Uploading Files.
//     Files must be attached using a multipart/form-data body as described in Uploading Files.
//     Note that when sending a message, you must provide a value for at least one of content, embeds, or files[n].
//
//     Discord may strip certain characters from message content, like invalid unicode characters or characters which cause unexpected message formatting. If you are passing user-generated strings into message content, consider sanitizing the data to prevent unexpected behavior and utilizing allowed_mentions to prevent unexpected mentions.
//
//     This endpoint supports the X-Audit-Log-Reason header.
func (c *Channel) StartThreadInForumChannel(payload StartThreadWithoutMessageJSON, reason *string) (*Channel, error) {
	u := parseRoute(fmt.Sprintf(startThreadInForumChannel, api, c.ID.String()))

	var channel Channel
	err := json.Unmarshal(firePostRequest(u, payload, reason), &channel)

	return &channel, err
}

// StartThreadInForumJSON - JSON payload structure
type StartThreadInForumJSON struct {
	Name                string                   `json:"name"`                          // 1-100 character channel name
	AutoArchiveDuration uint64                   `json:"auto_archive_duration"`         // duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	RateLimitPerUser    *uint64                  `json:"rate_limit_per_user,omitempty"` // amount of seconds a user has to wait before sending another message (0-21600)
	Message             ForumThreadMessageParams `json:"message"`                       // contents of the first message in the forum thread
}

// ForumThreadMessageParams - JSON for starting a new forum thread
//
// TODO: files[n]
type ForumThreadMessageParams struct {
	Content         string          `json:"content"`          // the message contents (up to 2000 characters)
	Embeds          []Embed         `json:"embeds"`           // embedded rich content (up to 6000 characters)
	AllowedMentions AllowedMentions `json:"allowed_mentions"` // allowed mentions for the message
	Components      []Component     `json:"components"`       // the components to include with the message
	StickerIDs      []Snowflake     `json:"sticker_ids"`      // IDs of up to 3 stickers in the server to send in the message
	PayloadJson     string          `json:"payload_json"`     // JSON encoded body of non-file params
	Attachments     []Attachment    `json:"attachments"`      // attachment objects with filename and description
	Flags           MessageFlags    `json:"flags"`            // message flags combined as a bitfield (only SUPPRESS_EMBEDS can be set)
}

// JoinThread - Adds the current user to a thread.
//
// Also requires the thread is not archived.
//
// Returns a 204 empty response on success.
//
// Fires a ThreadMembersUpdate Gateway event.
func (c *Channel) JoinThread() {
	u := parseRoute(fmt.Sprintf(joinThread, api, c.ID.String()))

	_ = firePutRequest(u, nil, nil)
}

// AddThreadMember - Adds another member to a thread.
//
// Requires the ability to send messages in the thread.
//
// Also requires the thread is not archived.
//
// Returns a 204 empty response if the member is successfully added or was already a member of the thread.
//
// Fires a Thread Members Update Gateway event.
func (c *Channel) AddThreadMember(userID Snowflake) {
	u := parseRoute(fmt.Sprintf(addThreadMember, api, c.ID.String(), userID.String()))

	_ = firePutRequest(u, nil, nil)
}

// LeaveThread - Removes the current user from a thread.
//
// Also requires the thread is not archived.
//
// Returns a 204 empty response on success.
//
// Fires a ThreadMembersUpdate Gateway event.
func (c *Channel) LeaveThread() error {
	u := parseRoute(fmt.Sprintf(leaveThread, api, c.ID.String()))

	return fireDeleteRequest(u, nil)
}

// RemoveThreadMember - Removes another member from a thread.
//
// Requires the ManageThreads permission, or the creator of the thread if it is a GuildPrivateThread.
//
// Also requires the thread is not archived.
//
// Returns a 204 empty response on success.
//
// Fires a Thread Members Update Gateway event.
func (c *Channel) RemoveThreadMember(userID Snowflake) error {
	u := parseRoute(fmt.Sprintf(removeThreadMember, api, c.ID.String(), userID.String()))

	return fireDeleteRequest(u, nil)
}

// GetThreadMember - Returns a thread member object for the specified user if they are a member of the thread, returns a 404 response otherwise.
func (c *Channel) GetThreadMember(userID Snowflake) (*ThreadMember, error) {
	u := parseRoute(fmt.Sprintf(getThreadMember, api, c.ID.String(), userID.String()))

	var threadMember *ThreadMember
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadMember)

	return threadMember, err
}

// ListThreadMembers - Returns array of thread members objects that are members of the thread.
//
// This endpoint is restricted according to whether the GuildMembers Privileged Intent is enabled for your application.
func (c *Channel) ListThreadMembers() ([]ThreadMember, error) {
	u := parseRoute(fmt.Sprintf(listThreadMembers, api, c.ID.String()))

	var threadMembers []ThreadMember
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadMembers)

	return threadMembers, err
}

// ListPublicArchivedThreads - Returns archived threads in the channel that are public.
//
// When called on a GuildText channel, returns threads of type GuildPublicThread.
//
// When called on a GuildNews channel returns threads of type GuildNewsThread.
//
// Threads are ordered by archive_timestamp, in descending order.
//
// Requires the ReadMessageHistory permission.
func (c *Channel) ListPublicArchivedThreads(before *time.Time, limit *int) (*ThreadListResponse, error) {
	u := parseRoute(fmt.Sprintf(listPublicArchivedThreads, api, c.ID.String()))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var threadListResponse *ThreadListResponse
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadListResponse)

	return threadListResponse, err
}

type ThreadListResponse struct {
	Threads []Channel      `json:"threads"`            // the archived threads
	Members []ThreadMember `json:"members"`            // a thread member object for each returned thread the current user has joined
	HasMore bool           `json:"has_more,omitempty"` // whether there are potentially additional threads that could be returned on a subsequent call
}

// ListPrivateArchivedThreads - Returns archived threads in the channel that are of type GuildPrivateThread.
//
// Threads are ordered by archive_timestamp, in descending order.
//
// Requires both the READ_MESSAGE_HISTORY and MANAGE_THREADS permissions.
func (c *Channel) ListPrivateArchivedThreads(before *time.Time, limit *int) (*ThreadListResponse, error) {
	u := parseRoute(fmt.Sprintf(listPrivateArchivedThreads, api, c.ID.String()))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var threadListResponse *ThreadListResponse
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadListResponse)

	return threadListResponse, err
}

// ListJoinedPrivateArchivedThreads - Returns archived threads in the channel that are of type GuildPrivateThread, and the user has joined.
//
// Threads are ordered by their id, in descending order.
//
// Requires the READ_MESSAGE_HISTORY permission.
func (c *Channel) ListJoinedPrivateArchivedThreads(before *Snowflake, limit *int) (*ThreadListResponse, error) {
	u := parseRoute(fmt.Sprintf(listJoinedPrivateArchivedThreads, api, c.ID.String()))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var threadListResponse *ThreadListResponse
	err := json.Unmarshal(fireGetRequest(u, nil, nil), &threadListResponse)

	return threadListResponse, err
}
