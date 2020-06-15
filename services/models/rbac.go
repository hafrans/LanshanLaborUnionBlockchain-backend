package models


type Role struct{
	Model
	Name        string `json:"name" gorm:"unique_index"`
	Descriptor  string `json:"descriptor" gorm:"unique_index"`
	Description string `json:"description"`
	Permissions []*Permission `json:"-" gorm:"many2many:role_permission"`
}

type Permission struct{
	Model
	Name        string  `json:"name" gorm:"unique_index"`
	Descriptor  string  `json:"descriptor" gorm:"unique_index"`
	Description string  `json:"description"`
	Roles       []*Role `json:"-" gorm:"many2many:role_permission"`
}


type RolePermission struct {
	Model
	RoleID int64
	PermissionID int64
}

type UserRole struct {
	Model
	UserID int64
	RoleID int64
}
