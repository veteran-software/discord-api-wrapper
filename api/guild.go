package api

import (
	"fmt"
	"net/http"
	"time"
)

/* GUILD OBJECT */

type Guild struct {
	ID                          Snowflake                       `json:"id"`
	Name                        string                          `json:"name"`
	Icon                        *string                         `json:"icon"`
	IconHash                    *string                         `json:"icon_hash,omitempty"`
	Splash                      *string                         `json:"splash,omitempty"`
	DiscoverySplash             *string                         `json:"discovery_splash"`
	OwnerID                     Snowflake                       `json:"owner_id"`
	AfkChannelID                Snowflake                       `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int64                           `json:"afk_timeout"`
	WidgetEnabled               bool                            `json:"widget_enabled,omitempty"`
	WidgetChannelID             *Snowflake                      `json:"widget_channel_id,omitempty"`
	VerificationLevel           VerificationLevel               `json:"verification_level"`
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`
	Roles                       []Role                          `json:"roles"`
	Emojis                      []Emoji                         `json:"emojis"`
	Features                    []GuildFeatures                 `json:"features"`
	MfaLevel                    MfaLevel                        `json:"mfa_level"`
	ApplicationID               *Snowflake                      `json:"application_id"`
	SystemChannelID             *Snowflake                      `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags"`
	RulesChannelID              *Snowflake                      `json:"rules_channel_id"`
	MaxPresences                *int64                          `json:"max_presences,omitempty"`
	MaxMembers                  int64                           `json:"max_members,omitempty"`
	VanityUrlCode               *string                         `json:"vanity_url_code"`
	Description                 *string                         `json:"description"`
	Banner                      *string                         `json:"banner"`
	PremiumTier                 PremiumTier                     `json:"premium_tier"`
	PremiumSubscriptionCount    uint64                          `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string                          `json:"preferred_locale"`
	PublicUpdatesChannelID      *Snowflake                      `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        uint64                          `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      uint64                          `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    uint64                          `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               WelcomeScreen                   `json:"welcome_screen,omitempty"`
	NsfwLevel                   GuildNsfwLevel                  `json:"nsfw_level"`
	// TODO: Stickers
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`

	// These fields are only sent within the GUILD_CREATE event

	JoinedAt       time.Time             `json:"joined_at,omitempty"`
	Large          bool                  `json:"large,omitempty"`
	Unavailable    bool                  `json:"unavailable,omitempty"`
	MemberCount    int64                 `json:"member_count,omitempty"`
	VoiceStates    []VoiceState          `json:"voice_states,omitempty"`
	Members        []GuildMember         `json:"members,omitempty"`
	Channel        []Channel             `json:"channel,omitempty"`
	Threads        []Channel             `json:"threads,omitempty"`
	Presences      []PresenceUpdateEvent `json:"presences,omitempty"`
	StageInstances []StageInstance       `json:"stage_instances,omitempty"`

	// These fields are only sent when using the GET Current user Guilds endpoint and are relative to the requested user

	Owner       bool   `json:"owner,omitempty"`
	Permissions string `json:"permissions,omitempty"`
}

type DefaultMessageNotificationLevel int

const (
	AllMessages DefaultMessageNotificationLevel = iota
	OnlyMentions
)

type ExplicitContentFilterLevel int

const (
	Disabled ExplicitContentFilterLevel = iota
	MembersWithoutRoles
	AllMembers
)

type MfaLevel int

const (
	MfaNone MfaLevel = iota
	MfaElevated
)

type VerificationLevel int

const (
	VerificationLevelNone VerificationLevel = iota
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)

type GuildNsfwLevel int

const (
	NsfwDefault GuildNsfwLevel = iota
	NsfwExplicit
	NsfwSafe
	NsfwAgeRestricted
)

type PremiumTier int

const (
	PremiumNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

type SystemChannelFlags int

const (
	SuppressJoinNotifications          SystemChannelFlags = 1 << 0
	SuppressPremiumSubscriptions       SystemChannelFlags = 1 << 1
	SuppressGuildReminderNotifications SystemChannelFlags = 1 << 2
	SuppressJoinNotificationReplies    SystemChannelFlags = 1 << 3
)

type GuildFeatures string

const (
	AnimatedIcon                  GuildFeatures = "ANIMATED_ICON"
	Banner                        GuildFeatures = "BANNER"
	Commerce                      GuildFeatures = "COMMERCE"
	Community                     GuildFeatures = "COMMUNITY"
	Discoverable                  GuildFeatures = "DISCOVERABLE"
	Featurable                    GuildFeatures = "FEATURABLE"
	InviteSplash                  GuildFeatures = "INVITE_SPLASH"
	MemberVerificationGateEnabled GuildFeatures = "MEMBER_VERIFICATION_GATE_ENABLED"
	MonetizationEnabled           GuildFeatures = "MONETIZATION_ENABLED"
	MoreStickers                  GuildFeatures = "MORE_STICKERS"
	News                          GuildFeatures = "NEWS"
	Partnered                     GuildFeatures = "PARTNERED"
	PreviewEnabled                GuildFeatures = "PREVIEW_ENABLED"
	PrivateThreads                GuildFeatures = "PRIVATE_THREADS"
	RoleIcons                     GuildFeatures = "ROLE_ICONS"
	SevenDayThreadArchive         GuildFeatures = "SEVEN_DAY_THREAD_ARCHIVE"
	ThreeDayThreadArchive         GuildFeatures = "THREE_DAY_THREAD_ARCHIVE"
	TicketedEventsEnabled         GuildFeatures = "TICKETED_EVENTS_ENABLED"
	VanityURL                     GuildFeatures = "VANITY_URL"
	Verified                      GuildFeatures = "VERIFIED"
	VipRegions                    GuildFeatures = "VIP_REGIONS"
	WelcomeScreenEnabled          GuildFeatures = "WELCOME_SCREEN_ENABLED"
)

/* GUILD WIDGET OBJECT */

type GuildWidget struct {
	Enabled   bool       `json:"enabled"`
	ChannelID *Snowflake `json:"channel_id"`
}

/* GUILD MEMBER OBJECT */

/*
GuildMember

Avatar: guild specific avatar

CommunicationDisabledUntil: when the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out

User: the user this guild member represents

Nick: the users' guild nickname

Roles: array of GuildRole id's

JoinedAt: when the user joined the guild

PremiumSince: when the user started boosting the guild

Deaf: whether the user is deafened in voice channels

Mute: whether the user is muted in voice channels

Pending: whether the user has not yet passed the guild's Membership Screening requirements

Permissions: total permissions of the member in the channel, including overwrites, returned when in the interaction object
*/
type GuildMember struct {
	Avatar                     *string     `json:"avatar,omitempty"`
	CommunicationDisabledUntil *time.Time  `json:"communication_disabled_until,omitempty"`
	Deaf                       bool        `json:"deaf"`
	JoinedAt                   time.Time   `json:"joined_at"`
	Mute                       bool        `json:"mute"`
	Nick                       *string     `json:"nick,omitempty"`
	Pending                    bool        `json:"pending,omitempty"`
	Permissions                string      `json:"permissions,omitempty"`
	PremiumSince               *time.Time  `json:"premium_since,omitempty"`
	Roles                      []Snowflake `json:"roles"`
	User                       User        `json:"user,omitempty"`
}

/* INTEGRATION OBJECT */

type Integration struct {
	ID                Snowflake                 `json:"id"`
	Name              string                    `json:"name"`
	Type              string                    `json:"type"`
	Enabled           bool                      `json:"enabled"`
	Syncing           bool                      `json:"syncing,omitempty"`
	RoleID            Snowflake                 `json:"role_id,omitempty"`
	EnableEmoticons   bool                      `json:"enable_emoticons,omitempty"`
	ExpireBehavior    IntegrationExpireBehavior `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int                       `json:"expire_grace_period,omitempty"`
	User              User                      `json:"user,omitempty"`
	Account           IntegrationAccount        `json:"account"`
	SyncedAt          time.Time                 `json:"synced_at,omitempty"`
	SubscriberCount   int                       `json:"subscriber_count,omitempty"`
	Revoked           bool                      `json:"revoked,omitempty"`
	Application       IntegrationApplication    `json:"application,omitempty"`
}

type IntegrationExpireBehavior int

const (
	RemoveRole IntegrationExpireBehavior = iota
	Kick
)

/* INTEGRATION ACCOUNT OBJECT */

type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* INTEGRATION APPLICATION OBJECT */

type IntegrationApplication struct {
	ID          Snowflake `json:"id"`
	Name        string    `json:"name"`
	Icon        *string   `json:"icon"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	Bot         User      `json:"bot,omitempty"`
}

type Ban struct {
	Reason *string `json:"reason"`
	User   User    `json:"user"`
}

/* WELCOME SCREEN OBJECT */

type WelcomeScreen struct {
	Description     *string                `json:"description,omitempty"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels,omitempty"`
}

type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id,omitempty"`
	Description string     `json:"description,omitempty"`
	EmojiID     *Snowflake `json:"emoji_id,omitempty"`
	EmojiName   *string    `json:"emoji_name,omitempty"`
}

/* HELPER FUNCTIONS */
func (g *Guild) String() string {
	return g.Name + "(" + g.ID.String() + ")"
}

/* ENDPOINTS */

// GetGuild
// Returns the guild object for the given id.
//
// If with_counts is set to true, this endpoint will also return approximate_member_count and approximate_presence_count for the guild.
func (g *Guild) GetGuild() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/guilds/%s?with_counts=true", api, g.ID.String())
}

// ListGuildMembers
// Returns a list of guild member objects that are members of the guild.
//
// This endpoint is restricted according to whether the GUILD_MEMBERS Privileged Intent is enabled for your application.
func (g *Guild) ListGuildMembers(after ...*Snowflake) (method string, route string) {
	var afterSnowflake string

	if len(after) != 0 && after[0].String() != "" {
		afterSnowflake = "&after=" + after[0].String()
	} else {
		afterSnowflake = ""
	}

	return http.MethodGet, fmt.Sprintf("%s/guilds/%s/members?limit=1000%s", api, g.ID.String(), afterSnowflake)
}

// AddGuildMemberRole
//
// Adds a role to a guild member.
//
// Requires the MANAGE_ROLES permission.
//
// Returns a 204 empty response on success.
//
// Fires a Guild Member Update Gateway event.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) AddGuildMemberRole(user *User, role *Snowflake) (method string, route string) {
	return http.MethodPut, fmt.Sprintf("%s/guilds/%s/members/%s/roles/%s", api, g.ID.String(), user.ID.String(), role.String())
}

// RemoveGuildMemberRole
//
// Removes a role from a guild member.
//
// Requires the MANAGE_ROLES permission.
//
// Returns a 204 empty response on success.
//
// Fires a Guild Member Update Gateway event.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) RemoveGuildMemberRole(user *User, role *Snowflake) (method string, route string) {
	return http.MethodDelete, fmt.Sprintf("%s/guilds/%s/members/%s/roles/%s", api, g.ID.String(), user.ID.String(), role.String())
}
