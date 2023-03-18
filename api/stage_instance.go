/*
 * Copyright (c) 2022-2023. Veteran Software
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

// StageInstance - A StageInstance holds information about a live stage.
type StageInstance struct {
	ID                    Snowflake    `json:"id"`                       // The id of this Stage instance
	GuildID               Snowflake    `json:"guild_id"`                 // The guild id of the associated Stage channel
	ChannelID             Snowflake    `json:"channel_id"`               // The id of the associated Stage channel
	Topic                 string       `json:"topic"`                    // The topic of the Stage instance (1-120 characters)
	PrivacyLevel          PrivacyLevel `json:"privacy_level"`            // The privacy level of the Stage instance
	GuildScheduledEventID *Snowflake   `json:"guild_scheduled_event_id"` // The id of the scheduled event for this Stage instance
}

// PrivacyLevel - The privacy level of the Stage instance
type PrivacyLevel int

//goland:noinspection GoUnusedConst
const (
	GuildOnly PrivacyLevel = iota + 2 // The Stage instance is visible to only guild members.
)
