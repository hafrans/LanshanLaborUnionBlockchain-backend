package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
	"log"
)

// 将 Material 的 vo 转 model
func PopulateMaterialFromVOToModel(material *vo.Material) *models.Material {

	model := models.Material{
		Name: material.Name,
		Path: material.Path,
	}
	return &model
}

// 将 Material 的 model 转 vo
func PopulateMaterialFromModelToVO(model *models.Material) *vo.Material {
	return &vo.Material{
		Name: model.Name,
		Path: model.Path,
		ID:   model.ID,
	}
}

// 将 Material 的 vo list转变为model ，一般是提交新的material
func PopulateMaterialListFromVOToModel(materials []*vo.Material) []*models.Material {

	length := len(materials)
	result := make([]*models.Material, 0, length)

	for _, v := range materials {
		result = append(result, PopulateMaterialFromVOToModel(v))
	}

	return result

}

// 将 Material 的 model list 转为 vo
func PopulateMaterialListFromModelToVO(model []*models.Material) []*vo.Material {

	length := len(model)
	log.Println(length)
	result := make([]*vo.Material, 0, length)
	for _, v := range model {
		result = append(result, PopulateMaterialFromModelToVO(v))
	}

	return result
}
