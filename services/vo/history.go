package vo

type SimplifiedHistory struct {
	User              string  `json:"user"`                                         // 修改人
	UserID            int64   `json:"user_id"`                                      // 修改人ID
	CaseID            string  `json:"case_id"`                                      // 案件号码
	OperationHash     string  `json:"operation_id" gorm:"varchar(64);unique_index"` // 操作id
	PrevOperationHash *string `json:"prev_operation_id"`                            // 上一操作id
	Operation         string  `json:"operation"`                                    // 操作类型
	Content           string  `json:"content"`                                      // 操作内容
}
