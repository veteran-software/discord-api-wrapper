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

package api

// Sticker - Represents a sticker that can be sent in messages.
type Sticker struct {
	ID          Snowflake   `json:"id"`                   // id of the sticker
	PackID      Snowflake   `json:"pack_id,omitempty"`    // for standard stickers, id of the pack the sticker is from
	Name        string      `json:"name"`                 // name of the sticker
	Description *string     `json:"description"`          // description of the sticker
	Tags        string      `json:"tags,omitempty"`       // autocomplete/suggestion tags for the sticker (max 200 characters)
	Type        StickerType `json:"type"`                 // type of sticker
	FormatType  int         `json:"format_type"`          // type of sticker format
	Available   bool        `json:"available,omitempty"`  // whether this guild sticker can be used, may be false due to loss of Server Boosts
	GuildID     Snowflake   `json:"guild_id,omitempty"`   // id of the guild that owns this sticker
	User        User        `json:"user,omitempty"`       // the user that uploaded the guild sticker
	SortValue   int         `json:"sort_value,omitempty"` // the standard sticker's sort order within its pack
}

// StickerType - type of sticker
type StickerType int

//goland:noinspection GoUnusedConst
const (
	StickerTypeStandard StickerType = iota + 1 // an official sticker in a pack, part of Nitro or in a removed purchasable pack
	StickerTypeGuild                           // a sticker uploaded to a Boosted guild for the guild's members
)

// StickerFormatType - The format of the Sticker
type StickerFormatType int

//goland:noinspection GoUnusedConst
const (
	StickerFormatTypePng         StickerFormatType = iota + 1 // PNG
	StickerFormatTypeAnimatedPng                              // APNG
	StickerFormatTypeLottie                                   // LOTTIE
	StickerFormatTypeGif                                      // GIF
)

// StickerItem - The smallest amount of data required to render a sticker. A partial sticker object.
type StickerItem struct {
	ID         Snowflake         `json:"id"`          // id of the sticker
	Name       string            `json:"name"`        // name of the sticker
	FormatType StickerFormatType `json:"format_type"` // type of sticker format
}

// StickerPack - Represents a pack of standard stickers.
type StickerPack struct {
	ID             Snowflake  `json:"id"`                         // id of the sticker pack
	Stickers       []*Sticker `json:"stickers"`                   // the stickers in the pack
	Name           string     `json:"name"`                       // name of the sticker pack
	SkuID          Snowflake  `json:"sku_id"`                     // id of the pack's SKU
	CoverStickerID Snowflake  `json:"cover_sticker_id,omitempty"` // id of a sticker in the pack which is shown as the pack's icon
	Description    string     `json:"description"`                // description of the sticker pack
	BannerAssetID  Snowflake  `json:"banner_asset_id,omitempty"`  // id of the sticker pack's banner image
}
