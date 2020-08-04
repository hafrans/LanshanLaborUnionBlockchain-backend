package qqmeeting_test

import (
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestNewRequest2(t *testing.T) {

	str := `GET
X-TC-Key=gfpuPNeBAq7jRL0hybQ3zrFVlM5ZwYsSmOTC&X-TC-Nonce=18081&X-TC-Timestamp=1596535863
/v1/users/list?page=1&page_size=1
`
	hm := hmac.New(crypto.SHA256.New,[]byte("zVMF9Z0erw1kSH54CBpNGu6cgxyRmDbX"))
	hm.Write([]byte(str))
	result := hm.Sum(nil)
	log.Printf("%x\n",result)
	log.Println(hex.EncodeToString(result))
}

func TestNewRequest(t *testing.T) {

	meeting := qqmeeting.Meeting{
		SdkId: "20xxxx011934",
		AppID: "2xxxx9",
		SecretID: "EXAMPLE",
		SecretKey: "EXAMPLE",
	}

	req, err := qqmeeting.NewRequest("GET","https://api.meeting.qq.com/v1/users/list?page=1&page_size=1","",meeting)
	if err != nil {
		t.Error(err)
	}else{

		client := http.DefaultClient
		t.Log(req.ContentLength)
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}else{
			content, _ := ioutil.ReadAll(resp.Body)
			t.Log(string(content))
		}
	}

}
