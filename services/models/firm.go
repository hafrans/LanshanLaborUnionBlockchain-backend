package models

import "RizhaoLanshanLabourUnion/utils"

// 类型/调解事项
type Category struct {
	Model
	Name        string `json:"name" gorm:"type:varchar(128);unique_index;not null;"`
	Description string `json:"description"`
}

// 申请调解人
type Applicant struct {
	Model
	Name           string     `json:"applicant_name" form:"applicant_name" gorm:"type:varchar(128);not null"`
	Birthday       utils.Time `json:"applicant_birth" form:"applicant_birth"`
	Nationality    string     `json:"applicant_nationality" form:"applicant_nationality" gorm:"type:varchar(32);not null"`
	IdentityNumber string     `json:"applicant_id" form:"applicant_id" gorm:"type:varchar(20);not null;unique_index"`
	Contact        string     `json:"applicant_contact" form:"applicant_contact" gorm:"type:varchar(32)"`
	Address        string     `json:"applicant_address" form:"applicant_address" gorm:"type:varchar(255)"`
}

// 用人单位
type Employer struct {
	Model
	Name                    string `json:"employer_name" form:"employer_name" gorm:"type:varchar(255);not null;unique_index"`
	LegalRepresentative     string `json:"employer_faren" form:"employer_faren" gorm:"type:varchar(128);not null"`
	UniformSocialCreditCode string `json:"employer_uscc" form:"employer_uscc" gorm:"type:varchar(32);not null;unique_index"`
	Contact                 string `json:"employer_contact" form:"employer_contact" gorm:"type:varchar(32)"`
	Address                 string `json:"employer_address" form:"employer_address" gorm:"type:varchar(255)"`
}
