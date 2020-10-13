package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/blockchain"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/smsqueue/smsrpc"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

// Create New Suggestion According to Case By caseID
// @Summary 根据 CaseID 创建部门建议
// @Description 由管理员或者部门人员创建部门建议
// @Tags case,suggestion
// @Accept json
// @Produce json
// @Param case body vo.Suggestion true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/suggestion/create [post]
func CreateSuggestion(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限创建部门建议"))
		return
	}

	var form vo.Suggestion
	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err)
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, err.Error()))
	} else {
		var _case *models.Case
		var cErr error
		// find case
		if _case, cErr = dao.GetCaseNotPreloadedModelByCaseID(form.CaseID); cErr != nil {
			if cErr == gorm.ErrRecordNotFound {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "关联案件号不存在"))
			} else {
				log.Println(cErr)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "关联案件状态异常"))
			}
			return
		}

		// just add suggestion

		model, cErr := dao.CreateSuggestion(form.Content, form.CaseID, claims.Id)
		if cErr != nil {
			log.Println(cErr)
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "部门建议创建失败"))
		} else {
			if suggestion, rErr := dao.GetPreloadedSuggestionById(model.ID); rErr != nil {
				log.Println(cErr)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "部门建议创建异常，但系统数据已成功提交"))
			} else {
				// 记录
				blockchain.CreateHistoryByUsingModel(suggestion.CaseID, "创建部门建议", suggestion, claims.Id)
				// 推送消息
				go smsrpc.SendSuggestion(_case)
				ctx.JSON(respcode.HttpOK, vo.CommonData{
					Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
					Data:   utils.PopulateSuggestionFromModelToVO(suggestion),
				})
			}

		}

	}

}

// Delete One Suggestion According to Case by suggestion id
// @Summary 根据部门建议 的id号码删除
// @Description 根据部门建议 的id号码删除，只有管理员、部门人员可以操作
// @Tags labor,suggestion
// @Produce json
// @Param id path number true "表单id"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/suggestion/delete/:id [get]
func DeleteSuggestion(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限删除部门建议"))
		return
	}

	// 解析record id

	if suggestionId, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法ID"))
	} else {
		if suggestion, me := dao.GetPreloadedSuggestionById(int64(suggestionId)); me != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "部门建议不存在"))
		} else {

			// 记录
			blockchain.CreateHistoryByUsingModel(suggestion.CaseID, "删除部门建议", suggestion, claims.Id)

			if dao.DeleteSuggestionById(int64(suggestionId)) {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "部门建议删除成功"))
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "部门建议删除失败"))
			}
		}
	}
}
