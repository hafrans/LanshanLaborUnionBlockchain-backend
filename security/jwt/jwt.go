package jwt

import (
	"RizhaoLanshanLabourUnion/security/jwt/handler"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var AuthMiddleWare *jwt.GinJWTMiddleware

func InitJwt() {
	var err error
	AuthMiddleWare, err = jwt.New(
		&jwt.GinJWTMiddleware{
			Key:             []byte(utils.JWTSettings.Key),
			Realm:           utils.JWTSettings.Realm,
			Timeout:         time.Hour * 24,
			MaxRefresh:      time.Hour,
			IdentityKey:     utils.JWTSettings.IdentityKey,
			Authenticator:   handler.Authenticator,
			IdentityHandler: handler.IdentityHandler,
			PayloadFunc:     handler.PayloadHandler,
			Authorizator: func(data interface{}, c *gin.Context) bool {
				return true
			},
			Unauthorized: func(context *gin.Context, code int, s string) {
				context.JSON(code, vo.Common{
					Status:    code,
					Message:   s,
					Timestamp: utils.NowTime(),
				})
			},
			LoginResponse: handler.LoginResponse,

			TokenLookup: "header: Authorization, query: token, cookie: jwt",
			// TokenLookup: "query:token",
			// TokenLookup: "cookie:token",

			// TokenHeadName is a string in the header. Default value is "Bearer"
			TokenHeadName: "Bearer",

			// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
			TimeFunc: time.Now,

			LogoutResponse: handler.LogoutResponse,

		})

	if err != nil {
		log.Fatalln("JWT Init failed: " + err.Error())
	}

}
