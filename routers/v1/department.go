package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// Get Department List
// @Summary 获取所有相关单位信息
// @Description 获取所有相关单位信息
// @Tags department
// @Produce json
// @Param page query number true "页码"
// @Param pageSize query number true "页大小"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/department/list [get]
func GetDepartmentList(ctx *gin.Context) {
	//claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageCount, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageCount < 1 {
		pageCount = 10
	}

	list, total, err := dao.GetDepartmentAllPaginated(pageNum, pageCount)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "列表为空"),
				Data: gin.H{
					"list":  []interface{}{},
					"total": total,
					"size":  pageCount,
					"page":  pageNum,
				},
			})
			return
		}
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "列表获取失败"))
		return
	} else {
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
			Data: gin.H{
				"list":        list,
				"total_count": total,
				"page_count":  pageCount,
				"page_num":    pageNum,
			},
		})
	}

}

// Create Department 创建department
// @Summary 创建单位
// @Description 创建单位
// @Tags department
// @Accept json
// @Produce json
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/department/create [post]
func CreateDepartment(ctx *gin.Context) {
	// 权限检查
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有创建单位的权限"))
		return
	}
	var form *vo.DepartmentVO
	if err := ctx.ShouldBindJSON(&form); err != nil {
		log.Println(err.Error())
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	} else {

	}

}
