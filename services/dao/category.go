package dao

import (
	"RizhaoLanshanLabourUnion/services/models"
	"errors"
	"log"
)

func InitCategory() {

	categories := make(map[string]string)

	categories["劳动报酬"] = "因用人单位拖欠职工的工资、加班费、奖金、津贴等而发生的争议类型"
	categories["社会保险"] = "因用人单位未给职工依法缴纳社会保险而发生的争议类型，也包含用人单位未为职工依法缴纳社会保险致使职工不能享受相关社会保险待遇，造成损失的争议类型"
	categories["经济补偿金"] = "用人单位与劳动者协议解除劳动关系、劳动者依法单方解除劳动关系等法定需要支付经济补偿金的情形"
	categories["双倍工资"] = "因用人单位未依法与劳动者签订劳动合同须支付双倍工资的争议类型"
	categories["工伤"] = "劳动者履行职务过程中发生工伤，因工伤赔偿待遇引起的争议类型"
	categories["赔偿金"] = "因用人单位违法解除劳动关系引起的赔偿争议"
	categories["劳动关系"] = "职工在履行职务过程中严重违反公司管理制度、存在营私舞弊等法定情形用人单位依法解除劳动关系引起的争议类型或劳动者需要与用人单位确认劳动关系的争议类型"
	categories["其他"] = "其他正义类型"

	for k, v := range categories {
		CreateCategory(k, v)
	}

}

func CreateCategory(name, description string) (*models.Category, error) {

	category := &models.Category{
		Name:        name,
		Description: description,
	}

	result := db.Create(category)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return category, nil
	}

}

func UpdateCategory(category *models.Category) bool {

	if category == nil {
		return false
	}

	result := db.Save(category)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}

}

func DeleteCategory(category *models.Category) bool {
	result := db.Delete(category)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func DeleteCategoryById(id int64) bool {
	result := db.Delete(&models.Category{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	} else {
		return true
	}
}

func GetCategoryById(id int64) (*models.Category, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	var category *models.Category
	result := db.FirstOrInit(category, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	} else {
		return category, nil
	}
}

func GetCategoryAllPaginated(pageNum, pageCount int) ([]*models.Category, int, error) {
	var categories []*models.Category
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&categories).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return categories, totalCounts, nil
	}
}

func GetCategoryAllLikedNamePaginated(name string, pageNum, pageCount int) ([]*models.Category, int, error) {
	var categories []*models.Category
	if pageNum < 0 {
		pageNum = 0
	}
	totalCounts := 0
	result := db.Where("name like ? ", "%"+name+"%").Limit(pageCount).Offset(pageCount * pageNum).Order("id desc").Find(&categories).Count(&totalCounts)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, 0, result.Error
	} else {
		return categories, totalCounts, nil
	}
}
