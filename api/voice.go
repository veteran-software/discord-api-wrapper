package api

import (
	"fmt"
	"net/http"
	"time"
)

/* VOICE STATE OBJECT */

type VoiceState struct {
	GuildID                 Snowflake   `json:"guild_id,omitempty"`
	ChannelID               *Snowflake  `json:"channel_id"`
	UserID                  Snowflake   `json:"user_id"`
	Member                  GuildMember `json:"member,omitempty"`
	SessionID               string      `json:"session_id"`
	Deaf                    bool        `json:"deaf"`
	Mute                    bool        `json:"mute"`
	SelfDeaf                bool        `json:"self_deaf"`
	SelfMute                bool        `json:"self_mute"`
	SelfStream              bool        `json:"self_stream,omitempty"`
	SelfVideo               bool        `json:"self_video"`
	Suppress                bool        `json:"suppress"`
	RequestToSpeakTimestamp *time.Time  `json:"request_to_speak_timestamp"`
}

/* VOICE REGION OBJECT */

type VoiceRegion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}

/* ENDPOINTS */

func ListVoiceRegions() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/voice/regions", api)
}
