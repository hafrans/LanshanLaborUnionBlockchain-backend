package utils

import "time"

func CreateCaseId(areaCode string) string {
	t := time.Now()
	return areaCode + t.Format("20060102150405") + string(t.UnixNano())
}
