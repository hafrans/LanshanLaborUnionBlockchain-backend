package middlewares

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckAdmin
// A Middleware in gin which checks user have admin role or not
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context){
		// User Must Logged on
		claim := jwtmodel.ExtractUserClaimsFromGinContext(c)
		if (claim.UserType == models.USER_TYPE_ADMIN){
			c.Next()
		}else{
			c.JSON(http.StatusOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您不是管理员，无权操作"))
		}
	}
}
