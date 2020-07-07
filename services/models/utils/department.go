package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
)

// 将Department的Model转成VO
func PopulateDepartmentFromModelToVO(model *models.Department) *vo.DepartmentVO {

	return &vo.DepartmentVO{
		Contact:     model.Contact,
		Service:     model.Service,
		Name:        model.Name,
		Description: model.Description,
	}
}

// 将Department的VO转成Model
func PopulateDepartmentFromVoToModel(vo *vo.DepartmentVO) *models.Department {
	return &models.Department{
		Description: vo.Description,
		Name:        vo.Name,
		Service:     vo.Service,
		Contact:     vo.Contact,
	}
}
