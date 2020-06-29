package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func CreateRecord(name, path string, caseId string) (*models.Record, error) {

	record := &models.Record{
		Name:   name,
		Path:   path,
		CaseID: caseId,
	}

	result := db.Create(record)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return record, nil
	}

}

func UpdateRecord(record *models.Record) bool {

	if record == nil {
		return false
	}

	result := db.Save(record)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteRecord(department *models.Record) bool {
	result := db.Delete(department)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteRecordById(id int64) bool {
	result := db.Delete(&models.Record{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetRecordById(id int64) (*models.Record, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var department *models.Record
	result := db.FirstOrInit(department, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return department, nil
	}
}

func GetRecordAllPaginated(pageNum, pageCount int) ([]*models.Record, int, error) {
	var records []*models.Record
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&records).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return records, totalCounts, nil
	}
}

func GetRecordAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Record, int, error) {
	var records []*models.Record
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&records).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return records, totalCounts, nil
	}
}
