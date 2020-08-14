package vo

import "RizhaoLanshanLabourUnion/utils"

type SimplifiedMeetingVO struct {
	ID          int64       `json:"id"`
	MeetingCode string      `json:"meeting_code"`
	CaseID      string      `json:"case_id"`
	Subject     string      `json:"subject"`
	StartTime   *utils.Time `json:"start_time"`
	EndTime     *utils.Time `json:"end_time"`
	JoinUrl     string      `json:"join_url"`
	Creator     string      `json:"creator"`
}

type MeetingVO struct {
	ID          int64       `json:"id"`
	MeetingCode string      `json:"meeting_code"`
	CaseID      string      `json:"case_id"`
	Subject     string      `json:"subject"`
	StartTime   *utils.Time `json:"start_time"`
	EndTime     *utils.Time `json:"end_time"`
	JoinUrl     string      `json:"join_url"`
	Creator     string      `json:"creator"`
	Host        []string    `json:"host"`
	Type        int         `json:"type"`
	Password    string      `json:"password"`
}

type MeetingCreateForm struct {
	Subject         string      `json:"subject" binding:"required"`
	CaseID          string      `json:"case_id" binding:"required"`
	StartTime       *utils.Time `json:"start_time" binding:"required"`
	EndTime         *utils.Time `json:"end_time"  binding:"required"`
	Hosts           []int64     `json:"hosts"`
	Invitees        []int64     `json:"invitee"`
	Password        string      `json:"password"`
	MuteEnableJoin  bool        `json:"mute_enable_join"`                 // 入会时静音
	AllowUnmuteSelf bool        `json:"allow_unmute_self"`                // 允许参会者取消静音
	MuteAll         bool        `json:"mute_all,omitempty"`               // 全体静音
	Type            int         `json:"meeting_type"  binding:"required"` // 会议类型
}
