package qqmeeting

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var ApiHost = "https://api.meeting.qq.com/v1"

type Meeting struct {
	SecretKey string
	SecretID  string
	AppID     string
	SdkId     string
}

type Request struct {
	Method        string
	URL           *url.URL
	Secret        string
	Body          string
	Key           string `json:"X-TC-Key"`
	Timestamp     int64  `json:"X-TC-Timestamp"`
	Nonce         int    `json:"X-TC-Nonce"`
	Signature     string `json:"X-TC-Signature"`
	AppID         string `json:"AppId"`
	SdkId         string `json:"SdkId"`
	Version       string `json:"X-TC-Version"`
	Registered    int    `json:"X-TC-Registered"`
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
}

func newMeetingRequest(method, path, body string, meeting Meeting) *Request {

	req := new(Request)

	req.ContentType = "application/json"
	req.Method = method
	req.URL, _ = url.Parse(path)
	req.Secret = meeting.SecretKey
	req.Body = body
	req.ContentLength = strconv.Itoa(len(body))

	req.Timestamp = time.Now().Unix()
	req.Key = meeting.SecretID
	req.Nonce = rand.Intn(10000) + 10000
	req.Version = "1.0.2"
	req.AppID = meeting.AppID
	req.Registered = 0
	req.SdkId = meeting.SdkId

	return req

}

func fillSignature(req *Request) {
	stringToSign := req.Method + "\n" + generateHeaderString(req) + "\n" + req.URL.Path + "?" + req.URL.RawQuery + "\n" + req.Body
	hm := hmac.New(crypto.SHA256.New, []byte(req.Secret))
	hm.Write([]byte(stringToSign))
	result := hm.Sum(nil)
	req.Signature = base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(result)))
}

func NewRequest(method, url, body string, meeting Meeting) (*http.Request, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	mReq := newMeetingRequest(method, url, body, meeting)
	fillSignature(mReq)
	fillHeader(mReq, &req.Header)

	return req, nil
}

func generateHeaderString(request *Request) string {

	var buf bytes.Buffer
	//counter := 0
	//callback := func(n, v string) {
	//	if counter > 0 {
	//		buf.WriteString("&")
	//	}
	//	buf.WriteString(fmt.Sprintf("%s=%s", n, v))
	//	counter++
	//}
	//fillFields(request, callback)
	buf.WriteString("X-TC-Key=" + request.Key +
		"&" + "X-TC-Nonce=" + strconv.Itoa(request.Nonce) +
		"&" + "X-TC-Timestamp=" + strconv.Itoa(int(request.Timestamp)))
	log.Println(buf.String())
	return buf.String()
}

func fillHeader(req *Request, header *http.Header) () {

	callback := func(n, v string) () {
		(*header)[n] = []string{v}
	}
	fillFields(req, callback)
}

func fillFields(req *Request, callback func(name, value string) ()) {

	ref := reflect.ValueOf(*req)
	typ := reflect.TypeOf(*req)

	numRef := ref.NumField()

	for i := 0; i < numRef; i++ {

		field := ref.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Tag.Get("json") != "" {
			switch field.Kind() {
			case reflect.String:
				callback(fieldType.Tag.Get("json"), field.String())
			case reflect.Int64:
				callback(fieldType.Tag.Get("json"), strconv.Itoa(int(field.Int())))
			case reflect.Int:
				callback(fieldType.Tag.Get("json"), strconv.Itoa(int(field.Int())))
			}
		}

	}

}

func SendRequest() {

}

func InitQQMeeting() {

}
