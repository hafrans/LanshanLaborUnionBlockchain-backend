package routers

import (
	v1 "RizhaoLanshanLabourUnion/routers/v1"
	"github.com/gin-gonic/gin"
)



func InitRouter(r *gin.Engine){


	//register static resource paths
	r.Static("/static","runtime/static")

	api := r.Group("/api")
	{
	   apiAuth := api.Group("/auth")
	   InitApiAuthRouters(apiAuth)

	   apiV1 := api.Group("/v1")
	   InitApiV1Routers(apiV1)
	}

}


func InitApiAuthRouters(apiAuth *gin.RouterGroup){

	apiAuth.Any("/authenticate",v1.LoginHandler)
	apiAuth.GET("/captcha/:id",v1.GetCaptcha)

}


func InitApiV1Routers(apiV1 *gin.RouterGroup){

	apiV1.GET("/", v1.ApiIndexHandler)

}

