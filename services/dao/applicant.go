package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"errors"
	"log"
)

func CreateApplicant(name, nationality, idNumber, contact, address string, birth *utils.Date) (*models.Applicant, error) {

	applicant := &models.Applicant{
		Name: name,
		Nationality: nationality,
		Address: address,
		Contact: contact,
		Birthday: birth,
		IdentityNumber: idNumber,
	}

	result := db.Create(applicant)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return applicant, nil
	}

}

func UpdateApplicant(applicant *models.Applicant) bool {

	if applicant == nil {
		return false
	}

	result := db.Save(applicant)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteApplicant(department *models.Applicant) bool {
	result := db.Delete(department)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteApplicantById(id int64) bool {
	result := db.Delete(&models.Applicant{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetApplicantById(id int64) (*models.Applicant, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var department *models.Applicant
	result := db.FirstOrInit(department, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return department, nil
	}
}

func GetApplicantAllPaginated(pageNum, pageCount int) ([]*models.Applicant, int, error) {
	var applicants []*models.Applicant
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&applicants).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return applicants, totalCounts, nil
	}
}

func GetApplicantAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Applicant, int, error) {
	var applicants []*models.Applicant
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&applicants).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return applicants, totalCounts, nil
	}
}
