package vo



type SelfCheckPhone struct{
	PhoneCaptcha string `json:"phone_captcha" binding:"required,gt=4"`
}
