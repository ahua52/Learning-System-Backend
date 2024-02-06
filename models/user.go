package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id       string
	Username string
	Password string
	Email    string
	Avatar   string
	// Permissions string
	Role       string
	CreatedAt  int32
	UpdatedAt  int32
	UserRoleId string
	UserRole   Role `gorm:"foreignKey:UserRoleId"`
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

func GetUser(username string) ([]Role, error) {
	// var user User
	// fmt.Println("are you comming")
	// err := DB.Preload("UserRole").Where("username = ? ", username).First(&user).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return nil, err
	// }
	// return &user, nil
	userRole := []Role{}
	fmt.Println("are you comming")
	err := DB.Preload("Permission").First(&userRole).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return userRole, nil
}
