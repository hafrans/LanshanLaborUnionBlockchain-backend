package models

import "time"

type User struct {
	Model
	UserName          string `json:"username" gorm:"type:varchar(128);unique_index"`
	Email             string `json:"email" gorm:"type:varchar(255);unique_index"`
	EmailChecked        bool `json:"email_checked" gorm:"default:false"`
	Phone             string `json:"phone" gorm:"type:varchar(32);unique_index"`
	PhoneChecked        bool `json:"phone_checked" gorm:"default:false"`
	Credentials       string `json:"-" gorm:"type:varchar(255);not null"`
	Activated           bool `json:"active" gorm:"not null"`
	LastLoginTime *time.Time `json:"last_login"`
}

type UserProfile struct {
	Model
}


type Department struct{
	Model
	Name string `json:"name" gorm:"unique_index"`
	Description string `json:"description"`
}

type Role struct{
	Model
	Name        string `json:"name" gorm:"unique_index"`
	Descriptor  string `json:"descriptor" gorm:"unique_index"`
	Description string `json:"description"`
}

type Permission struct{
	Model
	Name        string `json:"name" gorm:"unique_index"`
	Descriptor  string `json:"descriptor" gorm:"unique_index"`
	Description string `json:"description"`
}


type UserRole struct {

}
