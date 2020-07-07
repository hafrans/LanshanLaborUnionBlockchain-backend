package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func CreateMaterial(name, path string, caseId string) (*models.Material, error) {

	material := &models.Material{
		Name:   name,
		Path:   &path,
		CaseID: caseId,
	}

	result := db.Create(material)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return material, nil
	}

}

func UpdateMaterial(material *models.Material) bool {

	if material == nil {
		return false
	}

	result := db.Save(material)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteMaterial(department *models.Material) bool {
	result := db.Delete(department)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteUnscopedAllMaterialsByCaseId(caseId string) bool {

	result := db.Unscoped().Model(&models.Material{}).Where("case_id = ?", caseId).Delete(&models.Material{})

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		log.Println("delete ", result.RowsAffected, " row(s)")
		return true
	}

}

func DeleteMaterialById(id int64) bool {
	result := db.Delete(&models.Material{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetMaterialById(id int64) (*models.Material, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var department *models.Material
	result := db.FirstOrInit(department, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return department, nil
	}
}

func GetMaterialAllPaginated(pageNum, pageCount int) ([]*models.Material, int, error) {
	var materials []*models.Material
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&materials).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return materials, totalCounts, nil
	}
}

func GetMaterialAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Material, int, error) {
	var materials []*models.Material
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&materials).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return materials, totalCounts, nil
	}
}
