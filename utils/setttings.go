package utils

import (
	"github.com/go-ini/ini"
	"log"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}


type JWT struct{
	Realm   string
	Key     string
	IdentityKey string
}

var DatabaseSettings = &Database{}

var JWTSettings = &JWT{}


func mapTo(cfg *ini.File, section string, v interface{}){

	if err := cfg.Section(section).MapTo(v); err != nil{
		log.Fatalln("Config Settings Mapping "+section+" Failed: "+err.Error())
	}

}

func InitSettings(){

	if cfg, err := ini.Load("E:\\GolangProjects\\RizhaoLanshanLabourUnion\\conf\\conf.ini"); err == nil{
		mapTo(cfg,"database",DatabaseSettings)
		mapTo(cfg,"jwt",JWTSettings)
	}else{
		log.Fatalln("Config Settings Load Failed "+err.Error())
	}


}

