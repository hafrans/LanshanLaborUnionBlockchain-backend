package security

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"github.com/mikespook/gorbac"
	"log"
)

var rbac *gorbac.RBAC

func InitRBAC() {

	rbac = gorbac.New()
	log.Println("Initializing RBAC Model")
	roles, _ , err := dao.GetRoleAllPaginated(0, 10000)
	if err != nil {
		log.Panic("Can not fetch roles")
	}
	for _, role := range roles{
		log.Println("role "+role.Name+" found.")
		tmpRole := gorbac.NewStdRole(role.Descriptor)
		perms, permError := dao.GetPermissionsFromRole(role)
		if permError == nil {
			for _, perm := range perms{
				log.Println("role "+role.Name+" permission "+perm.Name+" registered.")
				tmpPerm := gorbac.NewStdPermission(perm.Descriptor)
				tmpRole.Assign(tmpPerm)
			}
		}

		rbac.Add(tmpRole)
	}
	log.Println("RBAC Model Initialized")
}



