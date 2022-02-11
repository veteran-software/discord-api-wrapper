package api

import "time"

type GuildScheduledEvent struct {
	ID                 Snowflake                         `json:"id"`
	GuildID            Snowflake                         `json:"guild_id"`
	ChannelID          *Snowflake                        `json:"channel_id"`
	CreatorID          *Snowflake                        `json:"creator_id"`
	Name               string                            `json:"name"`
	Description        string                            `json:"description,omitempty"`
	ScheduledStartTime time.Time                         `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                        `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel   `json:"privacy_level"`
	Status             GuildScheduledEventStatus         `json:"status"`
	EntityType         GuildScheduledEventType           `json:"entity_type"`
	EntityID           *Snowflake                        `json:"entity_id"`
	EntityMetadata     GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            User                              `json:"creator,omitempty"`
	UserCount          int64                             `json:"user_count,omitempty"`
}

type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = iota + 2
)

type GuildScheduledEventType int

const (
	GuildScheduledEventTypeStageInstance GuildScheduledEventType = iota + 1
	GuildScheduledEventTypeVoice
	GuildScheduledEventTypeExternal
)

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = iota + 1
	GuildScheduledEventStatusActive
	GuildScheduledEventStatusCompleted
	GuildScheduledEventStatusCancelled
)

type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"`
}
