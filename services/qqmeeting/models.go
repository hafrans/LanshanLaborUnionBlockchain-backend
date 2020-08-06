package qqmeeting

var (
	RequestDescriptorMeetingCreate      = MeetingRequestDescriptor{"/meetings", "POST", "Create"}
	RequestDescriptorQueryMeetingByID   = MeetingRequestDescriptor{"/meetings", "POST", "ID"}
	RequestDescriptorQueryMeetingByCode = MeetingRequestDescriptor{"/meetings", "POST", "Code"}
	RequestDescriptorUserCreate         = MeetingRequestDescriptor{"/users", "POST", "Create"}
	RequestDescriptorUserDetailQuery    = MeetingRequestDescriptor{"/users/%s", "GET", "QUERY"}
	RequestDescriptorUserUpdate         = MeetingRequestDescriptor{"/users/%s", "PUT", "UPDATE"}
	RequestDescriptorUserDelete         = MeetingRequestDescriptor{"/users/%s", "DELETE", "DELETE"}
	RequestDescriptorUserList           = MeetingRequestDescriptor{"/users/list", "GET", "QUERY"}
)

// UserObj  用户对象 https://cloud.tencent.com/document/product/1095/42417
type UserObj struct {
	UserID      string `json:"userid"`              // 调用方用于标示用户的唯一 ID
	IsAnonymous bool   `json:"is_anonymous"`        // 匿名入会
	NickName    string `json:"nick_name,omitempty"` // 用户匿名字符串
}

type Settings struct {
	MuteEnableJoin   bool `json:"mute_enable_join"`            // 入会时静音
	AllowUnmuteSelf  bool `json:"allow_unmute_self"`           // 允许参会者取消静音
	MuteAll          bool `json:"mute_all,omitempty"`          // 全体静音
	HostVideo        bool `json:"host_video"`                  // 入会时主持人视频是否开启，暂时不支持。
	ParticipantVideo bool `json:"participant_video,omitempty"` // 入会时参会者视频是否开启，暂时不支持。
	EnableRecord     bool `json:"enable_record,omitempty"`     // 开启录播，暂时不支持。
	PlayIVROnLeave   bool `json:"play_ivr_on_leave,omitempty"` // 参会者离开时播放提示音。
	PlayIVROnJoin    bool `json:"play_ivr_on_join,omitempty"`  // 有新的入会者加入时播放提示音。
	LiveUrl          bool `json:"live_url,omitempty"`          // 开启直播, 暂时不支持。
}

// User
type UserInfo struct {
	Email    string `json:"email" binding:"required"`    // 邮箱地址
	Phone    string `json:"phone" binding:"required"`    // 手机号码
	Username string `json:"username" binding:"required"` // 用户昵称
	UserID   string `json:"userid" binding:"required"`   // 调用方用于表示用户的唯一ID
}

type UserDetail struct {
	Username   string `json:"username"`    // 用户昵称
	UpdateTime string `json:"update_time"` // 更新时间
	Status     string `json:"status"`      // 用户状态，1：正常；2：删除
	Email      string `json:"email"`       // 邮箱地址
	Phone      string `json:"phone"`       // 手机号码
	UserID     string `json:"userid"`      // 调用方用于标示用户的唯一 ID
	AreaCode   string `json:"area"`        // 地区编码（国内默认86）
	AvatarUrl  string `json:"avatar_url"`  // 头像
}

// 创建会议
type MeetingCreateRequest struct {
	UserID     string     `json:"userid"`             // 调用方用于标示用户的唯一 ID
	InstanceID int        `json:"instanceid"`         // 用户的终端设备类型
	Subject    string     `json:"subject"`            // 会议主题
	Hosts      []*UserObj `json:"hosts,omitempty"`    // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	Type       int        `json:"type"`               // 会议类型
	Invitees   []*UserObj `json:"invitees,omitempty"` // 会议邀请的参会者，可为空
	StartTime  string     `json:"start_time"`         // 会议开始时间戳（单位秒）。
	EndTime    string     `json:"end_time"`           // 会议结束时间戳（单位秒）。
	Password   string     `json:"password,omitempty"` // 会议密码，可不填
	Settings   *Settings  `json:"settings,omitempty"` // 会议设置
}

func (req MeetingCreateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingCreate
}

func (req MeetingCreateRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

type MeetingCreationInfo struct {
	MeetingID    string     `json:"meeting_id"`             // 会议的唯一标示
	MeetingCode  string     `json:"meeting_code"`           // 会议 App 的呼入号码
	Subject      string     `json:"subject"`                // 会议主题
	Hosts        []*UserObj `json:"hosts,omitempty"`        // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	Participants []*UserObj `json:"participants,omitempty"` // 会议邀请的参会者，可为空
	StartTime    string     `json:"start_time"`             // 会议开始时间戳（单位秒）。
	EndTime      string     `json:"end_time"`               // 会议结束时间戳（单位秒）。
	Password     *string    `json:"password,omitempty"`     // 会议密码，可不填
	JoinUrl      string     `json:"join_url"`               // 加入会议　URL（点击链接直接加入会议）
	Settings     *Settings  `json:"settings,omitempty"`     // 会议设置
}

type MeetingCreateResponse struct {
	NextPos             int                    `json:"next_pos"`
	Remaining           int                    `json:"remaining"`
	MeetingNumber       int                    `json:"meeting_number"`    // 会议数量
	MeetingCreationInfo []*MeetingCreationInfo `json:"meeting_info_list"` // 预约会议列表
}

// 创建用户
type UserCreateRequest struct {
	UserInfo
}

func (req UserCreateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserCreate
}

func (req UserCreateRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

// 获取用户信息
type UserDetailQueryRequest struct {
	UserID string `json:"userid" param:"userid"`
}

func (req UserDetailQueryRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserDetailQuery
}

func (req UserDetailQueryRequest) fillPlaceholder(args ...interface{}) string {
	return defaultFillPlaceholder(req, args...)
}

type UserDetailQueryResponse struct {
	UserDetail
}

// 更新用户信息
type UserDetailUpdateRequest struct {
	UserID   string `json:"-" param:"userid"` // 调用方用于标示用户的唯一 ID
	Email    string `json:"email"`            // 新的邮箱地址
	Username string `json:"username"`         // 新的用户昵称
}

func (req UserDetailUpdateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserUpdate
}

func (req UserDetailUpdateRequest) fillPlaceholder(args ...interface{}) string {
	return defaultFillPlaceholder(req, args...)
}

// 删除用户
type UserDeleteRequest struct {
	UserID string `json:"-" param:"userid"` // 调用方用于标示用户的唯一 ID
}

func (req UserDeleteRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserDelete
}

func (req UserDeleteRequest) fillPlaceholder(args ...interface{}) string {
	return defaultFillPlaceholder(req, args...)
}

// 获取用户列表
type UserListRequest struct {
	Page     int `json:"page" query:"page"`           // 当前页
	PageSize int `json:"page_size" query:"page_size"` // 分页大小
}

func (req UserListRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserList
}

func (req UserListRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

type UserListResponse struct {
	Pager
	Users []*UserDetail `json:"users"` //
}
