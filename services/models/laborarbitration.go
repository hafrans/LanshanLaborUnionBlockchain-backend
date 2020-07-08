package models

import (
	"RizhaoLanshanLabourUnion/utils"
)

type LaborArbitration struct {
	Model

	Owner int64 `json:"owner"` // 创建者

	// 1. 主体性质
	Subject int `json:"subject" form:"subject" gorm:"type:tinyint(1);not null" example:"1" binding:"required"`

	// 2. 入职时间
	HireDate *utils.Date `json:"hire_date" form:"hire_date" example:"1999-12-31" time_format:"2006-01-02" gorm:"type:date" binding:"required"`

	// 3. 是否签订书面劳动合同以及次数
	LaborContractSigned       bool        `json:"labor_contract_signed" form:"labor_contract_signed" binding:"required"`
	LaborContractSigningTime  *utils.Date `json:"labor_contract_signing_time" form:"labor_contract_signing_time" gorm:"type:date"  example:"1999-12-31"`
	LaborContractSigningCount *int        `json:"labor_contract_signing_count" form:"labor_contract_signing_count" example:"1"`

	// 4. 劳动合同起止时间（多次的，填最后一次签订的）
	LaborContractType  *int        `json:"labor_contract_type" form:"labor_contract_type" gorm:"type:tinyint(1)" binding:"number"`
	LaborContractStart *utils.Date `json:"labor_contract_start" form:"labor_contract_start" gorm:"type:date" example:"1999-12-31" time_format:"2006-01-02"`
	LaborContractEnd   *utils.Date `json:"labor_contract_end" form:"labor_contract_end" gorm:"type:date" example:"1999-12-31" time_format:"2006-01-02"`

	// 5. 未签书面劳动合同工作期间
	LaborContractNotSignedStart *utils.Date `json:"labor_contract_not_signed_start" form:"labor_contract_not_signed_start" gorm:"type:date" example:"1999-12-31" time_format:"2006-01-02"`
	LaborContractNotSignedEnd   *utils.Date `json:"labor_contract_not_signed_end" form:"labor_contract_not_signed_end" gorm:"type:date" example:"1999-12-31" time_format:"2006-01-02"`

	// 6. 劳动合同约定的工作岗位
	Job *string `json:"job_name" form:"job_name" gorm:"type:varchar(255)"`

	// 7. 劳动合同约定的工作地点
	WorkPlace *string `json:"job_work_place" form:"job_work_place" gorm:"type:text"`

	// 8. 劳动合同约定的月工资数、工资构成、工时制
	ContractWage          *float64 `json:"contract_wage" form:"contract_wage" gorm:"type:varchar(64);default:'0.00'"`
	ContractWageComponent *string  `json:"contract_wage_component" form:"contract_wage_component"`
	ContractWageType      *int     `json:"contract_wage_type" form:"contract_wage_type" gorm:"type:tinyint(1)"`

	// 9. 实发月工资数及工资构成、发放形式、发放周期
	PaymentType   int     `json:"payment_type" form:"payment_type" gorm:"type:tinyint(1)" binding:"numeric,required"`
	PaymentCycle  int     `json:"payment_cycle" form:"payment_cycle"  gorm:"type:tinyint(1)" binding:"numeric,required"`
	Wage          float64 `json:"wage" form:"wage" gorm:"default:0"  binding:"number,required"`
	WageComponent string  `json:"wage_component" form:"wage_component" binding:"required" gorm:"type:text"`

	// 10. 最后一次支付工资时间
	LastPayment *utils.DateMonth `json:"last_payment" form:"last_payment" gorm:"type:date" time_format:"2006-01" example:"2010-01"`

	// 11. 欠发工资及加班费数额
	UnpaidWage         float64 `json:"unpaid_wage" form:"unpaid_wage" gorm:"default:0"  binding:"number"`
	UnpaidOvertimeWage float64 `json:"unpaid_overtime_wage" form:"unpaid_overtime_wage" gorm:"default:0"  binding:"number"`

	// 12.办理社会保险及险种
	SocialInsuranceApply bool             `json:"social_insurance_apply" form:"social_insurance_apply" binding:"required"`
	SocialInsuranceStart *utils.DateMonth `json:"social_insurance_start" form:"social_insurance_start" gorm:"type:date" time_format:"2006-01" example:"2010-01"`
	SocialInsuranceEnd   *utils.DateMonth `json:"social_insurance_end" form:"social_insurance_end" gorm:"type:date" time_format:"2006-01" example:"2010-01"`
	SocialInsuranceType  *int             `json:"social_insurance_type" form:"social_insurance_type" gorm:"type:tinyint(1)" binding:"number"`

	// 13. 发生工伤时间、工伤认定情况
	//WorkRelatedInjured               bool           `json:"work_related_injured" form:"work_related_injured" gorm:"type:tinyint(1)" binding:"required"`
	WorkRelatedInjuredDate           *utils.Date `json:"work_related_injured_date" form:"work_related_injured_date" gorm:"type:date" time_format:"2006-01-02" `
	WorkRelatedInjuredIdentification *bool       `json:"work_related_injured_identification" form:"work_related_injured_identification" gorm:"type:tinyint(1)" `

	// 14. 住院起止时间
	HospitalStart *utils.Date `json:"hospital_start" form:"hospital_start" gorm:"type:date" time_format:"2006-01-02" example:"2009-01-01"`
	HospitalEnd   *utils.Date `json:"hospital_end" form:"hospital_end" gorm:"type:date" time_format:"2006-01-02" example:"2009-02-03"`

	// 15. 伤残等级鉴定时间及结果
	DisabledAppraisalDate   *utils.Date `json:"disabled_appraisal_date" form:"disabled_appraisal_date" gorm:"type:date" time_format:"2006-01-02" `
	DisabledAppraisalResult *int        `json:"disabled_appraisal_result" form:"disabled_appraisal_result" gorm:"type:tinyint(2)" example:"10"`

	// 16. 停工留薪时间
	PayForWorkStoppageLabor    *float64 `json:"pay_for_work_stoppage_labor" form:"pay_for_work_stoppage_labor" example:"1"`
	PayForWorkStoppageEmployer *float64 `json:"pay_for_work_stoppage_employer" form:"pay_for_work_stoppage_employer" example:"1"`

	// 17. 需支付的工伤待遇项目及数额
	WorkRelatedTreatmentAmount

	// 18. 加班时间
	OvertimeWeekday *float64 `json:"overtime_weekday" form:"overtime_weekday" example:"10"` // 工作日
	OvertimeWeekend *float64 `json:"overtime_weekend" form:"overtime_weekend" example:"10"` // 法定休息
	OvertimeHoliday *float64 `json:"overtime_holiday" form:"overtime_holiday" example:"10"` // 法定节假日

	// 19. 加班工资计算基数
	OvertimeWageBase *float64 `json:"overtime_wage_base" form:"overtime_wage_base" example:"2"`

	// 20. 双方解除或终止劳动关系前12个月劳动者月平均工资额
	BeforeSeverLaborRelationshipAvgWage      *float64    `json:"before_sever_labor_relationship_avg_wage" form:"before_sever_labor_relationship_avg_wage" example:"1"`
	BeforeSeverLaborRelationshipAvgWageStart *utils.Date `json:"before_sever_labor_relationship_avg_wage_start" form:"before_sever_labor_relationship_avg_wage_start" gorm:"type:date" example:"2020-01-02"`
	BeforeSeverLaborRelationshipAvgWageEnd   *utils.Date `json:"before_sever_labor_relationship_avg_wage_end" form:"before_sever_labor_relationship_avg_wage_end" gorm:"type:date" example:"2020-01-02"`

	// 21. 劳动者在本单位工作年限
	WorkYear *int `json:"work_year" form:"work_year"`

	// 22. 未休带薪年休假天数
	PaidAnnualLeaveNotLeaveDay         bool `json:"paid_annual_leave_not_leave_day" form:"paid_annual_leave_not_leave" gorm:"type:tinyint(1)" binding:"required"`
	PaidAnnualLeaveNotLeaveDayShould   *int `json:"paid_annual_leave_not_leave_should" form:"paid_annual_leave_not_leave_should"`               // 应休
	PaidAnnualLeaveNotLeaveDayActual   *int `json:"paid_annual_leave_not_leave_actual" form:"paid_annual_leave_not_leave_actual"`               // 实休
	PaidAnnualLeaveNotLeaveDayNotLeave *int `json:"paid_annual_leave_not_leave_day_not_leave" form:"paid_annual_leave_not_leave_day_not_leave"` // 未休

	// 23. 扣除加班工资后十二个月劳动者月平均工资数额
	NoOvertimeAvgWage      *float64    `json:"no_overtime_avg_wage" form:"no_overtime_avg_wage" example:"1"`
	NoOvertimeAvgWageStart *utils.Date `json:"no_overtime_avg_wage_start" form:"no_overtime_avg_wage_start" gorm:"type:date" example:"2020-01-02"`
	NoOvertimeAvgWageEnd   *utils.Date `json:"no_overtime_avg_wage_end" form:"no_overtime_avg_wage_end" gorm:"type:date" example:"2020-01-02"`

	// 24. 双方发生劳动争议时间
	LaborDisputeDate *utils.Date `json:"labor_dispute_date" form:"labor_dispute_date" gorm:"type:date" example:"2020-01-02" binding:"required"`

	// 25. 双方解除或终止劳动关系的原因
	SeverLaborRelationshipEmployer       *bool   `json:"sever_labor_relationship_employer" form:"sever_labor_relationship_employer" gorm:"type:tinyint(1)"`
	SeverLaborRelationshipLabor          *bool   `json:"sever_labor_relationship_labor" form:"sever_labor_relationship_labor" gorm:"type:tinyint(1)"`
	SeverLaborRelationshipEmployerReason *string `json:"sever_labor_relationship_employer_reason" form:"sever_labor_relationship_employer_reason" gorm:"type:text"`
	SeverLaborRelationshipLaborReason    *string `json:"sever_labor_relationship_labor_reason" form:"sever_labor_relationship_labor_reason" gorm:"type:text"`

	// 26. 解除或终止劳动关系时间年月日
	SeverLaborRelationshipDate *utils.Date `json:"sever_labor_relationship_date" form:"sever_labor_relationship_date" gorm:"type:date" example:"2020-01-02"`

	// 27. 已办理劳动合同解除手续是□否□
	SeveredLaborRelationship *bool `json:"severed_labor_relationship" form:"severed_labor_relationship" gorm:"type:tinyint(1)"`

	// 28. 申请仲裁时间年月日
	LaborArbitrationDate *utils.Date `json:"labor_arbitration_date" form:"labor_arbitration_date" gorm:"type:date" example:"2020-01-02" binding:"required"`

	// 29. 涉及群体性是□否□
	MassDisturbance bool `json:"mass_disturbance" form:"mass_disturbance" gorm:"type:tinyint(1)" binding:"required"`

	// 30. 本表遗漏的其他项目
	Other string `json:"other_information" form:"other_information" gorm:"type:text"`
}

// 需支付的工伤待遇项目及数额
type WorkRelatedTreatmentAmount struct {

	// 1. 医疗费
	WorkRelatedTreatmentAmountYiliaofei *float64 `json:"work_related_treatment_yl" form:"work_related_treatment_yl" gorm:"default:0;comment:'医疗费'"`
	// 2. 假肢安装费
	WorkRelatedTreatmentAmountJiazhianzhuang *float64 `json:"work_related_treatment_jzaz" form:"work_related_treatment_jzaz" gorm:"default:0;comment:'假肢安装费'"`
	// 3. 住院期间伙食补助
	WorkRelatedTreatmentAmountHuoshibuzhu *float64 `json:"work_related_treatment_hsbz" form:"work_related_treatment_hsbz" gorm:"default:0;comment:'住院期间伙食补助费'"`
	// 4. 交通费
	WorkRelatedTreatmentAmountJiaotong *float64 `json:"work_related_treatment_jt" form:"work_related_treatment_jt" gorm:"default:0;comment:'交通费'"`
	// 5. 陪护费
	WorkRelatedTreatmentAmountPeihu *float64 `json:"work_related_treatment_ph" form:"work_related_treatment_ph" gorm:"default:0;comment:'陪护费'"`
	// 6. 生活护理费
	WorkRelatedTreatmentAmountShenghuohuli *float64 `json:"work_related_treatment_shhl" form:"work_related_treatment_shhl" gorm:"default:0;comment:'生活护理费'"`
	// 7. 伤残津贴
	WorkRelatedTreatmentAmountShangcanjintie *float64 `json:"work_related_treatment_scjt" form:"work_related_treatment_scjt" gorm:"default:0;comment:'伤残津贴'"`
	// 8. 一次性伤残补助金
	WorkRelatedTreatmentAmountYcxshangcanbuzhu *float64 `json:"work_related_treatment_scbz" form:"work_related_treatment_scbz" gorm:"default:0;comment:'一次性伤残补助金'"`
	// 9. 一次性工伤医疗补助金
	WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu *float64 `json:"work_related_treatment_gsylbz" form:"work_related_treatment_gsylbz" gorm:"default:0;comment:'一次性工伤医疗补助金'"`
	// 10. 一次性伤残就业补助金
	WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu *float64 `json:"work_related_treatment_scjybz" form:"work_related_treatment_scjybz" gorm:"default:0;comment:'一次性伤残就业补助金'"`
	// 11. 其他
	WorkRelatedTreatmentAmountOther *float64 `json:"work_related_treatment_other" form:"work_related_treatment_other" gorm:"default:0;comment:'其他'"`
}
