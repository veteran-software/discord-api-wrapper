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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/veteran-software/discord-api-wrapper/v10/logging"
)

// Guild - Guilds in Discord represent an isolated collection of users and channels, and are often referred to as "servers" in the UI.
type Guild struct {
	ID                          Snowflake                       `json:"id"`                                   // guild id
	Name                        string                          `json:"name"`                                 // guild name (2-100 characters, excluding trailing and leading whitespace)
	Icon                        *string                         `json:"icon"`                                 // icon hash
	IconHash                    *string                         `json:"icon_hash,omitempty"`                  // icon hash, returned when in the template object
	Splash                      *string                         `json:"splash,omitempty"`                     // splash hash
	DiscoverySplash             *string                         `json:"discovery_splash"`                     // discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	OwnerID                     Snowflake                       `json:"owner_id"`                             // id of owner
	AfkChannelID                Snowflake                       `json:"afk_channel_id,omitempty"`             // id of afk channel
	AfkTimeout                  int64                           `json:"afk_timeout"`                          // afk timeout in seconds
	WidgetEnabled               bool                            `json:"widget_enabled,omitempty"`             // true if the server widget is enabled
	WidgetChannelID             *Snowflake                      `json:"widget_channel_id,omitempty"`          // the channel id that the widget will generate an "invite" to, or null if set to no invite
	VerificationLevel           VerificationLevel               `json:"verification_level"`                   // verification level required for the guild
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"`        // default message notifications level
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`              // explicit content filter level
	Roles                       []Role                          `json:"roles"`                                // roles in the guild
	Emojis                      []Emoji                         `json:"emojis"`                               // custom guild emojis
	Features                    []GuildFeatures                 `json:"features"`                             // enabled guild features
	MfaLevel                    MfaLevel                        `json:"mfa_level"`                            // required MFA level for the guild
	ApplicationID               *Snowflake                      `json:"application_id"`                       // application id of the guild creator if it is bot-created
	SystemChannelID             *Snowflake                      `json:"system_channel_id"`                    // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags"`                 // system channel flags
	RulesChannelID              *Snowflake                      `json:"rules_channel_id"`                     // the id of the channel where Community guilds can display rules and/or guidelines
	MaxPresences                *int64                          `json:"max_presences,omitempty"`              // the maximum number of presences for the guild (null is always returned, apart from the largest of guilds)
	MaxMembers                  int64                           `json:"max_members,omitempty"`                // the maximum number of members for the guild
	VanityUrlCode               *string                         `json:"vanity_url_code"`                      // the vanity url code for the guild
	Description                 *string                         `json:"description"`                          // the description of a Community guild
	Banner                      *string                         `json:"banner"`                               // banner hash
	PremiumTier                 PremiumTier                     `json:"premium_tier"`                         // premium tier (Server Boost level)
	PremiumSubscriptionCount    uint64                          `json:"premium_subscription_count,omitempty"` // the number of boosts this guild currently has
	PreferredLocale             string                          `json:"preferred_locale"`                     // the preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	PublicUpdatesChannelID      *Snowflake                      `json:"public_updates_channel_id"`            // the id of the channel where admins and moderators of Community guilds receive notices from Discord
	MaxVideoChannelUsers        uint64                          `json:"max_video_channel_users,omitempty"`    // the maximum amount of users in a video channel
	ApproximateMemberCount      uint64                          `json:"approximate_member_count,omitempty"`   // approximate number of members in this guild, returned from the GET /guilds/<id> endpoint when with_counts is true
	ApproximatePresenceCount    uint64                          `json:"approximate_presence_count,omitempty"` // approximate number of non-offline members in this guild, returned from the GET /guilds/<id> endpoint when with_counts is true
	WelcomeScreen               WelcomeScreen                   `json:"welcome_screen,omitempty"`             // the welcome screen of a Community guild, shown to new members, returned in an Invite's guild object
	NsfwLevel                   GuildNsfwLevel                  `json:"nsfw_level"`                           // guild NSFW level
	Stickers                    []Sticker                       `json:"stickers,omitempty"`                   // custom guild stickers
	PremiumProgressBarEnabled   bool                            `json:"premium_progress_bar_enabled"`         // whether the guild has the boost progress bar enabled

	// These fields are only sent within the GUILD_CREATE event

	JoinedAt             time.Time             `json:"joined_at,omitempty"`              // when this guild was joined at
	Large                bool                  `json:"large,omitempty"`                  // true if this is considered a large guild
	Unavailable          bool                  `json:"unavailable,omitempty"`            // true if this guild is unavailable due to an outage
	MemberCount          int64                 `json:"member_count,omitempty"`           // total number of members in this guild
	VoiceStates          []VoiceState          `json:"voice_states,omitempty"`           // states of members currently in voice channels; lacks the guild_id key
	Members              []GuildMember         `json:"members,omitempty"`                // users in the guild
	Channels             []Channel             `json:"channels,omitempty"`               // channels in the guild
	Threads              []Channel             `json:"threads,omitempty"`                // all active threads in the guild that current user has permission to view
	Presences            []PresenceUpdateEvent `json:"presences,omitempty"`              // presences of the members in the guild, will only include non-offline members if the size is greater than large threshold
	StageInstances       []StageInstance       `json:"stage_instances,omitempty"`        // Stage instances in the guild
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events,omitempty"` // the scheduled events in the guild

	// These fields are only sent when using the GET Current user Guilds endpoint and are relative to the requested user

	Owner       bool   `json:"owner,omitempty"`       // true if the user is the owner of the guild
	Permissions string `json:"permissions,omitempty"` // total permissions for the user in the guild (excludes overwrites)
}

// DefaultMessageNotificationLevel - default message notifications level
type DefaultMessageNotificationLevel int

//goland:noinspection GoUnusedConst
const (
	AllMessages  DefaultMessageNotificationLevel = iota // members will receive notifications for all messages by default
	OnlyMentions                                        // members will receive notifications only for messages that @mention them by default
)

// ExplicitContentFilterLevel - explicit content filter level
type ExplicitContentFilterLevel int

//goland:noinspection GoUnusedConst
const (
	Disabled            ExplicitContentFilterLevel = iota // media content will not be scanned
	MembersWithoutRoles                                   // media content sent by members without roles will be scanned
	AllMembers                                            // media content sent by all members will be scanned
)

// MfaLevel - required MFA level for the guild
type MfaLevel int

//goland:noinspection GoUnusedConst
const (
	MfaNone     MfaLevel = iota // guild has no MFA/2FA requirement for moderation actions
	MfaElevated                 // guild has a 2FA requirement for moderation actions
)

// VerificationLevel - verification level required for the guild
type VerificationLevel int

//goland:noinspection GoUnusedConst
const (
	VerificationLevelNone     VerificationLevel = iota // unrestricted
	VerificationLevelLow                               // must have verified email on account
	VerificationLevelMedium                            // must be registered on Discord for longer than 5 minutes
	VerificationLevelHigh                              // must be a member of the server for longer than 10 minutes
	VerificationLevelVeryHigh                          // must have a verified phone number
)

// GuildNsfwLevel - guild NSFW level
type GuildNsfwLevel int

//goland:noinspection GoUnusedConst
const (
	NsfwDefault GuildNsfwLevel = iota
	NsfwExplicit
	NsfwSafe
	NsfwAgeRestricted
)

// PremiumTier - premium tier (Server Boost level)
type PremiumTier int

//goland:noinspection GoUnusedConst
const (
	PremiumNone  PremiumTier = iota // guild has not unlocked any Server Boost perks
	PremiumTier1                    // guild has unlocked Server Boost level 1 perks
	PremiumTier2                    // guild has unlocked Server Boost level 2 perks
	PremiumTier3                    // guild has unlocked Server Boost level 3 perks
)

// SystemChannelFlags - system channel flags
type SystemChannelFlags int

//goland:noinspection GoUnusedConst
const (
	SuppressJoinNotifications          SystemChannelFlags = 1 << 0 // Suppress member join notifications
	SuppressPremiumSubscriptions       SystemChannelFlags = 1 << 1 // Suppress server boost notifications
	SuppressGuildReminderNotifications SystemChannelFlags = 1 << 2 // Suppress server setup tips
	SuppressJoinNotificationReplies    SystemChannelFlags = 1 << 3 // Hide member join sticker reply buttons
)

// GuildFeatures - enabled guild features
type GuildFeatures string

//goland:noinspection SpellCheckingInspection,GrazieInspection,GoUnusedConst
const (
	AnimatedBanner                GuildFeatures = "ANIMATED_BANNER"                  // guild has access to set an animated guild banner image
	AnimatedIcon                  GuildFeatures = "ANIMATED_ICON"                    // guild has access to set an animated guild icon
	Banner                        GuildFeatures = "BANNER"                           // guild has access to set a guild banner image
	Commerce                      GuildFeatures = "COMMERCE"                         // guild has access to use commerce features (i.e. create store channels)
	Community                     GuildFeatures = "COMMUNITY"                        // guild can enable welcome screen, Membership Screening, stage channels and discovery, and receives community updates
	Discoverable                  GuildFeatures = "DISCOVERABLE"                     // guild is able to be discovered in the directory
	Featurable                    GuildFeatures = "FEATURABLE"                       // guild is able to be featured in the directory
	InviteSplash                  GuildFeatures = "INVITE_SPLASH"                    // guild has access to set an invite splash background
	MemberVerificationGateEnabled GuildFeatures = "MEMBER_VERIFICATION_GATE_ENABLED" // guild has enabled Membership Screening
	MonetizationEnabled           GuildFeatures = "MONETIZATION_ENABLED"             // guild has enabled monetization
	MoreStickers                  GuildFeatures = "MORE_STICKERS"                    // guild has increased custom sticker slots
	News                          GuildFeatures = "NEWS"                             // guild has access to create news channels
	Partnered                     GuildFeatures = "PARTNERED"                        // guild is partnered
	PreviewEnabled                GuildFeatures = "PREVIEW_ENABLED"                  // guild can be previewed before joining via Membership Screening or the directory
	PrivateThreads                GuildFeatures = "PRIVATE_THREADS"                  // guild has access to create private threads
	RoleIcons                     GuildFeatures = "ROLE_ICONS"                       // guild is able to set role icons
	SevenDayThreadArchive         GuildFeatures = "SEVEN_DAY_THREAD_ARCHIVE"         // guild has access to the seven day archive time for threads
	ThreeDayThreadArchive         GuildFeatures = "THREE_DAY_THREAD_ARCHIVE"         // guild has access to the three day archive time for threads
	TicketedEventsEnabled         GuildFeatures = "TICKETED_EVENTS_ENABLED"          // guild has enabled ticketed events
	VanityURL                     GuildFeatures = "VANITY_URL"                       // guild has access to set a vanity URL
	Verified                      GuildFeatures = "VERIFIED"                         // guild is verified
	VipRegions                    GuildFeatures = "VIP_REGIONS"                      // guild has access to set 384kbps bitrate in voice (previously VIP voice servers)
	WelcomeScreenEnabled          GuildFeatures = "WELCOME_SCREEN_ENABLED"           // guild has enabled the welcome screen
)

// UnavailableGuild - A partial guild object.
//
// Represents an Offline Guild, or a Guild whose information has not been provided through Guild Create events during the Gateway connect.
type UnavailableGuild struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

// GuildPreview - preview object
type GuildPreview struct {
	ID                       Snowflake       `json:"id"`                         // guild id
	Name                     string          `json:"name"`                       // guild name (2-100 characters)
	Icon                     *string         `json:"icon"`                       // icon hash
	Splash                   *string         `json:"splash"`                     // splash hash
	DiscoverySplash          *string         `json:"discovery_splash"`           // discovery splash hash
	Emojis                   []Emoji         `json:"emojis"`                     // custom guild emojis
	Features                 []GuildFeatures `json:"features"`                   // enabled guild features
	ApproximateMemberCount   int             `json:"approximate_member_count"`   // approximate number of members in this guild
	ApproximatePresenceCount int             `json:"approximate_presence_count"` // approximate number of online members in this guild
	Description              *string         `json:"description"`                // the description for the guild, if the guild is discoverable
	Stickers                 []Sticker       `json:"stickers"`                   // custom guild stickers
}

// GuildWidget - the guild widget status
type GuildWidget struct {
	Enabled   bool       `json:"enabled"`    // whether the widget is enabled
	ChannelID *Snowflake `json:"channel_id"` // the widget channel id
}

// GetGuildWidget - the guild widget
//
// The fields `id`, `discriminator` and `avatar` are anonymized to prevent abuse.
type GetGuildWidget struct {
	ID            Snowflake     `json:"id"`             // guild id
	Name          string        `json:"name"`           // guild name (2-100 characters)
	InstantInvite *string       `json:"instant_invite"` // instant invite for the guilds specified widget invite channel
	Channels      []Channel     `json:"channels"`       // voice and stage channels which are accessible by @everyone
	Members       []GuildMember `json:"members"`        // special widget user objects that includes users presence (Limit 100)
	PresenceCount int           `json:"presence_count"` // number of online members in this guild
}

// GuildMember - Represents a member of a Guild
//
//   The field `user` won't be included in the member object attached to MessageCreate and MessageUpdate gateway events.
//
//   In GUILD_ events, pending will always be included as true or false.
//   In non GUILD_ events which can only be triggered by non-pending users, pending will not be included.
type GuildMember struct {
	User                       User        `json:"user,omitempty"`                         // User - the user this guild member represents
	Nick                       *string     `json:"nick,omitempty"`                         // Nick - the users' guild nickname
	Avatar                     *string     `json:"avatar,omitempty"`                       // Avatar - guild specific avatar
	Roles                      []Snowflake `json:"roles"`                                  // Roles - array of GuildRole id's
	JoinedAt                   time.Time   `json:"joined_at"`                              // JoinedAt - when the user joined the guild
	PremiumSince               *time.Time  `json:"premium_since,omitempty"`                // PremiumSince - when the user started boosting the guild
	Deaf                       bool        `json:"deaf"`                                   // Deaf - whether the user is deafened in voice channels
	Mute                       bool        `json:"mute"`                                   // Mute - whether the user is muted in voice channels
	Pending                    bool        `json:"pending,omitempty"`                      // Pending - whether the user has not yet passed the guild's Membership Screening requirements
	Permissions                string      `json:"permissions,omitempty"`                  // Permissions - total permissions of the member in the channel, including overwrites, returned when in the interaction object
	CommunicationDisabledUntil *time.Time  `json:"communication_disabled_until,omitempty"` // CommunicationDisabledUntil - when the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
}

// Integration - a guild integration
type Integration struct {
	ID                Snowflake                 `json:"id"`                            // integration id
	Name              string                    `json:"name"`                          // integration name
	Type              string                    `json:"type"`                          // integration type (twitch, YouTube, or discord)
	Enabled           bool                      `json:"enabled"`                       // is this integration enabled
	Syncing           bool                      `json:"syncing,omitempty"`             // is this integration syncing
	RoleID            Snowflake                 `json:"role_id,omitempty"`             // id that this integration uses for "subscribers"
	EnableEmoticons   bool                      `json:"enable_emoticons,omitempty"`    // whether emoticons should be synced for this integration (twitch only currently)
	ExpireBehavior    IntegrationExpireBehavior `json:"expire_behavior,omitempty"`     // the behavior of expiring subscribers
	ExpireGracePeriod int                       `json:"expire_grace_period,omitempty"` // the grace period (in days) before expiring subscribers
	User              User                      `json:"user,omitempty"`                // user for this integration
	Account           IntegrationAccount        `json:"account"`                       // integration account information
	SyncedAt          time.Time                 `json:"synced_at,omitempty"`           // when this integration was last synced
	SubscriberCount   int                       `json:"subscriber_count,omitempty"`    // how many subscribers this integration has
	Revoked           bool                      `json:"revoked,omitempty"`             // has this integration been revoked
	Application       IntegrationApplication    `json:"application,omitempty"`         // The bot/OAuth2 application for discord integrations
}

// IntegrationExpireBehavior - the behavior of expiring subscribers
type IntegrationExpireBehavior int

//goland:noinspection GoUnusedConst
const (
	RemoveRole IntegrationExpireBehavior = iota // Remove role
	Kick                                        // Kick
)

// IntegrationAccount - integration account information
type IntegrationAccount struct {
	ID   string `json:"id"`   // id of the account
	Name string `json:"name"` // name of the account
}

// IntegrationApplication - The bot/OAuth2 application for discord integrations
type IntegrationApplication struct {
	ID          Snowflake `json:"id"`            // the id of the app
	Name        string    `json:"name"`          // the name of the app
	Icon        *string   `json:"icon"`          // the icon hash of the app
	Description string    `json:"description"`   // the description of the app
	Summary     string    `json:"summary"`       // the summary of the app
	Bot         User      `json:"bot,omitempty"` // the bot associated with this application
}

// Ban - represents a guild member ban object
type Ban struct {
	Reason *string `json:"reason"` // the reason for the ban
	User   User    `json:"user"`   // the banned user
}

// WelcomeScreen - the welcome screen object
type WelcomeScreen struct {
	Description     *string                `json:"description,omitempty"`      // the server description shown in the welcome screen
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels,omitempty"` // the channels shown in the welcome screen, up to 5
}

// WelcomeScreenChannel - the channels shown in the welcome screen, up to 5
type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id,omitempty"`  // the channel's id
	Description string     `json:"description,omitempty"` // the description shown for the channel
	EmojiID     *Snowflake `json:"emoji_id,omitempty"`    // the emoji id, if the emoji is custom
	EmojiName   *string    `json:"emoji_name,omitempty"`  // the emoji name if custom, the unicode character if standard, or null if no emoji is set
}

// String - Helper function to convert basic Guild data into string form
func (g *Guild) String() string {
	return g.Name + "(" + g.ID.String() + ")"
}

// CreateGuild
//
// Create a new guild. Returns a guild object on success. Fires a GuildCreate Gateway event.
//
//    This endpoint can be used only by bots in less than 10 guilds.
//
//    When using the roles parameter, the first member of the array is used to change properties of the guild's @everyone role. If you are trying to bootstrap a guild with additional roles, keep this in mind.
//
//    When using the roles parameter, the required id field within each role object is an integer placeholder, and will be replaced by the API upon consumption. Its purpose is to allow you to overwrite a role's permissions in a channel when also passing in channels with the channels array.
//
//    When using the channels parameter, the position field is ignored, and none of the default channels are created.
//
//    When using the channels parameter, the id field within each channel object may be set to an integer placeholder, and will be replaced by the API upon consumption. Its purpose is to allow you to create GUILD_CATEGORY channels by setting the parent_id field on any children to the category's id field. Category channels must be listed before any children.
//goland:noinspection GoUnusedExportedFunction
func CreateGuild(payload CreateGuildJSON) (*Guild, error) {
	u, err := url.Parse(fmt.Sprintf(createGuild, api))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodPost, u.String(), payload, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guild *Guild
	err = json.NewDecoder(resp.Body).Decode(&guild)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guild, nil
}

type CreateGuildJSON struct {
	Name                        string                          `json:"name"`                          // guild name (2-100 characters, excluding trailing and leading whitespace)
	Icon                        *string                         `json:"icon"`                          // icon hash
	VerificationLevel           VerificationLevel               `json:"verification_level"`            // verification level required for the guild
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"` // default message notifications level
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`       // explicit content filter level
	Roles                       []Role                          `json:"roles"`                         // roles in the guild
	Channels                    []Channel                       `json:"channels,omitempty"`            // channels in the guild
	AfkChannelID                Snowflake                       `json:"afk_channel_id,omitempty"`      // id of afk channel
	AfkTimeout                  int64                           `json:"afk_timeout"`                   // afk timeout in seconds
	SystemChannelID             *Snowflake                      `json:"system_channel_id"`             // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags"`          // system channel flags
}

// GetGuild - Returns the guild object for the given id.
//
// If with_counts is set to true, this endpoint will also return approximate_member_count and approximate_presence_count for the guild.
func (g *Guild) GetGuild(withCounts *bool) (*Guild, error) {
	u, err := url.Parse(fmt.Sprintf(getGuild, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	q := u.Query()
	if withCounts != nil {
		q.Set("with_counts", strconv.FormatBool(*withCounts))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guild *Guild
	err = json.NewDecoder(resp.Body).Decode(&guild)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guild, nil
}

// GetGuildPreview - Returns the guild preview object for the given id. If the user is not in the guild, then the guild must be lurkable.
func (g *Guild) GetGuildPreview() (*GuildPreview, error) {
	u, err := url.Parse(fmt.Sprintf(getGuildPreview, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guildPreview *GuildPreview
	err = json.NewDecoder(resp.Body).Decode(&guildPreview)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guildPreview, nil
}

// ModifyGuild - Modify a guild's settings.
//
// Requires the ManageGuild permission.
//
// Returns the updated guild object on success.
//
// Fires a GuildUpdate Gateway event.
//
//    All parameters to this endpoint are optional
//
//    This endpoint supports the X-Audit-Log-Reason header.
//
//    Attempting to add or remove the Community guild feature requires the Administrator permission.
func (g *Guild) ModifyGuild(payload ModifyGuildJSON, reason *string) (*Guild, error) {
	u, err := url.Parse(fmt.Sprintf(modifyGuild, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodPatch, u.String(), payload, reason)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guild *Guild
	err = json.NewDecoder(resp.Body).Decode(&guild)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guild, nil
}

type ModifyGuildJSON struct {
	Name                        string                           `json:"name"`                          // guild name (2-100 characters, excluding trailing and leading whitespace)
	VerificationLevel           *VerificationLevel               `json:"verification_level"`            // verification level required for the guild
	DefaultMessageNotifications *DefaultMessageNotificationLevel `json:"default_message_notifications"` // default message notifications level
	ExplicitContentFilter       *ExplicitContentFilterLevel      `json:"explicit_content_filter"`       // explicit content filter level
	AfkChannelID                *Snowflake                       `json:"afk_channel_id,omitempty"`      // id of afk channel
	AfkTimeout                  int64                            `json:"afk_timeout"`                   // afk timeout in seconds
	Icon                        *string                          `json:"icon"`                          // icon hash
	OwnerID                     Snowflake                        `json:"owner_id"`                      // id of owner
	Splash                      *string                          `json:"splash,omitempty"`              // splash hash
	DiscoverySplash             *string                          `json:"discovery_splash"`              // discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	Banner                      *string                          `json:"banner"`                        // banner hash
	SystemChannelID             *Snowflake                       `json:"system_channel_id"`             // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          SystemChannelFlags               `json:"system_channel_flags"`          // system channel flags
	RulesChannelID              *Snowflake                       `json:"rules_channel_id"`              // the id of the channel where Community guilds can display rules and/or guidelines
	PublicUpdatesChannelID      *Snowflake                       `json:"public_updates_channel_id"`     // the id of the channel where admins and moderators of Community guilds receive notices from Discord
	PreferredLocale             string                           `json:"preferred_locale"`              // the preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	Features                    []GuildFeatures                  `json:"features"`                      // enabled guild features
	Description                 *string                          `json:"description"`                   // the description of a Community guild
	PremiumProgressBarEnabled   bool                             `json:"premium_progress_bar_enabled"`  // whether the guild has the boost progress bar enabled
}

// DeleteGuild - Delete a guild permanently. User must be owner. Returns 204 No Content on success. Fires a GuildDelete Gateway event.
func (g *Guild) DeleteGuild() error {
	u, err := url.Parse(fmt.Sprintf(deleteGuild, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return err
	}

	resp, err := Rest.Request(http.MethodDelete, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}

// GetGuildChannels - Returns a list of guild Channel objects. Does not include threads.
func (g *Guild) GetGuildChannels() ([]Channel, error) {
	u, err := url.Parse(fmt.Sprintf(getGuildChannels, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channels []Channel
	err = json.NewDecoder(resp.Body).Decode(&channels)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return channels, nil
}

// CreateGuildChannel - Create a new channel object for the guild.
//
// Requires the ManageChannels permission.
//
// If setting permission overwrites, only permissions your bot has in the guild can be allowed/denied.
//
// Setting ManageRoles permission in channels is only possible for guild administrators.
//
// Returns the new channel object on success. Fires a Channel Create Gateway event.
//
//    All parameters to this endpoint are optional excluding name
//
//    This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) CreateGuildChannel(payload CreateGuildChannelJSON, reason *string) (*Channel, error) {
	u, err := url.Parse(fmt.Sprintf(createGuildChannel, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodPost, u.String(), payload, reason)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channel *Channel
	err = json.NewDecoder(resp.Body).Decode(&channel)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return channel, nil
}

type CreateGuildChannelJSON struct {
	Name                       string      `json:"name,omitempty"`                          // the name of the channel (1-100 characters)
	Type                       ChannelType `json:"type"`                                    // the ChannelType
	Topic                      *string     `json:"topic,omitempty"`                         // the channel topic (0-1024 characters)
	Bitrate                    int64       `json:"bitrate,omitempty"`                       // the bitrate (in bits) of the voice channel
	UserLimit                  int64       `json:"user_limit,omitempty"`                    // the user limit of the voice channel
	RateLimitPerUser           int64       `json:"rate_limit_per_user,omitempty"`           // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Position                   int         `json:"position,omitempty"`                      // sorting position of the channel
	PermissionOverwrites       []Overwrite `json:"permission_overwrites,omitempty"`         // explicit permission overwrites for members and roles
	ParentID                   *Snowflake  `json:"parent_id,omitempty"`                     // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Nsfw                       bool        `json:"nsfw,omitempty"`                          // whether the channel is nsfw
	DefaultAutoArchiveDuration int         `json:"default_auto_archive_duration,omitempty"` // default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
}

// ModifyGuildChannelPositions - Modify the positions of a set of channel objects for the guild.
//
// Requires ManageChannels permission. Returns a 204 empty response on success. Fires multiple ChannelUpdate Gateway events.
//
//    Only channels to be modified are required.
//
//    This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildChannelPositions(payload ModifyGuildChannelPositionsJSON, reason *string) error {
	u, err := url.Parse(fmt.Sprintf(modifyGuildChannelPositions, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return err
	}

	resp, err := Rest.Request(http.MethodPatch, u.String(), payload, reason)
	if err != nil {
		logging.Errorln(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}

// ModifyGuildChannelPositionsJSON - JSON payload
type ModifyGuildChannelPositionsJSON struct {
	ID              Snowflake  `json:"id"`               // channel id
	Position        *uint64    `json:"position"`         // sorting position of the channel
	LockPermissions *bool      `json:"lock_permissions"` // syncs the permission overwrites with the new parent, if moving to a new category
	ParentID        *Snowflake `json:"parent_id"`        // the new parent ID for the channel that is moved
}

// ListActiveThreads - Returns all active threads in the guild, including public and private threads. Threads are ordered by their id, in descending order.
func (g *Guild) ListActiveThreads() (*ThreadListResponse, error) {
	u, err := url.Parse(fmt.Sprintf(listActiveThreads, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var threadListResponse *ThreadListResponse
	err = json.NewDecoder(resp.Body).Decode(&threadListResponse)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return threadListResponse, nil
}

// GetGuildMember - Returns a GuildMember object for the specified User.
func (g *Guild) GetGuildMember(userID Snowflake) (*GuildMember, error) {
	u, err := url.Parse(fmt.Sprintf(getGuildMember, api, g.ID.String(), userID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guildMember *GuildMember
	err = json.NewDecoder(resp.Body).Decode(&guildMember)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guildMember, nil
}

// ListGuildMembers - Returns a list of guild member objects that are members of the guild.
//
// This endpoint is restricted according to whether the GuildMembers Privileged Intent is enabled for your application.
//
//     All parameters to this endpoint are optional
func (g *Guild) ListGuildMembers(limit *uint64, after *Snowflake) ([]GuildMember, error) {
	u, err := url.Parse(fmt.Sprintf(listGuildMembers, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	q := u.Query()
	if after != nil {
		q.Set("after", after.String())
	}
	if limit != nil {
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guildMembers []GuildMember
	err = json.NewDecoder(resp.Body).Decode(&guildMembers)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guildMembers, nil
}

// SearchGuildMembers - Returns a list of GuildMember objects whose username or nickname starts with a provided string.
//
//    All parameters to this endpoint except for `query` are optional
func (g *Guild) SearchGuildMembers(query string, limit *uint64) ([]GuildMember, error) {
	u, err := url.Parse(fmt.Sprintf(searchGuildMembers, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	q := u.Query()
	q.Set("query", query)
	if limit != nil {
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	u.RawQuery = q.Encode()

	resp, err := Rest.Request(http.MethodGet, u.String(), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var guildMembers []GuildMember
	err = json.NewDecoder(resp.Body).Decode(&guildMembers)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return guildMembers, nil
}

// TODO: Left off here https://discord.com/developers/docs/resources/guild#add-guild-member

// AddGuildMemberRole - Adds a role to a guild member.
//
// Requires the ManageRoles permission.
//
// Returns a 204 empty response on success.
//
// Fires a Guild Member Update Gateway event.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) AddGuildMemberRole(user *User, role *Snowflake) (method string, route string) {
	return http.MethodPut, fmt.Sprintf(addGuildMemberRole, api, g.ID.String(), user.ID.String(), role.String())
}

// RemoveGuildMemberRole - Removes a role from a guild member.
//
// Requires the ManageRoles permission.
//
// Returns a 204 empty response on success.
//
// Fires a GuildMember Update Gateway event.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) RemoveGuildMemberRole(user *User, role *Snowflake) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf(removeGuildMemberRole, api, g.ID.String(), user.ID.String(), role.String())
}
