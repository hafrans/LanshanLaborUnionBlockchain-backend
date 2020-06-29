package utils

import "time"

func GenerateCaseId(areaCode string) string {
	t := time.Now()
	return areaCode + t.Format("20060102150405") + string(t.UnixNano())
}
