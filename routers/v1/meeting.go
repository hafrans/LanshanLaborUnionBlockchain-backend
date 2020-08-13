package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateMeetingAccount
// @Summary 创建会议专用账户
// @Description 创建会议专用账户，只有管理员和部门人员可以创建
// @Tags meeting
// @Produce json
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common "未验证"
// @Failure 500 {object} vo.Common "服务器错误"
// @Router /api/v1/meeting/account/create [get]
func CreateMeetingAccount(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您无法创建会议专用账户"))
	} else {
		// 先获取用户信息

		user, err := dao.GetUserById(claims.Id)

		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "用户异常"))
			return
		}

		// 先查询
		_, err = qqmeeting.MeetingClient.Do(qqmeeting.UserDetailQueryRequest{
			UserID: user.Phone,
		})

		if err == nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您已创建会议专用账户，不可重复创建"))
			return
		}

		// 拼接用户名
		phone := user.Phone
		email := user.Email
		userId := user.Phone
		username := user.UserName

		if user.Email == "" {
			email = phone + "@content.com"
		}

		if claims.UserType == models.USER_TYPE_ADMIN {
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
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "会议专用用户创建成功"))
		} else {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "会议专用用户创建失败"+err.Error()))
		}
	}
}

// DeleteMeetingAccount
// @Summary 删除会议专用账户
// @Description 删除会议专用账户，只有管理员和部门人员可以操作
// @Tags meeting
// @Produce json
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common "未验证"
// @Failure 500 {object} vo.Common "服务器错误"
// @Router /api/v1/meeting/account/delete [get]
func DeleteAccount(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您无法删除会议专用账户"))
	} else {
		// 先获取用户信息
		user, err := dao.GetUserById(claims.Id)

		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "用户异常"))
			return
		}

		// 先查询
		_, err = qqmeeting.MeetingClient.Do(qqmeeting.UserDetailQueryRequest{
			UserID: user.Phone,
		})

		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您不存在可用的专用账户信息 "+err.Error()))
			return
		}

		_, err = qqmeeting.MeetingClient.Do(qqmeeting.UserDeleteRequest{
			UserID: user.Phone,
		})
		if err == nil { // success
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "会议专用用户删除成功"))
		} else {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "会议专用用户删除失败"+err.Error()))
		}

	}
}

// CreateMeeting
// @Summary 创建会议
// @Description 创建会议，只有管理员和部门人员可以创建
// @Tags meeting
// @Produce json
// @Accept json
// @Param meeting body vo.MeetingForm true "新建会议表单"
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common "未验证"
// @Failure 500 {object} vo.Common "服务器错误"
// @Router /api/v1/meeting/account/create [get]
func CreateMeeting(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_DEPARTMENT && claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有创建会议的权限"))
	} else {



	}

}

// ListMeeting
// @Summary 列出相关会议
// @Description 列出相关会议，管理员可以看到全部的，部门管理员只可以看到自己的
// @Tags meeting
// @Produce json
// @Param page query number true "页码"
// @Param pageSize query number true "页大小"
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common "未验证"
// @Failure 500 {object} vo.Common "服务器错误"
// @Router /api/v1/meeting/create [get]
func GetMyMeetingList(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageCount, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageCount < 1 {
		pageCount = 10
	}

	var list []*models.Meeting
	var totalCount int
	if claims.UserType == models.USER_TYPE_ADMIN { // 看到所有人的
		list, totalCount, err = dao.GetMeetingAllWithConditionPaginated(nil, nil, true, pageNum, pageCount)
	} else { // 其他人只看自己相关的
		list, totalCount, err = dao.GetMeetingAllRelatedPaginated(claims.Id, true, pageNum, pageCount)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "会议列表为空"),
				Data: gin.H{
					"list":  []interface{}{},
					"total": totalCount,
					"size":  pageCount,
					"page":  pageNum,
				},
			})
			return
		}
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "会议列表获取失败"))
		return
	} else {
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
			Data: gin.H{
				"list":        utils.PopulateMeetingListFromModelTOVO(list),
				"total_count": totalCount,
				"page_count":  pageCount,
				"page_num":    pageNum,
			},
		})
	}

}
