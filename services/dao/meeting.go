package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"log"
)

func CreateMeeting(model *models.Meeting) (*models.Meeting, error) {
	result := db.Create(model)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return model, nil
	}
}

func DeleteMeetingByID(id int64) bool {

	result := db.Delete(&models.Meeting{}, id)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func GetMeetingByID(id int64) (*models.Meeting, error) {

	var model models.Meeting
	result := db.Set("gorm:auto_preload", true).First(&model, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}
}

func GetMeetingAllWithConditionPaginated(caseId *string, userId *int64, filterOld bool, pageNum, pageSize int) ([]*models.Meeting, int, error) {
	var meetings []*models.Meeting
	var totalCount int

	pendingDB := db.Set("gorm:auto_preload", true).Model(&models.Meeting{})

	if caseId != nil && *caseId != "" {
		pendingDB = pendingDB.Where("case_id = ?", caseId)
	}

	if userId != nil && *userId != 0 {
		pendingDB = pendingDB.Where("user_id = ?", userId)
	}

	if filterOld {
		pendingDB = pendingDB.Where("end_time > NOW()")
	}

	result := pendingDB.Count(&totalCount).Count(pageSize).Offset((pageNum - 1) * pageSize).Find(&meetings)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCount, result.Error
	}else{
		return meetings, totalCount, nil
	}

}
