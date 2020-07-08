package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
)

func CreateNewHistory(caseId, operation, content string, userID int64, operationHash string, prevHash *string) (*models.HistoryV1, error) {

	user, _ := GetUserById(userID)

	model := models.HistoryV1{
		CaseID:            caseId,
		Content:           content,
		UserID:            userID,
		User:              user.UserName,
		OperationHash:     operationHash,
		PrevOperationHash: prevHash,
		Operation:         operation,
	}
	result := db.Create(&model)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}
}

func GetHistoryAllPaginatedByCaseId(pageNum, pageSize int, caseId *string) ([]*models.HistoryV1, int, error) {
	var histories []*models.HistoryV1
	var totalCount int

	pendingDb := db.Model(&models.HistoryV1{})

	if caseId != nil && *caseId != "" {
		pendingDb = pendingDb.Where("case_id like ?", "%"+*caseId+"%")
	}

	result := pendingDb.Count(&totalCount).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&histories)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCount, result.Error
	} else {
		return histories, totalCount, nil
	}

}

func GetLastHistoryOperationHashByCaseId(caseId string) (*string, error) {
	var history models.HistoryV1
	result := db.Model(&models.HistoryV1{}).Select("operation_hash").Where("case_id = ?", caseId).Order("id desc").First(&history)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows || result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, result.Error
		}
	} else {
		return &history.OperationHash, nil
	}
}
