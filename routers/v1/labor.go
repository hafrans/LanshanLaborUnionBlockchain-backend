package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	utils2 "RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

// GetLaborArbitrationFormInstructor
// @Summary 获取劳动争议表单模板
// @Description 获取劳动争议表单模板
// @Tags test,labor
// @Produce json
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common "未验证"
// @Failure 500 {object} vo.Common "服务器错误"
// @Router /api/v1/test/labor/template [get]
// @Router /api/v1/test/labor/template [post]
func LaborArbitrationFormInstructor(ctx *gin.Context) {

	var form vo.LaborArbitrationForm

	if err := ctx.ShouldBindJSON(&form); err == nil {
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
			Data:   form,
		})
	} else {
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success with error"+err.Error()),
			Data:   form,
		})
		lcsc := 1
		teststr := "测试文本"
		testwage := 12345.6
		testtrue := true
		testmonth := 1.5
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(0, "test"),
			Data: vo.LaborArbitrationForm{
				Subject:                                    0,
				HireDate:                                   utils.NowDateDay(),
				LaborContractSigned:                        true,
				LaborContractSigningTime:                   utils.NowDateDay(),
				LaborContractSigningCount:                  &lcsc,
				LaborContractRange:                         [2]*utils.Date{utils.NowDateDay(), utils.NowDateDay()},
				LaborContractNotSignedRange:                [2]*utils.Date{utils.NowDateDay(), utils.NowDateDay()},
				Job:                                        &teststr,
				WorkPlace:                                  &teststr,
				ContractWage:                               &testwage,
				ContractWageComponent:                      &teststr,
				ContractWageType:                           &lcsc,
				Wage:                                       123.45,
				WageComponent:                              "ceshi test",
				PaymentType:                                0,
				PaymentCycle:                               0,
				LastPayment:                                utils.NowDate(),
				UnpaidWage:                                 0,
				UnpaidOvertimeWage:                         0,
				SocialInsuranceApply:                       true,
				SocialInsuranceType:                        &lcsc,
				SocialInsuranceApplyRange:                  [2]*utils.DateMonth{utils.NowDate(), utils.NowDate()},
				WorkRelatedInjuredDate:                     utils.NowDateDay(),
				WorkRelatedInjuredIdentification:           nil,
				HospitalRange:                              [2]*utils.Date{utils.NowDateDay(), utils.NowDateDay()},
				DisabledAppraisalDate:                      utils.NowDateDay(),
				DisabledAppraisalResult:                    &lcsc,
				PayForWorkStoppageLabor:                    &testmonth,
				PayForWorkStoppageEmployer:                 &testmonth,
				WorkRelatedTreatmentAmountYiliaofei:        &testwage,
				WorkRelatedTreatmentAmountJiazhianzhuang:   &testwage,
				WorkRelatedTreatmentAmountHuoshibuzhu:      &testwage,
				WorkRelatedTreatmentAmountJiaotong:         &testwage,
				WorkRelatedTreatmentAmountPeihu:            &testwage,
				WorkRelatedTreatmentAmountShenghuohuli:     &testwage,
				WorkRelatedTreatmentAmountShangcanjintie:   &testwage,
				WorkRelatedTreatmentAmountYcxshangcanbuzhu: &testwage,
				WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu: &testwage,
				WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu:   &testwage,
				WorkRelatedTreatmentAmountOther:                   &testwage,
				OvertimeWeekday:                                   &testmonth,
				OvertimeWeekend:                                   &testmonth,
				OvertimeHoliday:                                   &testmonth,
				OvertimeWageBase:                                  &testmonth,
				BeforeSeverLaborRelationshipAvgWage:               &testwage,
				BeforeSeverLaborRelationshipRange:                 [2]*utils.Date{utils.NowDateDay(), utils.NowDateDay()},
				WorkYear:                                          &lcsc,
				PaidAnnualLeaveNotLeaveDay:                        true,
				PaidAnnualLeaveNotLeaveDayShould:                  &lcsc,
				PaidAnnualLeaveNotLeaveDayActual:                  &lcsc,
				PaidAnnualLeaveNotLeaveDayNotLeave:                &lcsc,
				NoOvertimeAvgWage:                                 &testwage,
				NoOvertimeAvgWageRange:                            [2]*utils.Date{utils.NowDateDay(), utils.NowDateDay()},
				LaborDisputeDate:                                  utils.NowDateDay(),
				SeverLaborRelationshipEmployer:                    &testtrue,
				SeverLaborRelationshipLabor:                       &testtrue,
				SeverLaborRelationshipEmployerReason:              &teststr,
				SeverLaborRelationshipLaborReason:                 &teststr,
				SeverLaborRelationshipDate:                        utils.NowDateDay(),
				SeveredLaborRelationship:                          &testtrue,
				LaborArbitrationDate:                              utils.NowDateDay(),
				MassDisturbance:                                   false,
				Other:                                             "无",
				LaborContractType:                                 &lcsc,
			},
		})
	}

}

// Create Labor Arbitration
// @Summary 创建劳动争议案件审判要素表
// @Description 劳动争议案件审判要素表新建
// @Tags labor
// @Accept json
// @Produce json
// @Param email body vo.LaborArbitrationForm true  "表单"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/v1/labor/create [post]
func CreateLaborArbitrationForm(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	var form vo.LaborArbitrationForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	} else {
		model, err := utils2.PopulateLaborArbitrationVOToModel(&form)
		if err != nil {
			ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		}
		model.Owner = claims.Id
		model, err = dao.CreateLaborArbitration(model)
		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		} else {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils2.PopulateLaborArbitrationModelToVO(model),
			})
		}

	}

}

// Get My Labor Arbitration Forms
// @Summary 获取自己所有的劳动争议案件审判要素表
// @Description 劳动争议案件审判要素表列表，记住，只能看到自己的！！！
// @Tags labor
// @Accept json
// @Produce json
// @Param page query number true "页码"
// @Param pageSize query number true "页大小"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/labor/list [get]
func GetMyLaborArbitrationFormList(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageCount, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageCount < 1 {
		pageCount = 10
	}

	var list []*models.LaborArbitration
	var total int

	// TODO 权限控制,暂时不用
	list, total, err = dao.GetLaborArbitrationAllPaginatedOwnByUser(pageNum, pageCount, claims.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "列表为空"),
				Data: gin.H{
					"list":  []interface{}{},
					"total": total,
					"size":  pageCount,
					"page":  pageNum,
				},
			})
			return
		}
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "列表获取失败"))
		return
	} else {
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
			Data: gin.H{
				"list":        utils2.SimplifyLaborArbitrationResult(list),
				"total_count": total,
				"page_count":  pageCount,
				"page_num":    pageNum,
			},
		})
	}

}

// Get One Labor Arbitration Forms By Id
// @Summary 获取指定劳动争议案件审判要素表列表
// @Description 劳动争议案件审判要素表获取，非管理员只能看自己的。
// @Tags labor
// @Accept json
// @Produce json
// @Param id path number true "表单id"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/labor/id/:id [get]
func GetOneLaborArbitrationFormById(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	if formId, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法ID"))
		return
	} else {
		model, err := dao.GetLaborArbitrationById(int64(formId))
		if err != nil {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "获取失败"))
			return
		} else {
			if model.Owner == claims.Id || claims.UserType == models.USER_TYPE_ADMIN { // 是否是自己的表单，管理员无视
				ctx.JSON(respcode.HttpOK, vo.CommonData{
					Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
					Data:   utils2.PopulateLaborArbitrationModelToVO(model),
				})
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限获取此内容"))
			}
		}
	}
}

// Delete One Labor Arbitration Forms By Id
// @Summary 删除单个劳动争议案件审判要素表
// @Description 劳动争议案件审判要素表删除，非管理员只能删除自己的，如果有案件依赖则不可以删除
// @Tags labor
// @Produce json
// @Param id path number true "表单id"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/labor/delete/:id [get]
func DeleteOneLaborArbitrationFormById(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	var targetModel *models.LaborArbitration

	if formId, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法ID"))
		return
	} else {

		_, caseCount, err := dao.GetCasesByFormId(int64(formId))

		if err != nil {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "其他错误"))
			return
		}

		if caseCount > 0 { // 说明有依赖
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "有调解案件依赖该表单"))
			return
		}

		targetModel, err = dao.GetLaborArbitrationById(int64(formId)) // 先获取有没有这个form
		if err != nil {
			if err == gorm.ErrRecordNotFound { // record not found
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "该表单不存在"))
				return
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
				return
			}
		} else {
			// 可以删除了,必须是自己的或者管理员
			if targetModel.Owner != claims.Id && claims.UserType != models.USER_TYPE_ADMIN { // 自己能删除自己的，否则是管理员手动删
				log.Printf("触发安全风险 %v 尝试修改 %v 的表单，id为 %v", claims.Id, targetModel.Owner, formId)
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限修改他人的表单！"))
				return
			}

			if dao.DeleteLaborArbitrationById(int64(formId)) {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "删除成功"))
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "删除失败"))
			}
			return
		}

	}
}

// Update Labor Arbitration
// @Summary 修改劳动争议案件审判要素表
// @Description 劳动争议案件审判要素表修改，只能修改未绑定到案件的表单，已绑定到案件的，不可修改，另外，每个人只能修改自己的表单，除了管理员，其他人不可以修改
// @Tags labor
// @Accept json
// @Produce json
// @Param id path integer true "表单id"
// @Param email body vo.LaborArbitrationForm true  "表单"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/v1/labor/update/:id [post]
func UpdateLaborArbitrationForm(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	var formVO vo.LaborArbitrationForm
	var newForm *models.LaborArbitration
	var targetForm *models.LaborArbitration

	if formId, err := strconv.Atoi(ctx.Param("id")); err != nil { // 获取要更新的formID
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法ID"))
		return
	} else {
		if err := ctx.ShouldBindJSON(&formVO); err != nil { // 表单绑定失败
			log.Println(err.Error())
			ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		} else {

			newForm, err = utils2.PopulateLaborArbitrationVOToModel(&formVO)
			if err != nil {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "表单信息迁移错误 "+err.Error()))
				return
			}

			_, caseCount, err := dao.GetCasesByFormId(int64(formId)) // 获取该formID相关的所有case

			if err != nil {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "其他错误"))
				return
			}

			if caseCount > 0 { // 说明有依赖 ， 有依赖不能修改
				//ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "有调解案件依赖该表单，不能修改"))
				//return
			}

			targetForm, err = dao.GetLaborArbitrationById(int64(formId)) // 先获取有没有这个form 一般是有的

			if err != nil {
				if err == gorm.ErrRecordNotFound { // record not found
					ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "该表单不存在"))
					return
				} else {
					ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
					return
				}
			} else {
				// 可以执行修改了
				if targetForm.Owner != claims.Id && claims.UserType != models.USER_TYPE_ADMIN { // 不是自己的 而且缺少管理员权限
					log.Printf("触发安全风险 %v 尝试修改 %v 的表单，id为 %v", claims.Id, targetForm.Owner, formId)
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限修改他人的表单！"))
					return
				}

				newForm.Model = targetForm.Model
				newForm.Owner = targetForm.Owner

				if dao.UpdateLaborArbitration(newForm) {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "修改成功"))
				} else {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "修改失败"))
				}
				return
			}
		}

	}

}
