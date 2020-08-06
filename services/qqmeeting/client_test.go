package qqmeeting_test

import (
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"
)

var meeting = qqmeeting.Meeting{

}

func TestNewRequest2(t *testing.T) {

	str := `GET
X-TC-Key=gfpuPNeBAq7jRL0hybQ3zrFVlM5ZwYsSmOTC&X-TC-Nonce=18081&X-TC-Timestamp=1596535863
/v1/users/list?page=1&page_size=1
`
	hm := hmac.New(crypto.SHA256.New, []byte("zVMF9Z0erw1kSH54CBpNGu6cgxyRmDbX"))
	hm.Write([]byte(str))
	result := hm.Sum(nil)
	log.Printf("%x\n", result)
	log.Println(hex.EncodeToString(result))
}

func TestNewRequest(t *testing.T) {

	req, err := qqmeeting.NewRequest("GET", "http://api.meeting.qq.com/v1/users/list?page=2&page_size=1", "", meeting)
	if err != nil {
		t.Error(err)
	} else {
		client := qqmeeting.GetHttpClient()
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		} else {
			content, _ := ioutil.ReadAll(resp.Body)
			t.Log(string(content))
		}
	}

}

func TestMeeting_Do_CreateMeeting(t *testing.T) {

	resp, err := meeting.Do(qqmeeting.MeetingCreateRequest{
		Settings: &qqmeeting.Settings{
			MuteAll:         true,
			AllowUnmuteSelf: true,
		},
		InstanceID: qqmeeting.InstancePC,
		Type:       qqmeeting.MeetingTypeBookingMeeting,
		UserID:     "17854176681",
		Hosts: []*qqmeeting.UserObj{

		},
		StartTime: strconv.Itoa(int(time.Now().Unix())),
		EndTime:   strconv.Itoa(int(time.Now().Unix() + 3600)),
		Subject:   "测试会议",
	})

	if err != nil {
		t.Error(err)
	} else {
		r := resp.(qqmeeting.MeetingCreateResponse)
		t.Log(r.MeetingCreationInfo[0].Subject)
		t.Log(r.MeetingCreationInfo[0].StartTime)
		t.Log(r.MeetingCreationInfo[0].EndTime)
		t.Log("meeting code:", r.MeetingCreationInfo[0].MeetingCode)
		t.Log("meeting id:", r.MeetingCreationInfo[0].MeetingID)
		t.Log("meeting join url:", r.MeetingCreationInfo[0].JoinUrl)
	}
}

func TestMeeting_Do_CreateUser(t *testing.T) {

	_, err := meeting.Do(qqmeeting.UserCreateRequest{
		UserInfo: qqmeeting.UserInfo{
			UserID:   "17854176681",
			Username: "石臼所",
			Phone:    "17854176681",
			Email:    "hafrans@163.com",
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户创建成功")
	}
}

func TestMeeting_Do_QueryUserDetail(t *testing.T) {

	detail, err := meeting.Do(qqmeeting.UserDetailQueryRequest{
		UserID: "17854176681",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户信息获取成功")
		d, ok := detail.(qqmeeting.UserDetailQueryResponse)
		if ok {
			t.Log(d.UserID)
			t.Log(d.Email)
			t.Log(d.Status)
			t.Log(d.Username)
			t.Log(d.AreaCode)
			t.Log(d.UpdateTime)
		} else {
			t.Error("判断错误")
		}
	}
}

func TestMeeting_Do_UpdateUser(t *testing.T) {

	_, err := meeting.Do(qqmeeting.UserDetailUpdateRequest{
		Username: "秘书长王五",
		Email:    "hafrans@test.com",
		UserID:   "17854176681",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户修改成功")
	}
}

func TestMeeting_Do_DeleteUser(t *testing.T) {

	_, err := meeting.Do(qqmeeting.UserDeleteRequest{
		UserID: "17854176682",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户修改成功")
	}
}

func TestMeeting_Do_ListUser(t *testing.T) {

	list, err := meeting.Do(qqmeeting.UserListRequest{
		Page:     1,
		PageSize: 20,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("数据获取成功")

		for _, d := range list.(qqmeeting.UserListResponse).Users {
			t.Log("-------------------------")
			t.Log(d.UserID)
			t.Log(d.Email)
			t.Log(d.Status)
			t.Log(d.Username)
			t.Log(d.AreaCode)
			t.Log(d.UpdateTime)
			t.Log("-------------------------")
		}
	}
}
