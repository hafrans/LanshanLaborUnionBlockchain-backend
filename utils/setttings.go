package utils

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type JWT struct {
	Realm       string
	Key         string
	IdentityKey string
}

type QQMeeting struct {
	AppID     string
	SDKID     string
	SecretID  string
	SecretKey string
}

var QQMeetingSettings = &QQMeeting{}

var DatabaseSettings = &Database{}

var JWTSettings = &JWT{}

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
	} else {
		log.Fatalln("Config Settings Load Failed " + err.Error())
	}

}

func InitTestSetting() {
	if cfg, err := ini.Load("E:\\GolangProjects\\RizhaoLanshanLabourUnion\\conf\\conf.ini"); err == nil {
		mapTo(cfg, "database", DatabaseSettings)
		mapTo(cfg, "jwt", JWTSettings)
	} else {
		log.Fatalln("Config Settings Load Failed " + err.Error())
	}
}
