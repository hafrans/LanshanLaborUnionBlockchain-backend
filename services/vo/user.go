package vo

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/models"
)

type UserDataPayload struct {
	User    *models.User         `json:"user"`
	Claims  *jwtmodel.UserClaims `json:"claims"`
	Profile *models.UserProfile  `json:"profile"`
}

type UserData struct {
	Common
	UserDataPayload
}



