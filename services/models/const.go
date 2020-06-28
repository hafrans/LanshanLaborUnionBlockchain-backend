package models

// 主体性质
const (
	SubjectLabor         = 1 + iota // 劳动者
	SubjectEmployer                 // 用人单位
	SubjectEmployingUnit            // 用工单位
	SubjectOther                    // 其他
)

// 劳动合同类型
const (
	LaborContractTypeFix    = 1 + iota // 固定期限
	LaborContractTypeNotFix            // 无固定期限
	LaborContractTypeOnce              // 完成一定工作任务为期限
	LaborContractTypeOther             // 其他
)

// 工资类型

const (
	WageTypeStandard      = 1 + iota // 标准工时制
	WageTypeComprehensive            // 综合工时制
	WageTypeIrregular                // 不定时工作制
	WageTypeOther                    // 其他
)

// 工资发放形式
const (
	PayoffTypeCash = 1 + iota
	PayoffTypeBank
	PayoffTypeOther
)

// 支付周期
const (
	PaymentMonth = 1 + iota
	PaymentNextMonth
	PaymentOther
)

// 社会保险险种
const (
	SocialInsuranceTypeWuxian    = 1 + iota // 五险
	SocialInsuranceTypeGongshang            // 工伤险
	SocialInsuranceTypeOther                // 其他
)

const (
	StatusSubmitted        = iota // 已提交
	StatusPending                 // 正在处理
	StatusResultConfirming        // 当事人确认调解结果
	StatusRefused                 // 拒绝调解
	StatusConfirmed               // 确认调解
	StatusCompleted               // 调解结束
)
