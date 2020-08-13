package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
	utils2 "RizhaoLanshanLabourUnion/utils"
)

func PopulateMeetingFromModelToVO(model *models.Meeting) *vo.MeetingVO {

	host := make([]string, 0)

	for _, v := range model.Personnel {
		if v.MeetingRole == models.MeetingRoleHost {
			host = append(host, v.Username)
		}
	}

	vo := &vo.MeetingVO{
		ID:          model.ID,
		CaseID:      model.CaseID,
		Creator:     model.User.UserName,
		EndTime:     utils2.GetTime(model.EndTime),
		StartTime:   utils2.GetTime(model.StartTime),
		Host:        host,
		JoinUrl:     model.JoinUrl,
		Subject:     model.Subject,
		MeetingCode: model.MeetingCode,
	}

	return vo
}

func PopulateMeetingListFromModelTOVO(model []*models.Meeting) []*vo.MeetingVO {
	list := make([]*vo.MeetingVO, 0, len(model))
	for _, v := range model {
		list = append(list, PopulateMeetingFromModelToVO(v))
	}
	return list
}
