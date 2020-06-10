package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"log"
)

func CreateRole(name, description, descriptor string) (*models.Role , error){
	role := &models.Role{
		Name: name,
		Description: description,
		Descriptor: descriptor,
	}
	result := db.Create(role)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return role, nil
	}
}

func UpdateRole(role *models.Role) bool {
	result := db.Save(role)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteRole(role *models.Role) bool {
	result := db.Delete(role)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteRoleById(id int64) bool {
	result := db.Delete(&models.Role{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetRoleAllPaginated(pageNum, pageCount int) ([]*models.Role, int, error) {
	var roles []*models.Role
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&roles).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return roles, totalCounts, nil
	}
}


func GetRoleAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Role, int, error){
	var roles []*models.Role
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&roles).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return roles, totalCounts, nil
	}

}

