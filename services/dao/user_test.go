package dao_test

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"encoding/json"
	"log"
	"math"
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

	user, err := dao.CreateUser("test", "test", "hafrans@163.com", "13800138000", false, false, true)

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



