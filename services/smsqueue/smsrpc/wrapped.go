package smsrpc

import (
	"RizhaoLanshanLabourUnion/services/smsqueue"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"html/template"
	"time"
)

var SMSCache = utils.NewLRUCache(32768)

// 发送验证码
func SendCaptcha(Phone string) (*vo.SMSCaptchaResponse, error) {
	// send buffer
	var buf bytes.Buffer

	// generate captcha
	captchaCode := captcha.RandomDigits(6)

	// inject template
	tpl, err := template.New("captcha").Parse(smsqueue.SMSContentCaptcha)
	if err != nil {
		return nil, err
	}
	tpl.Execute(&buf, struct {
		Code string
	}{
		Code: fmt.Sprintf(
			"%d%d%d%d%d%d", captchaCode[0],
			captchaCode[1],
			captchaCode[2],
			captchaCode[3],
			captchaCode[4],
			captchaCode[5]),
	})
	sendNowTime := time.Now()

	hash := sha256.New()
	hash.Write([]byte(sendNowTime.Format(time.RFC3339) + Phone + string(captchaCode) + utils.SMSSetting.Password))
	hashResult := hash.Sum(nil)
	challengeCode := base64.StdEncoding.EncodeToString(hashResult)

	// 异步发送成功
	SendMessage(utils.SMSSetting.Account, utils.SMSSetting.Password, Phone, buf.String())
	respTime := utils.Time(sendNowTime)

	// 注入数据
	SMSCache.Put(Phone, sendNowTime)
	return &vo.SMSCaptchaResponse{
		Identifier: challengeCode,
		Timestamp:  &respTime,
	}, nil

}
