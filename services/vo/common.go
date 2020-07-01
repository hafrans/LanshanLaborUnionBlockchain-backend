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

type Captcha struct {
	Captcha          string `json:"captcha_code" form:"captcha_code" binding:"required" example:"123456"`
	CaptchaTimestamp string `json:"captcha_time" form:"captcha_time" binding:"required"  time_format:"2006-01-02 15:04:05"`
	CaptchaChallenge string `json:"captcha_challenge" form:"captcha_challenge" binding:"required" example:"asbduiasdvasilvdiwlqdulisbdaiauldvil=="`
}

func GenerateCommonResponseHead(status int, message string) Common {
	return Common{
		Timestamp: utils.NowTime(),
		Status:    status,
		Message:   message,
	}
}
