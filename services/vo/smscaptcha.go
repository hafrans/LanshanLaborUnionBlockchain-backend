package vo

import "RizhaoLanshanLabourUnion/utils"

// 申请验证码
type SMSCaptchaRequest struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Captcha `binding:"required,dive"`
}


// 验证码挑战握手信息
type SMSCaptchaResponse struct {
	Identifier string      `json:"challenge_code"`
	Timestamp  *utils.Time `json:"challenge_time"`
}


// 验证码检查，脱离后台数据库
type SMSCaptchaCheckRequest struct {
	Identifier string      `json:"challenge_code" binding:"required"`
	Timestamp  *utils.Time `json:"challenge_time" binding:"required"`
	Phone      string      `json:"phone" binding:"required"`
	Code       string      `json:"captcha_code" binding:"required,len=6"`
}
