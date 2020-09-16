package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/blockchain"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	utils2 "RizhaoLanshanLabourUnion/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

// Get All Categories
// @Summary 获得所有类型
// @Description 获取所有类型
// @Tags category,case
// @Produce json
// @Success 200 {object} vo.CommonData
// @Failure 401 {object} vo.Common
// @Router /api/v1/category [get]
func GetAllCategories(ctx *gin.Context) {

	list, _, err := dao.GetCategoryAllPaginated(0, 50)
	if err != nil {
		log.Println(err)
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "出现错误："+err.Error()))
		return
	}

	ctx.JSON(respcode.HttpOK, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
		Data:   list,
	})
}

// Create New Case By Applicant
// @Summary 创建新调解案件
// @Description 由申请人填写创建新调解案件
// @Tags case
// @Accept json
// @Produce json
// @Param case body vo.CaseFirstSubmitForm true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/create [post]
func CreateNewCaseByApplicant(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	var form vo.CaseFirstSubmitForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		// 表单绑定失败
		log.Println(err.Error())
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	}

	newCase := utils.PopulateCaseBasicFromFormToModel(&form, claims.Id, "371100")

	// check form exists or else

	laborForm, err := dao.GetLaborArbitrationById(newCase.FormID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "表单不存在"))
			return
		} else {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		}
	}

	// 是否这个表单属于自己

	if laborForm.Owner != claims.Id { // 不是自己的表单
		log.Println("有人使用他人表单" + strconv.Itoa(int(laborForm.Owner)) + "," + strconv.Itoa(int(claims.Id)))
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "表单不存在"))
		return
	}

	model, err := dao.CreateCase(newCase)

	if err != nil {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	} else {
		// 注入material
		for _, v := range form.Materials {
			if v.Path != nil {
				if _, mErr := dao.CreateMaterial(v.Name, *v.Path, model.CaseID); mErr != nil {
					ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, mErr.Error()))
					return
				}
			}
		}

		result, err := dao.GetCasePreloadedModelById(model.ID)
		if err != nil {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		} else {
			// 记录
			blockchain.CreateHistoryByCase("创建调解案件", result, claims.Id)

			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils.PopulateCaseFullModelToFullForm(result),
			})
			return
		}

		// 提交成功 37110020200630134604159349596481
	}

}

// Update case by id
// @Summary 修改调解案件
// @Description 由申请人填写修改调解案件，申请人、管理员、部门人员可以修改，只可以修改案件信息，不可以修改案件状态
// @Tags case
// @Accept json
// @Produce json
// @Param id path integer true  "案件模型id，注意：不是案件号"
// @Param case body vo.CaseFirstSubmitForm true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/update/:id [post]
func UpdateCaseByApplicant(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)
	var form vo.CaseFirstSubmitForm
	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {

		log.Println(err.Error())
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
		return

	} else {

		if err := ctx.ShouldBindJSON(&form); err != nil {
			log.Println(err)
			ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		}

		newCase := utils.PopulateCaseBasicFromFormToModel(&form, claims.Id, "371100")
		currentCase, err := dao.GetCaseNotPreloadModelById(int64(id))

		if err != nil {
			log.Println(err)
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "case not found"+err.Error()))
			return
		}

		newCase.Model = currentCase.Model
		newCase.UserID = currentCase.UserID
		newCase.CaseID = currentCase.CaseID
		//newCase.Status = currentCase.Status // 不能修改案件状态

		// check labor arbitration form exists or else
		laborForm, err := dao.GetLaborArbitrationById(newCase.FormID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "表单不存在"))
				return
			} else {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
				return
			}
		}

		// 是否这个表单属于自己
		if laborForm.Owner != claims.Id && claims.UserType != models.USER_TYPE_ADMIN && claims.UserType != models.USER_TYPE_DEPARTMENT { // 不是自己的表单，管理员和部门人员无视
			log.Println("有人使用他人表单" + strconv.Itoa(int(laborForm.Owner)) + "," + strconv.Itoa(int(claims.Id)))
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "表单不存在"))
			return
		}

		// 检查可不可以修改该表单
		if newCase.UserID != claims.Id && claims.UserType != models.USER_TYPE_ADMIN && claims.UserType != models.USER_TYPE_DEPARTMENT {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您没有权限修改案件"))
			return
		}

		if !dao.UpdateCase(newCase) {
			ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "案件修改失败"))
			return
		} else {

			// 删除所有Material
			dao.DeleteUnscopedAllMaterialsByCaseId(newCase.CaseID)
			// 注入material
			for _, v := range form.Materials {
				if v.Path != nil {
					if _, mErr := dao.CreateMaterial(v.Name, *v.Path, newCase.CaseID); mErr != nil {
						ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, mErr.Error()))
						return
					}
				}
			}

			result, err := dao.GetCasePreloadedModelById(newCase.ID)
			if err != nil {
				ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
				return
			} else {
				// 记录
				blockchain.CreateHistoryByCase("修改调解案件", result, claims.Id)

				ctx.JSON(respcode.HttpOK, vo.CommonData{
					Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
					Data:   utils.PopulateCaseFullModelToFullForm(result),
				})
				return
			}

			// 修改成功 37110020200630134604159349596481
		}
	}

}

// Get Case First Submit Form Template
// @Summary 获取申请调解案件的上传模板
// @Description 获取申请调解案件的上传模板，测试用
// @Tags case,test
// @Produce json
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/test/case/template [get]
func GetCaseFirstSubmitFormTemplate(ctx *gin.Context) {

	s := new(vo.CaseFirstSubmitForm)
	marr := make([]*vo.Material, 0, 2)
	path := "/test/1.png"
	marr = append(marr, &vo.Material{Path: &path, Name: "欠条"})
	marr = append(marr, &vo.Material{Path: &path, Name: "老合同"})
	s.Materials = marr
	s.Applicant = vo.Applicant{Name: "张三", Contact: "10086", Address: "三体星", Nationality: "三体人", IdentityNumber: "1234567890123456789", Birthday: utils2.NowDateDay()}
	s.Respondent = vo.Employer{Name: "第三红岸基地", Address: "地球", Contact: "10010", UniformSocialCreditCode: "1234567889456123", LegalRepresentative: "李四"}
	s.Content = "一场简单的劳动纠纷"
	s.Title = "劳动纠纷2001"
	s.FormID = 1

	ctx.JSON(200, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(0, "success"),
		Data:   s,
	})

}

// Get Case By ID
// @Summary 通过ID（主键）获取case
// @Description 获取单一Case， 例子：9
// @Tags case
// @Produce json
// @Param id path integer true "案件id,不是案件号码"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/id/:id [get]
func GetCaseById(ctx *gin.Context) {

	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
		return
	} else {
		cases, cErr := dao.GetCasePreloadedModelById(int64(id))
		if cErr != nil {
			if cErr == sql.ErrNoRows {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "记录不存在"))
				return
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, cErr.Error()))
				return
			}
		}
		ctx.JSON(200, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
			Data:   utils.PopulateCaseFullModelToFullForm(cases),
		})
	}

}

// Get Case By Case ID
// @Summary 通过Case ID（调解申请号）获取case
// @Description 获取单一Case，通过CaseID 例子：3711002020063019254015935163407436142
// @Tags case
// @Produce json
// @Param caseId path string true "不是案件号码,不是案件id"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/caseId/:caseId [get]
func GetCaseByCaseID(ctx *gin.Context) {

	caseId := ctx.Param("caseId")

	cases, err := dao.GetCasePreloadedModelByCaseID(caseId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "记录不存在"))
			return
		} else {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		}
	}

	ctx.JSON(200, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
		Data:   utils.PopulateCaseFullModelToFullForm(cases),
	})

}

// Get Case list
// @Summary 获取所有的案件
// @Description 获取所有的案件，非管理员只能看到自己的，管理员能看到全部人的
// @Tags labor
// @Accept json
// @Produce json
// @Param page query number true "页码"
// @Param pageSize query number true "页大小"
// @Param search query string true "案件号（case id）模糊查询"
// @Success 200 {object} vo.CommonData "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Router /api/v1/case/ [get]
func GetCaseList(ctx *gin.Context) {
	var err error

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageCount, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageCount < 1 {
		pageCount = 10
	}
	caseId := ctx.DefaultQuery("search", "")

	var model []*models.Case
	var totalCount int

	if claims.UserType == models.USER_TYPE_LABOR { // 如果是普通用户
		model, totalCount, err = dao.GetCasesAllPaginatedByCaseId(&caseId, pageNum, pageCount, &claims.Id)
	} else if claims.UserType == models.USER_TYPE_EMPLOYER { // 单位用户
		user, _ := dao.GetUserById(claims.Id)
		model, totalCount, err = dao.GetCasesAllPaginatedByCaseIdAndUSSC(&caseId,
			pageNum,
			pageCount,
			&user.UserProfile.EmployerUniformSocialCreditCode)
	} else {
		model, totalCount, err = dao.GetCasesAllPaginatedByCaseId(&caseId, pageNum, pageCount, nil)
	}

	if err == nil {
		list := utils.SimplyCaseListItem(model)

		ctx.JSON(respcode.HttpOK, vo.CommonData{
			Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
			Data: gin.H{
				"list":        list,
				"total_count": totalCount,
				"page_count":  pageCount,
				"page_num":    pageNum,
			},
		})
	} else {
		if err == sql.ErrNoRows {
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "数据为空"),
				Data: gin.H{
					"list":        []interface{}{},
					"total_count": totalCount,
					"page_count":  pageCount,
					"page_num":    pageNum,
				},
			})
		} else {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		}
	}

}

// Delete Case By ID
// @Summary 通过 ID（调解案件ID，不是case_id）删除case
// @Description 删除单一Case，一般用户只可以删除自己的，特殊权限者可以删除任何人的，注意：如果案件正在处理中，则无法删除
// @Tags case
// @Produce json
// @Param id path integer true "案件id,不是案件号码"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/delete/:id [get]
func DeleteCaseById(ctx *gin.Context) {

	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
	} else {
		cases, cErr := dao.GetCaseNotPreloadModelById(int64(id))
		if cErr != nil {
			if cErr == sql.ErrNoRows {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "案件不存在"))
				return
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, cErr.Error()))
				return
			}
		}

		// 如果案件在处理，则无法删除
		if cases.Status != models.StatusSubmitted && cases.Status != models.StatusCompleted {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "该案件正在处理中，无法删除"))
			return
		}

		// 检查是否是管理员，是管理员就能删除其他人的
		if cases.UserID == claims.Id || claims.UserType == models.USER_TYPE_ADMIN { // 如果是个人或者 管理员
			// 执行删除
			if dao.DeleteCaseById(cases.ID) {

				// 记录
				blockchain.CreateHistoryByCase("删除调解案件", cases, claims.Id)

				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "删除成功"))
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "删除失败"))
			}

		} else {
			ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "删除失败,您没有删除该案件的权限"))
		}

	}

}

// Update Status case by id
// @Summary 修改调解案件的状态
// @Description 申请人只可以确认、拒绝状态；管理人员可以设置任何状态 StatusSubmitted= 0 已提交；StatusPending = 1 正在处理；StatusResultConfirming = 2 当事人等待确认调解结果；StatusRefused =3 拒绝调解；StatusConfirmed=4 确认调解；StatusCompleted=5；结束调解               // 调解结束
// @Tags case
// @Accept json
// @Produce json
// @Param id path integer true  "案件模型id，注意：不是案件号"
// @Param case body vo.CaseStatusChangeForm true "提交表单"
// @Success 200 {object} vo.CommonData "成功"
// @Failure 422 {object} vo.Common "绑定失败"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/case/status_change/:id [post]
func ChangeCaseStatusById(ctx *gin.Context) {
	claims := jwtmodel.ExtractUserClaimsFromGinContext(ctx)

	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
	} else {
		cases, cErr := dao.GetCaseNotPreloadModelById(int64(id))
		if cErr != nil {
			if cErr == sql.ErrNoRows {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "案件不存在"))
				return
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, cErr.Error()))
				return
			}
		}
		var form vo.CaseStatusChangeForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			log.Println(err)
			ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
			return
		} else {

			// 如果是用户
			if claims.UserType == models.USER_TYPE_LABOR && claims.Id == cases.UserID {
				if cases.Status == models.StatusResultConfirming || form.Status == models.StatusRefused || form.Status == models.StatusConfirmed { // 等待用户确认阶段，或者确认/拒绝反悔的时候
					if form.Status == models.StatusRefused || form.Status == models.StatusConfirmed { // 用户可以在确认接受或者拒绝调解的时候选择状态
						cases.Status = form.Status
					}
				} else {
					log.Println(err)
					ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您无法修改本案件的状态"))
					return
				}
			} else if claims.UserType == models.USER_TYPE_ADMIN || claims.UserType == models.USER_TYPE_DEPARTMENT {
				// 管理员或者部门人员可以修改
				cases.Status = form.Status
			} else {
				log.Println(err)
				ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, "您无法修改本案件的状态"))
				return
			}

			// 开始修改
			if dao.UpdateCase(cases) {
				// 记录
				blockchain.CreateHistoryByCase("修改调解案件状态", cases, claims.Id)

				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericSuccess, "状态修改成功"))
			} else {
				ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "状态修改失败"))
			}
		}

	}

}
