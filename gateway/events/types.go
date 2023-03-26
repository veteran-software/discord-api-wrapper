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

package events

type RawType string
type Type map[RawType]struct{}

//goland:noinspection GoUnusedConst
const (
	Ready                               RawType = "READY"
	Resumed                             RawType = "RESUMED"
	ApplicationCommandPermissionsUpdate RawType = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	AutoModerationRuleCreate            RawType = "AUTO_MODERATION_RULE_CREATE"
	AutoModerationRuleUpdate            RawType = "AUTO_MODERATION_RULE_UPDATE"
	AutoModerationRuleDelete            RawType = "AUTO_MODERATION_RULE_DELETE"
	AutoModerationActionExecution       RawType = "AUTO_MODERATION_ACTION_EXECUTION"
	ChannelCreate                       RawType = "CHANNEL_CREATE"
	ChannelUpdate                       RawType = "CHANNEL_UPDATE"
	ChannelDelete                       RawType = "CHANNEL_DELETE"
	ChannelPinsUpdate                   RawType = "CHANNEL_PINS_UPDATE"
	ThreadCreate                        RawType = "THREAD_CREATE"
	ThreadUpdate                        RawType = "THREAD_UPDATE"
	ThreadDelete                        RawType = "THREAD_DELETE"
	ThreadListSync                      RawType = "THREAD_LIST_SYNC"
	ThreadMemberUpdate                  RawType = "THREAD_MEMBER_UPDATE"
	ThreadMembersUpdate                 RawType = "THREAD_MEMBERS_UPDATE"
	GuildCreate                         RawType = "GUILD_CREATE"
	GuildUpdate                         RawType = "GUILD_UPDATE"
	GuildDelete                         RawType = "GUILD_DELETE"
	GuildAuditLogEntryCreate            RawType = "GUILD_AUDIT_LOG_ENTRY_CREATE"
	GuildBanAdd                         RawType = "GUILD_BAN_ADD"
	GuildBanRemove                      RawType = "GUILD_BAN_REMOVE"
	GuildEmojisUpdate                   RawType = "GUILD_EMOJIS_UPDATE"
	GuildStickersUpdate                 RawType = "GUILD_STICKERS_UPDATE"
	GuildIntegrationsUpdate             RawType = "GUILD_INTEGRATIONS_UPDATE"
	GuildMemberAdd                      RawType = "GUILD_MEMBER_ADD"
	GuildMemberRemove                   RawType = "GUILD_MEMBER_REMOVE"
	GuildMemberUpdate                   RawType = "GUILD_MEMBER_UPDATE"
	GuildMembersChunk                   RawType = "GUILD_MEMBERS_CHUNK"
	GuildRoleCreate                     RawType = "GUILD_ROLE_CREATE"
	GuildRoleUpdate                     RawType = "GUILD_ROLE_UPDATE"
	GuildRoleDelete                     RawType = "GUILD_ROLE_DELETE"
	GuildScheduledEventCreate           RawType = "GUILD_SCHEDULED_EVENT_CREATE"
	GuildScheduledEventUpdate           RawType = "GUILD_SCHEDULED_EVENT_UPDATE"
	GuildScheduledEventDelete           RawType = "GUILD_SCHEDULED_EVENT_DELETE"
	GuildScheduledEventUserAdd          RawType = "GUILD_SCHEDULED_EVENT_USER_ADD"
	GuildScheduledEventUserRemove       RawType = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	IntegrationCreate                   RawType = "INTEGRATION_CREATE"
	IntegrationUpdate                   RawType = "INTEGRATION_UPDATE"
	IntegrationDelete                   RawType = "INTEGRATION_DELETE"
	InteractionCreate                   RawType = "INTERACTION_CREATE"
	InviteCreate                        RawType = "INVITE_CREATE"
	InviteDelete                        RawType = "INVITE_DELETE"
	MessageCreate                       RawType = "MESSAGE_CREATE"
	MessageUpdate                       RawType = "MESSAGE_UPDATE"
	MessageDelete                       RawType = "MESSAGE_DELETE"
	MessageDeleteBulk                   RawType = "MESSAGE_DELETE_BULK"
	MessageReactionAdd                  RawType = "MESSAGE_REACTION_ADD"
	MessageReactionRemove               RawType = "MESSAGE_REACTION_REMOVE"
	MessageReactionRemoveAll            RawType = "MESSAGE_REACTION_REMOVE_ALL"
	MessageReactionRemoveEmoji          RawType = "MESSAGE_REACTION_REMOVE_EMOJI"
	PresenceUpdate                      RawType = "PRESENCE_UPDATE"
	StageInstanceCreate                 RawType = "STAGE_INSTANCE_CREATE"
	StageInstanceDelete                 RawType = "STAGE_INSTANCE_DELETE"
	StageInstanceUpdate                 RawType = "STAGE_INSTANCE_UPDATE"
	TypingStart                         RawType = "TYPING_START"
	UserUpdate                          RawType = "USER_UPDATE"
	VoiceStateUpdate                    RawType = "VOICE_STATE_UPDATE"
	VoiceServerUpdate                   RawType = "VOICE_SERVER_UPDATE"
	WebhooksUpdate                      RawType = "WEBHOOKS_UPDATE"
)
