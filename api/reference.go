/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
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
	"strconv"
	"time"
)

var (
	apiBase            = "https://discord.com/api"
	apiVersion         = fmt.Sprintf("/v%d", gatewayVersion)
	api                = apiBase + apiVersion
	discordEpoch int64 = 1420070400000
)

// Snowflake - Discord utilizes Twitter's snowflake format for uniquely identifiable descriptors (IDs).
//
// These IDs are guaranteed to be unique across all of Discord, except in some unique scenarios in which child objects share their parent's ID.
//
// Because Snowflake IDs are up to 64 bits in size (e.g. a uint64), they are always returned as strings in the HTTP API to prevent integer overflows in some languages.
//
// See Gateway ETF/JSON for more information regarding Gateway encoding.
type Snowflake string

// FormattedSnowflake - A breakdown of the data contained in a Snowflake
type FormattedSnowflake struct {
	Timestamp         int64
	InternalWorkerID  int64
	InternalProcessID int64
	Increment         int64
}

// String - Type converts a Snowflake into a string
func (s Snowflake) String() string {
	return string(s)
}

// ToBinary - Type converts a Snowflake into its binary representation
func (s Snowflake) ToBinary() string {
	var b []byte

	for _, c := range s {
		b = strconv.AppendInt(b, int64(c), 2)
	}

	return string(b)
}

// StringToSnowflake - Type converts a string into a Snowflake
func StringToSnowflake(s string) *Snowflake {
	q := Snowflake(s)
	return &q
}

// ParseSnowflake - Breaks down a Snowflake and assigns each value to the FormattedSnowflake struct
func (s Snowflake) ParseSnowflake() FormattedSnowflake {
	bin := s.ToBinary()
	tStamp, _ := strconv.ParseInt(bin[0:42], 2, 64)
	worker, _ := strconv.ParseInt(bin[42:47], 2, 64)
	process, _ := strconv.ParseInt(bin[47:52], 2, 64)
	incr, _ := strconv.ParseInt(bin[52:64], 2, 64)
	return FormattedSnowflake{
		Timestamp:         tStamp + discordEpoch,
		InternalWorkerID:  worker,
		InternalProcessID: process,
		Increment:         incr,
	}
}

// Timestamp - Extracts only the timestamp from a Snowflake
//
// Useful for determining when the object belonging to the Snowflake was created
func (s Snowflake) Timestamp() time.Time {
	return time.Unix(0, s.ParseSnowflake().Timestamp)
}

const (
	// UserAgent - header value to be sent with each API request
	UserAgent = "NowLiveCustomLib (https://nowlivebot.com, 1.0)"
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
