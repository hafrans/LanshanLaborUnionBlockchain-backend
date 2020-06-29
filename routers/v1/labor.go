package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	utils2 "RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
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
// @Router /api/v1/labor/arbitration_instructor [get]
// @Router /api/v1/labor/arbitration_instructor [post]
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
// @Router /api/v1/labor/arbitration/create [post]
func CreateLaborArbitrationForm(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	var form vo.LaborArbitrationForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	} else {
		model, err := utils2.PopulateLaborArbitrationFormToModel(&form)
		model.Owner = claims.Id
		if err != nil {
			ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		}
		model, err = dao.CreateLaborArbitration(model)
		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		} else {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils2.PopulateLaborArbitrationModelToForm(model),
			})
		}

	}

}

// Get My Labor Arbitration Forms
// @Summary 获取自己所有的劳动争议案件审判要素表
// @Description 劳动争议案件审判要素表列表
// @Tags labor
// @Accept json
// @Produce json
// @Param page query number true "页码"
// @Param pageSize query number true "页大小"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/labor/arbitration/ [get]
func GetMyLaborArbitrationForms(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageCount, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageCount < 1 {
		pageCount = 10
	}

	list, total, err := dao.GetLaborArbitrationAllPaginatedOwnByUser(pageNum, pageCount, claims.Id)

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
				"list":  utils2.SimplifyLaborArbitrationResult(list),
				"total": total,
				"size":  pageCount,
				"page":  pageNum,
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
// @Router /api/v1/labor/arbitration/:id [get]
func GetOneLaborArbitrationFormById(ctx *gin.Context) {

	// claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	if formId, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "非法ID"))
		return
	} else {

		model, err := dao.GetLaborArbitrationById(int64(formId))
		if err != nil {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "获取失败"))
			return
		} else {
			// TODO 在这里检查权限

			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils2.PopulateLaborArbitrationModelToForm(model),
			})
			return
		}

	}

}
