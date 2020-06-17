package handler

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// User Login
// @Summary 用户登录
// @Description 用户登录并获得token以及失效日期
// @Accept json
// @Produce json
// @Param username body string true  "用户名"
// @Param password  body string true "密码"
// @Param captcha_code  body string true "验证码"
// @Param captcha_time  body string true "验证码时间戳 （2006-01-02 15:04:05）"
// @Param captcha_challenge  body string true "验证码挑战指令"
// @Success 200 {object} vo.LoginResult
// @Failure 401 {object} vo.Common
// @Router /api/auth/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var login vo.UsernameLogin
	if err := c.ShouldBindJSON(&login); err == nil {

		// check captcha
		//result := utils.CheckCaptcha(captchaid.CAPTCHA_ID_LOGIN,
		//	login.Captcha,
		//	login.CaptchaTimestamp,
		//	login.CaptchaChallenge)
		//if !result {
		//	return "", errors.New("captcha is invalid")
		//}

		// check login
		user, err := dao.GetUserByUserName(login.Username)
		if err != nil {
			return "", errors.New("用户名或密码错误(1061)")
		} else {
			if utils.CheckHashedPassword(login.Password, user.Credentials) {
				return jwtmodel.PopulateUserToUserClaims(user), nil
			} else {
				return "", errors.New("用户名或密码错误(1062)")
			}
		}

	} else {
		return "", jwt.ErrMissingLoginValues
	}
}




func IdentityHandler(context *gin.Context) interface{} {
	claims := jwt.ExtractClaims(context)
	return &jwtmodel.UserClaims{
		Id:    int64(claims[utils.JWTSettings.IdentityKey].(float64)),
		Sub:   claims["sub"].(string),
		Roles: claims["roles"].([]interface{}),
		Iss:   claims["iss"].(string),
		Realm: claims["realm"].(string),
	}
}



func PayloadHandler(data interface{}) jwt.MapClaims {

	if v, ok := data.(*jwtmodel.UserClaims); ok {

		userClaims := v

		return jwt.MapClaims{
			utils.JWTSettings.IdentityKey: userClaims.Id,
			"sub":                         userClaims.Sub,
			"roles":                       userClaims.Roles,
			"iss":                         userClaims.Iss,
			"realm":                       userClaims.Realm,
		}
	} else {
		log.Println("payload function encountered an exception!")
		return jwt.MapClaims{}
	}

}


func LoginResponse(context *gin.Context, code int, s string, t time.Time) {
	context.JSON(code, vo.LoginResult{
		Common: vo.Common{
			Status:    0,
			Message:   "login successfully",
			Timestamp: utils.NowTime(),
		},
		TokenResult: vo.TokenResult{
			Expire:     utils.GetTime(t),
			Token:      s,
			RefreshUrl: "/api/auth/refresh_token",
		},
	})
}

