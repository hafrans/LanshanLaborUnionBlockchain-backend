package models


// 部门调解意见
type Suggestion struct {
	Model
	// 意见
	Content string `json:"suggestion_content" gorm:"type:text"`
	// caseID， 新建表单时不要上传该信息
	CaseID string `json:"case_id" gorm:"type:varchar(64);index"`

	// 谁录入的笔录
	UserID int64 `json:"-"`
	User   *User `json:"user,omitempty"`
}
