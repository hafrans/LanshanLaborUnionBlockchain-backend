package dao_test

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/utils"
	"strconv"
	"testing"
)

func TestCreateRole(t *testing.T) {

	if role, err := dao.CreateRole("管理员测试","Golang Test","admintest"); err == nil{
		t.Log(utils.GetStructJsonString(role))
	}

}

func TestCreatePermission(t *testing.T) {
	if role, err := dao.CreatePermission("管理员测试-权限","Golang Test","pm:admintest"); err == nil{
		t.Log(utils.GetStructJsonString(role))
	}
	if role, err := dao.CreatePermission("管理员测试-权限2","Golang Test","pm:admintest2"); err == nil{
		t.Log(utils.GetStructJsonString(role))
	}
}

func TestAddPermissionToRole(t *testing.T) {
	role,err := dao.GetRoleById(1)
	if err != nil{
		t.Failed()
	}
	perm, err := dao.GetPermissionById(1)
	if err != nil{
		t.Failed()
	}

	perm2, err := dao.GetPermissionById(2)
	if err != nil{
		t.Failed()
	}

	res := dao.AddPermissionToRole(perm,role)
	if !res {
		t.Error(utils.GetStructJsonString(perm))
		t.Error(utils.GetStructJsonString(role))
	}
	res2 := dao.AddPermissionToRole(perm2,role)
	if !res2 {
		t.Error(utils.GetStructJsonString(perm))
		t.Error(utils.GetStructJsonString(role))
	}
}

func TestGetPermissionsFromRole(t *testing.T) {

	role,err := dao.GetRoleById(1)
	if err != nil{
		t.Failed()
	}
	res , err := dao.GetPermissionsFromRole(role)
	if err != nil{
		t.Error(utils.GetStructJsonString(role))
	}
	if len(res) > 0{
		for i, perm := range res{
			t.Log("perm"+strconv.Itoa(i)+":"+utils.GetStructJsonString(perm))
		}
	}else{
		t.Error(" no perms!")
	}
}

