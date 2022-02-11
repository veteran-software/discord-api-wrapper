package api

import (
	"strconv"
	"time"
)

/* API Versioning */

const apiBase string = "https://discord.com/api"
const apiVersion string = "/v9"
const api = apiBase + apiVersion

/* Authentication */

/* Snowflakes */

const discordEpoch int64 = 1420070400000

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
