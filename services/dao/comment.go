package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"log"
)

func CreateComment(user *models.User, caseId string, content string) (*models.Comment, error) {

	model := models.Comment{
		UserID:  user.ID,
		CaseID:  caseId,
		Content: content,
	}

	result := db.Create(&model)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}
}

func DeleteComment(model *models.Comment) bool {

	result := db.Delete(&models.Comment{})

	if result.Error != nil {
		log.Println(result.Error)
		return false
	}else{
		return true
	}
}

func DeleteCommentById(id int64) bool{

	result := db.Delete(&models.Comment{},id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	}else{
		return true
	}

}

func GetCommentById(id int64) (*models.Comment, error) {
	var model models.Comment

	result := db.Set("gorm:auto_preload", true).First(&model,id)

	if result.Error != nil {
		return nil, result.Error
	}else{
		return &model, nil
	}
}