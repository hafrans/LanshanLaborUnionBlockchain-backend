package main

import (
	_ "RizhaoLanshanLabourUnion/docs"
	"RizhaoLanshanLabourUnion/routers"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/utils"
	"fmt"
	_ "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	"time"
)

// @title 区块链仲裁系统
// @version 0.0.1
// @description Hello
func main(){

	// init all components
	utils.InitSettings()
	dao.InitDB()
	dao.TryInitializeTables()


	// initialize api services
	port, ok:= os.LookupEnv("PORT")
	if !ok{
		fmt.Println("Using Default Port 8088. You can set environment value \"PORT\" to change the port")
		port = "8088"
	}

	router := gin.New()

	gin.DisableConsoleColor()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// register routers
	routers.InitRouter(router)

	/*
	 * Index
	 */
	router.GET("/", func(context *gin.Context) {
		context.String(200,"BUPT HAFRANS SERVER")
	})


	/*
	 * server status ping
	 */
	router.GET("/ping", pingHandler)


	/*
	 * remain for test
	 */

	router.GET("/test", func(context *gin.Context) {
		pendingCaptcha := utils.CreateCaptcha("hello")
		context.JSON(200,pendingCaptcha)
	})


	/*
	 * inject swagger
	 */
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	/*
	 * run server
	 */
	router.Run(":"+port)

}



// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func pingHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"status":0,
		"message":"pong",
		"data":map[string]interface{}{
			"timestamp":time.Now().Unix(),
			"datetime": time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}
