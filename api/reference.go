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
	"strconv"
	"time"
)

/* API Versioning */

var (
	apiBase            = "https://discord.com/api"
	apiVersion         = fmt.Sprintf("/v%d", gatewayVersion)
	api                = apiBase + apiVersion
	discordEpoch int64 = 1420070400000
)

type Snowflake string

type FormattedSnowflake struct {
	Timestamp         int64
	InternalWorkerID  int64
	InternalProcessID int64
	Increment         int64
}

func (s Snowflake) GetSnowflake() Snowflake {
	return s
}

func (s Snowflake) String() string {
	return string(s)
}

func (s Snowflake) ToBinary() string {
	var b []byte

	for _, c := range s {
		b = strconv.AppendInt(b, int64(c), 2)
	}

	return string(b)
}

func StringToSnowflake(s string) *Snowflake {
	q := Snowflake(s)
	return &q
}

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

func (s Snowflake) Timestamp() time.Time {
	return time.Unix(0, s.ParseSnowflake().Timestamp)
}

/* User Agent */

const (
	UserAgent = "NowLiveCustomLib (https://nowlivebot.com, 1.0)"
)

/* Message Formatting */

type Format string

const (
	userFormat                Format = "<@%s>"
	userNicknameFormat               = "<@!%s>"
	ChannelFormat                    = "<#%s>"
	roleFormat                       = "<@&%s>"
	customEmojiFormat                = "<:%s:%s>"
	customEmojiAnimatedFormat        = "<a:%s:%s>"
	unitTimestampFormat              = "<t:%s>"
	unixTimestampStyledFormat        = "<t:%s:%s>"
)

type TimestampStyle string

const (
	ShortTime     TimestampStyle = "t"
	LongTime                     = "T"
	ShortDate                    = "d"
	LongDate                     = "D"
	ShortDateTime                = "f" // default
	LongDateTime                 = "F"
	RelativeTime                 = "R"
)

/*  Image Formatting */

const (
	CdnBase string = "https://cdn.discordapp.com/"
)

// PtrStr converts a string pointer to a string
func PtrStr(s *string) string {
	return *s
}
