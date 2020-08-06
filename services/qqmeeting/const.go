package qqmeeting

// Entrance
const (
	ApiHost = "https://api.meeting.qq.com/v1"
)

// Meeting Type
const (
	MeetingTypeBookingMeeting = iota
	MeetingTypeQuickMeeting
)

// Instance Type
const (
	InstancePC = iota + 1
	InstanceMac
	InstanceAndroid
	InstanceIOS
	InstanceWeb
	InstanceIPad
	InstanceAndroidPad
	InstanceMicroProgram
)

// 错误列表
const (
	ErrTinyIdOrMeetingId       = 9002
	ErrMeetingNotExists        = 9003
	ErrMeetingCreateExceed     = 9008
	ErrMeetingQueryExceed      = 9061
	ErrApiCallUnknownType      = 10000
	ErrApiCallBadParameter     = 10001
	ErrAppVersionForbidden     = 10005
	ErrUserAlreadyExists       = 20002
	ErrUserUnavailable         = 20003
	ErrInvalidPhone            = 40000
	ErrInvalidEmail            = 41001
	ErrEmailUsed               = 41002
	ErrPhoneUsed               = 41003
	ErrCorpID                  = 50000
	ErrCorpUnavailable         = 50001
	ErrXTcTimestamp            = 190300
	ErrRequestReplay           = 190301
	ErrUnauthenticatedSecret   = 190303
	ErrCallMinuteExceed        = 190310
	ErrCallDayExceed           = 190311
	ErrCallParticularDayExceed = 190312
	ErrApiRequiredInfoNotFound = 200001
	ErrApiReplay               = 200002
	ErrApiBadSignature         = 200003
	ErrApiNotSupportRequest    = 200004
	ErrJsonSchemeInvalid       = 200005
	ErrApiBadRequestParameter  = 200006
)
