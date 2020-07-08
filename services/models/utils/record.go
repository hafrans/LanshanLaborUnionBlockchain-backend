package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
)

// 将Record 的vo转model
func PopulateRecordFromVOToModel(vo *vo.Record) *models.Record {
	return &models.Record{
		Path: vo.Path,
		Name: vo.Name,
	}
}

// 将Record 的 model 转vo
func PopulateRecordFromModelToVO(model *models.Record) *vo.Record {
	return &vo.Record{
		ID:             model.ID,
		Path:           model.Path,
		Name:           model.Name,
		Submitter:      model.User.UserName,
		CaseID:         model.CaseID,
		DepartmentInfo: PopulateDepartmentFromModelToVO(&model.User.Department),
		SubmitterPhone: model.User.Phone,
	}
}

// 将 Record 的 vo list转变为model ，一般是提交新的record
func PopulateRecordListFromVOToModel(materials []*vo.Record) []*models.Record {

	length := len(materials)
	result := make([]*models.Record, 0, length)

	for _, v := range materials {
		result = append(result, PopulateRecordFromVOToModel(v))
	}

	return result

}

// 将 Record 的 model list 转为 vo
func PopulateRecordListFromModelToVO(model []*models.Record) []*vo.Record {

	length := len(model)
	result := make([]*vo.Record, 0, length)
	for _, v := range model {
		result = append(result, PopulateRecordFromModelToVO(v))
	}

	return result
}
