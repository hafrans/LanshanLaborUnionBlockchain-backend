package security

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"github.com/mikespook/gorbac"
	"log"
)

var rbac *gorbac.RBAC

func InitRBAC() {

	rbac = gorbac.New()

	roles, _ , err := dao.GetRoleAllPaginated(0, 10000)
	if err != nil {
		log.Panic("Can not fetch roles")
	}
	for _, role := range roles{

	}



}
