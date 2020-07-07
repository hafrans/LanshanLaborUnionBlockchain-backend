package models

import (
	"time"
)

const (
	USER_TYPE_ADMIN      = 1
	USER_TYPE_LABOR      = 2
	USER_TYPE_EMPLOYER   = 3
	USER_TYPE_DEPARTMENT = 4
)

type User struct {
	Model
	UserName      string     `json:"username" gorm:"type:varchar(128);"`
	Email         string     `json:"email" gorm:"type:varchar(255);unique_index"`
	EmailChecked  bool       `json:"email_checked" gorm:"default:false"`
	Phone         string     `json:"phone" gorm:"type:varchar(32);unique_index"`
	PhoneChecked  bool       `json:"phone_checked" gorm:"default:false"`
	Credentials   string     `json:"-" gorm:"type:varchar(255);not null"`
	Activated     bool       `json:"active" gorm:"not null"`
	LastLoginTime *time.Time `json:"last_login"`
	Roles         []*Role    `json:"-" gorm:"many2many:user_role"`

	UserType    int         `json:"user_type" gorm:"size:1;not null"` // 约定使用UserProfile中的某些字段
	UserProfile UserProfile `json:"user_profile"`

	Department   Department `json:"department"` // 工会等用户
	DepartmentID *int64     `json:"-"`
}

type UserProfile struct {
	Model
	UserID int64
	Applicant
	Employer
}

type Department struct {
	Model
	Name        string `json:"name" gorm:"unique_index"` // 机构、单位名称
	Service     string `json:"service"`                  // 机构提供的服务
	Contact     string `json:"contact"`                  // 机构联系方式
	Description string `json:"description"`              // 机构介绍
}
