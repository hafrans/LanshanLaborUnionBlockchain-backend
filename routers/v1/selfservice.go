package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
	"log"
)


// 用户自认证手机号
// @Summary 自服务，实名制认证自己的手机号
// @Description 实名制认证，系统发送一个验证码，用户提交验证码后，如果验证通过，则将用户手机号实名标记设置为true
// @Tags user,selfserivce
// @Accept json
// @Produce json
// @Param email body vo.SelfCheckPhone true  "请求"
// @Success 200 {object} vo.Common "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/auth/employer/register [post]
func SelfCheckPhone(ctx *gin.Context){
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	var form vo.SelfCheckPhone
	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err)
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, err.Error()))
	}else{
		// get user entity
		user, _ := dao.GetUserById(claims.Id)
		// check sms captcha
		captcha,err :=  SMSCache.Get(user.Phone)
		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "验证异常，请稍后再试"))
		}else{
			// check captcha is equal to captcha in form
			if form.PhoneCaptcha == captcha{
				// manipulation procedure
				dao.SetUserEmailAndPhoneConfirmedFlag(user, true, user.EmailChecked)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "实名认证成功"))
			}else{
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "验证码有误"))
			}
		}
	}
}