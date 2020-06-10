package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"log"
)

func CreatePermission(name, description, descriptor string) (*models.Permission , error){
	permission := &models.Permission{
		Name: name,
		Description: description,
		Descriptor: descriptor,
	}
	result := db.Create(permission)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return permission, nil
	}
}

func UpdatePermission(permission *models.Permission) bool {
	result := db.Save(permission)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeletePermission(permission *models.Permission) bool {
	result := db.Delete(permission)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeletePermissionById(id int64) bool {
	result := db.Delete(&models.Permission{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetPermissionAllPaginated(pageNum, pageCount int) ([]*models.Permission, int, error) {
	var permissions []*models.Permission
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&permissions).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return permissions, totalCounts, nil
	}
}


func GetPermissionAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Permission, int, error){
	var permissions []*models.Permission
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&permissions).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return permissions, totalCounts, nil
	}

}

