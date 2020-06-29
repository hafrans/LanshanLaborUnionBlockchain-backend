package utils

import (
	"strconv"
	"time"
)

func GenerateCaseId(areaCode string) string {
	t := time.Now()
	return areaCode + t.Format("20060102150405") + strconv.Itoa(int(t.UnixNano()))
}
