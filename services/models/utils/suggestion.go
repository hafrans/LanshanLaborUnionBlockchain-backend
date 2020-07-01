package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
)

// 将Suggestion 的vo转model
func PopulateSuggestionFromVOToModel(vo *vo.Suggestion) *models.Suggestion {
	return &models.Suggestion{
		Content: vo.Content,
		Department: vo.Department,
	}
}

// 将Record 的 model 转vo
func PopulateSuggestionFromModelToVO(model *models.Suggestion) *vo.Suggestion {
	return &vo.Suggestion{
		ID:   model.ID,
		Department: model.Department,
		Content: model.Content,
		Submitter: model.User.UserName,
	}
}

// 将 Suggestion 的 vo list转变为model ，一般是提交新的record
func PopulateSuggestionListFromVOToModel(materials []*vo.Suggestion) []*models.Suggestion {

	length := len(materials)
	result := make([]*models.Suggestion, 0, length)

	for _, v := range materials {
		result = append(result, PopulateSuggestionFromVOToModel(v))
	}

	return result

}

// 将 Suggestion 的 model list 转为 vo
func PopulateSuggestionListFromModelToVO(model []*models.Suggestion) []*vo.Suggestion {

	length := len(model)
	result := make([]*vo.Suggestion, 0, length)
	for _, v := range model {
		result = append(result, PopulateSuggestionFromModelToVO(v))
	}

	return result
}





