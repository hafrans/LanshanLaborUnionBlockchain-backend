package models

import "RizhaoLanshanLabourUnion/utils"

// 申请调解人
type Applicant struct {
	Model
	Name           string      `json:"applicant_name" form:"applicant_name" gorm:"type:varchar(128);not null"`
	Birthday       *utils.Date `json:"applicant_birth" form:"applicant_birth" gorm:"type:date"`
	Nationality    string      `json:"applicant_nationality" form:"applicant_nationality" gorm:"type:varchar(32);not null"`
	IdentityNumber string      `json:"applicant_id" form:"applicant_id" gorm:"type:varchar(20);not null"`
	Contact        string      `json:"applicant_contact" form:"applicant_contact" gorm:"type:varchar(32)"`
	Address        string      `json:"applicant_address" form:"applicant_address" gorm:"type:varchar(255)"`
}

// 用人单位
type Employer struct {
	Model
	Name                    string `json:"employer_name" form:"employer_name" gorm:"type:varchar(255);not null"`
	LegalRepresentative     string `json:"employer_faren" form:"employer_faren" gorm:"type:varchar(128);not null"`
	UniformSocialCreditCode string `json:"employer_uscc" form:"employer_uscc" gorm:"type:varchar(32);not null"`
	Contact                 string `json:"employer_contact" form:"employer_contact" gorm:"type:varchar(32)"`
	Address                 string `json:"employer_address" form:"employer_address" gorm:"type:varchar(255)"`
}

// 相关证据材料
type Material struct {
	Model
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Path        *string `json:"path" binding:"omitempty"`
	CaseId      string  `json:"case_id"`
}

// 调解笔录
type Record struct {
	Model
	Name   string `json:"record_name"`
	CaseId string `json:"case_id"`
	Path   string `json:"record_path"`
}

// 部门调解意见
type Suggestion struct {
	Model
	Department string `json:"suggestion_department"` // 部门名称
	Content    string `suggestion_content`           // 意见
	CaseId string `json:"case_id"`
}

// 案件 many to many
type Case struct {
	Model

	CaseId string `json:"case_id" gorm:"type:varchar(32);unique_index"` // case id //要自己定

	UserId int64 `json:"user_id"` // 归属用户

	CategoryId int64    `json:"-"` // 争议类型
	Category   Category `json:"category"`

	Applicant   Applicant `json:"applicant"` // 申请人
	ApplicantId int64     `json:"-"`

	Respondent   Employer `json:"respondent" gorm:"foreignkey:RespondentId"` // 被申请用工单位
	RespondentId int64    `json:"-"`

	Title string `json:"title" binding:"required"` // 调解事项

	Content string `json:"content" binding:"required"` // 调解事实与理由

	FormId int64 `json:"-"` // 表单
	Form   *LaborArbitration

	Materials []*Material `json:"materials"` // 证据材料

	Suggestions []*Suggestion `json:"suggestions"` // 部门处理意见

}

// 案件
type CaseV2 struct {
	Model
	UserId      int64      `json:"user_id"`
	ApplicantId int64      `json:"applicant_id"`
	Applicant   *Applicant `json:"applicant"`
}
