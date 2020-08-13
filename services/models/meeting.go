package models

import (
	"time"
)

type Meeting struct {
	Model
	CaseID      string              `json:"case_id" gorm:"type:varchar(64);index:cid"`
	CreatorID   string              `json:"creator_id"` // 对应 QQMeeting 的UserID
	UserID      int64               `json:"-" gorm:"index:uid"`
	MeetingID   string              `json:"meeting_id" gorm:"unique_index:mid"`
	MeetingCode string              `json:"meeting_code" gorm:"index:mc"`
	InstanceID  int                 `json:"instance_id"`
	Subject     string              `json:"subject"`
	Type        int                 `json:"type"`
	StartTime   time.Time           `json:"start_time"`
	EndTime     time.Time           `json:"end_time"`
	Password    *string             `json:"password"`
	JoinUrl     string              `json:"join_url"`
	Personnel   []*MeetingPersonnel `json:"personnel"`
	User        User                `json:"user"`
}

type MeetingPersonnel struct {
	Model
	MeetingID   int64  `json:"meeting_id"` // 对应的是Meeting 中的ID，不是MeetingID
	Username    string `json:"username"`
	Userid      string `json:"meeting_userid"` // 对应腾讯会议的userid，一般是手机号
	UserID      int64  `json:"user_id"`
	MeetingRole int    `json:"meeting_role"` // 对应的会议的角色
}
