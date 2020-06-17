package vo

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/models"
)


type UserData struct {
	Common
	UserDataPayload
}


type UserDataPayload struct {
	User    *models.User         `json:"user"`
	Claims  *jwtmodel.UserClaims `json:"claims"`
	Profile *models.UserProfile  `json:"profile"`
}

type UserResetPassword struct {
	OldPassword string 	`json:"old_password" form:"old_password" binding:"required,min=3,max=20" `
	NewPassword string 	`json:"new_password" form:"new_password" binding:"required,min=3,max=20" `
	ConfirmPassword string  `json:"confirm_password" form:"confirm_password" binding:"required,min=3,max=20,eqfield=NewPassword" `
}



