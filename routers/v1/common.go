package v1

import (
	"RizhaoLanshanLabourUnion/services/respcode"
	"RizhaoLanshanLabourUnion/services/smsqueue"
	"RizhaoLanshanLabourUnion/services/smsqueue/smsrpc"
	"RizhaoLanshanLabourUnion/services/vo"
	"RizhaoLanshanLabourUnion/utils"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strconv"
	"text/template"
	"time"
)

var ContentAccept = make(map[string]string)
var SMSCache = utils.NewLRUCache(32768)

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
	targetDir := path.Join([]string{
		"static",
		strconv.Itoa(int(now.Year())),
		strconv.Itoa(int(now.Month())),
		strconv.Itoa(int(now.Day())),
	}...)

	os.MkdirAll("runtime"+"/"+targetDir, 0777)
	targetFile := targetDir + "/" + base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(time.Now().Second())+file.Filename)) + ext

	log.Println("save file:", path.Join("runtime", targetFile))
	err = ctx.SaveUploadedFile(file, path.Join("runtime", targetFile))

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

// User-End Triggered Captcha Sender
// @Summary 发送短信验证码，验证信息
// @Tags message
// @Accept json
// @Produce json
// @Param email body vo.SMSCaptchaRequest true  "请求"
// @Success 200 {object} vo.SMSCaptchaResponse "正常业务处理"
// @Failure 401 {object} vo.Common "未验证"
// @Failure 422 {object} vo.Common "表单绑定失败"
// @Failure 500 {object} vo.Common "表单绑定失败"
// @Router /api/auth/sms/captcha/request [post]
func SendShortMessages(ctx *gin.Context) {

	var form vo.SMSCaptchaRequest

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(respcode.HttpBindingFailed, vo.GenerateCommonResponseHead(respcode.FormBindingFailed, "bind form failed"+err.Error()))
		return
	}

	// check captcha

	if !utils.CheckCaptchaWithForm("captcha", form) {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "验证码无效"))
		return
	}

	// 需要查询手机号存不存在发送条件
	sendTime, err := SMSCache.Get(form.Phone)
	if err == nil && time.Now().Sub(sendTime.(time.Time)) < time.Duration(time.Minute-2*time.Second) {
		// 能获取到缓存、且距离上次发送时间不到58秒。
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "操作过于频繁，请稍后再试！"))
		return
	}
	// 尝试发送短信验证码
	// 生成短信验证码 6 位数字
	captchaCode := captcha.RandomDigits(6)

	// TODO 仅供初期实验代码，由于耦合度过高，需要将以下代码剥离出去。

	tpl, err := template.New("captcha").Parse(smsqueue.SMSContentCaptcha)
	if err != nil {
		ctx.JSON(respcode.HttpOK, vo.GenerateCommonResponseHead(respcode.GenericFailed, "验证码模板失效，请联系管理员！"))
		return
	}

	// compile template
	var buf bytes.Buffer

	_ = tpl.Execute(&buf, struct {
		Code string
	}{
		Code: fmt.Sprintf(
			"%d%d%d%d%d%d", captchaCode[0],
			captchaCode[1],
			captchaCode[2],
			captchaCode[3],
			captchaCode[4],
			captchaCode[5]),
	})
	// generate challenge

	sendNowTime := time.Now()

	hash := sha256.New()
	hash.Write([]byte(sendNowTime.Format(time.RFC3339) + form.Phone + string(captchaCode) + utils.SMSSetting.Password))
	hashResult := hash.Sum(nil)
	challengeCode := base64.StdEncoding.EncodeToString(hashResult)

	// 异步发送成功
	smsrpc.SendMessage(utils.SMSSetting.Account, utils.SMSSetting.Password, form.Phone, buf.String())
	respTime := utils.Time(sendNowTime)

	// 注入数据
	SMSCache.Put(form.Phone, sendNowTime)
	ctx.JSON(respcode.HttpOK, vo.CommonData{
		Common: vo.GenerateCommonResponseHead(respcode.GenericSuccess, "验证码已发送"),
		Data: vo.SMSCaptchaResponse{
			Identifier: challengeCode,
			Timestamp:  &respTime,
		},
	})
}
