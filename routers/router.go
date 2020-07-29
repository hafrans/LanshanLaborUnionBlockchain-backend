package routers

import (
	v1 "RizhaoLanshanLabourUnion/routers/v1"
	"RizhaoLanshanLabourUnion/security/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	//register static resource paths
	r.Static("/static", "runtime/static")

	api := r.Group("/api")
	{
		apiAuth := api.Group("/auth")
		InitApiAuthRouters(apiAuth)

		apiV1 := api.Group("/v1")
		apiV1.Use(jwt.AuthMiddleWare.MiddlewareFunc())
		InitApiV1Routers(apiV1)
	}

}

func InitApiAuthRouters(apiAuth *gin.RouterGroup) {

	apiAuth.POST("/login", jwt.AuthMiddleWare.LoginHandler)
	apiAuth.POST("/labor/register", v1.RegisterNewLaborUser)
	apiAuth.POST("/employer/register", v1.RegisterNewEmployerUser)
	apiAuth.GET("/logout", jwt.AuthMiddleWare.LogoutHandler)
	apiAuth.GET("/refresh_token", jwt.AuthMiddleWare.RefreshHandler)
	apiAuth.GET("/captcha/:id", v1.GetCaptcha)

}

func InitApiV1Routers(apiV1 *gin.RouterGroup) {

	apiV1.GET("/", v1.ApiIndexHandler)

	apiV1.POST("/upload", v1.UploadAssets)

	apiUser := apiV1.Group("/user")
	apiUser.GET("/info", v1.GetUserInfo)
	apiUser.POST("/reset_password", v1.ResetUserPassword)
	apiUser.POST("/update_info", v1.UpdateUserInfo)

	apiLabor := apiV1.Group("/labor")
	apiLabor.POST("/create", v1.CreateLaborArbitrationForm)
	apiLabor.POST("/update/:id", v1.UpdateLaborArbitrationForm)
	apiLabor.GET("/list", v1.GetMyLaborArbitrationFormList)
	apiLabor.GET("/id/:id", v1.GetOneLaborArbitrationFormById)
	apiLabor.GET("/delete/:id", v1.DeleteOneLaborArbitrationFormById)

	apiCase := apiV1.Group("/case/")
	apiCase.POST("/create", v1.CreateNewCaseByApplicant)
	apiCase.POST("/update/:id", v1.UpdateCaseByApplicant)
	apiCase.POST("/status_change/:id", v1.ChangeCaseStatusById)
	apiCase.GET("/id/:id", v1.GetCaseById)
	apiCase.GET("/caseId/:caseId", v1.GetCaseByCaseID)
	apiCase.GET("/", v1.GetCaseList)
	apiCase.GET("/delete/:id", v1.DeleteCaseById)

	apiCategory := apiV1.Group("/category")
	apiCategory.GET("/", v1.GetAllCategories)

	apiTest := apiV1.Group("/test")
	apiTest.GET("/labor/template", v1.LaborArbitrationFormInstructor)
	apiTest.POST("/labor/template", v1.LaborArbitrationFormInstructor)
	apiTest.GET("/case/template", v1.GetCaseFirstSubmitFormTemplate)

	apiDepartment := apiV1.Group("/department")
	apiDepartment.GET("/list", v1.GetDepartmentList)
	apiDepartment.POST("/create", v1.CreateDepartment)
	apiDepartment.GET("/delete/:id", v1.DeleteDepartmentById)
	apiDepartment.GET("/id/:id", v1.GetOneDepartmentById)

	apiRecord := apiV1.Group("/record")
	apiRecord.POST("/create", v1.AddRecord)
	apiRecord.GET("/delete/:id", v1.DeleteRecord)

	apiSuggestion := apiV1.Group("/suggestion")
	apiSuggestion.POST("/create", v1.CreateSuggestion)
	apiSuggestion.GET("/delete/:id", v1.DeleteSuggestion)

	apiComment := apiV1.Group("/comment")
	apiComment.POST("/create", v1.AddComment)
	apiComment.GET("/delete/:id", v1.DeleteComment)

	apiBlockChain := apiV1.Group("/blockchain")
	apiBlockChain.GET("/history/case/:caseId", v1.GetHistoryByCaseID)

}
