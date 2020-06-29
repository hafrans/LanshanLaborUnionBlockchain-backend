package utils

import (
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
)

func PopulateCaseBasicFromFormToModel(form *vo.CaseFirstSubmitForm, userId int64, areaCode string) *models.Case {

	cases := models.Case{
		Title:      form.Title,
		Content:    form.Content,
		Status:     models.StatusSubmitted,
		CategoryId: form.CategoryId,
		FormId:     form.FormId,
		Applicant: models.Applicant{
			ApplicantAddress:        form.Applicant.Address,
			ApplicantBirthday:       form.Applicant.Birthday,
			ApplicantContact:        form.Applicant.Contact,
			ApplicantName:           form.Applicant.Name,
			ApplicantIdentityNumber: form.Applicant.IdentityNumber,
			ApplicantNationality:    form.Applicant.Nationality,
		},
		Employer: models.Employer{
			EmployerAddress:                 form.Respondent.Address,
			EmployerContact:                 form.Respondent.Contact,
			EmployerLegalRepresentative:     form.Respondent.LegalRepresentative,
			EmployerUniformSocialCreditCode: form.Respondent.UniformSocialCreditCode,
			EmployerName:                    form.Respondent.Name,
		},
		UserId: userId,
		CaseId: utils.GenerateCaseId(areaCode),
	}
	return &cases
}

func PopulateMaterialFromFormToModel(material *vo.Material) *models.Material {

	model := models.Material{
		Name: material.Name,
		Path: material.Path,
	}
	return &model
}

func PopulateMaterialListFromFormToModel(materials []*vo.Material) []*models.Material {

	length := len(materials)
	result := make([]*models.Material, 0, length)

	for _, v := range materials {
		result = append(result, PopulateMaterialFromFormToModel(v))
	}

	return result

}

func PopulateApplicantFromFormToModel(applicant *vo.Applicant) *models.Applicant {
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

func PopulateEmployerFromFormToModel(employer *vo.Employer) *models.Employer {
	model := models.Employer{
		EmployerAddress:                 employer.Address,
		EmployerContact:                 employer.Contact,
		EmployerName:                    employer.Name,
		EmployerUniformSocialCreditCode: employer.UniformSocialCreditCode,
		EmployerLegalRepresentative:     employer.LegalRepresentative,
	}
	return &model
}
