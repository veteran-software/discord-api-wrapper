package api

/* Payloads */

// GatewayPayload
//
// s and t are null when op is not 0 (Gateway Dispatch Opcode).
type GatewayPayload struct {
	Op int          `json:"op"`          // opcode for the payload
	D  *interface{} `json:"d,omitempty"` // event data
	S  *int         `json:"s,omitempty"` // sequence number, used for resuming sessions and heartbeats
	T  *string      `json:"t,omitempty"` // the event name for this payload
}

/* Presence */

type PresenceStatus string

const (
	Idle    PresenceStatus = "idle"
	Dnd     PresenceStatus = "dnd"
	Online  PresenceStatus = "online"
	Offline PresenceStatus = "offline"
)

type ActivityTypes int8

const (
	Game ActivityTypes = iota
	Streaming
	Listening
	Watching
	Custom
	Competing
)

type ActivityTimestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type ActivityParty struct {
	ID   string  `json:"id,omitempty"`
	Size []int16 `json:"size,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type ActivityFlags int

const (
	Instance    ActivityFlags = 1 << 0
	Join        ActivityFlags = 1 << 1
	Spectate    ActivityFlags = 1 << 2
	JoinRequest ActivityFlags = 1 << 3
	Sync        ActivityFlags = 1 << 4
	Play        ActivityFlags = 1 << 5
)

type ActivityButtons struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Activity struct {
	Name          string             `json:"name"`
	Type          ActivityTypes      `json:"type"`
	URL           *string            `json:"url,omitempty"`
	CreatedAt     int64              `json:"created_at"`
	Timestamps    ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID Snowflake          `json:"application_id,omitempty"`
	Details       *string            `json:"details,omitempty"`
	State         *string            `json:"state,omitempty"`
	Emoji         *Emoji             `json:"emoji,omitempty"`
	Party         ActivityParty      `json:"party,omitempty"`
	Assets        ActivityAssets     `json:"assets,omitempty"`
	Secrets       ActivitySecrets    `json:"secrets,omitempty"`
	Instance      bool               `json:"instance,omitempty"`
	Flags         ActivityFlags      `json:"flags,omitempty"`
	Buttons       []ActivityButtons  `json:"buttons,omitempty"`
}

type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Web     string `json:"web,omitempty"`
}

type PresenceUpdateEvent struct {
	User         User           `json:"user"`
	GuildID      Snowflake      `json:"guild_id"`
	Status       PresenceStatus `json:"status"`
	Activities   []Activity     `json:"activities"`
	ClientStatus ClientStatus   `json:"client_status"`
}
