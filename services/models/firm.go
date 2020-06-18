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
	Name           string
	Birthday       utils.Time
	Nationality    string
	IdentityNumber string
	Contact        string
	Address        string
}

// 用人单位
type Employer struct {
	Model
	Name                    string
	LegalRepresentative     string
	UniformSocialCreditCode string
	Contact                 string
	Address                 string
}
