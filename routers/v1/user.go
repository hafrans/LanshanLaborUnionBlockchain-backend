package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// GetUserInfo
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags user
// @Produce json
// @Success 200 {object} vo.UserData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/user/info [get]
func GetUserInfo(c *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(c)
	user, _ := dao.GetUserById(claims.Id)
	c.JSON(200, vo.UserData{
		Common: vo.GenerateCommonResponseHead(0, "success"),
		UserDataPayload: vo.UserDataPayload{
			User:   user,
			Claims: claims,
		},
	})

}

// ResetUserPassword
// @Summary 修改密码
// @Description 用户自行修改密码
// @Tags user
// @Accept json
// @Produce json
// @Param old_password body string true  "原始密码"
// @Param new_password  body string true "新密码，最小长度3 最大长度20"
// @Param confirm_password  body string true "重新输入密码"
// @Success 200 {object} vo.Common "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/v1/user/reset_password [post]
func ResetUserPassword(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	user, err := dao.GetUserById(claims.Id)
	if err != nil {
		log.Println("ResetUserPassword," + err.Error())
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.UserInvalid, "invalid user"))
		return
	}

	// get form

	var resetsForm vo.UserResetPassword

	if err := ctx.ShouldBindJSON(&resetsForm); err == nil {
		// check old password
		if utils.CheckHashedPassword(resetsForm.OldPassword, user.Credentials) {
			user.Credentials, _ = utils.GenerateHashedPassword(resetsForm.NewPassword)
			if dao.UpdateUser(user) {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "密码修改成功"))
			}
		} else {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "原密码不正确"))
		}
	} else {
		log.Println("ResetUserPassword," + err.Error())
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, "bind form failed"))
	}

}

// UpdateUserBasicInfo
// @Summary 修改用户基础信息（邮箱和密码）
// @Description 修改用户邮箱和密码
// @Tags user
// @Accept json
// @Produce json
// @Param email body vo.UserUpdateInfo true  "请求"
// @Success 200 {object} vo.Common "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/v1/user/update_info [post]
func UpdateUserInfo(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	user, err := dao.GetUserById(claims.Id)
	if err != nil || !user.Activated {
		log.Println("UpdateUserInfo," + err.Error())
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.UserInvalid, "invalid user"))
		return
	}

	var userInfoVO vo.UserUpdateInfo

	if err := ctx.ShouldBindJSON(&userInfoVO); err == nil {

		if user.Email != userInfoVO.Email {
			user.Email = userInfoVO.Email
			user.EmailChecked = false
		}

		if user.Phone != userInfoVO.Phone {
			user.Phone = userInfoVO.Phone
			user.PhoneChecked = false
		}

		// save
		if dao.UpdateUser(user) {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "更新信息成功"))
		}

	} else {
		log.Println("ResetUserPassword," + err.Error())
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, "bind form failed"))
	}

}

// Labor User Register
// @Summary 劳动者 账户注册
// @Description 劳动者账户注册
// @Tags user,register
// @Accept json
// @Produce json
// @Param email body vo.UserRegisterLaborForm true  "请求"
// @Success 200 {object} vo.Common "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/auth/labor/register [post]
func RegisterNewLaborUser(ctx *gin.Context) {

	var registerForm vo.UserRegisterLaborForm

	if err := ctx.ShouldBindJSON(&registerForm); err != nil {
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, "bind form failed"+err.Error()))
	} else {

		var profile models.UserProfile
		profile.ApplicantContact = registerForm.Applicant.Contact
		profile.ApplicantAddress = registerForm.Applicant.Address
		profile.ApplicantBirthday = registerForm.Applicant.Birthday
		profile.ApplicantIdentityNumber = registerForm.Applicant.IdentityNumber
		profile.ApplicantNationality = registerForm.Applicant.Nationality
		profile.ApplicantName = registerForm.Applicant.Name

		result, err := dao.CreateUserWithProfile(registerForm.Applicant.Name, registerForm.Password, "", registerForm.Phone, models.USER_TYPE_LABOR, false, false, true, 0, &profile)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate") {
				if strings.Contains(err.Error(), "phone") {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "手机号已被注册"))
				} else if strings.Contains(err.Error(), "email") {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "E-mail已被注册"))
				}
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "注册失败"+err.Error()))
			}
		} else {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(0, "注册成功"),
				Data:   result,
			})
		}

	}

}

// Employer User Register
// @Summary 用人单位 账户注册
// @Description 用人单位账户注册
// @Tags user,register
// @Accept json
// @Produce json
// @Param email body vo.UserRegisterEmployerForm true  "请求"
// @Success 200 {object} vo.Common "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/auth/employer/register [post]
func RegisterNewEmployerUser(ctx *gin.Context) {

	var registerForm vo.UserRegisterEmployerForm

	if err := ctx.ShouldBindJSON(&registerForm); err != nil {
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, "bind form failed"+err.Error()))
	} else {

		var profile models.UserProfile
		profile.EmployerContact = registerForm.Employer.Contact
		profile.EmployerAddress = registerForm.Employer.Address
		profile.EmployerName = registerForm.Employer.Name
		profile.EmployerUniformSocialCreditCode = registerForm.Employer.UniformSocialCreditCode
		profile.EmployerLegalRepresentative = registerForm.Employer.LegalRepresentative

		result, err := dao.CreateUserWithProfile(registerForm.Employer.LegalRepresentative, registerForm.Password, "", registerForm.Phone, models.USER_TYPE_EMPLOYER, false, false, true, 0, &profile)

		if err != nil {
			if strings.Contains(err.Error(), "Duplicate") {
				if strings.Contains(err.Error(), "phone") {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "手机号已被注册"))
				} else if strings.Contains(err.Error(), "email") {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "E-mail已被注册"))
				}
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "注册失败"+err.Error()))
			}
		} else {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(0, "注册成功"),
				Data:   result,
			})
		}

	}

}
