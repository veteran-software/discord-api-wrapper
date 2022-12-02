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

// AutoModerationRule - Auto Moderation is a feature which allows each guild to set up rules that trigger based on some criteria.
// For example, a rule can trigger whenever a message contains a specific keyword.
//
// Rules can be configured to automatically execute actions whenever they trigger.
// For example, if a user tries to send a message which contains a certain keyword, a rule can trigger and block the message before it is sent.
type AutoModerationRule struct {
	ID              Snowflake       `json:"id"`               // the id of this rule
	GuildID         Snowflake       `json:"guild_id"`         // the id of the guild which this rule belongs to
	Name            string          `json:"name"`             // the rule name
	CreatorID       Snowflake       `json:"creator_id"`       // the user which first created this rule
	EventType       EventType       `json:"event_type"`       // the rule event type
	TriggerType     TriggerType     `json:"trigger_type"`     // the rule trigger type
	TriggerMetadata TriggerMetadata `json:"trigger_metadata"` // the rule trigger metadata
	Actions         []string        `json:"actions"`          // the actions which will execute when the rule is triggered
	Enabled         bool            `json:"enabled"`          // whether the rule is enabled
	ExemptRoles     []Snowflake     `json:"exempt_roles"`     // the role ids that should not be affected by the rule (Maximum of 20)
	ExemptChannels  []Snowflake     `json:"exempt_channels"`  // the channel ids that should not be affected by the rule (Maximum of 50)
}

// TriggerType - Characterizes the type of content which can trigger the rule.
type TriggerType int

//goland:noinspection GoUnusedConst
const (
	Keyword       TriggerType = iota + 1 // check if content contains words from a user defined list of keywords
	Spam          TriggerType = iota + 2 // check if content represents generic spam
	KeywordPreset                        // check if content contains words from internal pre-defined wordsets
	MentionSpam                          // check if content contains more unique mentions than allowed
)

// TriggerMetadata - Additional data used to determine whether a rule should be triggered. Different fields are relevant based on the value of trigger_type.
type TriggerMetadata struct {
	KeywordFilter     []string            `json:"keyword_filter"`      // substrings which will be searched for in content (Maximum of 1000)
	RegexPatterns     []string            `json:"regex_patterns"`      // regular expression patterns which will be matched against content (Maximum of 10)
	Presets           []KeyWordPresetType `json:"presets"`             // the internally pre-defined wordsets which will be searched for in content
	AllowList         []string            `json:"allow_list"`          // substrings which will be exempt from triggering the preset trigger type (Maximum of 1000)
	MentionTotalLimit int                 `json:"mention_total_limit"` // total number of unique role and user mentions allowed per message (Maximum of 50)
}

// KeyWordPresetType - the internally pre-defined wordsets which will be searched for in content
type KeyWordPresetType int

//goland:noinspection GoUnusedConst
const (
	Profanity     KeyWordPresetType = iota + 1 // Words that may be considered forms of swearing or cursing
	SexualContent                              // Words that refer to sexually explicit behavior or activity
	Slurs                                      // Personal insults or words that may be considered hate speech
)

// EventType - Indicates in what event context a rule should be checked.
type EventType int

//goland:noinspection GoUnusedConst
const (
	MessageSend EventType = iota + 1 // when a member sends or edits a message in the guild
)

// AutoModAction - An action which will execute whenever a rule is triggered.
type AutoModAction struct {
	Type     AutoModerationActionType `json:"type"`     // the type of action
	Metadata string                   `json:"metadata"` // additional metadata needed during execution for this specific action type
}

// AutoModerationActionType - the type of action
type AutoModerationActionType int

//goland:noinspection GoUnusedConst
const (
	BlockMessage AutoModerationActionType = iota + 1
	SendMessage
	Timeout
)

// AutoModerationActionMetadata - Additional data used when an action is executed. Different fields are relevant based on the value of action type.
type AutoModerationActionMetadata struct {
	ChannelID       Snowflake `json:"channel_id"`
	DurationSeconds int       `json:"duration_seconds"`
}

// TODO: Endpoints
