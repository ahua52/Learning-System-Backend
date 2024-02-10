package models

import (
	"gorm.io/gorm"
)

type Role struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Label       string       `json:"label"`
	Order       int          `json:"order"`
	Status      int          `json:"status"`
	Description string       `json:"description"`
	Permission  []Permission `gorm:"many2many:role_permission;"json:"permission"`
}

func (Role) TableName() string {
	return "role"
}

func GetRole(roleId string) (*Role, error) {

	var role Role
	err1 := DB.Preload("Permission").Where("id = ? ", roleId).First(&role).Error
	if err1 != nil && err != gorm.ErrRecordNotFound {
		return nil, err1
	}
	return &role, nil

}
