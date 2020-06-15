package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func CreateDepartment(name, description string) (*models.Department , error){
	department := &models.Department{
		Name: name,
		Description: description,
	}
	result := db.Create(department)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return department, nil
	}
}

func UpdateDepartment(department *models.Department) bool {
	result := db.Save(department)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteDepartment(department *models.Department) bool {
	result := db.Delete(department)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}


func DeleteDepartmentById(id int64) bool {
	result := db.Delete(&models.Department{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetDepartmentById(id int64) (*models.Department, error){
	if id <= 0{
		return nil, errors.New("invalid id")
	}
	var department *models.Department
	result := db.FirstOrInit(department,id)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return department, nil
	}
}

func GetDepartmentAllPaginated(pageNum, pageCount int) ([]*models.Department, int, error) {
	var departments []*models.Department
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&departments).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return departments, totalCounts, nil
	}
}


func GetDepartmentAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Department, int, error){
	var departments []*models.Department
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&departments).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return departments, totalCounts, nil
	}
}
