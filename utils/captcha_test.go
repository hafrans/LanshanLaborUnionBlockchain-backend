package utils_test

import (
	"RizhaoLanshanLabourUnion/utils"
	"encoding/json"
	"testing"
)

var pendingCaptcha *utils.PendingCaptcha

func TestCreateCaptcha(t *testing.T) {
	pendingCaptcha = utils.CreateCaptcha("hello")
}

func TestCheckCaptcha(t *testing.T) {
	text := pendingCaptcha.PlainText
	tt, _ := json.Marshal(*pendingCaptcha)
	t.Log(string(tt))
	if !utils.CheckCaptcha("hello",text,pendingCaptcha.Time.FormattedString(),pendingCaptcha.Challenge){
		t.Error("hello captcha failed 1")
	}

	if utils.CheckCaptcha("hellox",string(text),pendingCaptcha.Time.FormattedString(),pendingCaptcha.Challenge){
		t.Error("hello captcha failed 2")
	}

	if utils.CheckCaptcha("hello","asdsad",pendingCaptcha.Time.FormattedString(),pendingCaptcha.Challenge){
		t.Error("hello captcha failed 3")
	}

	if utils.CheckCaptcha("hello",string(text),"2099-08-11 12:34:56",pendingCaptcha.Challenge){
		t.Error("hello captcha failed")
	}

}
