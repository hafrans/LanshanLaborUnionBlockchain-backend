package smsqueue

var SMSContentCaptcha = `【岚山区总工会】您的验证码为：{{ .Code }}，有效期五分钟。如非本人操作，请忽略此短信。`
var SMSMeetingInvite = `【岚山区总工会】{{ .Name }}邀请您参加在线会议，请使用腾讯会议参会。会议ID:{{ .MeetingCode }}，或直接点击链接入会:{{.MeetingLink}}`
var SMSInfoChanged = `【岚山区总工会】您的案件信息有更新！您可登录平台查看更新信息`
var SMSCaseAccepted = `【岚山区总工会】您好，相关工作人员已经接受了您的案件请求，请登录平台查看案件处理进度。`
var SMSStatusChanged = `【岚山区总工会】您好，对于您的案件，相关工作人员已给出调解结果，请登录平台查看，并决定是否接受调解结果。`
