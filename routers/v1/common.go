package v1

import (
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"github.com/gin-gonic/gin"
)



// Api Index
// @Summary ApiIndex
// @Description 测试在登录情况下是否可以访问
// @Tags test,index
// @Accept json
// @Produce json
// @Success 200 {object} vo.Common
// @Failure 401 {object} vo.Common
// @Router /api/v1/ [get]
func ApiIndexHandler(ctx *gin.Context) {
	ctx.JSON(200, vo.Common{
		Status: 0,
		Message: "success",
		Timestamp: utils.NowTime(),
	})
}
