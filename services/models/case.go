package models

import "RizhaoLanshanLabourUnion/utils"

// 申请调解人
type Applicant struct {
	//Model
	// 姓名
	ApplicantName string `json:"applicant_name" form:"applicant_name" gorm:"type:varchar(128);not null"`
	// 生日
	ApplicantBirthday *utils.Date `json:"applicant_birth" form:"applicant_birth" gorm:"type:date"`
	// 民族
	ApplicantNationality string `json:"applicant_nationality" form:"applicant_nationality" gorm:"type:varchar(32);not null"`
	// 身份证号
	ApplicantIdentityNumber string `json:"applicant_id" form:"applicant_id" gorm:"type:varchar(20);not null"`
	// 联系方式
	ApplicantContact string `json:"applicant_contact" form:"applicant_contact" gorm:"type:varchar(32)"`
	// 地址
	ApplicantAddress string `json:"applicant_address" form:"applicant_address" gorm:"type:varchar(255)"`
}

// 用人单位
type Employer struct {
	//Model
	// 公司名
	EmployerName string `json:"employer_name" form:"employer_name" gorm:"type:varchar(255);not null"`
	// 法人
	EmployerLegalRepresentative string `json:"employer_faren" form:"employer_faren" gorm:"type:varchar(128);not null"`
	// 识别号
	EmployerUniformSocialCreditCode string `json:"employer_uscc" form:"employer_uscc" gorm:"type:varchar(32);not null"`
	// 联系方式
	EmployerContact string `json:"employer_contact" form:"employer_contact" gorm:"type:varchar(32)"`
	// 地址
	EmployerAddress string `json:"employer_address" form:"employer_address" gorm:"type:varchar(255)"`
}

// 相关证据材料
type Material struct {
	Model
	// 材料介绍
	Name string `json:"name" binding:"required"`
	// 材料资源path
	Path *string `json:"path" binding:"omitempty"`
	// caseID， 新建表单时不要上传该信息
	CaseID string `json:"case_id"`
}

// 调解笔录
type Record struct {
	Model
	// 记录介绍
	Name string `json:"record_name"`
	// caseID， 新建表单时不要上传该信息
	CaseID string `json:"case_id"`
	// 截图/材料等地址
	Path string `json:"record_path"`
}

// 部门调解意见
type Suggestion struct {
	Model
	// 部门名称
	Department string `json:"suggestion_department"`
	// 意见
	Content string `json:"suggestion_content"`
	// caseID， 新建表单时不要上传该信息
	CaseID string `json:"case_id"`
}

// 案件 many to many
type Case struct {
	Model

	CaseID string `json:"case_id" gorm:"type:varchar(64);unique_index"` // case id //要自己定

	Status int `json:"status" gorm:"type:tinyint(1)"`

	UserID int64 `json:"user_id"` // 归属用户

	CategoryID int64    `json:"-"` // 争议类型
	Category   Category `json:"category"`

	Applicant // 申请人

	Employer // 用工单位

	Title string `json:"title" binding:"required"` // 调解事项

	Content string `json:"content" binding:"required"` // 调解事实与理由

	FormID int64             `json:"-"` // 表单
	Form   *LaborArbitration `gorm:"foreignkey:FormID"`

	Materials []*Material `json:"materials" gorm:"foreignkey:CaseID;association_foreignkey:CaseID"` // 证据材料

	Records []*Record `json:"records" gorm:"foreignkey:CaseID;association_foreignkey:CaseID"` // 笔录

	Suggestions []*Suggestion `json:"suggestions" gorm:"foreignkey:CaseID;association_foreignkey:CaseID"` // 部门处理意见

}

// 案件
type CaseV2 struct {
	Model
	UserId      int64      `json:"user_id"`
	ApplicantId int64      `json:"applicant_id"`
	Applicant   *Applicant `json:"applicant"`
}
