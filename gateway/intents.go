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

package gateway

// Intents - Maintaining a stateful application can be difficult when it comes to the amount of data you're expected to process, especially at scale.
//
// Gateway Intents are a system to help you lower that computational burden.
//
// When identifying to the gateway, you can specify an intents parameter which allows you to conditionally subscribe to pre-defined "intents", groups of events defined by Discord.
//
// If you do not specify a certain intent, you will not receive any of the gateway events that are batched into that group.
type Intents uint64

//goland:noinspection GoUnusedConst
const (
	Guilds                      Intents = 1 << 0
	GuildMembers                Intents = 1 << 1
	GuildBans                   Intents = 1 << 2
	GuildEmojisAndStickers      Intents = 1 << 3
	GuildIntegrations           Intents = 1 << 4
	GuildWebhooks               Intents = 1 << 5
	GuildInvites                Intents = 1 << 6
	GuildVoiceStates            Intents = 1 << 7
	GuildPresences              Intents = 1 << 8
	GuildMessages               Intents = 1 << 9
	GuildMessageReactions       Intents = 1 << 10
	GuildMessageTyping          Intents = 1 << 11
	DirectMessages              Intents = 1 << 12
	DirectMessageReactions      Intents = 1 << 13
	DirectMessageTyping         Intents = 1 << 14
	MessageContent              Intents = 1 << 15
	GuildScheduleEvents         Intents = 1 << 16
	AutoModerationConfiguration Intents = 1 << 20
	AutoModerationExecution     Intents = 1 << 21
)
