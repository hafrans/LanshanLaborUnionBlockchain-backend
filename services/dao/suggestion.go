package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func CreateSuggestion(depart, content string, caseId string) (*models.Suggestion, error) {

	suggestion := &models.Suggestion{
		Department: depart,
		Content:    content,
		CaseId:     caseId,
	}

	result := db.Create(suggestion)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return suggestion, nil
	}

}

func UpdateSuggestion(suggestion *models.Suggestion) bool {

	if suggestion == nil {
		return false
	}

	result := db.Save(suggestion)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteSuggestion(department *models.Suggestion) bool {
	result := db.Delete(department)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteSuggestionById(id int64) bool {
	result := db.Delete(&models.Suggestion{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetSuggestionById(id int64) (*models.Suggestion, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var department *models.Suggestion
	result := db.FirstOrInit(department, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return department, nil
	}
}

func GetSuggestionAllPaginated(pageNum, pageCount int) ([]*models.Suggestion, int, error) {
	var suggestions []*models.Suggestion
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&suggestions).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return suggestions, totalCounts, nil
	}
}
