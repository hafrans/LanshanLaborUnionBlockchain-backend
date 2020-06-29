package vo

import (
	"RizhaoLanshanLabourUnion/utils"
)

type LaborArbitrationForm struct {

	// 1. 主体性质
	Subject int `json:"main" binding:"required,number,gte=0,lte=2"`
	// 2. 入职时间
	HireDate *utils.Date `json:"hire_date" time_format:"2006-01-02" binding:"required"`

	// 3.1 是否签订书面劳动合同
	LaborContractSigned bool `json:"labor_contract_signed""`
	// 3.2 签订时间
	LaborContractSigningTime *utils.Date `json:"labor_contract_signing_date" json:"omitempty"` // 没找到
	// 3.3 签订次数
	LaborContractSigningCount *int `json:"number" binding:"omitempty,number,gte=0"`

	// 4.劳动合同起止时间
	LaborContractType  *int           `json:"labor_contract_type" binding:"omitempty,number,gte=0"`
	LaborContractRange [2]*utils.Date `json:"labor_contract" time_format:"2006-01-02" binding:"omitempty"`

	// 5. 未签书面劳动合同工作期间
	LaborContractNotSignedRange [2]*utils.Date `json:"labor_contract_not_signed" time_format:"2006-01-02" binding:"omitempty"`

	// 6. 劳动合同约定的工作岗位
	Job *string `json:"job_name" `

	// 7. 劳动合同约定的工作地点
	WorkPlace *string `json:"job_work_place" form:"job_work_place"`

	// 8. 劳动合同约定的月工资数、工资构成、工时制
	// 8.1 劳动合同约定的月工资数
	ContractWage *float64 `json:"contract_wage" form:"contract_wage"  binding:"omitempty,number,gte=0"`
	// 8.2 劳动合同约定的工资构成
	ContractWageComponent *string `json:"contract_wage_component" binding:"omitempty"`
	// 8.3 工时制
	ContractWageType *int `json:"contract_wage_type" binding:"omitempty,number,gte=0,lte=2"`

	// 9. 实发月工资数及工资构成、发放形式、发放周期
	// 9.1 月工资数
	Wage float64 `json:"wage" binding:"number,required,gte=0"`
	// 9.2 工资构成
	WageComponent string `json:"wage_component" form:"wage_component"`
	// 9.3 发放形式
	PaymentType int `json:"payment_type" binding:"required,gte=0,lte=2"`
	// 9.4 发放周期
	PaymentCycle int `json:"payment_cycle" binding:"required,gte=0,lte=2"` // 没找到

	// 10. 最后一次支付工资时间
	LastPayment *utils.DateMonth `json:"last_payment" time_format:"2006-01" example:"2010-01"` //格式不对

	// 11. 欠发工资及加班费数额
	// 11.1 欠发工资数
	UnpaidWage float64 `json:"unpaid_wages"  binding:"number,gte=0"`
	// 11.2 欠发加班费奖金
	UnpaidOvertimeWage float64 `json:"overtime_bonus"   binding:"number,gte=0"`

	// 12.办理社会保险及险种
	// 12.1 是否进行社会保险
	SocialInsuranceApply bool `json:"social_insurance"`
	// 12.2 社会保险险种
	SocialInsuranceType *int `json:"social_insurance_type" binding:"omitempty,gte=0,lte=1"`
	// 12.3 社会保险缴险时间
	SocialInsuranceApplyRange [2]*utils.DateMonth `json:"social_insurance_payment_time" time_format:"2006-01"` //格式不对

	// 13. 发生工伤时间、工伤认定情况
	// 13.1 发生工伤时间
	WorkRelatedInjuredDate *utils.Date `json:"injury_time" time_format:"2006-01-02" `
	// 13.2 是否经过工伤认定
	WorkRelatedInjuredIdentification *bool `json:"industrial_injury" `

	// 14. 住院起止时间
	HospitalRange [2]*utils.Date `json:"hospitalization_time" time_format:"2006-01-02"`

	// 15. 伤残等级鉴定时间及结果
	// 15.1 伤残等级鉴定时间
	DisabledAppraisalDate *utils.Date `json:"disability_rating_time"`
	// 15.2 伤残等级鉴定结果
	DisabledAppraisalResult *int `json:"disable_results" binding:"omitempty,number,gte=1,lte=10"`

	// 16. 停工留薪期限时间  //格式都不对
	// 16.1 劳动者
	PayForWorkStoppageLabor *float64 `json:"laborer_time" binding:"omitempty,number,gte=0" example:"1"`
	// 16.2 用人单位
	PayForWorkStoppageEmployer *float64 `json:"employer_time" binding:"omitempty,number,gte=0" example:"1"`

	// 17.需支付的工伤待遇项目及数额
	// 17.1 医疗费
	WorkRelatedTreatmentAmountYiliaofei *float64 `json:"work_related_treatment_ylf" binding:"omitempty,number,gte=0"`
	// 17.2 假肢安装费
	WorkRelatedTreatmentAmountJiazhianzhuang *float64 `json:"work_related_treatment_jzaz"  binding:"omitempty,number,gte=0" `
	// 17.3 住院期间伙食补助
	WorkRelatedTreatmentAmountHuoshibuzhu *float64 `json:"work_related_treatment_hsbz"  binding:"omitempty,number,gte=0" ` // 没找到
	// 17.4 交通费
	WorkRelatedTreatmentAmountJiaotong *float64 `json:"work_related_treatment_jt" binding:"omitempty,number,gte=0" `
	// 17.5 陪护费
	WorkRelatedTreatmentAmountPeihu *float64 `json:"work_related_treatment_ph"  binding:"omitempty,number,gte=0"`
	// 17.6 生活护理费
	WorkRelatedTreatmentAmountShenghuohuli *float64 `json:"work_related_treatment_shhl" binding:"omitempty,number,gte=0"`
	// 17.7 伤残津贴
	WorkRelatedTreatmentAmountShangcanjintie *float64 `json:"work_related_treatment_scjt" binding:"omitempty,number,gte=0"`
	// 17.8 一次性伤残补助金
	WorkRelatedTreatmentAmountYcxshangcanbuzhu *float64 `json:"work_related_treatment_scbz" binding:"omitempty,number,gte=0"`
	// 17.9. 一次性工伤医疗补助金
	WorkRelatedTreatmentAmountYcxgongshangyiliaobuzhu *float64 `json:"work_related_treatment_gsylbz" binding:"omitempty,number,gte=0"`
	// 17.10. 一次性伤残就业补助金
	WorkRelatedTreatmentAmountYcxshangcanjiuyebuzhu *float64 `json:"work_related_treatment_scjybz" binding:"omitempty,number,gte=0"`
	// 17.11. 其他
	WorkRelatedTreatmentAmountOther *float64 `json:"work_related_treatment_other" binding:"omitempty,number,gte=0"`

	// 18. 加班时间
	OvertimeWeekday *float64 `json:"normal_overtime"  example:"10"`  // 正常工作日加班小时
	OvertimeWeekend *float64 `json:"statutory_rest"  example:"10"`   // 法定休息日加班小时
	OvertimeHoliday *float64 `json:"statutory_holidays"example:"10"` // 法定节假日加班小时

	// 19. 加班工资计算基数
	OvertimeWageBase *float64 `json:"overtime_wage_base" example:"2"`

	// 20. 双方解除或终止劳动关系前12个月劳动者月平均工资额
	BeforeSeverLaborRelationshipAvgWage *float64       `json:"before_sever_labor_relationship_avg_wage" example:"1"`
	BeforeSeverLaborRelationshipRange   [2]*utils.Date `json:"before_sever_labor_relationship_range" time_format:"2006-01-02" binding:"omitempty"`

	// 21. 劳动者在本单位工作年限
	WorkYear *int `json:"work_year" form:"work_year" binding:"number,gte=0"`

	// 22. 未休带薪年休假天数
	PaidAnnualLeaveNotLeaveDay         bool `json:"paid_annual_leave_not_leave_day"`
	PaidAnnualLeaveNotLeaveDayShould   *int `json:"should_rest"  binding:"required,number,gte=0"`   // 应休
	PaidAnnualLeaveNotLeaveDayActual   *int `json:"real_rest" binding:"required,number,gte=0" `     // 实休
	PaidAnnualLeaveNotLeaveDayNotLeave *int `json:"not_take_days" binding:"required,number,gte=0" ` // 未休

	// 23. 扣除加班工资后十二个月劳动者月平均工资数额
	NoOvertimeAvgWage      *float64       `json:"no_overtime_avg_wage"  example:"1"`
	NoOvertimeAvgWageRange [2]*utils.Date `json:"no_overtime_avg_wage_range" time_format:"2006-01-02" binding:"omitempty"`

	// 24.双方发生劳动争议时间
	LaborDisputeDate *utils.Date `json:"parties_dispute_time" gorm:"type:date" example:"2020-01-02" binding:"required"`

	// 25. 双方解除或终止劳动关系的原因
	SeverLaborRelationshipEmployer       *bool   `json:"sever_labor_relationship_employer" form:"sever_labor_relationship_employer" gorm:"type:tinyint(1)"`
	SeverLaborRelationshipLabor          *bool   `json:"sever_labor_relationship_labor" form:"sever_labor_relationship_labor" gorm:"type:tinyint(1)"`
	SeverLaborRelationshipEmployerReason *string `json:"sever_labor_relationship_employer_reason" form:"sever_labor_relationship_employer_reason"`
	SeverLaborRelationshipLaborReason    *string `json:"sever_labor_relationship_labor_reason" form:"sever_labor_relationship_labor_reason"`

	// 26.解除或终止劳动关系时间年月日
	SeverLaborRelationshipDate *utils.Date `json:"sever_labor_relationship_date" example:"2020-01-02"`

	// 27. 已办理劳动合同解除手续
	SeveredLaborRelationship *bool `json:"labor_contract_completed"`

	// 28. 申请仲裁时间年月日
	LaborArbitrationDate *utils.Date `json:"arbitration_time" binding:"required"`

	// 29. 涉及群体性是□否□
	MassDisturbance bool `json:"involve_group"`

	// 30. 本表遗漏的其他项目
	Other string `json:"other_information" form:"other_information"`
}


type SimplifiedLaborArbitrationResult struct {
	ID        int64       `json:"id"`
	CreatedAt *utils.Time `json:"created_at"`
	UpdatedAt *utils.Time `json:"updated_at"`
	Owner     int64       `json:"owner"`
}
