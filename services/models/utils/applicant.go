package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
)

func PopulateApplicantFromVOToModel(applicant *vo.Applicant) *models.Applicant {
	model := models.Applicant{
		ApplicantName:           applicant.Name,
		ApplicantIdentityNumber: applicant.IdentityNumber,
		ApplicantBirthday:       applicant.Birthday,
		ApplicantContact:        applicant.Contact,
		ApplicantAddress:        applicant.Address,
		ApplicantNationality:    applicant.Nationality,
	}

	return &model
}
