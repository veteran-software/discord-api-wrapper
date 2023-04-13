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

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	log "github.com/veteran-software/nowlive-logging"
)

// CreateGuild
//
// Create a new guild. Returns a guild object on success. Fires a GuildCreate Gateway event.
//
//	This endpoint can be used only by bots in less than 10 guilds.
//
//	When using the roles parameter, the first member of the array is used to change properties of the guild's @everyone role. If you are trying to bootstrap a guild with additional roles, keep this in mind.
//
//	When using the roles parameter, the required id field within each role object is an integer placeholder, and will be replaced by the API upon consumption. Its purpose is to allow you to overwrite a role's permissions in a channel when also passing in channels with the channels array.
//
//	When using the channels parameter, the position field is ignored, and none of the default channels are created.
//
//	When using the channels parameter, the id field within each channel object may be set to an integer placeholder, and will be replaced by the API upon consumption. Its purpose is to allow you to create GUILD_CATEGORY channels by setting the parent_id field on any children to the category's id field. Category channels must be listed before any children.
//
//goland:noinspection GoUnusedExportedFunction
func CreateGuild(payload *CreateGuildJSON) (*Guild, error) {
	u := parseRoute(fmt.Sprintf(createGuild, api))

	var guild *Guild
	responseBytes, err := firePostRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guild)

	return guild, err
}

type CreateGuildJSON struct {
	Name                        string                          `json:"name"`                          // guild name (2-100 characters, excluding trailing and leading whitespace)
	Icon                        *string                         `json:"icon"`                          // icon hash
	VerificationLevel           VerificationLevel               `json:"verification_level"`            // verification level required for the guild
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"` // default message notifications level
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`       // explicit content filter level
	Roles                       []*Role                         `json:"roles"`                         // roles in the guild
	Channels                    []*Channel                      `json:"channels,omitempty"`            // channels in the guild
	AfkChannelID                Snowflake                       `json:"afk_channel_id,omitempty"`      // id of afk channel
	AfkTimeout                  int64                           `json:"afk_timeout"`                   // afk timeout in seconds
	SystemChannelID             *Snowflake                      `json:"system_channel_id"`             // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags"`          // system channel flags
}

// GetGuild - Returns the guild object for the given id.
//
// If with_counts is set to true, this endpoint will also return approximate_member_count and approximate_presence_count for the guild.
func (g *Guild) GetGuild(withCounts *bool) (*Guild, error) {
	u := parseRoute(fmt.Sprintf(getGuild, api, g.ID.String()))

	q := u.Query()
	if withCounts != nil {
		q.Set("with_counts", strconv.FormatBool(*withCounts))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	// Some testing script for moving to a channel based request system
	// Going to use this on internal calls first to make sure it works fine
	//ch := make(chan []byte)
	//
	//go fireGetRequest(u, nil, nil, ch)
	//
	//by := <-ch

	var guild *Guild
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guild)

	return guild, err
}

// GetGuildPreview - Returns the guild preview object for the given id. If the user is not in the guild, then the guild must be lurkable.
func (g *Guild) GetGuildPreview() (*GuildPreview, error) {
	u := parseRoute(fmt.Sprintf(getGuildPreview, api, g.ID.String()))

	var guildPreview *GuildPreview
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildPreview)

	return guildPreview, err
}

// ModifyGuild - Modify a guild's settings.
//
// Requires the ManageGuild permission.
//
// Returns the updated guild object on success.
//
// Fires a GuildUpdate Gateway event.
//
//	All parameters to this endpoint are optional
//
//	This endpoint supports the X-Audit-Log-Reason header.
//
//	Attempting to add or remove the Community guild feature requires the Administrator permission.
func (g *Guild) ModifyGuild(payload ModifyGuildJSON, reason *string) (*Guild, error) {
	u := parseRoute(fmt.Sprintf(modifyGuild, api, g.ID.String()))

	var guild *Guild
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guild)

	return guild, err
}

// ModifyGuildJSON
// All parameters to this endpoint are optional
//
// This endpoint supports the `X-Audit-Log-Reason` header.
//
// Attempting to add or remove the COMMUNITY guild feature requires the Administrator permission.
type ModifyGuildJSON struct {
	Name                        string                           `json:"name,omitempty"`                          // guild name (2-100 characters, excluding trailing and leading whitespace)
	VerificationLevel           *VerificationLevel               `json:"verification_level,omitempty"`            // verification level required for the guild
	DefaultMessageNotifications *DefaultMessageNotificationLevel `json:"default_message_notifications,omitempty"` // default message notifications level
	ExplicitContentFilter       *ExplicitContentFilterLevel      `json:"explicit_content_filter,omitempty"`       // explicit content filter level
	AfkChannelID                *Snowflake                       `json:"afk_channel_id,omitempty,omitempty"`      // id of afk channel
	AfkTimeout                  int64                            `json:"afk_timeout,omitempty"`                   // afk timeout in seconds
	Icon                        *string                          `json:"icon,omitempty"`                          // icon hash
	OwnerID                     Snowflake                        `json:"owner_id,omitempty"`                      // id of owner
	Splash                      *string                          `json:"splash,omitempty,omitempty"`              // splash hash
	DiscoverySplash             *string                          `json:"discovery_splash,omitempty"`              // discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	Banner                      *string                          `json:"banner,omitempty"`                        // banner hash
	SystemChannelID             *Snowflake                       `json:"system_channel_id,omitempty"`             // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          SystemChannelFlags               `json:"system_channel_flags,omitempty"`          // system channel flags
	RulesChannelID              *Snowflake                       `json:"rules_channel_id,omitempty"`              // the id of the channel where Community guilds can display rules and/or guidelines
	PublicUpdatesChannelID      *Snowflake                       `json:"public_updates_channel_id,omitempty"`     // the id of the channel where admins and moderators of Community guilds receive notices from Discord
	PreferredLocale             string                           `json:"preferred_locale,omitempty"`              // the preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	Features                    []*GuildFeatures                 `json:"features,omitempty"`                      // enabled guild features
	Description                 *string                          `json:"description,omitempty"`                   // the description of a Community guild
	PremiumProgressBarEnabled   bool                             `json:"premium_progress_bar_enabled,omitempty"`  // whether the guild has the boost progress bar enabled
}

// DeleteGuild - Delete a guild permanently. User must be Owner. Returns `204 No Content` on success. Fires a GuildDelete Gateway event.
func (g *Guild) DeleteGuild() error {
	u := parseRoute(fmt.Sprintf(deleteGuild, api, g.ID.String()))

	return fireDeleteRequest(u, nil)
}

// GetGuildChannels - Returns a list of guild Channel objects. Does not include threads.
func (g *Guild) GetGuildChannels() ([]*Channel, error) {
	u := parseRoute(fmt.Sprintf(getGuildChannels, api, g.ID.String()))

	var channels []*Channel
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &channels)

	return channels, err
}

// CreateGuildChannel - Create a new channel object for the guild.
//
// Requires the ManageChannels permission.
//
// If setting permission overwrites, only permissions your bot has in the guild can be allowed/denied.
//
// Setting ManageRoles permission in channels is only possible for guild administrators.
//
// Returns the new channel object on success. Fires a ChannelCreate Gateway event.
//
//	All parameters to this endpoint are optional excluding name
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) CreateGuildChannel(payload CreateGuildChannelJSON, reason *string) (*Channel, error) {
	u := parseRoute(fmt.Sprintf(createGuildChannel, api, g.ID.String()))

	var channel *Channel
	responseBytes, err := firePostRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &channel)

	return channel, err
}

type CreateGuildChannelJSON struct {
	Name                       *string          `json:"name"`                                    // the name of the channel (1-100 characters)
	Type                       *ChannelType     `json:"type,omitempty"`                          // the ChannelType
	Topic                      *string          `json:"topic,omitempty"`                         // the channel topic (0-1024 characters)
	Bitrate                    *int64           `json:"bitrate,omitempty"`                       // the bitrate (in bits) of the voice channel
	UserLimit                  *int64           `json:"user_limit,omitempty"`                    // the user limit of the voice channel
	RateLimitPerUser           *int64           `json:"rate_limit_per_user,omitempty"`           // amount of seconds a user has to wait before sending another Message (0-21600); bots, as well as users with the permission ManageMessages or ManageChannels, are unaffected
	Position                   *int             `json:"position,omitempty"`                      // sorting position of the channel
	PermissionOverwrites       []*Overwrite     `json:"permission_overwrites,omitempty"`         // explicit permission overwrites for members and roles
	ParentID                   *Snowflake       `json:"parent_id,omitempty"`                     // for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	Nsfw                       *bool            `json:"nsfw,omitempty"`                          // whether the channel is nsfw
	RtcRegion                  *string          `json:"rtc_region,omitempty"`                    // channel voice region id of the voice or stage channel, automatic when set to null
	VideoQualityMode           VideoQualityMode `json:"video_quality_mode,omitempty"`            // the camera video quality mode of the voice channel
	DefaultAutoArchiveDuration *int             `json:"default_auto_archive_duration,omitempty"` // default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	DefaultReactionEmoji       DefaultReaction  `json:"default_reaction_emoji,omitempty"`        // Emoji to show in the add Reaction button on a thread in a GuildForum channel
	AvailableTags              []*ForumTag      `json:"available_tags,omitempty"`                // set of tags that can be used in a GuildForum channel
	DefaultSortOrder           SortOrderType    `json:"default_sort_order,omitempty"`            // the default SortOrderType used to order posts in GuildForum channels
}

// ModifyGuildChannelPositions - Modify the positions of a set of channel objects for the guild.
//
// Requires ManageChannels permission. Returns a 204 empty response on success. Fires multiple ChannelUpdate Gateway events.
//
//	Only channels to be modified are required.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildChannelPositions(payload *ModifyGuildChannelPositionsJSON, reason *string) error {
	u := parseRoute(fmt.Sprintf(modifyGuildChannelPositions, api, g.ID.String()))
	_, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return err
	}

	return nil
}

// ModifyGuildChannelPositionsJSON - JSON payload
type ModifyGuildChannelPositionsJSON struct {
	ID              Snowflake  `json:"id"`               // channel id
	Position        *uint64    `json:"position"`         // sorting position of the channel
	LockPermissions *bool      `json:"lock_permissions"` // syncs the permission overwrites with the new parent, if moving to a new category
	ParentID        *Snowflake `json:"parent_id"`        // the new parent ID for the channel that is moved
}

// ListActiveThreads - Returns all active threads in the guild, including public and private threads. Threads are ordered by their id, in descending order.
func (g *Guild) ListActiveThreads() (*ThreadListResponse, error) {
	u := parseRoute(fmt.Sprintf(listActiveThreads, api, g.ID.String()))

	var threadListResponse *ThreadListResponse
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &threadListResponse)

	return threadListResponse, err
}

// GetGuildMember - Returns a GuildMember object for the specified User.
func (g *Guild) GetGuildMember(userID *Snowflake) (*GuildMember, error) {
	u := parseRoute(fmt.Sprintf(getGuildMember, api, g.ID.String(), userID.String()))

	var guildMember *GuildMember
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildMember)

	return guildMember, err
}

// ListGuildMembers - Returns a list of guild member objects that are members of the guild.
//
// This endpoint is restricted according to whether the GuildMembers Privileged Intent is enabled for your application.
func (g *Guild) ListGuildMembers(limit *uint64, after *Snowflake) ([]*GuildMember, error) {
	u := parseRoute(fmt.Sprintf(listGuildMembers, api, g.ID.String()))

	q := u.Query()
	if after != nil {
		q.Set("after", after.String())
	}
	if limit != nil {
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var guildMembers []*GuildMember
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildMembers)

	return guildMembers, err
}

// SearchGuildMembers - Returns a list of GuildMember objects whose username or nickname starts with a provided string.
//
//	All parameters to this endpoint except for `query` are optional
func (g *Guild) SearchGuildMembers(query string, limit *uint64) ([]*GuildMember, error) {
	u := parseRoute(fmt.Sprintf(searchGuildMembers, api, g.ID.String()))

	q := u.Query()
	q.Set("query", query)
	if limit != nil {
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	u.RawQuery = q.Encode()

	var guildMembers []*GuildMember
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildMembers)

	return guildMembers, err
}

// AddGuildMember - Adds a user to the guild, provided you have a valid oauth2 access token for the user with the `guilds.join` scope.
//
// Returns a 201 Created with the guild member as the body, or 204 No Content if the user is already a member of the guild.
//
// Fires a GuildMemberAdd Gateway event.
//
//	For guilds with MembershipScreening enabled, this endpoint will default to adding new members as pending in the guild member object.
//
// Members that are pending will have to complete membership screening before they become full members that can talk.
//
//	All parameters to this endpoint except for access_token are optional.
//
//	The Authorization header must be a Bot token (belonging to the same application used for authorization), and the bot must be a member of the guild with CreateInstantInvite permission.
//
// For guilds with Membership Screening enabled, assigning a role using the `roles` parameter will add the user to the guild as a full member (pending is false in the member object).
//
// A member with a role will bypass membership screening and the guild's verification level, and get immediate access to chat.
//
// Therefore, instead of assigning a role when the member joins, it is recommended to grant roles only after the user completes screening.
func (g *Guild) AddGuildMember(userID *Snowflake, payload *AddGuildMemberJSON) (*GuildMember, error) {
	u := parseRoute(fmt.Sprintf(addGuildMember, api, g.ID.String(), userID.String()))

	var guildMember *GuildMember
	responseBytes, err := firePutRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildMember)

	return guildMember, err
}

// AddGuildMemberJSON - JSON payload
type AddGuildMemberJSON struct {
	AccessToken string       `json:"access_token"`    // an oauth2 access token granted with the `guilds.join` to the bots' application for the user you want to add to the guild
	Nick        string       `json:"nick,omitempty"`  // value to set user's nickname to
	Roles       []*Snowflake `json:"roles,omitempty"` // array of role ids the member is assigned
	Mute        bool         `json:"mute,omitempty"`  // whether the user is muted in voice channels
	Deaf        bool         `json:"deaf,omitempty"`  // whether the user is deafened in voice channels
}

// ModifyGuildMember - Modify attributes of a guild member.
//
// Returns a 200 OK with the guild member as the body. Fires a GuildMember Update Gateway event.
//
// If the channel_id is set to null, this will force the target user to be disconnected from voice.
//
//	All parameters to this endpoint are optional and nullable. When moving members to channels, the API user must have permissions to both connect to the channel and have the MoveMembers permission.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildMember(userID *Snowflake, payload *ModifyGuildMemberJSON, reason *string) (
	*GuildMember,
	error,
) {
	u := parseRoute(fmt.Sprintf(modifyGuildMember, api, g.ID.String(), userID.String()))

	var guildMember *GuildMember
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildMember)

	return guildMember, err
}

// ModifyGuildMemberJSON - JSON payload
type ModifyGuildMemberJSON struct {
	Nick                       *string         `json:"nick,omitempty"`                         // value to set user's nickname to
	Roles                      []*Snowflake    `json:"roles,omitempty"`                        // array of role ids the member is assigned
	Mute                       *bool           `json:"mute,omitempty"`                         // whether the user is muted in voice channels. Will throw a 400 error if the user is not in a voice channel
	Deaf                       *bool           `json:"deaf,omitempty"`                         // whether the user is deafened in voice channels. Will throw a 400 error if the user is not in a voice channel
	ChannelID                  *Snowflake      `json:"channel_id,omitempty"`                   // id of channel to move user to (if they are connected to voice)
	CommunicationDisabledUntil *time.Time      `json:"communication_disabled_until,omitempty"` // when the user's timeout will expire and the User will be able to communicate in the guild again (up to 28 days in the future), set to null to remove timeout. Will throw a 403 error if the user has the Administrator permission or is the owner of the guild
	Flags                      GuildMemberFlag `json:"flags"`                                  // guild member flags
}

// ModifyCurrentMember - Modifies the current member in a guild. Returns a 200 with the updated member object on success. Fires a Guild Member Update Gateway event.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyCurrentMember(nick *string, reason *string) (*GuildMember, error) {
	u := parseRoute(fmt.Sprintf(modifyCurrentMember, api, g.ID.String()))

	payload := struct {
		Nick *string
	}{
		nick,
	}

	var guildMember *GuildMember
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildMember)

	return guildMember, err
}

// AddGuildMemberRole - Adds a role to a guild member.
//
// Requires the ManageRoles permission.
//
// Returns a 204 empty response on success.
//
// Fires a Guild Member Update Gateway event.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) AddGuildMemberRole(user *User, roleID *Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(addGuildMemberRole, api, g.ID.String(), user.ID.String(), roleID.String()))

	_, err := firePutRequest(u, nil, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return err
	}

	return nil
}

// RemoveGuildMemberRole - Removes a role from a guild member.
//
// Requires the ManageRoles permission.
//
// Returns a 204 empty response on success.
//
// Fires a GuildMember Update Gateway event.
//
// This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) RemoveGuildMemberRole(user *User, role *Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(removeGuildMemberRole, api, g.ID.String(), user.ID.String(), role.String()))

	return fireDeleteRequest(u, reason)
}

// RemoveGuildMember - Remove a member from a guild.
//
// Requires KickMembers permission.
//
// Returns a 204 empty response on success.
//
// Fires a GuildMemberRemove Gateway event.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) RemoveGuildMember(user *User, reason *string) error {
	u := parseRoute(fmt.Sprintf(removeGuildMember, api, g.ID.String(), user.ID.String()))

	return fireDeleteRequest(u, reason)
}

// GetGuildBans - Returns a list of Ban objects for the users banned from this guild. Requires the BanMembers permission.
func (g *Guild) GetGuildBans(limit *uint64, before *Snowflake, after *Snowflake) ([]*Ban, error) {
	u := parseRoute(fmt.Sprintf(getGuildBans, api, g.ID.String()))

	q := u.Query()
	if before != nil {
		q.Set("before", before.String())
	}
	if after != nil {
		q.Set("after", after.String())
	}
	if limit != nil {
		q.Set("limit", strconv.FormatUint(*limit, 10))
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var bans []*Ban
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &bans)

	return bans, err
}

// GetGuildBan - Returns a ban object for the given user or a 404 not found if the ban cannot be found. Requires the BanMembers permission.
func (g *Guild) GetGuildBan(userID Snowflake) (*Ban, error) {
	u := parseRoute(fmt.Sprintf(getGuildBan, api, g.ID.String(), userID.String()))

	var ban *Ban
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &ban)

	return ban, err
}

// CreateGuildBan - Create a guild ban, and optionally delete previous messages sent by the banned user. Requires the BanMembers permission. Returns a 204 empty response on success. Fires a GuildBanAdd Gateway event.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) CreateGuildBan(userID *Snowflake, deleteMessageSeconds *uint64, reason *string) error {
	u := parseRoute(fmt.Sprintf(createGuildBan, api, g.ID.String(), userID.String()))

	payload := struct {
		DeleteMessageSeconds *uint64 `json:"delete_message_seconds"`
	}{
		deleteMessageSeconds,
	}

	_, err := firePutRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return err
	}

	return nil
}

// RemoveGuildBan - Remove the ban for a user. Requires the BanMembers permissions. Returns a 204 empty response on success. Fires a GuildBanRemove Gateway event.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) RemoveGuildBan(userID *Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(removeGuildBan, api, g.ID.String(), userID.String()))

	return fireDeleteRequest(u, reason)
}

// GetGuildRoles - Returns a list of role objects for the guild.
func (g *Guild) GetGuildRoles() ([]*Role, error) {
	u := parseRoute(fmt.Sprintf(getGuildRoles, api, g.ID.String()))

	var roles []*Role
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &roles)

	return roles, err
}

func (g *Guild) CreateGuildRole(payload *CreateGuildRoleJSON, reason *string) ([]*Role, error) {
	u := parseRoute(fmt.Sprintf(createGuildRole, api, g.ID.String()))

	var roles []*Role
	responseBytes, err := firePostRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &roles)

	return roles, err
}

// CreateGuildRoleJSON - JSON payload
type CreateGuildRoleJSON struct {
	Name         string  `json:"name"`
	Permissions  string  `json:"permissions"`
	Color        uint64  `json:"color"`
	Hoist        bool    `json:"hoist"`
	Icon         *string `json:"icon"`
	UnicodeEmoji *string `json:"unicode_emoji"`
	Mentionable  bool    `json:"mentionable"`
}

// ModifyGuildRolePositions - Modify the positions of a set of role objects for the guild.
//
// Requires the ManageRoles permission.
//
// Returns a list of all the guild's role objects on success.
//
// Fires multiple Guild Role Update Gateway events.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildRolePositions(payload []*ModifyGuildRolePositionsJSON, reason *string) ([]*Role, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildRolePositions, api, g.ID.String()))

	var roles []*Role
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &roles)

	return roles, err
}

// ModifyGuildRolePositionsJSON - JSON payload
type ModifyGuildRolePositionsJSON struct {
	ID       Snowflake `json:"id"`                 // role
	Position *uint64   `json:"position,omitempty"` // sorting position of the role
}

// ModifyGuildRole - Modify a Guild Role.
//
// Requires the ManageRoles permission.
//
// Returns the updated role on success.
//
// Fires a GuildRoleUpdate Gateway event.
//
//	All parameters to this endpoint are optional and nullable.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildRole(roleID *Snowflake, payload *ModifyGuildRoleJSON, reason *string) (*Role,
	error) {
	u := parseRoute(fmt.Sprintf(modifyGuildRole, api, g.ID.String(), roleID.String()))

	var roles *Role
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &roles)

	return roles, err
}

// ModifyGuildRoleJSON - JSON payload
type ModifyGuildRoleJSON struct {
	Name         *string `json:"name,omitempty"`
	Permissions  *string `json:"permissions,omitempty"`
	Color        *uint64 `json:"color,omitempty"`
	Hoist        *bool   `json:"hoist,omitempty"`
	Icon         *string `json:"icon,omitempty"`
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
	Mentionable  *bool   `json:"mentionable,omitempty"`
}

// ModifyGuildMfaLevel - Modify a guild's MFA level.
//
// Requires guild ownership.
//
// Returns the updated level on success.
//
// Fires a GuildUpdate Gateway event.
func (g *Guild) ModifyGuildMfaLevel(level MfaLevel, reason *string) (*MfaLevel, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildMfaLevel, api, g.ID.String()))

	payload := struct {
		Level MfaLevel `json:"level"`
	}{
		level,
	}

	var mfaLevel *MfaLevel
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &mfaLevel)

	return mfaLevel, err
}

// DeleteGuildRole - Delete a guild role.
//
// Requires the ManageRoles permission.
//
// Returns a 204 empty response on success.
//
// Fires a GuildRoleDelete Gateway event.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) DeleteGuildRole(roleID *Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteGuildRole, api, g.ID.String(), roleID.String()))

	return fireDeleteRequest(u, reason)
}

// GetGuildPruneCount - Returns an object with one pruned key indicating the number of members that would be removed in a prune operation.
//
// Requires the KickMembers permission.
//
// By default, prune will not remove users with roles.
//
// You can optionally include specific roles in your prune by providing the `include_roles` parameter.
//
// Any inactive user that has a subset of the provided role(s) will be counted in the prune and users with additional roles will not.
func (g *Guild) GetGuildPruneCount(days uint, includeRoles *string) (*GetGuildPruneCountResponse, error) {
	if days < 1 || days > 30 {
		return nil, errors.New("the number of days to prune must be >= 1 && <= 30")
	}

	u := parseRoute(fmt.Sprintf(getGuildPruneCount, api, g.ID.String()))

	q := u.Query()
	q.Set("days", strconv.Itoa(int(days)))
	if includeRoles != nil {
		q.Set("include_Roles", *includeRoles)
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	var pruneCountResponse *GetGuildPruneCountResponse
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &pruneCountResponse)

	return pruneCountResponse, err
}

type GetGuildPruneCountResponse struct {
	Pruned *uint64 `json:"pruned"`
}

// BeginGuildPrune - Begin a prune operation.
//
// Requires the KickMembers permission.
//
// Returns an object with one pruned key indicating the number of members that were removed in the prune operation.
//
// For large guilds it's recommended to set the `compute_prune_count` option to false, forcing pruned to null.
//
// Fires multiple GuildMemberRemove Gateway events.
//
// By default, prune will not remove users with roles.
//
// You can optionally include specific roles in your prune by providing the include_roles parameter.
//
// Any inactive user that has a subset of the provided role(s) will be included in the prune and users with additional roles will not.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) BeginGuildPrune(payload *BeginGuildPruneJSON, reason *string) (*GetGuildPruneCountResponse, error) {
	if payload.Days < 1 || payload.Days > 30 {
		return nil, errors.New("the number of days to prune must be >= 1 && <= 30")
	}

	u := parseRoute(fmt.Sprintf(beginGuildPrune, api, g.ID.String()))

	var response *GetGuildPruneCountResponse
	responseBytes, err := firePostRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &response)

	return response, err
}

// BeginGuildPruneJSON - JSON payload
type BeginGuildPruneJSON struct {
	Days              uint         `json:"days"`                // number of days to prune (1-30)
	ComputePruneCount bool         `json:"compute_prune_count"` // whether `pruned` is returned, discouraged for large guilds
	IncludeRoles      []*Snowflake `json:"include_roles"`       // role(s) to include
}

// GetGuildVoiceRegions - Returns a list of voice region objects for the guild.
//
// Unlike the similar `/voice` route, this returns VIP servers when the guild is VIP-enabled.
func (g *Guild) GetGuildVoiceRegions() ([]*VoiceRegion, error) {
	u := parseRoute(fmt.Sprintf(getGuildVoiceRegions, api, g.ID.String()))

	var voiceRegions []*VoiceRegion
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &voiceRegions)

	return voiceRegions, err
}

// GetGuildInvites - Returns a list of invite objects (with invite metadata) for the guild.
//
// Requires the ManageGuild permission.
func (g *Guild) GetGuildInvites() ([]*Invite, error) {
	u := parseRoute(fmt.Sprintf(getGuildInvites, api, g.ID.String()))

	var invites []*Invite
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &invites)

	return invites, err
}

// GetGuildIntegrations - Returns a list of integration objects for the guild.
//
// Requires the ManageGuild permission.
func (g *Guild) GetGuildIntegrations() ([]*Integration, error) {
	u := parseRoute(fmt.Sprintf(getGuildIntegrations, api, g.ID.String()))

	var integrations []*Integration
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &integrations)

	return integrations, err
}

// DeleteGuildIntegration - Delete the attached integration object for the guild.
//
// Deletes any associated webhooks and kicks the associated bot if there is one.
//
// Requires the ManageGuild permission.
//
// Returns a 204 empty response on success.
//
// Fires a GuildIntegrationsUpdate Gateway event.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) DeleteGuildIntegration(integrationID *Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteGuildIntegration, api, g.ID.String(), integrationID.String()))

	return fireDeleteRequest(u, reason)
}

// GetGuildWidgetSettings - Returns a guild widget settings object.
//
// Requires the ManageGuild permission.
func (g *Guild) GetGuildWidgetSettings() (*GuildWidgetSettings, error) {
	u := parseRoute(fmt.Sprintf(getGuildWidgetSettings, api, g.ID.String()))

	var guildWidgetSettings *GuildWidgetSettings
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildWidgetSettings)

	return guildWidgetSettings, err
}

// ModifyGuildWidget - Modify a guild widget settings object for the guild.
//
// All attributes may be passed in with JSON and modified.
//
// Requires the ManageGuild permission.
//
// Returns the updated GuildWidgetSettings object.
//
//	This endpoint supports the X-Audit-Log-Reason header.
func (g *Guild) ModifyGuildWidget(payload *GuildWidgetSettings, reason *string) (*GuildWidgetSettings, error) {
	u := parseRoute(fmt.Sprintf(modifyGuildWidget, api, g.ID.String()))

	var guildWidgetSettings *GuildWidgetSettings
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildWidgetSettings)

	return guildWidgetSettings, err
}

// GetGuildWidget - Returns the widget for the guild.
func (g *Guild) GetGuildWidget() (*GetGuildWidget, error) {
	u := parseRoute(fmt.Sprintf(getGuildWidget, api, g.ID.String()))

	var guildWidget *GetGuildWidget
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &guildWidget)

	return guildWidget, err
}

// GetGuildVanityURL - Returns a partial invite object for guilds with that feature enabled.
//
// Requires the ManageGuild permission.
//
// `code` will be null if a vanity url for the Guild is not set.
func (g *Guild) GetGuildVanityURL() (*Invite, error) {
	u := parseRoute(fmt.Sprintf(getGuildVanityURL, api, g.ID.String()))

	var invite *Invite
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &invite)

	return invite, err
}

// GetGuildWelcomeScreen - Returns the WelcomeScreen object for the guild.
//
// If the welcome screen is not enabled, the ManageGuild permission is required.
func (g *Guild) GetGuildWelcomeScreen() (*WelcomeScreen, error) {
	u := parseRoute(fmt.Sprintf(getGuildWelcomeScreen, api, g.ID.String()))

	var welcomeScreen *WelcomeScreen
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &welcomeScreen)

	return welcomeScreen, err
}

// ModifyGuildWelcomeScreen - Modify the guild's Welcome Screen. Requires the ManageGuild permission. Returns the updated WelcomeScreen object.
//
//	All parameters to this endpoint are optional and nullable
//
//	This endpoint supports the `X-Audit-Log-Reason` header.
func (g *Guild) ModifyGuildWelcomeScreen(payload *ModifyGuildWelcomeScreenJSON, reason *string) (
	*WelcomeScreen,
	error,
) {
	u := parseRoute(fmt.Sprintf(modifyGuildWelcomeScreen, api, g.ID.String()))

	var welcomeScreen *WelcomeScreen
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &welcomeScreen)

	return welcomeScreen, err
}

// ModifyGuildWelcomeScreenJSON - JSON payload
type ModifyGuildWelcomeScreenJSON struct {
	Enabled         *bool                   `json:"enabled,omitempty"`          // whether the welcome screen is enabled
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels,omitempty"` // channels linked in the welcome screen and their display options
	Description     *string                 `json:"description,omitempty"`      // the server description to show in the welcome screen
}

func (g *Guild) GetGuildOnboarding() (*GuildOnboarding, error) {
	u := parseRoute(fmt.Sprintf(getGuildOnboarding, api, g.ID.String()))

	var onboarding *GuildOnboarding
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &onboarding)

	return onboarding, err
}

// ModifyCurrentUserVoiceState - Updates the current user's voice state. Returns 204 No Content on success.
//
// There are currently several caveats for this endpoint:
//
//   - `channel_id` must currently point to a stage channel.
//   - current user must already have joined `channel_id`.
//   - You must have the MuteMembers permission to unsuppress yourself. You can always suppress yourself.
//   - You must have the RequestToSpeak permission to request to speak. You can always clear your own request to speak.
//   - You are able to set `request_to_speak_timestamp` to any present or future time.
func (g *Guild) ModifyCurrentUserVoiceState(payload *ModifyCurrentUserVoiceStateJSON) error {
	u := parseRoute(fmt.Sprintf(modifyCurrentUserVoiceState, api, g.ID.String()))

	_, err := firePatchRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return err
	}

	return nil
}

// ModifyCurrentUserVoiceStateJSON - JSON payload
type ModifyCurrentUserVoiceStateJSON struct {
	ChannelID               Snowflake  `json:"channel_id,omitempty"`                 // the id of the channel the user is currently in
	Suppress                bool       `json:"suppress,omitempty"`                   // toggles the user's suppress state
	RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"` // sets the user's request to speak
}

// ModifyUserVoiceState - Updates another user's voice state.
//
// There are currently several caveats for this endpoint:
//
//	`channel_id` must currently point to a stage channel.
//	User must already have joined `channel_id`.
//	You must have the MuteMembers permission. (Since suppression is the only thing that is available currently.)
//	When unsuppressed, non-bot users will have their `request_to_speak_timestamp` set to the current time. Bot users will not.
//	When suppressed, the user will have their `request_to_speak_timestamp` removed.
func (g *Guild) ModifyUserVoiceState(userID *Snowflake, payload *ModifyUserVoiceStateJSON) error {
	u := parseRoute(fmt.Sprintf(modifyUserVoiceState, api, g.ID.String(), userID.String()))

	_, err := firePatchRequest(u, payload, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return err
	}

	return nil
}

// ModifyUserVoiceStateJSON - JSON payload
type ModifyUserVoiceStateJSON struct {
	ChannelID Snowflake `json:"channel_id"`         // the id of the channel the user is currently in
	Suppress  bool      `json:"suppress,omitempty"` // toggles the user's suppress state
}

func (g *Guild) getSelfMember() (*GuildMember, error) {
	return g.GetGuildMember(&ApplicationID)
}

func (g *Guild) GetGuildWidgetImage() ([]byte, error) {
	u := parseRoute(fmt.Sprintf("%s/guilds/%s/widget.png", api, g.ID.String()))

	b, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	return b, nil
}
