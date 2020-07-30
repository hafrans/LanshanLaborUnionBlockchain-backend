package vo

type Comment struct {
	// 对质ID
	ID int64 `json:"id,omitempty"`

	// 关联案件号
	CaseID string `json:"case_id" binding:"required"`

	// 对质内容
	Content string `json:"content" binding:"required"`

	// 提交者
	Submitter string `json:"submitter,omitempty"`

	// 提交者类型
	SubmitterType int `json:"submitter_type,omitempty"`

	// 提交者联系方式
	SubmitterPhone string `json:"submitter_phone,omitempty"`
}
