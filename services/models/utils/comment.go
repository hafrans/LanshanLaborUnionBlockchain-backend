package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
)

func PopulateCommentFromModelToVO(comment *models.Comment) *vo.Comment{

	var vo *vo.Comment = &vo.Comment{
		ID: comment.ID,
		Content: comment.Content,
		CaseID: comment.CaseID,
		SubmitterPhone: comment.User.Phone,
		SubmitterType: comment.User.UserType,
	}

	if comment.User.UserType == models.USER_TYPE_LABOR {// 劳动者
		vo.Submitter = comment.User.UserProfile.ApplicantName
	}else if comment.User.UserType == models.USER_TYPE_EMPLOYER { // 用人单位
		vo.Submitter = comment.User.UserProfile.EmployerName
	}else{
		vo.Submitter = comment.User.UserName // 其他人员(管理员或者部门工作人员)
	}

	return vo
}
