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
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/veteran-software/discord-api-wrapper/v10/logging"
)

/* Whenever an admin action is performed on the API, an entry is added to the respective guild's audit log.
You can specify the reason by attaching the X-Audit-Log-Reason request header.
This header supports url encoded utf8 characters.
*/

// AuditLog - Whenever an admin action is performed on the API, an entry is added to the respective guild's audit log.
//
// You can specify the reason by attaching the "X-Audit-Log-Reason" request header.
//
// This header supports url encoded utf8 characters.
type AuditLog struct {
	AuditLogEntries      []AuditLogEntry       `json:"audit_log_entries"`      // list of audit log entries
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events"` // list of GuildScheduledEvent found in the audit log
	Integrations         []Integration         `json:"integrations"`           // list of partial integration objects
	Threads              []Channel             `json:"threads"`                // list of threads found in the audit log
	Users                []User                `json:"users"`                  // list of users found in the audit log
	Webhooks             []Webhook             `json:"webhooks"`               // list of webhooks found in the audit log
}

// AuditLogEntry - Representation of a single Audit Log
type AuditLogEntry struct {
	TargetID   *string            `json:"target_id"`         // id of the affected entity (webhook, user, role, etc.)
	Changes    []AuditLogChange   `json:"changes,omitempty"` // changes made to the target_id
	UserID     *Snowflake         `json:"user_id"`           // the user who made the changes
	ID         Snowflake          `json:"id"`                // id of the entry
	ActionType AuditLogEvent      `json:"action_type"`       // type of action that occurred
	Options    OptionalAuditEntry `json:"options,omitempty"` // additional info for certain action types
	Reason     string             `json:"reason,omitempty"`  // the reason for the change (0-512 characters)
}

// AuditLogEvent - The event type that triggered the log action
type AuditLogEvent int

//goland:noinspection GoUnusedConst
const (
	// GuildUpdate - Guild update Events
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

	GuildScheduledEventCreate AuditLogEvent = iota + 59
	GuildScheduledEventUpdate
	GuildScheduledEventDelete

	/* Thread Events */

	ThreadCreate AuditLogEvent = iota + 66
	ThreadUpdate
	ThreadDelete

	/* Application Command Events */

	ApplicationCommandPermissionUpdate AuditLogEvent = iota + 74
)

// OptionalAuditEntry - Information that is specific to certain events
type OptionalAuditEntry struct {
	ChannelID        Snowflake `json:"channel_id"`         // channel in which the entities were targeted
	Count            string    `json:"count"`              // number of entities that were targeted
	DeleteMemberDays string    `json:"delete_member_days"` // number of days after which inactive members were kicked
	ID               Snowflake `json:"id"`                 // id of the overwritten entity
	MembersRemoved   string    `json:"members_removed"`    // number of members removed by the prune
	MessageID        Snowflake `json:"message_id"`         // id of the message that was targeted
	RoleName         string    `json:"role_name"`          // name of the role if type is "0" (not present if type is "1")
	Type             string    `json:"type"`               // type of overwritten entity - "0" for "role" or "1" for "member"
}

// AuditLogChange - If new_value is not present in the change object, while old_value is, that means the property that was changed has been reset, or set to null
type AuditLogChange struct {
	NewValue interface{} `json:"new_value,omitempty"` // new value of the key
	OldValue interface{} `json:"old_value,omitempty"` // old value of the key
	Key      string      `json:"key"`                 // name of audit log change key
}

// GetGuildAuditLog - Returns an audit log object for the guild.
//
// Requires the ViewAuditLog permission.
func (g *Guild) GetGuildAuditLog(userID *Snowflake, actionType *uint64, before *Snowflake, limit *uint64) (*AuditLog, error) {
	u, err := url.Parse(fmt.Sprintf(getGuildAuditLog, api, g.ID.String()))
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	// Set the optional qsp
	q := u.Query()
	if userID != nil {
		q.Set("user_id", userID.String())
	}
	if actionType != nil {
		q.Set("action_type", strconv.FormatUint(*actionType, 10))
	}
	if before != nil {
		q.Set("before", before.String())
	}
	if limit != nil {
		if *limit >= 1 && *limit <= 100 {
			q.Set("limit", strconv.FormatUint(*limit, 10))
		} else {
			return nil, errors.New("the limit filter must be >= 1 && <= 100")
		}
	}

	// If there's any of the optional qsp present, encode and add to the URL
	if len(q) != 0 {
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

	var log AuditLog
	err = json.NewDecoder(resp.Body).Decode(&log)
	if err != nil {
		logging.Errorln(err)
		return nil, err
	}

	return &log, nil
}
