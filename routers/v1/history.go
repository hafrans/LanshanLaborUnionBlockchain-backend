package v1

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Get Paginated History By CaseID
// @Summary 通过case 案件号码（37开头）获取其历史信息，可支持分页
// @Description 通过case 案件号码（37开头）获取其历史信息，可支持分页
// @Tags case,blockchain
// @Produce json
// @Param page query number true "页码"
// @Param pageSize query number true "页大小"
// @Param caseId path string true "案件号(37开头)"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/blockchain/history/case/:caseId [get]
func GetHistoryByCaseID(ctx *gin.Context) {

	//claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageCount, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "25"))
	if err != nil || pageCount < 1 {
		pageCount = 25
	}

	caseId := ctx.Param("caseId")

	var list []*models.HistoryV1
	var total int

	// TODO 权限控制,暂时不用
	list, total, err = dao.GetHistoryAllPaginatedByCaseId(pageNum, pageCount, &caseId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericFailed, "无相关案件信息存档"),
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
				"list":        utils.PopulateHistoryV1ListFromModelToVO(list),
				"total_count": total,
				"page_count":  pageCount,
				"page_num":    pageNum,
			},
		})

	}

}
