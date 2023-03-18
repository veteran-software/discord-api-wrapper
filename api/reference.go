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
	"fmt"
)

var (
	apiBase    = "https://discord.com/api"
	apiVersion = fmt.Sprintf("/v%d", gatewayVersion)
	api        = apiBase + apiVersion
)

const (
	// UserAgent - header value to be sent with each API request
	UserAgent = "NowLiveCustomLib (https://nowlivebot.com, 10.0.19)"
)

// Format - Discord utilizes a subset of markdown for rendering message content on its clients, while also adding some custom functionality to enable things like mentioning users and channels.
type Format string

//goland:noinspection GoUnusedConst
const (
	userFormat                Format = "<@%s>"     // userFormat - <@USER_ID>
	userNicknameFormat        Format = "<@!%s>"    // userNicknameFormat - <@!USER_ID>
	ChannelFormat             Format = "<#%s>"     // ChannelFormat - <#CHANNEL_ID>
	roleFormat                Format = "<@&%s>"    // roleFormat - <@&ROLE_ID>
	customEmojiFormat         Format = "<:%s:%s>"  // customEmojiFormat - <:NAME:ID>
	customEmojiAnimatedFormat Format = "<a:%s:%s>" // customEmojiAnimatedFormat - <a:NAME:ID>
	unixTimestampFormat       Format = "<t:%s>"    // unixTimestampFormat - <t:TIMESTAMP>
	unixTimestampStyledFormat Format = "<t:%s:%s>" // unixTimestampStyledFormat - <t:TIMESTAMP:STYLE>
)

// TimestampStyle - Timestamps will display the given timestamp in the user's timezone and locale.
type TimestampStyle string

//goland:noinspection GoUnusedConst
const (
	ShortTime     TimestampStyle = "t" // ShortTime - 16:20
	LongTime      TimestampStyle = "T" // LongTime - 16:20:30
	ShortDate     TimestampStyle = "d" // ShortDate - 20/04/2021
	LongDate      TimestampStyle = "D" // LongDate - 20 April 2021
	ShortDateTime TimestampStyle = "f" // ShortDateTime - 20 April 2021 16:20; default
	LongDateTime  TimestampStyle = "F" // LongDateTime - Tuesday, 20 April 2021 16:20
	RelativeTime  TimestampStyle = "R" // RelativeTime - 2 months ago
)

const (
	// ImageBaseURL - The root URL for image links
	ImageBaseURL string = "https://cdn.discordapp.com/"
)

// PtrStr converts a string pointer to a string
func PtrStr(s *string) string {
	return *s
}

// LocalizationDict - officially supported languages by Discord
type LocalizationDict struct {
	Danish              string `json:"da,omitempty"`
	German              string `json:"de,omitempty"`
	EnglishUK           string `json:"en-GB,omitempty"`
	EnglishUS           string `json:"en-US,omitempty"`
	Spanish             string `json:"es-ES,omitempty"`
	French              string `json:"fr,omitempty"`
	Croatian            string `json:"hr,omitempty"`
	Italian             string `json:"it,omitempty"`
	Lithuanian          string `json:"lt,omitempty"`
	Hungarian           string `json:"hu,omitempty"`
	Dutch               string `json:"nl,omitempty"`
	Norwegian           string `json:"no,omitempty"`
	Polish              string `json:"pl,omitempty"`
	PortugueseBrazilian string `json:"pt-BR,omitempty"`
	Romanian            string `json:"ro,omitempty"`
	Finnish             string `json:"fi,omitempty"`
	Swedish             string `json:"sv-SE,omitempty"`
	Vietnamese          string `json:"vi,omitempty"`
	Turkish             string `json:"tr,omitempty"`
	Czech               string `json:"cs,omitempty"`
	Greek               string `json:"el,omitempty"`
	Bulgarian           string `json:"bg,omitempty"`
	Russian             string `json:"ru,omitempty"`
	Ukrainian           string `json:"uk,omitempty"`
	Hindi               string `json:"hi,omitempty"`
	Thai                string `json:"th,omitempty"`
	ChineseChina        string `json:"zh-CN,omitempty"`
	Japanese            string `json:"ja,omitempty"`
	ChineseTaiwan       string `json:"zh-TW,omitempty"`
	Korean              string `json:"ko,omitempty"`
}
