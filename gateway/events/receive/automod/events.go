/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package automod

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

/* AUTO MODERATION */

// All Auto Moderation related events are only sent to bot users which have the MANAGE_GUILD permission.
type (
	// AutoModerationRuleCreate - Sent when a rule is created. The inner payload is an auto moderation rule object.
	AutoModerationRuleCreate api.AutoModerationRule

	// AutoModerationRuleUpdate - Sent when a rule is updated. The inner payload is an auto moderation rule object.
	AutoModerationRuleUpdate api.AutoModerationRule

	// AutoModerationRuleDelete - Sent when a rule is deleted. The inner payload is an auto moderation rule object.
	AutoModerationRuleDelete api.AutoModerationRule

	// AutoModerationRuleExecution - Sent when a rule is triggered and an action is executed (e.g. when a message is blocked).
	AutoModerationRuleExecution struct {
		GuildID              api.Snowflake     `json:"guild_id"`
		Action               api.AutoModAction `json:"action"`
		RuleID               api.Snowflake     `json:"rule_id"`
		RuleTriggerType      api.TriggerType   `json:"rule_trigger_type"`
		UserID               api.Snowflake     `json:"user_id"`
		ChannelID            api.Snowflake     `json:"channel_id,omitempty"`
		MessageID            api.Snowflake     `json:"message_id"`
		AlertSystemMessageID api.Snowflake     `json:"alert_system_message_id"`
		Content              string            `json:"content,omitempty"` // MESSAGE_CONTENT (1 << 15) gateway intent is required to receive the `content` and `matched_content` fields
		MatchedKeyword       *string           `json:"matched_keyword"`
		MatchedContent       *string           `json:"matched_content"` // MESSAGE_CONTENT (1 << 15) gateway intent is required to receive the `content` and `matched_content` fields
	}
)
