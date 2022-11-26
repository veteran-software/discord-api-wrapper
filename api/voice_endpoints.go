package api

import (
	"encoding/json"
	"fmt"
	"github.com/veteran-software/discord-api-wrapper/v10/logging"
	"io"
	"net/http"
)

// ListVoiceRegions - Returns an array of voice region objects that can be used when setting a voice or stage channel's rtc_region.
//
//goland:noinspection GoUnusedExportedFunction
func ListVoiceRegions() *[]VoiceRegion {
	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(listVoiceRegions, api), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var voiceRegions *[]VoiceRegion
	err = json.NewDecoder(resp.Body).Decode(&voiceRegions)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return voiceRegions
}
