package smsrpc

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/smsqueue"
	"log"
)

// SendCreateMeetingInfo
// 将会议创建结果发送到所有相关人员的手机上。如果没有验证自己的手机号
// 系统不会发送消息
func SendCreateMeetingInfo(meeting *models.Meeting) {

	// meeting 变量逃逸

	personnel, err := dao.GetMeetingPersonnelsWithUserByMeetingID((meeting).ID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, val := range personnel {
		if val.User.PhoneChecked {
			details := struct {
				Name        string
				MeetingCode string
				MeetingPass *string
				MeetingLink string
			}{
				val.Username,
				meeting.MeetingCode,
				meeting.Password,
				meeting.JoinUrl,
			}
			if meeting.Password != nil {
				SendTemplateMessage(meeting.User.Phone, smsqueue.SMSMeetingInvitePassword, details)
			} else {
				SendTemplateMessage(meeting.User.Phone, smsqueue.SMSMeetingInviteWithoutPassword, details)
			}
		} // 用户手机号可用性检测
	}

}

// 取消会议
func SendMeetingCancelNotification(meeting *models.Meeting) {
	// meeting 变量逃逸

	personnel, err := dao.GetMeetingPersonnelsWithUserByMeetingID((meeting).ID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, val := range personnel {
		if val.User.PhoneChecked {
			details := struct {
				Name        string
				MeetingCode string
			}{
				val.Username,
				meeting.MeetingCode,
			}

			SendTemplateMessage(meeting.User.Phone, smsqueue.SMSMeetingCancelNotification, details)

		} // 用户手机号可用性检测
	}
}
// 案件状态更新
func SendCaseInfoChanged(_case *models.Case) {

	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSInfoChanged, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSInfoChanged,struct{Title string}{_case.Title})
		//}

	}

}

// 案件状态更新
func SendCaseAccepted(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSCaseAccepted, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 案件调解结果通知。
func SendStatusChanged(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSStatusChanged, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 部门调解意见
func SendSuggestion(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSSuggestion, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 提交笔录
func SendRecord(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSRecord, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 确认状态
func SendResultConfirming(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSCResultConfirming, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 拒绝调解
func SendCaseRejected(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSCaseRejected, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 接受调解
func SendCaseConfirmed(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSCaseConfirmed, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 调解结束
func SendCaseCompleted(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSCaseCompleted, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}

// 添加质证信息
func SendAddComment(_case *models.Case) {
	summitUser, err := dao.GetUserById(_case.UserID) // 获取提交者信息

	if err != nil {
		log.Println(err)
		return
	}

	if summitUser.PhoneChecked {
		SendTemplateMessage(summitUser.Phone, smsqueue.SMSAddComment, struct{ Title string }{_case.Title})
		// 同时向当事人发送短信？
		//if len(_case.ApplicantContact) == 11 { // 手机号
		//	SendTemplateMessage(_case.ApplicantContact,smsqueue.SMSCaseAccepted,struct{Title string}{_case.Title})
		//}
	}
}





