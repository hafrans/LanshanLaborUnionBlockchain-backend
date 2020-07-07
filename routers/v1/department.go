package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

// Create Department
// @Summary 创建department
// @Description 创建单位，只有管理员可以使用
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
		model, err := dao.CreateDepartment(form.Name, form.Description, form.Service, form.Contact)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		} else {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils.PopulateDepartmentFromModelToVO(model),
			})
		}
	}
}

// Delete Department By ID
// @Summary 通过 ID 删除 department
// @Description 删除单一单位，只有管理员可以操作
// @Tags department
// @Produce json
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/department/delete/:id [get]
func DeleteDepartmentById(ctx *gin.Context) {
	// 权限检查
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	if claims.UserType != models.USER_TYPE_ADMIN {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有删除该单位的权限"))
		return
	}

	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
	} else {
		if dao.DeleteDepartmentById(int64(id)) {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "删除成功"))
		} else {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "删除失败"))
		}
	}
}

// Get One Department By Id
// @Summary 获取指定单位的信息
// @Description 获取指定单位的信息，所有人都可以使用
// @Tags department
// @Produce json
// @Param id path number true "表单id"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/department/id/:id [get]
func GetOneDepartmentById(ctx *gin.Context) {
	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
	} else {
		model, err := dao.GetDepartmentById(int64(id))
		if err != nil {
			if err == sql.ErrNoRows || err == gorm.ErrRecordNotFound {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "没有该单位"))
			}
		} else {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils.PopulateDepartmentFromModelToVO(model),
			})
		}
	}
}
