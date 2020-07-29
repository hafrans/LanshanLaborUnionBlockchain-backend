package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/blockchain"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

// Create New Comment According to Case By caseID
// @Summary 根据 CaseID 创建双方对质信息
// @Description 由双方人员创建对质信息，管理员不干涉，但是可以进行添加
// @Tags case,comment
// @Accept json
// @Produce json
// @Param case body vo.Comment true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/comment/create [post]
func AddComment(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	var form vo.Comment
	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err)
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, err.Error()))
	} else {

		// find case
		if myCase, cErr := dao.GetCaseNotPreloadedModelByCaseID(form.CaseID); cErr != nil {
			if cErr == gorm.ErrRecordNotFound {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "关联案件号不存在"))
			} else {
				log.Println(cErr)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "关联案件状态异常"))
			}
			return
		} else {

			user, err := dao.GetUserById(claims.Id)
			if err != nil {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "用户状态获取异常"))
				return
			}

			// 如果是当事人才可以添加内容
			if claims.UserType == models.USER_TYPE_LABOR || claims.UserType == models.USER_TYPE_EMPLOYER {
				// 需要检查劳动者身份证
				if claims.UserType == models.USER_TYPE_LABOR {
					if user.UserProfile.ApplicantIdentityNumber != myCase.ApplicantIdentityNumber {
						ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非本人填写对质信息"))
						return
					}
				}
				// 检查用工单位ussc
				if claims.UserType == models.USER_TYPE_EMPLOYER {
					if user.UserProfile.EmployerUniformSocialCreditCode != myCase.EmployerUniformSocialCreditCode {
						ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非本人填写对质信息"))
						return
					}
				}

			} else {
				// 如果取消注释，则管理员用户无法添加
				///ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非本案件关联人员无法添加该内容"))
				//return
			} // 检查完毕

			// just add comment
			model, cErr := dao.CreateComment(user, form.CaseID, form.Content)
			if cErr != nil {
				log.Println(cErr)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "信息创建失败"))
			} else {

				if comment, rErr := dao.GetCommentById(model.ID); rErr != nil {
					log.Println(cErr)
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "信息创建异常，但系统数据已成功提交"))
				} else {

					// 记录
					blockchain.CreateHistoryByUsingModel(myCase.CaseID, "添加对质信息", comment, claims.Id)

					ctx.JSON(respcode.HttpOK, vo.CommonData{
						Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
						Data:   utils.PopulateCommentFromModelToVO(comment),
					})
				}

			}
		}

	}

}

// Delete One Comment According to Case by comment id
// @Summary 根据comment 的id号码删除comment
// @Description 根据comment 的id号码删除comment，只有管理员、部门人员、以及创建者可以操作
// @Tags labor,comment
// @Produce json
// @Param id path number true "comment id"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/comment/delete/:id [get]
func DeleteComment(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	//if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
	//	ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限删除笔录"))
	//	return
	//}

	// 解析record id

	if commentId, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法信息ID"))
	} else {
		if comment, me := dao.GetCommentById(int64(commentId)); me != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "该信息不存在"))
		} else {

			// 非管理员要验证身份
			if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN && claims.Id != comment.UserID {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限删除该信息"))
				return
			}

			// 记录
			blockchain.CreateHistoryByUsingModel(comment.CaseID, "删除对质笔录", comment, claims.Id)

			if dao.DeleteCommentById(int64(commentId)) {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "信息删除成功"))
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "信息删除失败"))
			}
		}

	}
}
