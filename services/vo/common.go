package vo

import (
	"RizhaoLanshanLabourUnion/utils"
)

type Common struct {
	Status    int         `json:"status" example:"401"`
	Message   string      `json:"message" example:"unauthorized"`
	Timestamp *utils.Time `json:"timestamp" example:"2048-05-06 12:34:56"`
}

type CommonData struct {
	Common
	Data interface{} `json:"data"`
}



func GenerateCommonResponseHead(status int, message string) Common{
	return Common{
		Timestamp: utils.NowTime(),
		Status: status,
		Message: message,
	}
}
