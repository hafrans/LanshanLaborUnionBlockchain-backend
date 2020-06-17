package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"errors"
	"log"
	"strconv"
	"time"
)


func CreateUser(username, password, email, phone string, userType int, emailChecked, phoneChecked, active bool) (*models.User, error) {

	hashedPassword, ok := utils.GenerateHashedPassword(password)
	if !ok {
		log.Println("generate password failed")
		return nil, errors.New("can not generate password")
	}

	user := &models.User {
		UserName:     username,
		Credentials:  hashedPassword,
		Email:        email,
		EmailChecked: emailChecked,
		Phone:        phone,
		PhoneChecked: phoneChecked,
		UserType:     userType,
	}

	if err := db.Create(user).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func GetUserById(id int64) (*models.User, error) {
	user := &models.User{}
	result := db.First(user, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return user, nil
	}
}

func GetUserByUserName(username string) (*models.User, error) {
	user := &models.User{}
	result := db.Where("user_name = ?", username).First(user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return user, nil
	}
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := db.Where("email = ?", email).First(user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return user, nil
	}
}

func GetUserByPhone(phone string) (*models.User, error) {
	user := &models.User{}
	result := db.Where("phone = ?", phone).First(user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return user, nil
	}
}

func UpdateUser(user *models.User) bool {
	result := db.Save(user)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteUser(user *models.User) bool {
	result := db.Delete(user)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteUserById(id int64) bool {
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func LoginUser(user *models.User) bool {
	now := time.Now()
	result := db.Model(user).Update(models.User{LastLoginTime: &now})
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else if result.RowsAffected == 1 {
		return true
	} else {
		log.Println("[LoginUser] multi row affected" + strconv.Itoa(int(result.RowsAffected)))
		return false
	}
}

func SetUserActivateFlag(user *models.User, active bool) bool {
	result := db.Model(user).Update(models.User{Activated: active})
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func SetUserEmailAndPhoneConfirmedFlag(user *models.User, phone, email bool) bool {
	result := db.Model(user).Update(models.User{PhoneChecked: phone, EmailChecked: email})
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetUserAllPaginated(pageNum, pageCount int) ([]*models.User, int, error) {
	var users []*models.User
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&users).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return users, totalCounts, nil
	}
}

func GetUserAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.User, int, error) {
	var users []*models.User
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&users).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return users, totalCounts, nil
	}
}



func AddRoleToUser(role *models.Role, user *models.User) bool {

	if role == nil || user == nil{
		log.Println("add role failed, because of nullptr")
		return false
	}

	userRole := &models.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}

	result := db.Create(userRole)

	if result.Error != nil{
		log.Println(result.Error)
		return false
	}else{
		return true
	}


}


func GetRolesFromUser(user *models.User)([]*models.Role, error){
	var roles []*models.Role
	result := db.Model(user).Related(&roles, "Roles")
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}else{
		return roles, nil
	}
}
