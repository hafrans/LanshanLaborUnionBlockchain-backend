package dao_test

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"RizhaoLanshanLabourUnion/utils"
	"testing"
)

func init(){
	utils.InitTestSetting()
	qqmeeting.InitMeeting()
	println("meeting initialized")
}

func TestCreateMeeting(t *testing.T) {
	meeting := &models.Meeting{
		JoinUrl: "asdsad",
		EndTime: utils.NowTime(),
		StartTime: utils.NowTime(),
		MeetingCode: "213214214",
		MeetingID: "21181651",
		Subject: "asdbiwuqbdouqw",
		UserID: 1,
		InstanceID: 1,
		Type: 1,
		CaseID: "281896",
		CreatorID: "213213124213",
	}
	dao.CreateMeeting(meeting)
}

