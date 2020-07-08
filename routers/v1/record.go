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

// Create New Record According to Case By caseID
// @Summary 根据 CaseID 创建记录
// @Description 由管理员或者部门人员创建记录
// @Tags case,record
// @Accept json
// @Produce json
// @Param case body vo.Record true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/record/create [post]
func AddRecord(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限创建笔录"))
		return
	}

	var form vo.Record
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
			// just add record
			model, cErr := dao.CreateRecord(form.Name, form.Path, form.CaseID, claims.Id)
			if cErr != nil {
				log.Println(cErr)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "记录创建失败"))
			} else {

				if record, rErr := dao.GetRecordById(model.ID); rErr != nil {
					log.Println(cErr)
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "笔录创建异常，但系统数据已成功提交"))
				} else {

					// 记录
					blockchain.CreateHistoryByUsingModel(myCase.CaseID, "添加笔录", record, claims.Id)

					ctx.JSON(respcode.HttpOK, vo.CommonData{
						Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
						Data:   utils.PopulateRecordFromModelToVO(record),
					})
				}

			}
		}

	}

}

// Delete One Record According to Case by record id
// @Summary 根据record 的id号码删除record
// @Description 根据record 的id号码删除record，只有管理员、部门人员可以操作
// @Tags labor,record
// @Produce json
// @Param id path number true "表单id"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/record/delete/:id [get]
func DeleteRecord(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限删除笔录"))
		return
	}

	// 解析record id

	if recordId, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法笔录ID"))
	} else {
		if record, me := dao.GetRecordById(int64(recordId)); me != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "该笔录不存在"))
		} else {

			// 记录
			blockchain.CreateHistoryByUsingModel(record.CaseID, "删除笔录", record, claims.Id)

			if dao.DeleteRecordById(int64(recordId)) {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "笔录删除成功"))
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "笔录删除失败"))
			}
		}

	}
}
