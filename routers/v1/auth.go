package v1

import (
	"RizhaoLanshanLabourUnion/utils"
	"github.com/gin-gonic/gin"
)




// Get Captcha
// @Summary 获得验证码
// @Description 获取验证码
// @Tags captcha
// @Produce json
// @Param id path string true "验证码id"
// @Success 200 {object} vo.CommonData "同时返回验证码的时间戳、base64化图片、验证码challenge指令"
// @Failure 401 {object} vo.Common
// @Router /api/auth/captcha/:id [get]
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
