package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"errors"
)

func PopulateLaborArbitrationVOToModel(form *vo.LaborArbitrationForm) (*models.LaborArbitration, error) {

	if form == nil {
		return nil, errors.New("no form found")
	}

	model := &models.LaborArbitration{
		BeforeSeverLaborRelationshipAvgWage:      form.BeforeSeverLaborRelationshipAvgWage,
		BeforeSeverLaborRelationshipAvgWageStart: form.BeforeSeverLaborRelationshipRange[0],
		BeforeSeverLaborRelationshipAvgWageEnd:   form.BeforeSeverLaborRelationshipRange[1],
		ContractWage:                             form.ContractWage,
		ContractWageComponent:                    form.ContractWageComponent,
		ContractWageType:                         form.ContractWageType,
		DisabledAppraisalDate:                    form.DisabledAppraisalDate,
		DisabledAppraisalResult:                  form.DisabledAppraisalResult,
		HireDate:                                 form.HireDate,
		HospitalStart:                            form.HospitalRange[0],
		HospitalEnd:                              form.HospitalRange[1],
		Job:                                      form.Job,
		LaborArbitrationDate:                     form.LaborArbitrationDate,
		LaborContractStart:                       form.LaborContractRange[0],
		LaborContractEnd:                         form.LaborContractRange[1],
		LaborContractNotSignedStart:              form.LaborContractNotSignedRange[0],
		LaborContractNotSignedEnd:                form.LaborContractNotSignedRange[1],
		LaborContractSigned:                      form.LaborContractSigned,
		LaborContractSigningCount:                form.LaborContractSigningCount,
		LaborContractSigningTime:                 form.LaborContractSigningTime,
		LaborContractType:                        form.LaborContractType,
		LaborDisputeDate:                         form.LaborDisputeDate,
		LastPayment:                              form.LastPayment,
		MassDisturbance:                          form.MassDisturbance,
		NoOvertimeAvgWage:                        form.NoOvertimeAvgWage,
		NoOvertimeAvgWageStart:                   form.NoOvertimeAvgWageRange[0],
		NoOvertimeAvgWageEnd:                     form.NoOvertimeAvgWageRange[1],
		Other:                                    form.Other,
		OvertimeHoliday:                          form.OvertimeHoliday,
		OvertimeWageBase:                         form.OvertimeWageBase,
		OvertimeWeekday:                          form.OvertimeWeekday,
		OvertimeWeekend:                          form.OvertimeWeekend,
		PaidAnnualLeaveNotLeaveDay:               form.PaidAnnualLeaveNotLeaveDay,
		PaidAnnualLeaveNotLeaveDayActual:         form.PaidAnnualLeaveNotLeaveDayActual,
		PaidAnnualLeaveNotLeaveDayNotLeave:       form.PaidAnnualLeaveNotLeaveDayNotLeave,
		PaidAnnualLeaveNotLeaveDayShould:         form.PaidAnnualLeaveNotLeaveDayShould,
		PayForWorkStoppageEmployer:               form.PayForWorkStoppageEmployer,
		PayForWorkStoppageLabor:                  form.PayForWorkStoppageLabor,
		PaymentCycle:                             form.PaymentCycle,
		PaymentType:                              form.PaymentType,
		SeveredLaborRelationship:                 form.SeveredLaborRelationship,
		SeverLaborRelationshipDate:               form.SeverLaborRelationshipDate,
		SeverLaborRelationshipEmployer:           form.SeverLaborRelationshipEmployer,
		SeverLaborRelationshipEmployerReason:     form.SeverLaborRelationshipEmployerReason,
		SeverLaborRelationshipLabor:              form.SeverLaborRelationshipLabor,
		SeverLaborRelationshipLaborReason:        form.SeverLaborRelationshipLaborReason,
		SocialInsuranceApply:                     form.SocialInsuranceApply,
		SocialInsuranceStart:                     form.SocialInsuranceApplyRange[0],
		SocialInsuranceEnd:                       form.SocialInsuranceApplyRange[1],
		SocialInsuranceType:                      form.SocialInsuranceType,
		Subject:                                  form.Subject,
		UnpaidOvertimeWage:                       form.UnpaidOvertimeWage,
		UnpaidWage:                               form.UnpaidWage,
		Wage:                                     form.Wage,
		WageComponent:                            form.WageComponent,
		WorkPlace:                                form.WorkPlace,
		WorkRelatedInjuredDate:                   form.WorkRelatedInjuredDate,
		WorkRelatedInjuredIdentification:         form.WorkRelatedInjuredIdentification,
		WorkYear:                                 form.WorkYear,
		WorkRelatedTreatmentAmount: models.WorkRelatedTreatmentAmount{
			WorkRelatedTreatmentAmountHuoshibuzhu:             form.WorkRelatedTreatmentAmountHuoshibuzhu,
			WorkRelatedTreatmentAmountJiaotong:                form.WorkRelatedTreatmentAmountJiaotong,
			WorkRelatedTreatmentAmountJiazhianzhuang:          form.WorkRelatedTreatmentAmountJiazhianzhuang,
			WorkRelatedTreatmentAmountOther:                   form.WorkRelatedTreatmentAmountOther,
			WorkRelatedTreatmentAmountPeihu:                   form.WorkRelatedTreatmentAmountPeihu,
			WorkRelatedTreatmentAmountShangcanjintie:          form.WorkRelatedTreatmentAmountShangcanjintie,
			WorkRelatedTreatmentAmountShenghuohuli:            form.WorkRelatedTreatmentAmountShenghuohuli,
			WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu: form.WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu,
			WorkRelatedTreatmentAmountYcxshangcanbuzhu:        form.WorkRelatedTreatmentAmountYcxshangcanbuzhu,
			WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu:   form.WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu,
			WorkRelatedTreatmentAmountYiliaofei:               form.WorkRelatedTreatmentAmountYiliaofei,
		},
	}

	return model, nil

}

func PopulateLaborArbitrationModelToVO(model *models.LaborArbitration) *vo.LaborArbitrationForm {

	form := &vo.LaborArbitrationForm{
		LaborContractType:                                 model.LaborContractType,
		WorkRelatedTreatmentAmountYiliaofei:               model.WorkRelatedTreatmentAmountYiliaofei,
		WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu:   model.WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu,
		WorkRelatedTreatmentAmountYcxshangcanbuzhu:        model.WorkRelatedTreatmentAmountYcxshangcanbuzhu,
		WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu: model.WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu,
		WorkRelatedTreatmentAmountShenghuohuli:            model.WorkRelatedTreatmentAmountShenghuohuli,
		WorkRelatedTreatmentAmountShangcanjintie:          model.WorkRelatedTreatmentAmountShangcanjintie,
		WorkRelatedTreatmentAmountPeihu:                   model.WorkRelatedTreatmentAmountPeihu,
		WorkRelatedTreatmentAmountOther:                   model.WorkRelatedTreatmentAmountOther,
		WorkRelatedTreatmentAmountJiazhianzhuang:          model.WorkRelatedTreatmentAmountJiazhianzhuang,
		WorkRelatedTreatmentAmountJiaotong:                model.WorkRelatedTreatmentAmountJiaotong,
		WorkRelatedTreatmentAmountHuoshibuzhu:             model.WorkRelatedTreatmentAmountHuoshibuzhu,
		WorkYear:                                          model.WorkYear,
		WorkRelatedInjuredIdentification:                  model.WorkRelatedInjuredIdentification,
		WorkRelatedInjuredDate:                            model.WorkRelatedInjuredDate,
		WorkPlace:                                         model.WorkPlace,
		WageComponent:                                     model.WageComponent,
		Wage:                                              model.Wage,
		UnpaidWage:                                        model.UnpaidWage,
		UnpaidOvertimeWage:                                model.UnpaidOvertimeWage,
		NoOvertimeAvgWage:                                 model.NoOvertimeAvgWage,
		NoOvertimeAvgWageRange: [2]*utils.Date{
			model.NoOvertimeAvgWageStart, model.NoOvertimeAvgWageEnd,
		},
		Subject:             model.Subject,
		SocialInsuranceType: model.SocialInsuranceType,
		SocialInsuranceApplyRange: [2]*utils.DateMonth{
			model.SocialInsuranceStart, model.SocialInsuranceEnd,
		},
		SocialInsuranceApply:                 model.SocialInsuranceApply,
		SeverLaborRelationshipLaborReason:    model.SeverLaborRelationshipLaborReason,
		SeverLaborRelationshipLabor:          model.SeverLaborRelationshipLabor,
		SeverLaborRelationshipEmployerReason: model.SeverLaborRelationshipEmployerReason,
		SeverLaborRelationshipEmployer:       model.SeverLaborRelationshipEmployer,
		SeverLaborRelationshipDate:           model.SeverLaborRelationshipDate,
		SeveredLaborRelationship:             model.SeveredLaborRelationship,
		PaymentType:                          model.PaymentType,
		PaymentCycle:                         model.PaymentCycle,
		PayForWorkStoppageLabor:              model.PayForWorkStoppageLabor,
		PayForWorkStoppageEmployer:           model.PayForWorkStoppageEmployer,
		PaidAnnualLeaveNotLeaveDayShould:     model.PaidAnnualLeaveNotLeaveDayShould,
		PaidAnnualLeaveNotLeaveDayNotLeave:   model.PaidAnnualLeaveNotLeaveDayNotLeave,
		PaidAnnualLeaveNotLeaveDayActual:     model.PaidAnnualLeaveNotLeaveDayActual,
		PaidAnnualLeaveNotLeaveDay:           model.PaidAnnualLeaveNotLeaveDay,
		OvertimeWeekend:                      model.OvertimeWeekend,
		OvertimeWeekday:                      model.OvertimeWeekday,
		OvertimeWageBase:                     model.OvertimeWageBase,
		OvertimeHoliday:                      model.OvertimeHoliday,
		Other:                                model.Other,
		MassDisturbance:                      model.MassDisturbance,
		LastPayment:                          model.LastPayment,
		LaborDisputeDate:                     model.LaborDisputeDate,
		LaborContractSigningTime:             model.LaborContractSigningTime,
		LaborContractSigningCount:            model.LaborContractSigningCount,
		LaborContractSigned:                  model.LaborContractSigned,
		LaborArbitrationDate:                 model.LaborArbitrationDate,
		Job:                                  model.Job,
		HireDate:                             model.HireDate,
		DisabledAppraisalResult:              model.DisabledAppraisalResult,
		DisabledAppraisalDate:                model.DisabledAppraisalDate,
		ContractWageType:                     model.ContractWageType,
		ContractWageComponent:                model.ContractWageComponent,
		ContractWage:                         model.ContractWage,
		BeforeSeverLaborRelationshipAvgWage:  model.BeforeSeverLaborRelationshipAvgWage,
		LaborContractNotSignedRange:          [2]*utils.Date{model.LaborContractNotSignedStart, model.LaborContractNotSignedEnd},
		LaborContractRange:                   [2]*utils.Date{model.LaborContractStart, model.LaborContractEnd},
		HospitalRange:                        [2]*utils.Date{model.HospitalStart, model.HospitalEnd},
		BeforeSeverLaborRelationshipRange:    [2]*utils.Date{model.BeforeSeverLaborRelationshipAvgWageStart, model.BeforeSeverLaborRelationshipAvgWageEnd},
		UpdatedAt:                            utils.GetTime(model.UpdatedAt),
		CreatedAt:                            utils.GetTime(model.CreatedAt),
		ID:                                   model.ID,
		Owner:                                model.Owner,
	}

	return form

}

func SimplifyLaborArbitrationResult(list []*models.LaborArbitration) []*vo.SimplifiedLaborArbitrationResult {

	length := len(list)
	result := make([]*vo.SimplifiedLaborArbitrationResult, 0, length)

	for _, v := range list {

		tmp := new(vo.SimplifiedLaborArbitrationResult)
		tmp.Owner = v.Owner
		tmp.ID = v.ID
		d := utils.Time(v.CreatedAt)
		tmp.CreatedAt = &d
		t := utils.Time(v.UpdatedAt)
		tmp.UpdatedAt = &t

		result = append(result, tmp)
	}

	return result
}
