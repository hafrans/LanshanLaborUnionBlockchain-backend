package vo

type DepartmentVO struct {
	Name        string `json:"name" gorm:"unique_index"` // 机构、单位名称
	Service     string `json:"service"`                  // 机构提供的服务
	Contact     string `json:"contact"`                  // 机构联系方式
	Description string `json:"description"`              // 机构介绍
}
