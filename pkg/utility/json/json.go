package json

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
)

func ToJsonString(any interface{}) string {
	if any == nil {
		log.Debugf("The source any param is empty, return empty string")
		return ""
	}

	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Errorf("Failed to convert any param to json, return empty string, source any: %+v", any)
		return ""
	}

	return string(jsonBytes)
}
