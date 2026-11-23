package net

import "encoding/json"

func IsJSONResponseOK(response string) bool {
	var array []json.RawMessage
	err := json.Unmarshal([]byte(response), &array)
	if err != nil {
		return false
	}
	if len(array) == 0 {
		return false
	}

	var status string
	err = json.Unmarshal(array[0], &status)
	if err != nil {
		return false
	}

	return status == "ok"
}
