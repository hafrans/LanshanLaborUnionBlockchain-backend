package models

import "time"

type Model struct {
	// 模型ID，提交表单时不要上传该信息
	ID        int64      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	// 创建日期，提交表单时不要上传该信息
	CreatedAt time.Time  `json:"created_at"`
	// 更新日期，提交表单时不要上传该信息
	UpdatedAt time.Time  `json:"updated_at"`
	// 删除日期 提交表单时不要上传该信息
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:NULL"`
}
