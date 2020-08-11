package qqmeeting

import (
	"RizhaoLanshanLabourUnion/utils"
)

var MeetingClient *Meeting

func InitMeeting() *Meeting{

	MeetingClient = &Meeting{
		Version:    "0.9.8",
		Registered: EnableRegistered,
		SecretKey:  utils.QQMeetingSettings.SecretKey,
		AppID:      utils.QQMeetingSettings.AppID,
		SecretID:   utils.QQMeetingSettings.SecretID,
		SdkId:      utils.QQMeetingSettings.SDKID,
	}

	return MeetingClient
}
