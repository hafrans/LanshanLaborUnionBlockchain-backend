package smsqueue

import (
	"github.com/gorilla/schema"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type A map[string]string

var encoder *schema.Encoder = schema.NewEncoder()

func (sender *UnicomSMSSender) SendRequest(request *SendRequest) (string, error) {

	qvs := QueryValues{}

	qvs.Add("action", "send")
	qvs.Add("account", sender.Authenticator.Account)
	qvs.Add("password", sender.Authenticator.Password)
	qvs.Add("mobile", request.Mobile)
	qvs.Add("content", "【岚山区总工会】"+request.Content)

	client := http.DefaultClient

	req, err := http.NewRequest("POST", SmsServer, strings.NewReader(qvs.Encode()))
	if err != nil {
		return "", err
	}

	// set form content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf8")

	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(res))
	return "", nil
}
