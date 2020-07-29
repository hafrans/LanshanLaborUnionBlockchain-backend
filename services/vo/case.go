package vo

import (
	"RizhaoLanshanLabourUnion/utils"
)

// 类型

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"category_name"`
	Description string `json:"category_description"`
}

// 申请调解人
type Applicant struct {
	// 姓名
	Name string `json:"applicant_name" form:"applicant_name" gorm:"type:varchar(128);not null" binding:"required"`
	// 生日
	Birthday *utils.Date `json:"applicant_birth" form:"applicant_birth" gorm:"type:date" binding:"required"`
	// 民族
	Nationality string `json:"applicant_nationality" form:"applicant_nationality" gorm:"type:varchar(32);not null" binding:"required"`
	// 身份证号
	IdentityNumber string `json:"applicant_id" form:"applicant_id" gorm:"type:varchar(20);not null" binding:"required"`
	// 联系方式
	Contact string `json:"applicant_contact" form:"applicant_contact" gorm:"type:varchar(32)" binding:"required"`
	// 地址
	Address string `json:"applicant_address" form:"applicant_address" gorm:"type:varchar(255)" binding:"required"`
}

// 用人单位
type Employer struct {
	// 公司名
	Name string `json:"employer_name" form:"employer_name" gorm:"type:varchar(255);not null" binding:"required"`
	// 法人
	LegalRepresentative string `json:"employer_faren" form:"employer_faren" gorm:"type:varchar(128);not null" binding:"required"`
	// 识别号
	UniformSocialCreditCode string `json:"employer_uscc" form:"employer_uscc" gorm:"type:varchar(32);not null" binding:"required"`
	// 联系方式
	Contact string `json:"employer_contact" form:"employer_contact" gorm:"type:varchar(32)" binding:"required"`
	// 地址
	Address string `json:"employer_address" form:"employer_address" gorm:"type:varchar(255)" binding:"required"`
}

// 相关证据材料
type Material struct {
	ID int64 `json:"id,omitempty"`
	// 材料介绍
	Name string `json:"name" binding:"required"`
	// 材料资源path
	Path *string `json:"path" binding:"omitempty"`
}

type CaseFirstSubmitForm struct {

	// 类型id
	CategoryID int64 `json:"category_id" binding:"required"`

	// 1.申请人
	Applicant Applicant `json:"applicant" binding:"required,dive"`

	// 2.被申请人
	Respondent Employer `json:"respondent" binding:"required,dive"`

	// 3.调解事项
	Title string `json:"title" binding:"required"`

	// 4.调解事项所依据的事实与理由
	Content string `json:"content" binding:"required"`

	// 5.劳动者填写的表单 id
	FormID int64 `json:"form_id" binding:"required"`

	// 6. 证据材料
	Materials []*Material `json:"materials"`

	// 7.调解笔录， 8. 调解意见 不用填写

}

type CaseFullResultForm struct {
	ID        int64       `json:"id"`
	CaseID    string      `json:"case_id"`
	CreatedAt *utils.Time `json:"created_at"`
	UpdateAt  *utils.Time `json:"updated_at"`
	Owner     int64       `json:"owner"`

	Status int `json:"status"`

	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Category   Category  `json:"category"`
	Applicant  Applicant `json:"applicant"`
	Respondent Employer  `json:"respondent"`

	Form        *LaborArbitrationForm `json:"form"`
	Materials   []*Material           `json:"materials"`
	Records     []*Record             `json:"records"`
	Suggestions []*Suggestion         `json:"suggestions"`
	Comments    []*Comment            `json:"comments"`
}

type SimplifiedCaseListItem struct {
	ID             int64       `json:"id"`
	CaseID         string      `json:"case_id"`
	ApplicantName  string      `json:"applicant_name"`  // 申请人
	RespondentName string      `json:"respondent_name"` // 被申请人
	Title          string      `json:"title"`
	CreatedAt      *utils.Time `json:"created_at"`
	UpdateAt       *utils.Time `json:"updated_at"`
	Status         int         `json:"status"`
	Owner          int64       `json:"owner"`
}

type CaseStatusChangeForm struct {
	Status int `json:"status" binding:"required"`
}
