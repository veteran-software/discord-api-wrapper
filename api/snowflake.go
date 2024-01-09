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

import (
	"strconv"
	"time"
)

var discordEpoch int64 = 1420070400000

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
