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
		CategoryID: form.CategoryID,
		FormID:     form.FormID,
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
		UserID: userId,
		CaseID: utils.GenerateCaseId(areaCode),
	}
	return &cases
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

func SimplyCaseListItem(list []*models.Case) []*vo.SimplifiedCaseListItem {

	length := len(list)
	arr := make([]*vo.SimplifiedCaseListItem, 0, length)

	for _, v := range list {
		tmp := new(vo.SimplifiedCaseListItem)
		tmp.Owner = v.UserID
		tmp.ID = v.ID
		tmp.CreatedAt = utils.GetTime(v.CreatedAt)
		tmp.UpdateAt = utils.GetTime(v.UpdatedAt)
		tmp.Title = v.Title
		tmp.Status = v.Status
		tmp.CaseID = v.CaseID
		tmp.ApplicantName = v.ApplicantName
		tmp.RespondentName = v.EmployerName

		arr = append(arr, tmp)
	}

	return arr
}

func PopulateCaseFullModelToFullForm(model *models.Case) *vo.CaseFullResultForm {

	form := &vo.CaseFullResultForm{
		ID:        model.ID,
		CaseID:    model.CaseID,
		Status:    model.Status,
		Title:     model.Title,
		UpdateAt:  utils.GetTime(model.UpdatedAt),
		CreatedAt: utils.GetTime(model.CreatedAt),
		Owner:     model.UserID,
		Content:   model.Content,
		Form:      PopulateLaborArbitrationModelToVO(model.Form),
		Applicant: vo.Applicant{
			Name:           model.ApplicantName,
			Contact:        model.ApplicantContact,
			Address:        model.ApplicantAddress,
			Birthday:       model.ApplicantBirthday,
			IdentityNumber: model.ApplicantIdentityNumber,
			Nationality:    model.ApplicantNationality,
		},
		Respondent: vo.Employer{
			Name:                    model.EmployerName,
			Address:                 model.EmployerAddress,
			Contact:                 model.EmployerContact,
			LegalRepresentative:     model.EmployerLegalRepresentative,
			UniformSocialCreditCode: model.EmployerUniformSocialCreditCode,
		},
		Category: vo.Category{
			Name:        model.Category.Name,
			Description: model.Category.Description,
			ID:          model.CategoryID,
		},
		Materials:   PopulateMaterialListFromModelToVO(model.Materials),
		Records:     PopulateRecordListFromModelToVO(model.Records),
		Suggestions: PopulateSuggestionListFromModelToVO(model.Suggestions),
	}

	return form

}
