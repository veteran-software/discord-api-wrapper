/*
 * Copyright (c) 2023. Veteran Software
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

	utils "github.com/veteran-software/discord-api-wrapper/v10/utilities"
)

// ListAutoModerationRulesForGuild - Get a list of all rules currently configured for the guild. Returns a list of auto moderation rule objects for the given guild.
//
// This endpoint requires the ManageGuild permission.
//
//goland:noinspection GoUnusedExportedFunction
func ListAutoModerationRulesForGuild(guildID string,
	channel *Channel,
	userID *Snowflake) ([]*AutoModerationRule,
	error) {
	g := &Guild{ID: *StringToSnowflake(guildID)}

	member, err := g.GetGuildMember(userID)
	if err != nil {
		return nil, err
	}

	if !CanManageGuild(member, channel) {
		return nil, errors.New(utils.ManageGuildPermissionsAreRequired)
	}

	u := parseRoute(fmt.Sprintf(listAutoModerationRulesForGuild, api, guildID))

	var rules []*AutoModerationRule
	err = json.Unmarshal(fireGetRequest(u, nil, nil), &rules)

	return rules, err
}

// GetAutoModerationRule - Get a single rule. Returns an auto moderation rule object.
//
// This endpoint requires the ManageGuild permission.
//
//goland:noinspection GoUnusedExportedFunction
func GetAutoModerationRule(guildID string, channel *Channel, userID, ruleID *Snowflake) (*AutoModerationRule,
	error) {
	g := &Guild{ID: *StringToSnowflake(guildID)}

	member, err := g.GetGuildMember(userID)
	if err != nil {
		return nil, err
	}

	if !CanManageGuild(member, channel) {
		return nil, errors.New(utils.ManageGuildPermissionsAreRequired)
	}

	u := parseRoute(fmt.Sprintf(getAutoModerationRule, api, guildID, ruleID.String()))

	var rule *AutoModerationRule
	err = json.Unmarshal(fireGetRequest(u, nil, nil), &rule)

	return rule, err
}

// CreateAutoModerationRule - Create a new rule. Returns an auto moderation rule on success. Fires an Auto Moderation Rule Create Gateway event.
//
// This endpoint requires the ManageGuild permission.
//
//goland:noinspection GoUnusedExportedFunction
func CreateAutoModerationRule(guildID string,
	channel *Channel,
	userID,
	ruleID *Snowflake,
	payload AutoModerationRuleJSON,
	reason *string) (*AutoModerationRule, error) {
	g := &Guild{ID: *StringToSnowflake(guildID)}

	member, err := g.GetGuildMember(userID)
	if err != nil {
		return nil, err
	}

	if !CanManageGuild(member, channel) {
		return nil, errors.New(utils.ManageGuildPermissionsAreRequired)
	}

	u := parseRoute(fmt.Sprintf(getAutoModerationRule, api, guildID, ruleID.String()))

	var rule *AutoModerationRule
	err = json.Unmarshal(firePostRequest(u, payload, reason), &rule)

	return rule, err
}

// AutoModerationRuleJSON - JSON payload for AutoMod actions
type AutoModerationRuleJSON struct {
	Name            string           `json:"name"`                       // the rule name
	EventType       EventType        `json:"event_type"`                 // the event type
	TriggerType     TriggerType      `json:"trigger_type,omitempty"`     // the trigger metadata; only req for CreateAutoModerationRule
	TriggerMetadata TriggerMetadata  `json:"trigger_metadata,omitempty"` // the actions which will execute when the rule is triggered
	Actions         []*AutoModAction `json:"actions"`                    // whether the rule is enabled
	Enabled         bool             `json:"enabled,omitempty"`
	ExemptRoles     []*Snowflake     `json:"exempt_roles,omitempty"`
	ExemptChannels  []*Snowflake     `json:"exempt_channels,omitempty"`
}

// ModifyAutoModerationRule - Modify an existing rule. Returns an auto moderation rule on success. Fires an Auto Moderation Rule Update Gateway event.
//
// This endpoint requires the ManageGuild permission.
//
//goland:noinspection GoUnusedExportedFunction
func ModifyAutoModerationRule(guildID string,
	channel *Channel,
	userID, ruleID *Snowflake,
	payload AutoModerationRuleJSON,
	reason *string) (*AutoModerationRule, error) {
	g := &Guild{ID: *StringToSnowflake(guildID)}

	member, err := g.GetGuildMember(userID)
	if err != nil {
		return nil, err
	}

	if !CanManageGuild(member, channel) {
		return nil, errors.New(utils.ManageGuildPermissionsAreRequired)
	}

	u := parseRoute(fmt.Sprintf(modifyAutoModerationRule, api, guildID, ruleID.String()))

	var rule *AutoModerationRule
	err = json.Unmarshal(firePatchRequest(u, payload, reason), &rule)

	return rule, err
}

// DeleteAutoModerationRule - Delete a rule. Returns a 204 on success. Fires an AutoModerationRuleDelete Gateway event.
//
// This endpoint requires the ManageGuild permission.
//
//goland:noinspection GoUnusedExportedFunction
func DeleteAutoModerationRule(guildID string,
	channel *Channel,
	userID,
	ruleID *Snowflake,
	reason *string) (*AutoModerationRule, error) {
	g := &Guild{ID: *StringToSnowflake(guildID)}

	member, err := g.GetGuildMember(userID)
	if err != nil {
		return nil, err
	}

	if !CanManageGuild(member, channel) {
		return nil, errors.New(utils.ManageGuildPermissionsAreRequired)
	}

	u := parseRoute(fmt.Sprintf(deleteAutoModerationRule, api, guildID, ruleID.String()))

	var rule *AutoModerationRule
	err = json.Unmarshal(firePatchRequest(u, nil, reason), &rule)

	return rule, err
}
