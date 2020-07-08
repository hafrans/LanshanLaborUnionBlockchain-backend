package models

// 调解笔录
type Record struct {
	Model
	// 记录介绍
	Name string `json:"record_name"`
	// caseID， 新建表单时不要上传该信息
	CaseID string `json:"case_id"  gorm:"type:varchar(64);index"`
	// 截图/材料等地址
	Path string `json:"record_path"`

	// 谁录入的笔录
	UserID int64 `json:"-" `
	User   *User `json:"user,omitempty" gorm:"PRELOAD:true"`
}
