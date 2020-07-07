package vo

import "RizhaoLanshanLabourUnion/utils"

type PhoneLogin struct {
	Phone            string `json:"phone" form:"phone" binding:"required" `
	Password         string `json:"password" form:"password" binding:"min=3,max=20,required"`
	Captcha          string `json:"captcha_code" form:"captcha_code" binding:"required"`
	CaptchaTimestamp string `json:"captcha_time" form:"captcha_time" binding:"required"  time_format:"2006-01-02 15:04:05"`
	CaptchaChallenge string `json:"captcha_challenge" form:"captcha_challenge" binding:"required"`
}

type UsernameLogin struct {
	Username         string `json:"username" form:"username" binding:"min=3,max=20,required" `
	Password         string `json:"password" form:"password" binding:"min=3,max=20,required"`
	Captcha          string `json:"captcha_code" form:"captcha_code" binding:"required"`
	CaptchaTimestamp string `json:"captcha_time" form:"captcha_time" binding:"required"  time_format:"2006-01-02 15:04:05"`
	CaptchaChallenge string `json:"captcha_challenge" form:"captcha_challenge" binding:"required"`
}

type TokenResult struct {
	Expire     *utils.Time `json"expire" example:"2048-01-02 15:04:05"`
	Token      string      `json:"token" example:"asdsadasdasdasd"`
	RefreshUrl string      `json:"refresh_url" example:"/api/auth/refreshToken"`
}

type LoginResult struct {
	Common
	TokenResult
}
