package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/qqmeeting"
	"log"
)

func CreateMeetingPersonnelByModel(model *models.MeetingPersonnel) (*models.MeetingPersonnel, error) {

	result := db.Create(model)
	if result.Error != nil {
		log.Print(result.Error)
		return nil, result.Error
	} else {
		return model, nil
	}
}

func CreateMeetingPersonnel(meetingId int64, user *models.User, meetingUserResp qqmeeting.UserDetailQueryResponse, meetingRole int) (*models.MeetingPersonnel, error) {

	model := &models.MeetingPersonnel{
		MeetingID:   meetingId,
		MeetingRole: meetingRole,
		Userid:      meetingUserResp.UserID,
		UserID:      user.ID,
		Username:    meetingUserResp.Username,
	}

	result := db.Create(model)
	if result.Error != nil {
		log.Print(result.Error)
		return nil, result.Error
	} else {
		return model, nil
	}
}

func DeleteMeetingPersonnel(id int64) bool {

	result := db.Delete(&models.MeetingPersonnel{}, id)

	if result.Error != nil {
		log.Print(result.Error)
		return false
	} else {
		return true
	}
}

func GetMeetingPersonnelById(id int64) (*models.MeetingPersonnel, error) {

	var model models.MeetingPersonnel

	result := db.Set("gorm:auto_preload", true).First(&model, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}
}

func UpdateMeetingPersonnel(model *models.MeetingPersonnel) bool {

	result := db.Save(model)

	if result.Error != nil {
		log.Print(result.Error)
		return false
	} else {
		return true
	}
}


