package models

import "RizhaoLanshanLabourUnion/utils"

type Meeting struct {
	Model
	CaseID      string      `json:"case_id" gorm:"type:varchar(64);index:cid"`
	CreatorID   string      `json:"creator_id"` // 对应 QQMeeting 的UserID
	UserID      int64       `json:"-" gorm:"index:uid"`
	MeetingID   string      `json:"meeting_id" gorm:"unique_index:mid"`
	MeetingCode string      `json:"meeting_code" gorm:"index:mc"`
	InstanceID  int         `json:"instance_id"`
	Subject     string      `json:"subject"`
	Type        int         `json:"type"`
	StartTime   *utils.Time `json:"start_time"`
	EndTime     *utils.Time `json:"end_time"`
	Password    *string     `json:"password"`
	JoinUrl     string      `json:"join_url"`
}
