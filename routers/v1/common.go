package v1

import (
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

var ContentAccept = make(map[string]string)

func init() {
	ContentAccept["application/pdf"] = ".pdf"
	ContentAccept["image/png"] = ".png"
	ContentAccept["image/jpg"] = ".jpg"
	ContentAccept["image/gif"] = ".gif"
	ContentAccept["image/jpeg"] = ".jpeg"
	ContentAccept["text/plain"] = ".txt"
	ContentAccept["application/vnd.openxmlformats-officedocument.wordprocessingml.document"] = ".docx"
	ContentAccept["application/msword"] = ".doc"
	ContentAccept["application/vnd.openxmlformats-officedocument.presentationml.presentation"] = ".pptx"
	ContentAccept["application/zip"] = ".zip"
	ContentAccept["audio/mpeg"] = ".mp3"
}

// Api Index
// @Summary ApiIndex
// @Description 测试在登录情况下是否可以访问
// @Tags utils
// @Accept json
// @Produce json
// @Success 200 {object} vo.Common
// @Failure 401 {object} vo.Common
// @Router /api/v1/ [get]
func ApiIndexHandler(ctx *gin.Context) {
	ctx.JSON(200, vo.Common{
		Status:    0,
		Message:   "success",
		Timestamp: utils.NowTime(),
	})
}

// Upload Assets
// @Summary 上传资源
// @Description 上传资源 ，支持 pdf,doc,docx,pdfx,txt,mp3,zip,jpg,gif,jpeg,png格式
// @Tags utils
// @Accept mpfd
// @Produce json
// @Param file query string true "要上传的资源"
// @Success 200 {object} vo.Common
// @Failure 401 {object} vo.Common
// @Router /api/v1/upload [get]
func UploadAssets(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	}

	ext, ok := ContentAccept[file.Header.Get("Content-Type")]

	if !ok {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, "上传文件格式不允许！"))
		return
	}

	// store
	now := time.Now()
	targetDir := "/static/" + strconv.Itoa(int(now.Year())) + "/" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(int(now.Day()))
	os.MkdirAll("runtime"+targetDir,0777)
	targetFile := targetDir +  "/" + base64.StdEncoding.EncodeToString([]byte("1111"+file.Filename)) + ext
	err = ctx.SaveUploadedFile(file, "runtime"+targetFile)

	if err != nil {
		ctx.JSON(200, vo.GenerateCommonResponseHead(respcode.GenericFailed, err.Error()))
		return
	}

	ctx.JSON(200, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(0, "success"),
		Data: gin.H{
			"filename": file.Filename,
			"path":     targetFile,
		},
	})
}
