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
	AfkTimeout                  int                             `json:"afk_timeout"`                          // afk timeout in seconds
	WidgetEnabled               bool                            `json:"widget_enabled,omitempty"`             // true if the server widget is enabled
	WidgetChannelID             *Snowflake                      `json:"widget_channel_id,omitempty"`          // the channel id that the widget will generate an "invite" to, or null if set to no invite
	VerificationLevel           VerificationLevel               `json:"verification_level"`                   // verification level required for the guild
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"`        // default message notifications level
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`              // explicit content filter level
	Roles                       []*Role                         `json:"roles"`                                // roles in the guild
	Emojis                      []*Emoji                        `json:"emojis"`                               // custom guild emojis
	Features                    []*GuildFeatures                `json:"features"`                             // enabled guild features
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
	Stickers                    []*Sticker                      `json:"stickers,omitempty"`                   // custom guild stickers
	PremiumProgressBarEnabled   bool                            `json:"premium_progress_bar_enabled"`         // whether the guild has the boost progress bar enabled

	// TODO: These fields are only sent within the GUILD_CREATE event; move to Gateway.go

	JoinedAt             time.Time              `json:"joined_at,omitempty"`              // when this guild was joined at
	Large                bool                   `json:"large,omitempty"`                  // true if this is considered a large guild
	Unavailable          bool                   `json:"unavailable,omitempty"`            // true if this guild is unavailable due to an outage
	MemberCount          int64                  `json:"member_count,omitempty"`           // total number of members in this guild
	VoiceStates          []*VoiceState          `json:"voice_states,omitempty"`           // states of members currently in voice channels; lacks the guild_id key
	Members              []*GuildMember         `json:"members,omitempty"`                // users in the guild
	Channels             []*Channel             `json:"channels,omitempty"`               // channels in the guild
	Threads              []*Channel             `json:"threads,omitempty"`                // all active threads in the guild that current user has permission to view
	Presences            []*PresenceUpdateEvent `json:"presences,omitempty"`              // presences of the members in the guild, will only include non-offline members if the size is greater than large threshold
	StageInstances       []*StageInstance       `json:"stage_instances,omitempty"`        // Stage instances in the guild
	GuildScheduledEvents []*GuildScheduledEvent `json:"guild_scheduled_events,omitempty"` // the scheduled events in the guild

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
	SuppressJoinNotifications                           SystemChannelFlags = 1 << 0 // Suppress member join notifications
	SuppressPremiumSubscriptions                        SystemChannelFlags = 1 << 1 // Suppress server boost notifications
	SuppressGuildReminderNotifications                  SystemChannelFlags = 1 << 2 // Suppress server setup tips
	SuppressJoinNotificationReplies                     SystemChannelFlags = 1 << 3 // Hide member join sticker reply buttons
	SuppressRoleSubscriptionPurchaseNotifications       SystemChannelFlags = 1 << 4 // Suppress role subscription purchase and renewal notifications
	SuppressRoleSubscriptionPurchaseNotificationReplies SystemChannelFlags = 1 << 5 // Hide role subscription sticker reply buttons
)

// GuildFeatures - enabled guild features
type GuildFeatures string

//goland:noinspection SpellCheckingInspection,GrazieInspection,GoUnusedConst
const (
	AnimatedBanner                        GuildFeatures = "ANIMATED_BANNER"                           // guild has access to set an animated guild banner image
	AnimatedIcon                          GuildFeatures = "ANIMATED_ICON"                             // guild has access to set an animated guild icon
	ApplicationCommandPermissionsV2       GuildFeatures = "APPLICATION_COMMAND_PERMISSIONS_V2"        // guild is using the old permissions configuration behavior
	AutoModeration                        GuildFeatures = "AUTO_MODERATION"                           // guild has set up auto moderation rules
	Banner                                GuildFeatures = "BANNER"                                    // guild has access to set a guild banner image
	Community                             GuildFeatures = "COMMUNITY"                                 // Mutable; guild can enable welcome screen, Membership Screening, stage channels and discovery, and receives community updates
	CreatorMonetizableProvisional         GuildFeatures = "CREATOR_MONETIZABLE_PROVISIONAL"           // guild has enabled monetization
	CreatorStorePage                      GuildFeatures = "CREATOR_STORE_PAGE"                        // guild has enabled the role subscription promo page
	DeveloperSupportServer                GuildFeatures = "DEVELOPER_SUPPORT_SERVER"                  //	guild has been set as a support server on the App Directory
	Discoverable                          GuildFeatures = "DISCOVERABLE"                              // Mutable; guild is able to be discovered in the directory
	Featurable                            GuildFeatures = "FEATURABLE"                                // guild is able to be featured in the directory
	InvitesDisabled                       GuildFeatures = "INVITES_DISABLED"                          // Mutable; Pauses all invites/access to the server
	InviteSplash                          GuildFeatures = "INVITE_SPLASH"                             // guild has access to set an invite splash background
	MemberVerificationGateEnabled         GuildFeatures = "MEMBER_VERIFICATION_GATE_ENABLED"          // guild has enabled Membership Screening
	MonetizationEnabled                   GuildFeatures = "MONETIZATION_ENABLED"                      // guild has enabled monetization
	MoreStickers                          GuildFeatures = "MORE_STICKERS"                             // guild has increased custom sticker slots
	News                                  GuildFeatures = "NEWS"                                      // guild has access to create news channels
	Partnered                             GuildFeatures = "PARTNERED"                                 // guild is partnered
	PreviewEnabled                        GuildFeatures = "PREVIEW_ENABLED"                           // guild can be previewed before joining via Membership Screening or the directory
	RoleIcons                             GuildFeatures = "ROLE_ICONS"                                // guild is able to set role icons
	RoleSubscriptionsAvailableForPurchase GuildFeatures = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE" // guild has role subscriptions that can be purchased
	RoleSubscriptionsEnabled              GuildFeatures = "ROLE_SUBSCRIPTIONS_ENABLED"                // guild has enabled role subscriptions
	TicketedEventsEnabled                 GuildFeatures = "TICKETED_EVENTS_ENABLED"                   // guild has enabled ticketed events
	VanityURL                             GuildFeatures = "VANITY_URL"                                // guild has access to set a vanity URL
	Verified                              GuildFeatures = "VERIFIED"                                  // guild is verified
	VipRegions                            GuildFeatures = "VIP_REGIONS"                               // guild has access to set 384kbps bitrate in voice (previously VIP voice servers)
	WelcomeScreenEnabled                  GuildFeatures = "WELCOME_SCREEN_ENABLED"                    // guild has enabled the welcome screen
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
	ID                       Snowflake        `json:"id"`                         // guild id
	Name                     string           `json:"name"`                       // guild name (2-100 characters)
	Icon                     *string          `json:"icon"`                       // icon hash
	Splash                   *string          `json:"splash"`                     // splash hash
	DiscoverySplash          *string          `json:"discovery_splash"`           // discovery splash hash
	Emojis                   []*Emoji         `json:"emojis"`                     // custom guild emojis
	Features                 []*GuildFeatures `json:"features"`                   // enabled guild features
	ApproximateMemberCount   int              `json:"approximate_member_count"`   // approximate number of members in this guild
	ApproximatePresenceCount int              `json:"approximate_presence_count"` // approximate number of online members in this guild
	Description              *string          `json:"description"`                // the description for the guild, if the guild is discoverable
	Stickers                 []*Sticker       `json:"stickers"`                   // custom guild stickers
}

// GuildWidgetSettings - the guild widget status
type GuildWidgetSettings struct {
	Enabled   bool       `json:"enabled"`    // whether the widget is enabled
	ChannelID *Snowflake `json:"channel_id"` // the widget channel id
}

// GetGuildWidget - the guild widget
//
// The fields `id`, `discriminator` and `avatar` are anonymized to prevent abuse.
type GetGuildWidget struct {
	ID            Snowflake      `json:"id"`             // guild id
	Name          string         `json:"name"`           // guild name (2-100 characters)
	InstantInvite *string        `json:"instant_invite"` // instant invite for the guilds specified widget invite channel
	Channels      []*Channel     `json:"channels"`       // voice and stage channels which are accessible by @everyone
	Members       []*GuildMember `json:"members"`        // special widget user objects that includes users presence (Limit 100)
	PresenceCount int            `json:"presence_count"` // number of online members in this guild
}

// GuildMember - Represents a member of a Guild
//
//	The field `user` won't be included in the member object attached to MessageCreate and MessageUpdate gateway events.
//
//	In GUILD_ events, pending will always be included as true or false.
//	In non GUILD_ events which can only be triggered by non-pending users, pending will not be included.
type GuildMember struct {
	User                       User         `json:"user,omitempty"`                         // User - the user this guild member represents
	Nick                       *string      `json:"nick,omitempty"`                         // Nick - the users' guild nickname
	Avatar                     *string      `json:"avatar,omitempty"`                       // Avatar - guild specific avatar
	Roles                      []*Snowflake `json:"roles"`                                  // Roles - array of GuildRole id's
	JoinedAt                   time.Time    `json:"joined_at"`                              // JoinedAt - when the user joined the guild
	PremiumSince               *time.Time   `json:"premium_since,omitempty"`                // PremiumSince - when the user started boosting the guild
	Deaf                       bool         `json:"deaf"`                                   // Deaf - whether the user is deafened in voice channels
	Mute                       bool         `json:"mute"`                                   // Mute - whether the user is muted in voice channels
	Pending                    bool         `json:"pending,omitempty"`                      // Pending - whether the user has not yet passed the guild's Membership Screening requirements
	CommunicationDisabledUntil *time.Time   `json:"communication_disabled_until,omitempty"` // CommunicationDisabledUntil - when the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out

	// Undocumented as of 12/3/2022
	Flags     UserFlags `json:"flags,omitempty"`
	IsPending bool      `json:"is_pending,omitempty"`
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
	Scopes            []string                  `json:"scopes,omitempty"`              // the scopes the application has been authorized for
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
	Description     *string                 `json:"description,omitempty"`      // the server description shown in the welcome screen
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels,omitempty"` // the channels shown in the welcome screen, up to 5
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
