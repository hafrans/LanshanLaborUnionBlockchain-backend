package smsqueue


var SMSContentCaptcha = `您的验证码为：{{ .Code }}，有效期五分钟。如非本人操作，请忽略此短信。`
var SMSMeetingInviteWithoutPassword = `{{ .Name }}邀请您参加在线会议，请使用腾讯会议参会。会议ID:{{ .MeetingCode }}，或直接点击链接入会:{{.MeetingLink}}`
var SMSMeetingCancelNotification = `{{ .Name }} 因故取消会议ID为 {{ .MeetingCode }} 的在线会议，请知悉。`
var SMSMeetingInvitePassword = `{{ .Name }}邀请您参加在线会议，请使用腾讯会议参会。会议ID:{{ .MeetingCode }}，密码:{{ .MeetingPass }}，或直接点击链接入会:{{.MeetingLink}}`
var SMSInfoChanged = `您的案件【{{ .Title }}】信息有更新！您可登录平台查看更新信息`
var SMSCaseAccepted = `您好，相关工作人员已经接受了您的案件【{{ .Title }}】的调解请求，请登录平台查看案件处理进度。`
var SMSCResultConfirming = `您好，您的案件【{{ .Title }}】的最终调节意见已提交，请登录平台查看案件，并选择接受或拒绝调解。`

var SMSStatusChanged = `您好，您的案件【{{ .Title }}】，相关工作人员已给出调解结果，请登录平台查看。`
var SMSSuggestion = `您好，您的案件【{{ .Title }}】，相关工作人员已给出部门建议，请登录平台查看。`
var SMSRecord = `您好，您的案件【{{ .Title }}】，相关工作人员提交了新的笔录信息，请登录平台查看。`
var SMSCaseRejected = `您好，您已经拒绝了案件【{{ .Title }}】的调解结果。`
var SMSCaseConfirmed = `您好，您已经接受了案件【{{ .Title }}】的调解结果。`
var SMSCaseCompleted = `您好，您的案件【{{ .Title }}】调解结束。`

var SMSAddComment = `您好，您的案件【{{ .Title }}】，一方人员添加了质证信息，请登录平台查看。`



