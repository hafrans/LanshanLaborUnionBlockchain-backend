package serviceimpl

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	"log"
)

func GetTwoParticipantsOfCase(caseId string) ([]*models.User, error){
	var list []*models.User
	tablePrefix := utils.DatabaseSettings.TablePrefix
	result := dao.GetExternalDB().
		Model(&models.User{}).
		Joins("LEFT JOIN "+tablePrefix+"user_profile on "+tablePrefix+"user.id = "+tablePrefix+"user_profile.user_id").
		Where("ls_user.id = (select user_id from ls_case where case_id = ? ) and ls_user.user_type = 2 "+" or "+
			  "ls_user_profile.employer_uniform_social_credit_code = (select ls_case.employer_uniform_social_credit_code from ls_case where case_id = ? ) and ls_user.user_type = 3", caseId,caseId).
		Find(&list)


	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}else {
		return list, nil
	}
}