package dao_test

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"math"
	"testing"
	"time"
)

func init() {
	utils.InitSettings()
	dao.InitDB()
	dao.CreateTables()
}

var id int64

func TestCreateUserWithProfile(t *testing.T) {
	profile := &models.UserProfile{}
	profile.ApplicantName = "张三"
	profile.ApplicantNationality = "三体人"
	profile.ApplicantIdentityNumber = "12345678891234848"
	profile.ApplicantBirthday = utils.NowDateDay()
	profile.ApplicantAddress = "三体星"
	profile.ApplicantContact = "13800138001"

	_, err := dao.CreateUserWithProfile("张三", "test", "test1@163.com",
		"13800138001", models.USER_TYPE_LABOR, false, false, true,
		0, profile)

	if err != nil {
		t.Error(err.Error())
	}

	profile2 := &models.UserProfile{}
	profile2.Employer.EmployerLegalRepresentative = "李四"
	profile2.Employer.EmployerUniformSocialCreditCode = "G12345678901231456"
	profile2.Employer.EmployerName = "测试公司"
	profile2.Employer.EmployerContact = "13800138002"
	profile2.Employer.EmployerAddress = "山东省日照市东港区迎宾路XXX号"

	_, err = dao.CreateUserWithProfile("李四", "test", "test2@163.com",
		"13800138002", models.USER_TYPE_EMPLOYER, false, false, true,
		0, profile2)

	if err != nil {
		t.Error(err.Error())
	}

	profile3 := &models.UserProfile{}

	_, err = dao.CreateUserWithProfile("王五", "test", "test3@163.com",
		"13800138011", models.USER_TYPE_DEPARTMENT, false, false, true,
		1, profile3)
	_, err = dao.CreateUserWithProfile("王五", "test", "test4@163.com",
		"13800138012", models.USER_TYPE_DEPARTMENT, false, false, true,
		2, profile3)
	_, err = dao.CreateUserWithProfile("王五2", "test", "test5@163.com",
		"13800138013", models.USER_TYPE_DEPARTMENT, false, false, true,
		1, profile3)
	_, err = dao.CreateUserWithProfile("王五2", "test", "test6@163.com",
		"13800138014", models.USER_TYPE_DEPARTMENT, false, false, true,
		2, profile3)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestCreateUser(t *testing.T) {

	user, err := dao.CreateUser("test2", "test2", "hafrans@1632.com", "138001328000", models.USER_TYPE_ADMIN, false, false, true)

	if err != nil {
		t.Error(err)
	} else {
		id = user.ID
		//b,_ := json.Marshal(user)
		////t.Log(string(b))
	}

}

func TestLoginUser(t *testing.T) {
	user := &models.User{}

	if !dao.LoginUser(user) {
		t.Error("login failed")
	}

	if math.Abs(user.LastLoginTime.Sub(time.Now()).Seconds()) > 2 {
		t.Error("abs check failed")
	}

}

func TestGetUserById(t *testing.T) {
	user, _ := dao.GetUserById(id)
	if user.ID != id {
		t.Error("User id is not id")
	}
}

func TestGetAllUserPaginated(t *testing.T) {
	users, count, err := dao.GetUserAllPaginated(0, 50)
	t.Log("count:", count)
	if err != nil {
		//log.Println(err)
		t.Failed()
	} else {
		if users[0].ID != id {
			t.Error(users[0].ID)
		}
	}
}

func TestDeleteUserById(t *testing.T) {
	//if !dao.DeleteUserById(1000){
	//	t.Error("delete ok?")
	//}
}

func TestAddRoleToUser(t *testing.T) {
	user, err := dao.GetUserById(1)
	if err != nil {
		t.Error(err)
	}
	role, err := dao.GetRoleById(1)
	if err != nil {
		t.Error(err)
	}
	res := dao.AddRoleToUser(role, user)
	if !res {
		t.Error("failed")
	}
}

func TestGetRolesFromUser(t *testing.T) {
	user, err := dao.GetUserById(1)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := dao.GetRolesFromUser(user)

	if len(res) == 0 {
		t.Error("NO ROLES")
	} else {
		//for i, role := range res{
		//	//t.Log("role"+strconv.Itoa(i)+":"+utils.GetStructJsonString(role))
		//}
	}
}
