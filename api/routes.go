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

//goland:noinspection SpellCheckingInspection
const (
	getGlobalApplicationCommands           = "%s/applications/%s/commands"
	createGlobalApplicationCommand         = getGlobalApplicationCommands
	bulkOverwriteGlobalApplicationCommands = getGlobalApplicationCommands
	getGlobalApplicationCommand            = "%s/applications/%s/commands/%s"
	editGlobalApplicationCommand           = getGlobalApplicationCommand
	deleteGlobalApplicationCommand         = getGlobalApplicationCommand
	getGuildApplicationCommands            = "%s/applications/%s/guilds/%s/commands"
	createGuildApplicationCommand          = getGuildApplicationCommands
	bulkOverwriteGuildApplicationCommands  = getGuildApplicationCommands
	getGuildApplicationCommand             = "%s/applications/%s/guilds/%s/commands/%s"
	editGuildApplicationCommand            = getGuildApplicationCommand
	deleteGuildApplicationCommand          = getGuildApplicationCommand
	getApplicationCommandPermissions       = "%s/applications/%s/guilds/%s/command/%s/permissions"
	editApplicationCommandPermissions      = getApplicationCommandPermissions
	getGuildApplicationCommandPermissions  = "%s/applications/%s/guilds/%s/command/permissions"
	batchEditApplicationCommandPermissions = getGuildApplicationCommandPermissions
	createInteractionResponse              = "%s/interactions/%s/%s/callback"
	getGuildAuditLog                       = "%s/guilds/%s/audit-logs"
	getChannel                             = "%s/channels/%s"
	modifyChannel                          = getChannel
	deleteChannel                          = getChannel
	followNewsChannel                      = "%s/channels/%s/followers"
	getChannelInvites                      = "%s/channels/%s/invites"
	createMessage                          = "%s/channels/%s/messages"
	getChannelMessages                     = createMessage
	getChannelMessage                      = "%s/channels/%s/messages/%s"
	editMessage                            = getChannelMessage
	deleteMessage                          = getChannelMessage
	crosspostMessage                       = "%s/channels/%s/messages/%s/crosspost"
	deleteAllReactions                     = "%s/channels/%s/messages/%s/reactions"
	deleteAllReactionsForEmoji             = "%s/channels/%s/messages/%s/reactions/%s"
	deleteUserReaction                     = "%s/channels/%s/messages/%s/reactions/%s/%s"
	getReactions                           = deleteAllReactionsForEmoji
	createReaction                         = "%s/channels/%s/messages/%s/reactions/%s/@me"
	deleteOwnReaction                      = createReaction
	startThreadWithMessage                 = "%s/channels/%s/messages/%s/threads"
	bulkDeleteMessages                     = "%s/channels/%s/messages/bulk-delete"
	editChannelPermissions                 = "%s/channels/%s/permissions/%s"
	deleteChannelPermission                = editChannelPermissions
	getPinnedMessages                      = "%s/channels/%s/pins"
	pinMessage                             = "%s/channels/%s/pins/%s"
	unpinMessage                           = pinMessage
	groupDmAddRecipient                    = "%s/channels/%s/recipients/%s"
	groupDmRemoveRecipient                 = groupDmAddRecipient
	listThreadMembers                      = "%s/channels/%s/thread-members"
	addThreadMember                        = "%s/channels/%s/thread-members/%s"
	removeThreadMember                     = addThreadMember
	getThreadMember                        = addThreadMember
	joinThread                             = "%s/channels/%s/thread-members/@me"
	leaveThread                            = joinThread
	startThreadWithoutMessage              = "%s/channels/%s/threads"
	startThreadInForumChannel              = startThreadWithoutMessage
	listPrivateArchivedThreads             = "%s/channels/%s/threads/archived/private"
	listPublicArchivedThreads              = "%s/channels/%s/threads/archived/public"
	triggerTypingIndicator                 = "%s/channels/%s/typing"
	listJoinedPrivateArchivedThreads       = "%s/channels/%s/users/@me/threads/archived/private"
	listGuildEmojis                        = "%s/guilds/%s/emojis"
	createGuildEmoji                       = listGuildEmojis
	getGuildEmoji                          = "%s/guilds/%s/emojis/%s"
	modifyGuildEmoji                       = getGuildEmoji
	deleteGuildEmoji                       = getGuildEmoji
	createGuild                            = "%s/guilds"
	getGuildTemplate                       = "%s/guilds/templates/%s"
	createGuildFromGuildTemplate           = getGuildTemplate
	getGuild                               = "%s/guilds/%s"
	modifyGuild                            = getGuild
	deleteGuild                            = modifyGuild
	getGuildBans                           = "%s/guilds/%s/bans"
	getGuildBan                            = "%s/guilds/%s/bans/%s"
	createGuildBan                         = getGuildBan
	removeGuildBan                         = createGuildBan
	getGuildChannels                       = "%s/guilds/%s/channels"
	createGuildChannel                     = getGuildChannels
	modifyGuildChannelPositions            = createGuildChannel
	getGuildIntegrations                   = "%s/guilds/%s/integrations"
	deleteGuildIntegration                 = "%s/guilds/%s/integrations/%s"
	getGuildInvites                        = "%s/guilds/%s/invites"
	getGuildPreview                        = "%s/guilds/%s/preview"
	listGuildMembers                       = "%s/guilds/%s/members"
	searchGuildMembers                     = "%s/guilds/%s/members/search"
	getGuildMember                         = "%s/guilds/%s/members/%s"
	addGuildMember                         = getGuildMember
	modifyGuildMember                      = addGuildMember
	removeGuildMember                      = modifyGuildMember
	modifyCurrentMember                    = "%s/guilds/%s/members/@me"
	addGuildMemberRole                     = "%s/guilds/%s/members/%s/roles/%s"
	getGuildPruneCount                     = "%s/guilds/%s/prune"
	beginGuildPrune                        = getGuildPruneCount
	getGuildVoiceRegions                   = "%s/guilds/%s/regions"
	getGuildRoles                          = "%s/guilds/%s/roles"
	createGuildRole                        = getGuildRoles
	modifyGuildRolePositions               = createGuildRole
	modifyGuildRole                        = "%s/guilds/%s/roles/%s"
	deleteGuildRole                        = modifyGuildRole
	listGuildScheduledEvents               = "%s/guilds/%s/scheduled-events"
	createGuildScheduledEvent              = listGuildScheduledEvents
	getGuildScheduledEvent                 = "%s/guilds/%s/scheduled-events/%s"
	modifyGuildScheduledEvent              = getGuildScheduledEvent
	deleteGuildScheduledEvent              = modifyGuildScheduledEvent
	getGuildScheduledEventUsers            = "%s/guilds/%s/scheduled-events/%s/users"
	getGuildTemplates                      = "%s/guilds/%s/templates"
	createGuildTemplate                    = getGuildTemplates
	syncGuildTemplate                      = "%s/guilds/%s/templates/%s"
	modifyGuildTemplate                    = syncGuildTemplate
	deleteGuildTemplate                    = modifyGuildTemplate
	listActiveThreads                      = "%s/guilds/%s/threads/active"
	removeGuildMemberRole                  = addGuildMemberRole
	getGuildVanityURL                      = "%s/guilds/%s/vanity-url"
	modifyUserVoiceState                   = "%s/guilds/%s/voice-states/%s"
	modifyCurrentUserVoiceState            = "%s/guilds/%s/voice-states/@me"
	getGuildWelcomeScreen                  = "%s/guilds/%s/welcome-screen"
	modifyGuildWelcomeScreen               = getGuildWelcomeScreen
	getGuildWidgetSettings                 = "%s/guilds/%s/widget"
	modifyGuildWidget                      = getGuildWidgetSettings
	getGuildWidget                         = "%s/guilds/%s/widget.json"
	deleteInvite                           = "%s/invites/%s"
	getInvite                              = "%s/invites/%s%s"
	listGuildStickers                      = "%s/guilds/%s/stickers"
	createGuildSticker                     = listGuildStickers
	getGuildSticker                        = "%s/guilds/%s/stickers/%s"
	modifyGuildSticker                     = getGuildSticker
	deleteGuildSticker                     = getGuildSticker
	getSticker                             = "%s/stickers/%s"
	listNitroStickerPacks                  = "%s/sticker-packs"
	getAvatarUrlGif                        = "avatars/%s/%s.gif"
	getAvatarUrlPng                        = "avatars/%s/%s.png"
	getDefaultUserAvatarUrl                = "embed/avatars/%s.png"
	getCurrentUser                         = "%s/users/@me"
	listVoiceRegions                       = "%s/voice/regions"
	createWebhook                          = "%s/channels/%s/webhooks"
	getChannelWebhooks                     = createWebhook
	getGuildWebhooks                       = "%s/guilds/%s/webhooks"
	getWebhook                             = "%s/webhooks/%s"
	modifyWebhook                          = getWebhook
	deleteWebhook                          = getWebhook
	createFollowupMessage                  = "%s/webhooks/%s/%s"
	modifyWebhookWithToken                 = createFollowupMessage
	deleteWebhookWithToken                 = createFollowupMessage
	executeWebhook                         = "%s/webhooks/%s/%s%s"
	executeGitHubCompatibleWebhook         = "%s/webhooks/%s/%s/github%s"
	getWebhookMessage                      = "%s/webhooks/%s/%s/messages/%s%s"
	editWebhookMessage                     = getWebhookMessage
	deleteWebhookMessage                   = getWebhookMessage
	getFollowupMessage                     = "%s/webhooks/%s/%s/messages/%s"
	editFollowupMessage                    = getFollowupMessage
	deleteFollowupMessage                  = editFollowupMessage
	getOriginalInteractionResponse         = "%s/webhooks/%s/%s/messages/@original"
	editOriginalInteractionResponse        = getOriginalInteractionResponse
	deleteOriginalInteractionResponse      = getOriginalInteractionResponse
	executeSlackCompatibleWebhook          = "%s/webhooks/%s/%s/slack%s"
	createStageInstance                    = "%s/stage-instances"
	getStageInstance                       = "%s/stage-instances/%s"
	modifyStageInstance                    = getStageInstance
	deleteStageInstance                    = getStageInstance
)
