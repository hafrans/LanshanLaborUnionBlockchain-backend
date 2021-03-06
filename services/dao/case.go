package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"database/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

func CreateCase(model *models.Case) (*models.Case, error) {

	if model == nil {
		return nil, errors.New("no model found")
	}

	result := db.Create(model)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return model, nil
	}

}

func GetCasePreloadedModelById(id int64) (*models.Case, error) {

	var model models.Case

	result := db.Set("gorm:auto_preload", true).First(&model, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}

}

func GetCaseNotPreloadModelById(id int64) (*models.Case, error) {

	var model models.Case
	result := db.First(&model, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}

}

func GetCasePreloadedModelByCaseID(caseId string) (*models.Case, error) {
	var model models.Case
	result := db.Set("gorm:auto_preload", true).Where("case_id = ?", caseId).First(&model)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}
}

func GetCaseNotPreloadedModelByCaseID(caseId string) (*models.Case, error) {
	var model models.Case
	result := db.Where("case_id = ?", caseId).First(&model)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return &model, nil
	}
}

func GetCasesAllPaginatedOwnByUserId(pageNum, pageSize int, userId *int64) ([]*models.Case, int, error) {

	var cases []*models.Case
	var totalCount int

	pendingDb := db.Model(&models.Case{})

	if userId != nil && *userId != 0 {
		pendingDb = pendingDb.Where("user_id = ?", userId)
	}

	result := pendingDb.Count(&totalCount).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cases)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCount, result.Error
	} else {
		return cases, 0, nil
	}
}

// 劳动者、管理员等获取自己的信息
func GetCasesAllPaginatedByCaseId(caseId *string, pageNum, pageSize int, userId *int64) ([]*models.Case, int, error) {

	var cases []*models.Case
	var totalCount int

	pendingDb := db.Model(&models.Case{})

	if caseId != nil && *caseId != "" {
		pendingDb = pendingDb.Where("case_id like ?", "%"+*caseId+"%")
	}

	if userId != nil && *userId != 0 {
		pendingDb = pendingDb.Where("user_id = ?", userId)
	}

	result := pendingDb.Count(&totalCount).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cases)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCount, result.Error
	} else {
		return cases, totalCount, nil
	}

}

// 用工单位获取自己的案件信息
func GetCasesAllPaginatedByCaseIdAndUSSC(caseId *string, pageNum, pageSize int, ussc *string) ([]*models.Case, int, error) {

	var cases []*models.Case
	var totalCount int

	pendingDb := db.Model(&models.Case{})

	if caseId != nil && *caseId != "" {
		pendingDb = pendingDb.Where("case_id like ?", "%"+*caseId+"%")
	}

	if ussc != nil && *ussc != "" {
		pendingDb = pendingDb.Where("employer_uniform_social_credit_code = ?", ussc)
	}

	result := pendingDb.Count(&totalCount).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cases)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, totalCount, result.Error
	} else {
		return cases, totalCount, nil
	}

}

func GetCasesByFormId(formId int64) ([]*models.Case, int, error) {

	var cases []*models.Case

	result := db.Model(&models.Case{}).Where("form_id = ? ", formId).Find(&cases)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows || result.Error == gorm.ErrRecordNotFound {
			log.Println(result.Error)
			return cases, 0, nil
		} else {
			log.Println(result.Error)
			return cases, 0, result.Error
		}
	} else {
		return cases, len(cases), nil
	}
}

func DeleteCaseById(id int64) bool {

	result := db.Model(&models.Case{}).Delete(&models.Case{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return true
		} else {
			return false
		}
	} else {
		return true
	}

}

func DeleteCaseByCaseId(caseId string) bool {
	result := db.Model(&models.Case{}).Where("caseId = ?", caseId).Delete(&models.Case{})
	if result.Error != nil {
		log.Println(result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func UpdateCase(model *models.Case) bool {

	result := db.Save(model)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else if result.RowsAffected == 1 {
		return true
	} else {
		log.Println(result.RowsAffected)
		return false
	}

}
