package api

type PrivacyLevel int

const (
	_ PrivacyLevel = iota
	Public
	GuildOnly
)

type StageInstance struct {
	ID           Snowflake    `json:"id"`
	GuildID      Snowflake    `json:"guild_id"`
	ChannelID    Snowflake    `json:"channel_id"`
	Topic        string       `json:"topic"`
	PrivacyLevel PrivacyLevel `json:"privacy_level"`
}
