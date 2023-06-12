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

package api

import (
	"time"
)

// Channel - Represents a guild or DM channel within Discord.
type Channel struct {
	rest *httpData

	ID                            Snowflake        `json:"id"`                                           // the id of this channel
	Type                          ChannelType      `json:"type"`                                         // the ChannelType
	GuildID                       Snowflake        `json:"guild_id,omitempty"`                           // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	Position                      int              `json:"position,omitempty"`                           // sorting position of the channel
	PermissionOverwrites          []*Overwrite     `json:"permission_overwrites,omitempty"`              // explicit permission overwrites for members and roles
	Name                          string           `json:"name,omitempty"`                               // the name of the channel (1-100 characters)
	Topic                         *string          `json:"topic,omitempty"`                              // the channel topic (0-1024 characters)
	Nsfw                          bool             `json:"nsfw,omitempty"`                               // whether the channel is nsfw
	LastMessageID                 *Snowflake       `json:"last_message_id,omitempty"`                    // the id of the last message sent in this channel (may not point to an existing or valid message)
	Bitrate                       int64            `json:"bitrate,omitempty"`                            // the bitrate (in bits) of the voice channel
	UserLimit                     int64            `json:"user_limit,omitempty"`                         // the user limit of the voice channel
	RateLimitPerUser              int64            `json:"rate_limit_per_user,omitempty"`                // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Recipients                    []*User          `json:"recipients,omitempty"`                         // the recipients of the DM
	Icon                          *string          `json:"icon,omitempty"`                               // icon hash of the group DM
	OwnerID                       Snowflake        `json:"owner_id,omitempty"`                           // id of the creator of the group DM or thread
	ApplicationID                 Snowflake        `json:"application_id,omitempty"`                     // application id of the group DM creator if it is bot-created
	Managed                       bool             `json:"managed,omitempty"`                            // for group DM channels: whether the channel is managed by an application via the gdm.join OAuth2 scope
	ParentID                      *Snowflake       `json:"parent_id,omitempty"`                          // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	LastPinTimestamp              *time.Time       `json:"last_pin_timestamp,omitempty"`                 // when the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned.
	RtcRegion                     *string          `json:"rtc_region,omitempty"`                         // voice region id for the voice channel, automatic when set to null
	VideoQualityMode              int64            `json:"video_quality_mode,omitempty"`                 // the camera video quality mode of the voice channel, 1 when not present
	MessageCount                  int64            `json:"message_count,omitempty"`                      // an approximate count of messages in a thread, stops counting at 50
	MemberCount                   int64            `json:"member_count,omitempty"`                       // an approximate count of users in a thread, stops counting at 50
	ThreadMetadata                ThreadMetadata   `json:"thread_metadata,omitempty"`                    // thread-specific fields not needed by other channels
	Member                        ThreadMember     `json:"member,omitempty"`                             // ThreadMember for the current User, if they have joined the thread, only included on certain API endpoints
	DefaultAutoArchiveDuration    int              `json:"default_auto_archive_duration,omitempty"`      // default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	Permissions                   string           `json:"permissions"`                                  // computed permissions for the invoking user in the channel, including overwrites, only included when part of the resolved data received on a slash command interaction
	Flags                         ChannelFlag      `json:"flags,omitempty"`                              // channel flags combined as a bitfield
	TotalMessagesSent             int64            `json:"total_messages_sent,omitempty"`                // number of messages ever sent in a thread, it's similar to MessageCount on message creation, but will not decrement the number when a message is deleted
	AvailableTags                 []*ForumTag      `json:"available_tags,omitempty"`                     // the set of tags that can be used in a GuildForum channel
	AppliedTags                   []*Snowflake     `json:"applied_tags,omitempty"`                       // the IDs of the set of tags that have been applied to a thread in a GuildForum channel
	DefaultReactionEmoji          *DefaultReaction `json:"default_reaction_emoji,omitempty"`             // the emoji to show in the add reaction button on a thread in a GuildForum channel
	DefaultThreadRateLimitPerUser uint             `json:"default_thread_rate_limit_per_user,omitempty"` // the initial RateLimitPerUser to set on newly created threads in a channel. this field is copied to the thread at creation time and does not live update.
	DefaultSortOrder              *SortOrderType   `json:"default_sort_order,omitempty"`                 // the default sort order type used to order posts in GuildForum channels. Defaults to null, which indicates a preferred sort order hasn't been set by a channel admin
	DefaultForumLayout            *ForumLayoutType `json:"default_forum_layout,omitempty"`               // the default forum layout view used to display posts in GuildForum channels. Defaults to NotSet (0), which indicates a layout view has not been set by a channel admin
}

type GuildTextChannel struct {
	*Channel

	LastMessageID        *Snowflake   `json:"last_message_id,omitempty"`       // the id of the last message sent in this channel (may not point to an existing or valid message)
	Position             int          `json:"position,omitempty"`              // sorting position of the channel
	Flags                ChannelFlag  `json:"flags,omitempty"`                 // channel flags combined as a bitfield
	ParentID             *Snowflake   `json:"parent_id,omitempty"`             // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Topic                *string      `json:"topic,omitempty"`                 // the channel topic (0-1024 characters)
	GuildID              Snowflake    `json:"guild_id,omitempty"`              // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"` // explicit permission overwrites for members and roles
	LastPinTimestamp     *time.Time   `json:"last_pin_timestamp,omitempty"`    // when the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned.
	RateLimitPerUser     int64        `json:"rate_limit_per_user,omitempty"`   // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Nsfw                 bool         `json:"nsfw,omitempty"`                  // whether the channel is nsfw
}

type GuildVoiceChannel struct {
	*Channel

	LastMessageID        *Snowflake   `json:"last_message_id,omitempty"`       // the id of the last message sent in this channel (may not point to an existing or valid message)
	Position             int          `json:"position,omitempty"`              // sorting position of the channel
	Flags                ChannelFlag  `json:"flags,omitempty"`                 // channel flags combined as a bitfield
	ParentID             *Snowflake   `json:"parent_id,omitempty"`             // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Bitrate              int64        `json:"bitrate,omitempty"`               // the bitrate (in bits) of the voice channel
	UserLimit            int64        `json:"user_limit,omitempty"`            // the user limit of the voice channel
	RtcRegion            *string      `json:"rtc_region,omitempty"`            // voice region id for the voice channel, automatic when set to null
	GuildID              Snowflake    `json:"guild_id,omitempty"`              // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"` // explicit permission overwrites for members and roles
	RateLimitPerUser     int64        `json:"rate_limit_per_user,omitempty"`   // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Nsfw                 bool         `json:"nsfw,omitempty"`                  // whether the channel is nsfw
}

type GuildCategoryChannel struct {
	*Channel

	Position             int          `json:"position,omitempty"`              // sorting position of the channel
	Flags                ChannelFlag  `json:"flags,omitempty"`                 // channel flags combined as a bitfield
	ParentID             *Snowflake   `json:"parent_id,omitempty"`             // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	GuildID              Snowflake    `json:"guild_id,omitempty"`              // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"` // explicit permission overwrites for members and roles
}

type GuildAnnouncementChannel struct {
	*Channel

	LastMessageID        *Snowflake   `json:"last_message_id,omitempty"`       // the id of the last message sent in this channel (may not point to an existing or valid message)
	Position             int          `json:"position,omitempty"`              // sorting position of the channel
	Flags                ChannelFlag  `json:"flags,omitempty"`                 // channel flags combined as a bitfield
	ParentID             *Snowflake   `json:"parent_id,omitempty"`             // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Topic                *string      `json:"topic,omitempty"`                 // the channel topic (0-1024 characters)
	GuildID              Snowflake    `json:"guild_id,omitempty"`              // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"` // explicit permission overwrites for members and roles
	LastPinTimestamp     *time.Time   `json:"last_pin_timestamp,omitempty"`    // when the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned.
	RateLimitPerUser     int64        `json:"rate_limit_per_user,omitempty"`   // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Nsfw                 bool         `json:"nsfw,omitempty"`                  // whether the channel is nsfw
}

type GuildAnnouncementThreadChannel struct {
	*Channel

	GuildID           Snowflake      `json:"guild_id,omitempty"`            // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	ParentID          *Snowflake     `json:"parent_id,omitempty"`           // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	OwnerID           Snowflake      `json:"owner_id,omitempty"`            // id of the creator of the group DM or thread
	LastMessageID     *Snowflake     `json:"last_message_id,omitempty"`     // the id of the last message sent in this channel (may not point to an existing or valid message)
	ThreadMetadata    ThreadMetadata `json:"thread_metadata,omitempty"`     // thread-specific fields not needed by other channels
	MessageCount      int64          `json:"message_count,omitempty"`       // an approximate count of messages in a thread, stops counting at 50
	MemberCount       int64          `json:"member_count,omitempty"`        // an approximate count of users in a thread, stops counting at 50
	RateLimitPerUser  int64          `json:"rate_limit_per_user,omitempty"` // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Flags             ChannelFlag    `json:"flags,omitempty"`               // channel flags combined as a bitfield
	TotalMessagesSent int64          `json:"total_messages_sent,omitempty"` // number of messages ever sent in a thread, it's similar to MessageCount on message creation, but will not decrement the number when a message is deleted
}

type GuildPublicThreadChannel struct {
	*Channel

	GuildID           Snowflake      `json:"guild_id,omitempty"`            // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	ParentID          *Snowflake     `json:"parent_id,omitempty"`           // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	OwnerID           Snowflake      `json:"owner_id,omitempty"`            // id of the creator of the group DM or thread
	LastMessageID     *Snowflake     `json:"last_message_id,omitempty"`     // the id of the last message sent in this channel (may not point to an existing or valid message)
	ThreadMetadata    ThreadMetadata `json:"thread_metadata,omitempty"`     // thread-specific fields not needed by other channels
	MessageCount      int64          `json:"message_count,omitempty"`       // an approximate count of messages in a thread, stops counting at 50
	MemberCount       int64          `json:"member_count,omitempty"`        // an approximate count of users in a thread, stops counting at 50
	RateLimitPerUser  int64          `json:"rate_limit_per_user,omitempty"` // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Flags             ChannelFlag    `json:"flags,omitempty"`               // channel flags combined as a bitfield
	TotalMessagesSent int64          `json:"total_messages_sent,omitempty"` // number of messages ever sent in a thread, it's similar to MessageCount on message creation, but will not decrement the number when a message is deleted
	AppliedTags       []*Snowflake   `json:"applied_tags,omitempty"`        // GuildForum only : the IDs of the set of tags that have been applied to a thread in a GuildForum channel
}

type GuildPrivateThreadChannel struct {
	*Channel

	GuildID           Snowflake      `json:"guild_id,omitempty"`            // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	ParentID          *Snowflake     `json:"parent_id,omitempty"`           // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	OwnerID           Snowflake      `json:"owner_id,omitempty"`            // id of the creator of the group DM or thread
	LastMessageID     *Snowflake     `json:"last_message_id,omitempty"`     // the id of the last message sent in this channel (may not point to an existing or valid message)
	ThreadMetadata    ThreadMetadata `json:"thread_metadata,omitempty"`     // thread-specific fields not needed by other channels
	MessageCount      int64          `json:"message_count,omitempty"`       // an approximate count of messages in a thread, stops counting at 50
	MemberCount       int64          `json:"member_count,omitempty"`        // an approximate count of users in a thread, stops counting at 50
	RateLimitPerUser  int64          `json:"rate_limit_per_user,omitempty"` // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Flags             ChannelFlag    `json:"flags,omitempty"`               // channel flags combined as a bitfield
	TotalMessagesSent int64          `json:"total_messages_sent,omitempty"` // number of messages ever sent in a thread, it's similar to MessageCount on message creation, but will not decrement the number when a message is deleted
}

type GuildStageVoiceChannel struct {
	*Channel

	LastMessageID        *Snowflake   `json:"last_message_id,omitempty"`       // the id of the last message sent in this channel (may not point to an existing or valid message)
	Position             int          `json:"position,omitempty"`              // sorting position of the channel
	Flags                ChannelFlag  `json:"flags,omitempty"`                 // channel flags combined as a bitfield
	ParentID             *Snowflake   `json:"parent_id,omitempty"`             // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Topic                *string      `json:"topic,omitempty"`                 // the channel topic (0-1024 characters)
	Bitrate              int64        `json:"bitrate,omitempty"`               // the bitrate (in bits) of the voice channel
	UserLimit            int64        `json:"user_limit,omitempty"`            // the user limit of the voice channel
	RtcRegion            *string      `json:"rtc_region,omitempty"`            // voice region id for the voice channel, automatic when set to null
	GuildID              Snowflake    `json:"guild_id,omitempty"`              // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"` // explicit permission overwrites for members and roles
	RateLimitPerUser     int64        `json:"rate_limit_per_user,omitempty"`   // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Nsfw                 bool         `json:"nsfw,omitempty"`                  // whether the channel is nsfw
}

type GuildForumChannel struct {
	*Channel

	LastMessageID                 *Snowflake       `json:"last_message_id,omitempty"`                    // the id of the last message sent in this channel (may not point to an existing or valid message)
	DefaultAutoArchiveDuration    int              `json:"default_auto_archive_duration,omitempty"`      // default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	DefaultThreadRateLimitPerUser uint             `json:"default_thread_rate_limit_per_user,omitempty"` // the initial RateLimitPerUser to set on newly created threads in a channel. this field is copied to the thread at creation time and does not live update.
	Position                      int              `json:"position,omitempty"`                           // sorting position of the channel
	Flags                         ChannelFlag      `json:"flags,omitempty"`                              // channel flags combined as a bitfield
	ParentID                      *Snowflake       `json:"parent_id,omitempty"`                          // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Topic                         *string          `json:"topic,omitempty"`                              // the channel topic (0-1024 characters)
	GuildID                       Snowflake        `json:"guild_id,omitempty"`                           // the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	PermissionOverwrites          []*Overwrite     `json:"permission_overwrites,omitempty"`              // explicit permission overwrites for members and roles
	RateLimitPerUser              int64            `json:"rate_limit_per_user,omitempty"`                // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Nsfw                          bool             `json:"nsfw,omitempty"`                               // whether the channel is nsfw
	AvailableTags                 []*ForumTag      `json:"available_tags,omitempty"`                     // the set of tags that can be used in a GuildForum channel
	Template                      string           `json:"template,omitempty"`                           // Undocumented as of 3/15/2023
	DefaultReactionEmoji          *DefaultReaction `json:"default_reaction_emoji,omitempty"`             // the emoji to show in the add reaction button on a thread in a GuildForum channel
	DefaultSortOrder              *SortOrderType   `json:"default_sort_order,omitempty"`                 // the default sort order type used to order posts in GuildForum channels. Defaults to null, which indicates a preferred sort order hasn't been set by a channel admin
	DefaultForumLayout            *ForumLayoutType `json:"default_forum_layout,omitempty"`               // the default forum layout view used to display posts in GuildForum channels. Defaults to NotSet (0), which indicates a layout view has not been set by a channel admin
}

// ChannelType - the type of channel
type ChannelType int

//goland:noinspection SpellCheckingInspection,GoUnusedConst
const (
	GuildText               ChannelType = iota     // a text channel within a server
	DM                                             // a direct message between users
	GuildVoice                                     // a voice channel within a server
	GroupDM                                        // a direct message between multiple users
	GuildCategory                                  // an organizational category that contains up to 50 channels
	GuildAnnouncement                              // a channel that users can follow and crosspost into their own server (formerly news channels)
	GuildAnnouncementThread ChannelType = iota + 4 // a temporary sub-channel within a GuildAnnouncement channel
	GuildPublicThread                              // a temporary sub-channel within a GuildText channel
	GuildPrivateThread                             // a temporary sub-channel within a GuildText channel that is only viewable by those invited and those with the ManageThreads permission
	GuildStageVoice                                // a voice channel for hosting events with an audience
	GuildDirectory                                 // the channel in a hub containing the listed servers
	GuildForum                                     // Channel that can only contain threads
)

func isTextChannel(channel *Channel) bool {
	return channel.Type == GuildText || channel.Type == GuildAnnouncement || channel.Type == GuildAnnouncementThread || channel.Type == GuildPublicThread ||
		channel.Type == GuildPrivateThread || channel.Type == GuildDirectory || channel.Type == GuildForum
}

// VideoQualityMode - the camera video quality mode of the voice channel, 1 when not present
type VideoQualityMode int

//goland:noinspection GoUnusedConst
const (
	Auto VideoQualityMode = iota + 1 // Discord chooses the quality for optimal performance
	Full                             // 720p
)

// ChannelFlag - channel flags combined as a bitfield
type ChannelFlag int

//goland:noinspection GoUnusedConst
const (
	Pinned     ChannelFlag = 1 << 1 // this thread is pinned to the top of its parent GuildForum channel
	RequireTag ChannelFlag = 1 << 4 // whether a tag is required to be specified when creating a thread in a GuildForum channel. Tags are specified in the AppliedTags field.
)

// SortOrderType - the default sort order type used to order posts in GuildForum channels.
type SortOrderType int

//goland:noinspection GoUnusedConst
const (
	LatestActivity SortOrderType = iota // Sort forum posts by activity
	CreationDate                        // Sort forum posts by creation time (from most recent to oldest)
)

// ForumLayoutType - the default forum layout view used to display posts in GuildForum channels.
type ForumLayoutType int

//goland:noinspection GoUnusedConst
const (
	NotSet      ForumLayoutType = iota // No default has been set for forum channel
	ListView                           // Display posts as a list
	GalleryView                        // Display posts as a collection of tiles
)

// Message - Represents a message sent in a channel within Discord.
//
//goland:noinspection SpellCheckingInspection
type Message struct {
	ID                   Snowflake          `json:"id,omitempty"`                     // id of the message
	ChannelID            Snowflake          `json:"channel_id,omitempty"`             // id of the Channel the message was sent in
	Author               User               `json:"author,omitempty"`                 // the author of this message (not guaranteed to be a valid user)
	Content              string             `json:"content,omitempty"`                // contents of the message
	Timestamp            time.Time          `json:"timestamp,omitempty"`              // when this message was sent
	EditedTimestamp      *time.Time         `json:"edited_timestamp,omitempty"`       // when this message was edited (or null if never)
	TTS                  bool               `json:"tts,omitempty"`                    // whether this was a TTS message
	MentionEveryone      bool               `json:"mention_everyone,omitempty"`       // whether this message mentions everyone
	Mentions             []*User            `json:"mentions,omitempty"`               // users specifically mentioned in the message
	MentionRoles         []*Snowflake       `json:"mention_roles,omitempty"`          // roles specifically mentioned in this message
	MentionChannels      []*Channel         `json:"mention_channels,omitempty"`       // channels specifically mentioned in this message
	Attachments          []*Attachment      `json:"attachments,omitempty"`            // any attached files
	Embeds               []*Embed           `json:"embeds,omitempty"`                 // any embedded content
	Reactions            []*Reaction        `json:"reactions,omitempty"`              // reactions to the message
	Nonce                any                `json:"nonce,omitempty"`                  // used for validating a message was sent
	Pinned               bool               `json:"pinned,omitempty"`                 // whether this message is pinned
	WebhookID            Snowflake          `json:"webhook_id,omitempty"`             // if the message is generated by a Webhook, this is the webhook's id
	Type                 MessageType        `json:"type,omitempty"`                   // the MessageType
	Activity             MessageActivity    `json:"activity,omitempty"`               // sent with Rich Presence-related chat embeds
	Application          Application        `json:"application,omitempty"`            // sent with Rich Presence-related chat embeds
	ApplicationID        Snowflake          `json:"application_id,omitempty"`         // if the message is an Interaction or application-owned webhook, this is the id of the application
	MessageReference     MessageReference   `json:"message_reference,omitempty"`      // data showing the source of a crosspost, channel follow add, pin, or reply message
	Flags                MessageFlags       `json:"flags,omitempty"`                  // MessageFlags combined as a bitfield
	ReferencedMessage    *Message           `json:"referenced_message,omitempty"`     // the message associated with the MessageReference
	Interaction          MessageInteraction `json:"interaction,omitempty"`            // sent if the message is a response to an Interaction
	Thread               Channel            `json:"thread,omitempty"`                 // the thread that was started from this message, includes ThreadMember object
	Components           []*Component       `json:"components,omitempty"`             // sent if the message contains components like buttons, action rows, or other interactive components
	StickerItems         []string           `json:"sticker_items,omitempty"`          // sent if the message contains stickers
	Stickers             []string           `json:"stickers,omitempty"`               // Deprecated: the stickers sent with the message
	Position             int                `json:"position,omitempty"`               // A generally increasing integer (there may be gaps or duplicates) that represents the approximate position of the message in a thread, it can be used to estimate the relative position of the message in a thread in company with total_message_sent on parent thread
	RoleSubscriptionData any                `json:"role_subscription_data,omitempty"` // data of the role subscription purchase or renewal that prompted this RoleSubscriptionPurchase message
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
	UserJoin                                                       // USER_JOIN
	GuildBoost                                                     // GUILD_BOOST
	GuildBoostTier1                                                // GUILD_BOOST_TIER_1
	GuildBoostTier2                                                // GUILD_BOOST_TIER_2
	GuildBoostTier3                                                // GUILD_BOOST_TIER_3
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
	AutoModerationAction                                           // AUTO_MODERATION_ACTION
	RoleSubscriptionPurchase                                       // ROLE_SUBSCRIPTION_PURCHASE
	InteractionPremiumUpsell                                       // INTERACTION_PREMIUM_UPSELL
	StageStart                                                     // STAGE_START
	StageEnd                                                       // STAGE_END
	StageSpeaker                                                   // STAGE_SPEAKER
	StageTopic                              MessageType = iota + 2 // STAGE_TOPIC
	GuildApplicationPremiumSubscription                            // GUILD_APPLICATION_PREMIUM_SUBSCRIPTION
)

// MessageActivity - sent with Rich Presence-related chat embeds
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`               // type of message activity
	PartyID string              `json:"party_id,omitempty"` // party_id from a Rich Presence event
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
	CrossPosted                      MessageFlags = 1 << 0  // this message has been published to subscribed channels (via Channel Following)
	IsCrossPost                      MessageFlags = 1 << 1  // this message originated from a message in another channel (via Channel Following)
	SuppressEmbeds                   MessageFlags = 1 << 2  // do not include any embeds when serializing this message
	SourceMessageDeleted             MessageFlags = 1 << 3  // the source message for this crosspost has been deleted (via Channel Following)
	Urgent                           MessageFlags = 1 << 4  // this message came from the urgent message system
	HasThread                        MessageFlags = 1 << 5  // this message has an associated thread, with the same id as the message
	Ephemeral                        MessageFlags = 1 << 6  // this message is only visible to the user who invoked the Interaction
	Loading                          MessageFlags = 1 << 7  // this message is an Interaction Response and the bot is "thinking"
	FailedToMentionSomeRolesInThread MessageFlags = 1 << 8  // this message failed to mention some roles and add their members to the thread
	SuppressNotifications            MessageFlags = 1 << 12 // this message will not trigger push and desktop notifications
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
	ID    Snowflake     `json:"id"`    // Role or User id
	Type  OverwriteType `json:"type"`  // either PermissionRole or PermissionMember
	Allow string        `json:"allow"` // permission bit set
	Deny  string        `json:"deny"`  // permission bit set
}

// OverwriteType - either PermissionRole or PermissionMember
type OverwriteType int

//goland:noinspection GoUnusedConst
const (
	PermissionRole   OverwriteType = iota // 0 (role)
	PermissionMember                      // 1 (member)
)

// ThreadMetadata - The thread metadata object contains a number of thread-specific channel fields that are not needed by other channel types.
//
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
	ID            Snowflake   `json:"id,omitempty"`      // ID of the thread
	UserID        Snowflake   `json:"user_id,omitempty"` // ID of the user
	JoinTimestamp time.Time   `json:"join_timestamp"`    // Time the user last joined the thread
	Flags         int64       `json:"flags"`             // Any user-thread settings, currently only used for notifications
	Member        GuildMember `json:"member,omitempty"`  // Additional information about the user
}

// DefaultReaction - An object that specifies the emoji to use as the default way to react to a forum post. Exactly one of emoji_id and emoji_name must be set.
type DefaultReaction struct {
	EmojiID   *Snowflake `json:"emoji_id"`   // the id of a guild's custom emoji
	EmojiName *string    `json:"emoji_name"` // the unicode character of the emoji
}

// ForumTag - An object that represents a tag that is able to be applied to a thread in a GuildForum channel.
//
// When updating a GuildForum channel, tag objects in available_tags only require the Name field.
//
// At most one of EmojiID and EmojiName may be set.
type ForumTag struct {
	ID        Snowflake `json:"id"`         // the id of the tag
	Name      string    `json:"name"`       // the name of the tag (0-20 characters)
	Moderated bool      `json:"moderated"`  // whether this tag can only be added to or removed from threads by a member with the ManageThreads permission
	EmojiID   Snowflake `json:"emoji_id"`   // the id of a guild's custom emoji
	EmojiName *string   `json:"emoji_name"` // the unicode character of the emoji
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
//
//goland:noinspection SpellCheckingInspection
type Attachment struct {
	ID          Snowflake `json:"id"`                     // attachment id
	Filename    string    `json:"filename"`               // name of file attached
	Description string    `json:"description,omitempty"`  // description for the file
	ContentType string    `json:"content_type,omitempty"` // the attachment's media type
	Size        int       `json:"size"`                   // size of file in bytes
	URL         string    `json:"url"`                    // source url of file
	ProxyURL    string    `json:"proxy_url"`              // a proxied url of file
	Height      *int      `json:"height,omitempty"`       // height of file (if image)
	Width       *int      `json:"width,omitempty"`        // width of file (if image)
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
	Parse       []*AllowedMentionType `json:"parse"`                  // An array of AllowedMentionType to parse from the content.
	Roles       []*Snowflake          `json:"roles,omitempty"`        // Array of role_ids to mention (Max size of 100)
	Users       []*Snowflake          `json:"users,omitempty"`        // Array of user_ids to mention (Max size of 100)
	RepliedUser bool                  `json:"replied_user,omitempty"` // For replies, whether to mention the author of the message being replied to (default false)
}

// AllowedMentionType - the type of mention allowed
type AllowedMentionType string

//goland:noinspection GoUnusedConst
const (
	RoleMentions     AllowedMentionType = "roles"    // Controls role mentions
	UserMentions     AllowedMentionType = "users"    // Controls user mentions
	EveryoneMentions AllowedMentionType = "everyone" // Controls @everyone and @here mentions
)

// RoleSubscriptionData - data of the role subscription purchase or renewal that prompted this RoleSubscriptionData message
type RoleSubscriptionData struct {
	RoleSubscriptionListingID Snowflake `json:"role_subscription_listing_id"` // the id of the sku and listing that the user is subscribed to
	TierName                  string    `json:"tier_name"`                    // the name of the tier that the user is subscribed to
	TotalMonthsSubscribed     uint      `json:"total_months_subscribed"`      // the cumulative number of months that the user has been subscribed for
	IsRenewal                 bool      `json:"is_renewal"`                   // whether this notification is for a renewal rather than a new purchase
}

// Additionally, the combined sum of characters in all title, description, field.name, field.value, footer.text, and author.name fields across all embeds attached to a message must not exceed 6000 characters.
//
// Violating any of these constraints will result in a Bad Request response.
const (
	TitleLimit       = 256  // 256 characters
	DescriptionLimit = 4096 // 4096 characters
	FieldCount       = 25   // Up to 25 field objects
	FieldNameLimit   = 256  // 256 characters
	FieldValueLimit  = 1024 // 1024 characters
	FooterTextLimit  = 2048 // 2048 characters
	AuthorNameLimit  = 256  // 256 characters
)
