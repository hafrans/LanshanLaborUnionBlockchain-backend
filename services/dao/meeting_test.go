package dao_test

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"RizhaoLanshanLabourUnion/utils"
	"log"
	"testing"
)

func init(){
	utils.InitTestSetting()
	qqmeeting.InitMeeting()
	println("meeting initialized")
}

func TestCreateMeeting(t *testing.T) {
	//meeting := &models.Meeting{
	//	JoinUrl: "asdsad",
	//	EndTime: time.Now(),
	//	StartTime: time.Now(),
	//	MeetingCode: "213214214",
	//	MeetingID: "21181651",
	//	Subject: "asdbiwuqbdouqw",
	//	UserID: 1,
	//	InstanceID: 1,
	//	Type: 1,
	//	CaseID: "281896",
	//	CreatorID: "213213124213",
	//}
	//dao.CreateMeeting(meeting)
}


func TestGetMeetingAllWithConditionPaginated(t *testing.T) {

	_,b,c := dao.GetMeetingAllWithConditionPaginated(nil, nil, true, 1, 22)
	if c != nil {
		t.Error(c)
	}else{
		t.Log("TOTAL:",b)
	}
}

func TestGetMeetingPersonnelsByMeetingID(t *testing.T) {
	person, err := dao.GetMeetingPersonnelsByMeetingID(8);
	if err != nil {
		log.Println(err)
		t.Fail()
	}else{
		for _, v := range person{
			t.Log(v.UserID,v.Username,v.User.Phone)
		}
	}

}
