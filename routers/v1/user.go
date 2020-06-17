package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context){

	claims := jwtmodel.ExtractUserClaimsFromGinContext(c)
	user, _ := dao.GetUserById(claims.Id)
	c.JSON(200, vo.UserData{
		Common: vo.GenerateCommonResponseHead(0,"success"),
		UserDataPayload:vo.UserDataPayload{
			User: user,
			Claims: claims,
		},
	})

}
