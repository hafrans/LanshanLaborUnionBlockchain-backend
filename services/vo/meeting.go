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
}
