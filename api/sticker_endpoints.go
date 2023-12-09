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
	"fmt"

	log "github.com/veteran-software/nowlive-logging"
)

// GetSticker - Returns a sticker object for the given sticker ID.
func (s *Sticker) GetSticker() (*Sticker, error) {
	u := parseRoute(fmt.Sprintf(getSticker, api, s.ID.String()))

	var sticker *Sticker
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &sticker)

	return sticker, err
}

// ListStickerPacks - Returns the list of sticker packs available to Nitro subscribers.
//
//goland:noinspection GoUnusedExportedFunction
func ListStickerPacks() (*ListStickerPacksResponse, error) {
	u := parseRoute(fmt.Sprintf(listNitroStickerPacks, api))

	var listStickerPacksResponse *ListStickerPacksResponse
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &listStickerPacksResponse)

	return listStickerPacksResponse, err
}

// ListStickerPacksResponse - JSON response
type ListStickerPacksResponse struct {
	StickerPacks []*StickerPack `json:"sticker_packs"`
}

// ListGuildStickers - Returns an array of sticker objects for the given guild.
//
// Includes user fields if the bot has the CreateGuildExpressions or CreateGuildExpressions permission.
func (g *Guild) ListGuildStickers() ([]*Sticker, error) {
	u := parseRoute(fmt.Sprintf(listGuildStickers, api, g.ID.String()))

	var stickers []*Sticker
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &stickers)

	return stickers, err
}

// GetGuildSticker - Returns a sticker object for the given guild and sticker IDs.
//
// Includes the user field if the bot has the CreateGuildExpressions or CreateGuildExpressions permission.
func (g *Guild) GetGuildSticker(stickerID Snowflake) (*Sticker, error) {
	u := parseRoute(fmt.Sprintf(getGuildSticker, api, g.ID.String(), stickerID.String()))

	var sticker *Sticker
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &sticker)

	return sticker, err
}

// CreateGuildSticker - Create a new sticker for the guild.
//
// Send a multipart/form-data body.
//
// Requires the CreateGuildExpressions permission.
//
// Returns the new sticker object on success.
//
// Fires a GuildStickersUpdateGateway event.
//
// This endpoint supports the "X-Audit-Log-Reason" header.
//
// Lottie stickers can only be uploaded on guilds that have either the Verified and/or the Partnered guild feature.
//
// Uploaded stickers are constrained to 5 seconds in length for animated stickers, and 320 x 320 pixels.
// TODO: FormData fields
func (g *Guild) CreateGuildSticker() (*Sticker, error) {
	u := parseRoute(fmt.Sprintf(createGuildSticker, api, g.ID.String()))

	var sticker *Sticker
	responseBytes, err := firePostRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &sticker)

	return sticker, err
}

// ModifyGuildSticker - Modify the given sticker.
//
// For stickers created by the current user, requires either the CreateGuildExpressions or ManageGuildExpressions permission.
//
// For other stickers, requires the ManageGuildExpressions permission.
//
// Returns the updated Sticker object on success.
//
// Fires a Guild Stickers Update Gateway event.
//
// All parameters to this endpoint are optional.
//
// This endpoint supports the `X-Audit-Log-Reason` header.
func (g *Guild) ModifyGuildSticker(stickerID Snowflake,
	payload ModifyGuildStickerJSON,
	reason *string) (*Sticker,
	error) {
	u := parseRoute(fmt.Sprintf(modifyGuildSticker, api, g.ID.String(), stickerID.String()))

	var sticker *Sticker
	responseBytes, err := firePatchRequest(u, payload, reason)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &sticker)

	return sticker, err
}

// ModifyGuildStickerJSON - JSON payload
type ModifyGuildStickerJSON struct {
	Name        string  `json:"name"`        // name of the sticker (2-30 characters)
	Description *string `json:"description"` // description of the sticker (2-100 characters)
	Tags        string  `json:"tags"`        // autocomplete/suggestion tags for the sticker (max 200 characters)
}

// DeleteGuildSticker - Delete the given sticker.
//
// For stickers created by the current user, requires either the CreateGuildExpressions or ManageGuildExpressions permission.
//
// For other stickers, requires the ManageGuildExpressions permission.
//
// Returns "204 No Content" on success.
//
// This endpoint supports the `X-Audit-Log-Reason` header.
func (g *Guild) DeleteGuildSticker(stickerID Snowflake, reason *string) error {
	u := parseRoute(fmt.Sprintf(deleteGuildSticker, api, g.ID.String(), stickerID.String()))

	return fireDeleteRequest(u, reason)
}
