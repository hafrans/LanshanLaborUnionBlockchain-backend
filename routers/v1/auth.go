package v1

import (
	"RizhaoLanshanLabourUnion/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context){

	ctx.String(200,"UserLogin")

}


func GetCaptcha(ctx *gin.Context){

	captchaId := ctx.Param("id")

	if len(captchaId) == 0 {
		ctx.String(200,"captchaId is too short")
	}else{
		pendingCaptcha := utils.CreateCaptcha(captchaId)
		if pendingCaptcha != nil{
			ctx.JSON(200,gin.H{
				"status": 0,
				"message":"success",
				"data":gin.H{
					"captcha":pendingCaptcha,
				},
			})
		}
	}

}
