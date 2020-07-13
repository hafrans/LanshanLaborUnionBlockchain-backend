package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
)

func PopulateHistoryV1FromModelToVO(model *models.HistoryV1) *vo.SimplifiedHistory {
	return &vo.SimplifiedHistory{
		User:              model.User,
		PrevOperationHash: model.PrevOperationHash,
		OperationHash:     model.OperationHash,
		UserID:            model.UserID,
		Content:           model.Content,
		CaseID:            model.CaseID,
		Operation:         model.Operation,
		CreatedAt: utils.GetTime(model.CreatedAt),
	}
}

func PopulateHistoryV1ListFromModelToVO(histories []*models.HistoryV1) []*vo.SimplifiedHistory {
	length := len(histories)
	result := make([]*vo.SimplifiedHistory, 0, length)
	for _, v := range histories {
		result = append(result, PopulateHistoryV1FromModelToVO(v))
	}
	return result
}
