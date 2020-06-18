package models

import "time"

type Model struct {
	ID        int64      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `json: "created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:default:NULL`
}
