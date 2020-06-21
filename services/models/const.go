package models

// 主体性质
const (
	SubjectLabor         = iota // 劳动者
	SubjectEmployer             // 用人单位
	SubjectEmployingUnit        // 用工单位
	SubjectOther                // 其他
)

// 劳动合同类型
const (
	LaborContractTypeFix    = iota // 固定期限
	LaborContractTypeNotFix        // 无固定期限
	LaborContractTypeOnce          // 完成一定工作任务为期限
	LaborContractTypeOther         // 其他
)

// 工资类型

const (
	WageTypeStandard      = iota // 标准工时制
	WageTypeComprehensive        // 综合工时制
	WageTypeIrregular            // 不定时工作制
	WageTypeOther                // 其他
)

// 工资发放形式
const (
	PayoffTypeCash = iota
	PayoffTypeBank
	PayoffTypeOther
)

// 支付周期
const (
	PaymentMonth = iota
	PaymentNextMonth
	PaymentOther
)

// 社会保险险种
const (
	SocialInsuranceTypeWuxian    = iota // 五险
	SocialInsuranceTypeGongshang        // 工伤险
	SocialInsuranceTypeOther            // 其他
)
