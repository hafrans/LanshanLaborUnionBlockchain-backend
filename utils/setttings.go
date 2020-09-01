package utils

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

// 短信发送接口
type SmsSender struct {
	Account  string
	Password string
}

// 数据库
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// jwt token 相关
type JWT struct {
	Realm       string
	Key         string
	IdentityKey string
}

// 腾讯会议相关
type QQMeeting struct {
	AppID     string
	SDKID     string
	SecretID  string
	SecretKey string
}

var QQMeetingSettings = &QQMeeting{}

var DatabaseSettings = &Database{}

var JWTSettings = &JWT{}

var SMSSetting = &SmsSender{}

func mapTo(cfg *ini.File, section string, v interface{}) {

	if err := cfg.Section(section).MapTo(v); err != nil {
		log.Fatalln("Config Settings Mapping " + section + " Failed: " + err.Error())
	}

}

func InitSettings() {

	if env, t := os.LookupEnv("TEST"); t && env == "on" {
		log.Println("+++++++++++++++  TEST CONF ++++++++++++++++++")
		//TODO 部署的时候关掉
		InitTestSetting()
		return
	} else {
		log.Println("==============  NO TEST CONF ==============")
	}

	if cfg, err := ini.Load("conf/conf.ini"); err == nil {
		mapTo(cfg, "database", DatabaseSettings)
		mapTo(cfg, "jwt", JWTSettings)
		mapTo(cfg, "QQMeeting", QQMeetingSettings)
		mapTo(cfg, "sms", SMSSetting)
	} else {
		log.Fatalln("Config Settings Load Failed " + err.Error())
	}

}

func InitTestSetting() {
	if cfg, err := ini.Load("E:\\GolangProjects\\RizhaoLanshanLabourUnion\\conf\\conf.ini"); err == nil {
		mapTo(cfg, "database", DatabaseSettings)
		mapTo(cfg, "jwt", JWTSettings)
		mapTo(cfg, "sms", SMSSetting)
	} else {
		log.Fatalln("Config Settings Load Failed " + err.Error())
	}
}
