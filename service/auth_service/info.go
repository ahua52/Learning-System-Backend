package auth_service

import (
	"HMS/payment/models"
)

type Info struct {
	// Username string
	// Password string
	id          string
	username    string
	password    string
	email       string
	avatar      string
	roleId      string
	role        *models.Role
	permissions []models.Permission
}

// func GetInfo(username string) (*Info, error) {
// 	user, role, err1 := models.GetUser(username)
// 	if err1 != nil && err1 != gorm.ErrRecordNotFound {
// 		return nil, err1
// 	}
// 	role, err2 := models.GetRole(user.Id)
// 	if err2 != nil && err2 != gorm.ErrRecordNotFound {
// 		return nil, err1
// 	}

// 	var info = Info{
// 		id:          user.Id,
// 		username:    user.Username,
// 		password:    user.Password,
// 		email:       user.Email,
// 		avatar:      user.Avatar,
// 		role:        role,
// 		permissions: role.Permission,
// 	}
// 	fmt.Println(info, 22222)
// 	return &info, nil

// }
