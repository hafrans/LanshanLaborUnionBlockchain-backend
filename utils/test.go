package utils

import "encoding/json"

func GetStructJsonString(value interface{}) string{

	result, _ := json.Marshal(value)
	return string(result)
}