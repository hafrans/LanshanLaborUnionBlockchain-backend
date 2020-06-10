package v1

import (
	"RizhaoLanshanLabourUnion/utils"
	"github.com/gin-gonic/gin"
)

func ApiIndexHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":    0,
		"message":   "ok",
		"timestamp": utils.CurrentTimeString(),
	})
}
