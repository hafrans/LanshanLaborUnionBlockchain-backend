package dao_test

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"encoding/json"
	"log"
	"math"
	"strconv"
	"testing"
	"time"
)

func init(){
	utils.InitSettings()
	dao.InitDB()
	dao.CreateTables()
}

var id int64

func TestCreateUser(t *testing.T) {

	user, err := dao.CreateUser("test", "test", "hafrans@163.com", "13800138000",models.USER_TYPE_ADMIN, false, false, true)

	if err != nil{
		t.Error(err)
	}else{
		id = user.ID
		b,_ := json.Marshal(user)
		t.Log(string(b))
	}

}

func TestLoginUser(t *testing.T) {
	user := &models.User{}

	if !dao.LoginUser(user){
		t.Error("login failed")
	}

	if math.Abs(user.LastLoginTime.Sub(time.Now()).Seconds()) > 2{
		t.Error("abs check failed")
	}

}

func TestGetUserById(t *testing.T) {
	user , _ := dao.GetUserById(id)
	if user.ID != id{
		t.Error("User id is not id")
	}
}

func TestGetAllUserPaginated(t *testing.T) {
	users, count, err := dao.GetUserAllPaginated(0,50)
	t.Log("count:",count)
	if err != nil{
		log.Println(err)
		t.Failed()
	}else{
		if users[0].ID != id{
			t.Error(users[0].ID)
		}
	}
}

func TestDeleteUserById(t *testing.T) {
	if !dao.DeleteUserById(1000){
		t.Error("delete ok?")
	}
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
	res := dao.AddRoleToUser(role,user)
	if !res{
		t.Error("failed")
	}
}

func TestGetRolesFromUser(t *testing.T) {
	user, err := dao.GetUserById(1)
	if err != nil {
		t.Error(err)
		return
	}

	res , err := dao.GetRolesFromUser(user)

	if len(res) == 0 {
		t.Error("NO ROLES")
	}else{
		for i, role := range res{
			t.Log("role"+strconv.Itoa(i)+":"+utils.GetStructJsonString(role))
		}
	}
}

