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
	"fmt"
	"net/http"

	"github.com/veteran-software/discord-api-wrapper/routes"
)

// Sticker - Represents a sticker that can be sent in messages.
type Sticker struct {
	ID          Snowflake   `json:"id"`                   // ID - id of the sticker
	PackID      Snowflake   `json:"pack_id,omitempty"`    // PackID - for standard stickers, id of the pack the sticker is from
	Name        string      `json:"name"`                 // Name - name of the sticker
	Description *string     `json:"description"`          // Description - description of the sticker
	Tags        string      `json:"tags,omitempty"`       // Tags - autocomplete/suggestion tags for the sticker (max 200 characters)
	Asset       string      `json:"asset,omitempty"`      // Asset - Deprecated: previously the sticker asset hash, now an empty string
	Type        StickerType `json:"type"`                 // Type - type of sticker
	FormatType  int         `json:"format_type"`          // FormatType - type of sticker format
	Available   bool        `json:"available,omitempty"`  // Available - whether this guild sticker can be used, may be false due to loss of Server Boosts
	GuildID     Snowflake   `json:"guild_id,omitempty"`   // GuildID - id of the guild that owns this sticker
	User        User        `json:"user,omitempty"`       // User - the user that uploaded the guild sticker
	SortValue   int         `json:"sort_value,omitempty"` // SortValue - the standard sticker's sort order within its pack
}

type StickerType int

const (
	StickerTypeStandard StickerType = iota + 1 // StickerTypeStandard - an official sticker in a pack, part of Nitro or in a removed purchasable pack
	StickerTypeGuild                           // StickerTypeGuild - a sticker uploaded to a Boosted guild for the guild's members
)

type StickerFormatType int

const (
	StickerFormatTypePng StickerFormatType = iota + 1
	StickerFormatTypeAnimatedPng
	StickerFormatTypeLottie
)

// StickerItem - The smallest amount of data required to render a sticker. A partial sticker object.
type StickerItem struct {
	ID         Snowflake         `json:"id"`          // ID - id of the sticker
	Name       string            `json:"name"`        // Name - name of the sticker
	FormatType StickerFormatType `json:"format_type"` // FormatType - type of sticker format
}

// StickerPack - Represents a pack of standard stickers.
type StickerPack struct {
	ID             Snowflake `json:"id"`                         // ID - id of the sticker pack
	Stickers       []Sticker `json:"stickers"`                   // Stickers - the stickers in the pack
	Name           string    `json:"name"`                       // Name - name of the sticker pack
	SkuID          Snowflake `json:"sku_id"`                     // SkuID - id of the pack's SKU
	CoverStickerID Snowflake `json:"cover_sticker_id,omitempty"` // CoverStickerID - id of a sticker in the pack which is shown as the pack's icon
	Description    string    `json:"description"`                // Description - description of the sticker pack
	BannerAssetID  Snowflake `json:"banner_asset_id,omitempty"`  // BannerAssetID - id of the sticker pack's banner image
}

// GetSticker - Returns a sticker object for the given sticker ID.
func (s *Sticker) GetSticker() (string, string) {
	return http.MethodGet, fmt.Sprintf(routes.Stickers_, api, s.ID.String())
}

// ListNitroStickerPacks - Returns the list of sticker packs available to Nitro subscribers.
func ListNitroStickerPacks() (string, string) {
	return http.MethodGet, fmt.Sprintf(routes.StickerPacks, api)
}

// ListGuildStickers - Returns an array of sticker objects for the given guild.
//
// Includes user fields if the bot has the ManageEmojisAndStickers permission.
func (g *Guild) ListGuildStickers() (string, string) {
	return http.MethodGet, fmt.Sprintf(routes.Guilds_Stickers, api, g.ID.String())
}

func (g *Guild) GetGuildSticker(stickerID Snowflake) (string, string) {
	return http.MethodGet, fmt.Sprintf(routes.Guilds_Stickers_, api, g.ID.String(), stickerID.String())
}

// CreateGuildSticker - Create a new sticker for the guild.
//
// Send a multipart/form-data body.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns the new sticker object on success.
func (g *Guild) CreateGuildSticker() (string, string) {
	return http.MethodPost, fmt.Sprintf(routes.Guilds_Stickers, api, g.ID.String())
}

// ModifyGuildSticker - Modify the given sticker.
//
// Requires the ManageEmojisAndStickers permission.
//
// Returns the updated sticker object on success.
//
// All parameters to this endpoint are optional.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) ModifyGuildSticker(stickerID Snowflake) (string, string) {
	return http.MethodPatch, fmt.Sprintf(routes.Guilds_Stickers_, api, g.ID.String(), stickerID.String())
}

// DeleteGuildSticker - Delete the given sticker.
//
// Requires the MANAGE_EMOJIS_AND_STICKERS permission.
//
// Returns "204 No Content" on success.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
func (g *Guild) DeleteGuildSticker(stickerID Snowflake) (string, string) {
	return http.MethodDelete, fmt.Sprintf(routes.Guilds_Stickers_, api, g.ID.String(), stickerID.String())
}
