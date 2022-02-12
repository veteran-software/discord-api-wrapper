package api

import (
	"fmt"
	"net/http"
)

/* Whenever an admin action is performed on the API, an entry is added to the respective guild's audit log.
You can specify the reason by attaching the X-Audit-Log-Reason request header.
This header supports url encoded utf8 characters.
*/

type AuditLog struct {
	AuditLogEntries      []AuditLogEntry       `json:"audit_log_entries"`
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events"`
	Integrations         []Integration         `json:"integrations"`
	Threads              []Channel             `json:"threads"`
	Users                []User                `json:"users"`
	// TODO: Webhooks []Webhook `json:"webhooks"`
}

type AuditLogEntry struct {
	TargetID   *string            `json:"target_id"`
	Changes    []AuditLogChange   `json:"changes,omitempty"`
	UserID     *Snowflake         `json:"user_id"`
	ID         Snowflake          `json:"id"`
	ActionType AuditLogEvent      `json:"action_type"`
	Options    OptionalAuditEntry `json:"options,omitempty"`
	Reason     string             `json:"reason,omitempty"`
}

type AuditLogEvent int

const (
	GuildUpdate AuditLogEvent = iota + 1

	/* Channel Events */

	ChannelCreate AuditLogEvent = iota + 9
	ChannelUpdate
	ChannelDelete
	ChannelOverwriteCreate
	ChannelOverwriteUpdate
	ChannelOverwriteDelete

	/* Member Events */

	MemberKick AuditLogEvent = iota + 13
	MemberPrune
	MemberBanAdd
	MemberBanRemove
	MemberUpdate
	MemberRoleUpdate
	MemberMove
	MemberDisconnect
	BotAdd

	/* Role Events */

	RoleCreate AuditLogEvent = iota + 14
	RoleUpdate
	RoleDelete

	/* Invite Events */

	InviteCreate AuditLogEvent = iota + 21
	InviteUpdate
	InviteDelete

	/* Webhook Events */

	WebhookCreate AuditLogEvent = iota + 28
	WebhookUpdate
	WebhookDelete

	/* Emoji Events */

	EmojiCreate AuditLogEvent = iota + 35
	EmojiUpdate
	EmojiDelete

	/* Message Events */

	MessageDelete AuditLogEvent = iota + 44
	MessageBulkDelete
	MessagePin
	MessageUnpin

	/* Integration & Stage Instance Events */

	IntegrationCreate AuditLogEvent = iota + 48
	IntegrationUpdate
	IntegrationDelete
	StageInstanceCreate
	StageInstanceUpdate
	StageInstanceDelete

	/* Sticker Events */

	StickerCreate AuditLogEvent = iota + 52
	StickerUpdate
	StickerDelete

	/* Guild Scheduled Event Events */

	GuildScheduledEventCreate = iota + 59
	GuildScheduledEventUpdate
	GuildScheduledEventDelete

	/* Thread Events */

	ThreadCreate AuditLogEvent = iota + 66
	ThreadUpdate
	ThreadDelete
)

type OptionalAuditEntry struct {
	ChannelID        Snowflake `json:"channel_id"`
	Count            string    `json:"count"`
	DeleteMemberDays string    `json:"delete_member_days"`
	ID               Snowflake `json:"id"`
	MembersRemoved   string    `json:"members_removed"`
	MessageID        Snowflake `json:"message_id"`
	RoleName         string    `json:"role_name"`
	Type             string    `json:"type"`
}

// AuditLogChange
// If new_value is not present in the change object, while old_value is, that means the property that was changed has
// been reset, or set to null
type AuditLogChange struct {
	NewValue interface{} `json:"new_value,omitempty"`
	OldValue interface{} `json:"old_value,omitempty"`
	Key      string      `json:"key"`
}

// GetGuildAuditLog
// Returns an audit log object for the guild. Requires the 'VIEW_AUDIT_LOG' permission.
func (g *Guild) GetGuildAuditLog() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/guilds/%s/audit-logs", api, g.ID.String())
}
