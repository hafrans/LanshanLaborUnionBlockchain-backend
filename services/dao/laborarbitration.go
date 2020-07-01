package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

func CreateLaborArbitration(model *models.LaborArbitration) (*models.LaborArbitration, error) {

	if model == nil {
		return nil, errors.New("form is not found")
	}

	result := db.Create(model)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return model, nil
	}

}

func UpdateLaborArbitration(labor *models.LaborArbitration) bool {

	if labor == nil {
		return false
	}

	result := db.Save(labor)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteLaborArbitration(labor *models.LaborArbitration) bool {
	result := db.Delete(labor)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteLaborArbitrationById(id int64) bool {
	result := db.Delete(&models.LaborArbitration{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetLaborArbitrationById(id int64) (*models.LaborArbitration, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var labor models.LaborArbitration
	result := db.First(&labor, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &labor, nil
	}
}

func GetLaborArbitrationAllPaginated(pageNum, pageCount int) ([]*models.LaborArbitration, int, error) {
	var labors []*models.LaborArbitration
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Model(&models.LaborArbitration{}).Count(&totalCounts).Limit(pageCount).Offset(pageCount * (pageNum - 1)).Find(&labors)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCounts, result.Error
	} else {
		//db.Model(&models.LaborArbitration{}).Where("Owner = ?" , userId)
		return labors, totalCounts, nil
	}
}

func GetLaborArbitrationAllPaginatedOwnByUser(pageNum, pageCount int, userId int64) ([]*models.LaborArbitration, int, error) {
	var labors []*models.LaborArbitration
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0

	var result *gorm.DB

	if userId != 0 {
		result = db.Model(&models.LaborArbitration{}).Where("Owner = ?", userId).Count(&totalCounts).Limit(pageCount).Offset(pageCount * (pageNum - 1)).Find(&labors)
	} else {
		result = db.Model(&models.LaborArbitration{}).Count(&totalCounts).Limit(pageCount).Offset(pageCount * (pageNum - 1)).Find(&labors)
	}

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCounts, result.Error
	} else {
		//db.Model(&models.LaborArbitration{}).Where("Owner = ?" , userId)
		return labors, totalCounts, nil
	}
}
