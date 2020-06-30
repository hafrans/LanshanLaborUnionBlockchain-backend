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
	apiLabor.POST("/arbitration/create", v1.CreateLaborArbitrationForm)
	apiLabor.GET("/arbitration/", v1.GetMyLaborArbitrationForms)
	apiLabor.GET("/arbitration/:id", v1.GetOneLaborArbitrationFormById)

	apiCase := apiV1.Group("/case/")
	apiCase.POST("/create", v1.CreateNewCaseByApplicant)
	apiCase.GET("/id/:id", v1.GetCaseById)

	apiCategory := apiV1.Group("/category")
	apiCategory.GET("/", v1.GetAllCategories)

	apiTest := apiV1.Group("/test")
	apiTest.GET("/labor/arbitration/template", v1.LaborArbitrationFormInstructor)
	apiTest.POST("/labor/arbitration/template", v1.LaborArbitrationFormInstructor)
	apiTest.GET("/case/template", v1.GetCaseFirstSubmitFormTemplate)

}
