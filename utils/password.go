package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GenerateHashedPassword(password string) (string, bool){

	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil{
		log.Println("generate hashed password failed" + err.Error())
		return "", false
	}else{
		return string(hash),true
	}

}


func CheckHashedPassword(password, hashed string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(password))
	if err != nil {
		return false
	}else{
		return true
	}
}

