package models


// 类型/调解事项
type Category struct {
	Model
	Name        string `json:"name" gorm:"type:varchar(128);unique_index;not null;"`
	Description string `json:"description" gorm:"type:text"`
}





