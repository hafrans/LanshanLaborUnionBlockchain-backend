package models

type Comment struct {
	Model
	CaseID  string `json:"case_id" gorm:"type:varchar(64);index"`
	UserID  int64  `json:"user_id"`
	User    *User  `json:"user"`
	Content string `json:"content"`
}
