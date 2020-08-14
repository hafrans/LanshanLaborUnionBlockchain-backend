package serviceimpl

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"log"
)

func TryCreateNewMeetingAccount(mUserId int64) (string, *models.User, *qqmeeting.UserInfo, bool) {
	// 先获取用户信息

	user, err := dao.GetUserById(mUserId)

	if err != nil {
		return "", nil, nil, false
	}

	// 先查询
	userDetail, err := qqmeeting.MeetingClient.Do(qqmeeting.UserDetailQueryRequest{
		UserID: user.Phone,
	})
	detailInfo := userDetail.(qqmeeting.UserDetailQueryResponse)

	if err == nil { // 存在用户了
		return detailInfo.UserID, user, &qqmeeting.UserInfo{UserID: detailInfo.UserID, Username: detailInfo.Username, Phone: detailInfo.Phone, Email: detailInfo.Email}, true
	}

	// 拼接用户名
	phone := user.Phone
	email := user.Email
	userId := user.Phone
	username := user.UserName

	if user.Email == "" {
		email = phone + "@content.com"
	}

	if user.UserType == models.USER_TYPE_ADMIN {
		username = "管理员 " + username
	} else {
		username = user.Department.Name + " " + username
	}
	// 尝试创建
	_, err = qqmeeting.MeetingClient.Do(qqmeeting.UserCreateRequest{
		UserInfo: qqmeeting.UserInfo{
			UserID:   userId,
			Email:    email,
			Phone:    phone,
			Username: username,
		},
	})
	if err == nil { // success
		return userId, user, &qqmeeting.UserInfo{UserID: userId, Email: email, Phone: phone, Username: username}, true
	} else {
		log.Println(err)
		return "", nil, nil, false
	}
}
