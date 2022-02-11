package utilities

import (
	"fmt"
	"strings"
)

func ParseQueryString(opts *map[string]interface{}) *string {
	var qsp []string
	for k, v := range *opts {
		if strings.ToLower(k) == "around" ||
			strings.ToLower(k) == "before" ||
			strings.ToLower(k) == "after" ||
			strings.ToLower(k) == "limit" ||
			strings.ToLower(k) == "with_counts" ||
			strings.ToLower(k) == "with_expiration" ||
			strings.ToLower(k) == "guild_scheduled_event_id" {
			qsp = append(qsp, fmt.Sprintf("%s=%v", strings.ToLower(k), v))
		}
	}

	var q string
	if len(qsp) > 0 {
		q = "?" + strings.Join(qsp, "&")
		return &q
	}

	return nil
}
