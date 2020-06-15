package security

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models/vo"
	"RizhaoLanshanLabourUnion/utils"
	"RizhaoLanshanLabourUnion/utils/captchaid"
	"errors"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)



func InitJwt(){

	jwt.New(
		&jwt.GinJWTMiddleware{
			Key: []byte(utils.JWTSettings.Key),
			Realm: utils.JWTSettings.Realm,
			Timeout:     time.Hour,
			MaxRefresh:  time.Hour,
			IdentityKey: utils.JWTSettings.IdentityKey,
			Authenticator: func(c *gin.Context) (interface{}, error) {
				var login vo.UsernameLogin
				if err := c.ShouldBindJSON(&login); err == nil {

					// check captcha
					result := utils.CheckCaptcha(captchaid.CAPTCHA_ID_LOGIN,login.Captcha,login.CaptchaTimestamp,login.CaptchaChallenge)
					if !result {
						return "", errors.New("captcha is invalid")
					}

					// check login
					user, err := dao.GetUserByUserName(login.Username)
					if err != nil {
						return "", errors.New("用户名或密码错误(1061)")
					}else{
						encodedPassword, _ := utils.GenerateHashedPassword(login.Password)
						if user.Credentials == encodedPassword {
							return user, nil
						}else{
							return "", errors.New("用户名或密码错误(1062)")
						}
					}

				}else{
					return "",jwt.ErrMissingLoginValues
				}
			},

		})



}