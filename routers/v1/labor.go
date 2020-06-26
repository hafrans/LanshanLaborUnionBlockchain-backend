package v1

import (
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"github.com/gin-gonic/gin"
)

// GetLaborArbitrationFormInstructor
// @Summary 获取劳动争议表单模板
// @Description 获取劳动争议表单模板
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
