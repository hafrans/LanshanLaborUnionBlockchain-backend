package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
	"log"
)

// Get All Categories
// @Summary 获得所有类型
// @Description 获取所有类型
// @Tags category,case
// @Produce json
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common
// @Router /api/v1/category [get]
func GetAllCategories(ctx *gin.Context) {

	list, _, err := dao.GetCategoryAllPaginated(0, 50)
	if err != nil {
		log.Println(err)
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "出现错误："+err.Error()))
		return
	}

	ctx.JSON(respcode.HttpOK, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
		Data:   list,
	})
}

// Create New Case By Applicant
// @Summary 创建新调解案件
// @Description 由申请人填写创建新调解案件
// @Tags case
// @Accept json
// @Produce json
// @Param case body vo.CaseFirstSubmitForm true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/create [get]
func CreateNewCaseByApplicant(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	var form vo.CaseFirstSubmitForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		// 表单绑定失败
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	}

	_ = utils.PopulateCaseBasicFromFormToModel(&form, claims.Id, "371100")


}
