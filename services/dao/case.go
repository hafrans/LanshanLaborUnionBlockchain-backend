package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func CreateCase(model *models.Case) (*models.Case, error) {

	if model == nil {
		return nil, errors.New("no model found")
	}

	result := db.Create(model)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return model, nil
	}

}

func GetCaseFullModelById(id int64) (*models.Case, error) {

	var model models.Case

	result := db.Preload("Category").Preload("Form").Preload("Materials").Preload("Records").Preload("Suggestions").First(&model, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}

}
