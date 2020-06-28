package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func CreateEmployer(name, lr, uscc, contact, address string) (*models.Employer, error) {

	employer := &models.Employer{
		Name:                    name,
		LegalRepresentative:     lr,
		Address:                 address,
		Contact:                 contact,
		UniformSocialCreditCode: uscc,
	}

	result := db.Create(employer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return employer, nil
	}

}

func UpdateEmployer(employer *models.Employer) bool {

	if employer == nil {
		return false
	}

	result := db.Save(employer)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteEmployer(employer *models.Employer) bool {
	result := db.Delete(employer)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteEmployerById(id int64) bool {
	result := db.Delete(&models.Employer{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetEmployerById(id int64) (*models.Employer, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var employer *models.Employer
	result := db.FirstOrInit(employer, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return employer, nil
	}
}

func GetEmployerAllPaginated(pageNum, pageCount int) ([]*models.Employer, int, error) {
	var employers []*models.Employer
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&employers).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return employers, totalCounts, nil
	}
}

func GetEmployerAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Employer, int, error) {
	var employers []*models.Employer
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&employers).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return employers, totalCounts, nil
	}
}
