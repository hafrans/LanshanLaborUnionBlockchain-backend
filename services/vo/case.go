package vo

import "RizhaoLanshanLabourUnion/services/models"

type CaseFirstSubmitForm struct {

	// 类型id
	CategoryId int64 `json:"category_id"`

	// 1.申请人
	Applicant *models.Applicant `json:"applicant" binding:"required,dive"`

	// 2.被申请人
	Respondent *models.Employer `json:"respondent" binding:"required,dive"`

	// 3.调解事项
	Title string `json:"title" binding:"required"`

	// 4.调解事项所依据的事实与理由
	Content string `json:"content" binding:"required"`

	// 5.劳动者填写的表单 id
	FormId int64 `json:"form_id" binding:"required"`

	// 6. 证据材料 各种id
	Materials []int64 `json:"materials"`

	// 7.调解笔录， 8. 调解意见 不用填写

}
