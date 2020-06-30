package v1

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models/utils"
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	utils2 "RizhaoLanshanLabourUnion/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	}

	newCase := utils.PopulateCaseBasicFromFormToModel(&form, claims.Id, "371100")

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
			ctx.JSON(respcode.HttpOK, vo.CommonData{
				Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "success"),
				Data:   utils.PopulateCaseFullModelToFullForm(result),
			})
			return
		}

		// 提交成功 37110020200630134604159349596481
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
	marr := make([]*vo.Material,0,2)
	path := "/test/1.png"
	marr = append(marr, &vo.Material{Path: &path ,Name:"欠条"})
	marr = append(marr, &vo.Material{Path: &path ,Name:"老合同"})
	s.Materials = marr
	s.Applicant = vo.Applicant{Name: "张三",Contact: "10086",Address: "三体星",Nationality: "三体人",IdentityNumber: "1234567890123456789",Birthday: utils2.NowDateDay()}
	s.Respondent = vo.Employer{Name: "第三红岸基地",Address: "地球",Contact: "10010",UniformSocialCreditCode: "1234567889456123",LegalRepresentative: "李四"}
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
// @Description 获取单一Case
// @Tags case
// @Produce json
// @Success 200 {object} vo.CommonData "成功"
// @Failure 401 {object} vo.Common "没有认证"
// @Router /api/v1/test/case/id/:id [get]
func GetCaseById(ctx *gin.Context){

	if id,err := strconv.Atoi(ctx.Param("id")); err != nil{
		log.Println(err.Error())
		ctx.JSON(200,vo.GenerateCommonResponseHead(respcode.GenericFailed, "invalid id"))
		return
	}else{
		cases, cErr := dao.GetCasePreloadedModelById(int64(id))
		if cErr != nil{
			if cErr == sql.ErrNoRows {
				ctx.JSON(200,vo.GenerateCommonResponseHead(respcode.GenericFailed, "记录不存在"))
				return
			}else{
				ctx.JSON(200,vo.GenerateCommonResponseHead(respcode.GenericFailed, cErr.Error()))
				return
			}
		}
		ctx.JSON(200,vo.CommonData{
			Common:vo.GenerateCommonResponseHead(respcode.GenericSuccess,"success"),
			Data:utils.PopulateCaseFullModelToFullForm(cases),
		})
	}

}
