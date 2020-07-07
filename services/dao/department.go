package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func InitDepartment() {
	CreateDepartment("[测试]莒州市总工会", "暂时无描述", "测试、调解、测试、调解", "0633-8888888")
	CreateDepartment("[测试]莒州市中级人民法院", "暂时无描述", "测试、调解、测试、调解", "0633-7777777")
	CreateDepartment("[测试]莒州市劳动仲裁委员会", "暂时无描述", "测试、调解、测试、调解", "0633-6666666")
	CreateDepartment("[测试]莒州市人社局", "暂时无描述", "测试、调解、测试、调解", "0633-5555555")
}

func CreateDepartment(name, description, service, contact string) (*models.Department, error) {
	department := &models.Department{
		Name:        name,
		Description: description,
		Service:     service,
		Contact:     contact,
	}
	result := db.Create(department)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
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

func GetDepartmentById(id int64) (*models.Department, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var department *models.Department
	result := db.FirstOrInit(department, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return department, nil
	}
}

func GetDepartmentAllPaginated(pageNum, pageCount int) ([]*models.Department, int, error) {
	var departments []*models.Department
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Model(&models.Department{}).Count(&totalCounts).Limit(pageCount).Offset(pageCount * (pageNum - 1)).Order("id desc").Find(&departments)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return departments, totalCounts, nil
	}
}

func GetDepartmentAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Department, int, error) {
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
