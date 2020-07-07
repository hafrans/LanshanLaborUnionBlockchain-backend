package jwtmodel

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"RizhaoLanshanLabourUnion/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserClaims struct {
	Id       int64         `json:"id"`
	Sub      string        `json:"sub"`
	Roles    []interface{} `json:"roles"`
	Iss      string        `json:"iss"`
	Realm    string        `json:"realm"`
	UserType int           `json:"type"`
}

func ExtractUserClaimsFromGinContext(ctx *gin.Context) *UserClaims {

	claims := jwt.ExtractClaims(ctx)
	return &UserClaims{
		Id:       int64(claims[utils.JWTSettings.IdentityKey].(float64)),
		Sub:      claims["sub"].(string),
		Roles:    claims["roles"].([]interface{}),
		Iss:      claims["iss"].(string),
		Realm:    claims["realm"].(string),
		UserType: int(claims["type"].(float64)),
	}
}

func PopulateUserToUserClaims(user *models.User) *UserClaims {

	if user == nil {
		return &UserClaims{}
	}
	var roles []interface{} = make([]interface{}, 0, 4)
	if rs, err := dao.GetRolesFromUser(user); err == nil {
		for _, role := range rs {
			roles = append(roles, role.Descriptor)
		}
	}

	return &UserClaims{
		Id:       user.ID,
		Sub:      user.UserName,
		Roles:    roles,
		Iss:      "hafrans",
		Realm:    utils.JWTSettings.Realm,
		UserType: user.UserType,
	}

}
