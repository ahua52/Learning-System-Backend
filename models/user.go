package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	RoleId   string `json:"role_id"`
	Role     Role   `json:"role";gorm:"foreignKey:RoleId"`
	// permissions []Permission
	// Permissions []Permission
}

func (User) TableName() string {
	return "user"
}

func CheckAuth(username, password string) (bool, error) {
	var auth User
	err := DB.Select("id").Where(User{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.Id != "" {
		return true, nil
	}

	return false, nil
}

func GetUser(username string) (*User, *Role, error) {
	var user User
	fmt.Println("are you comming")
	err := DB.Preload("Role").Where("username = ? ", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}

	var role Role
	fmt.Println("444", user)
	err1 := DB.Preload("Permission").Where("id = ? ", user.RoleId).First(&role).Error
	if err1 != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err1
	}

	return &user, &role, nil

}
