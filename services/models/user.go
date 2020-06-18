package models

import "time"

const (
	USER_TYPE_ADMIN = 1
	USER_TYPE_LABOR = 2
	USER_TYPE_FIRM  = 3
)

type User struct {
	Model
	UserName      string     `json:"username" gorm:"type:varchar(128);unique_index"`
	Email         string     `json:"email" gorm:"type:varchar(255);unique_index"`
	EmailChecked  bool       `json:"email_checked" gorm:"default:false"`
	Phone         string     `json:"phone" gorm:"type:varchar(32);unique_index"`
	PhoneChecked  bool       `json:"phone_checked" gorm:"default:false"`
	Credentials   string     `json:"-" gorm:"type:varchar(255);not null"`
	Activated     bool       `json:"active" gorm:"not null"`
	LastLoginTime *time.Time `json:"last_login"`
	Roles         []*Role    `json:"-" gorm:"many2many:user_role"`
	UserType      int        `json:"user_type" gorm:"size:1;not null"`
}

type UserProfile struct {
	Model
}

type Department struct {
	Model
	Name        string `json:"name" gorm:"unique_index"`
	Description string `json:"description"`
}
