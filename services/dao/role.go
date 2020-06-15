package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
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


func GetRoleById(id int64) (*models.Role, error){
	if id <= 0{
		return nil, errors.New("invalid id")
	}
	var role *models.Role = &models.Role{}
	result := db.FirstOrInit(role,id)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return role, nil
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


func AddPermissionToRole(permission *models.Permission, role *models.Role) bool {

	if permission == nil || role == nil{
		log.Println("add permission failed, because of nullptr")
		return false
	}

	rolePermission := &models.RolePermission{
		PermissionID: permission.ID,
		RoleID: role.ID,
	}

	result := db.Create(rolePermission)

	if result.Error != nil{
		log.Println(result.Error)
		return false
	}else{
		return true
	}


}


func GetPermissionsFromRole(role *models.Role)([]*models.Permission, error){
	var permissions []*models.Permission
	result := db.Model(role).Related(&permissions, "Permissions")
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return permissions, nil
	}
}
