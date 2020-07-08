package vo

// 调解笔录
type Record struct {
	// ID 新建不需要填写
	ID int64 `json:"id,omitempty"`

	// 记录介绍
	Name string `json:"name" binding:"required"`

	// 截图/材料等地址
	Path string `json:"path" binding:"required"`

	// 案件号，37xxx开头
	CaseID string `json:"case_id" binding:"required"`

	// 录入人, 新建不需要填写
	Submitter string `json:"submitter,omitempty"`

	// 录入人电话，新建不需要填写
	SubmitterPhone string `json:"submitter_phone,omitempty"`

	// 部门信息，新建不需要填写
	DepartmentInfo *DepartmentVO `json:"department_info"`
}
