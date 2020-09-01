package smsqueue_test

import (
	"RizhaoLanshanLabourUnion/services/smsqueue"
	"RizhaoLanshanLabourUnion/utils"
	"net/url"
	"testing"
)
import "github.com/gorilla/schema"

var sms *smsqueue.UnicomSMSSender

func init() {

	utils.InitTestSetting()
	sms = &smsqueue.UnicomSMSSender{
		Authenticator: &smsqueue.Authenticator{
			Account:  utils.SMSSetting.Account,
			Password: utils.SMSSetting.Password,
		},
	}
}

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func TestGorilla(t *testing.T) {

	encoder := schema.NewEncoder()
	form := url.Values{}
	form.Set("action", "send")

	acc := smsqueue.Authenticator{
		Password: "15158",
		Account:  "2+965218",
	}

	if err := encoder.Encode(&acc, form); err != nil {
		t.Error(err)
	} else {

		req := smsqueue.SendRequest{Content: "FFox", Mobile: "1380000000"}
		_ = encoder.Encode(&req, form)

		t.Log(form.Encode())
	}

}

func TestUnicomSMSSender_SendRequest(t *testing.T) {
	resp, err := sms.SendRequest(&smsqueue.SendRequest{
		Mobile:  "15163372783",
		Content: "DebugMessage",
	})
	t.Log("Account:", sms.Authenticator.Account)
	t.Log("Pwd:", sms.Authenticator.Password)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}
