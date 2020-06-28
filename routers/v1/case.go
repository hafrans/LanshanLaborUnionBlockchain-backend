package v1

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"github.com/gin-gonic/gin"
	"log"
)



// Get All Categories
// @Summary 获得所有类型
// @Description 获取所有类型
// @Tags category,case
// @Produce json
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common
// @Router /api/v1/category [get]
func GetAllCategories(ctx *gin.Context){

	list, _, err := dao.GetCategoryAllPaginated(0,50)
	if err != nil {
		log.Println(err)
		ctx.JSON(respcode.HttpOK,vo.GenerateCommonResponseHead(respcode.GenericFailed, "出现错误："+err.Error()))
		return
	}

	ctx.JSON(respcode.HttpOK, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
		Data: list,
	})
}
