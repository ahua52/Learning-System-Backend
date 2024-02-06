package models

type Role struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Label       string       `json:"label"`
	Order       int          `json:"order"`
	Status      int          `json:"status"`
	Description string       `json:"description"`
	Permission  []Permission `gorm:"many2many:role_permission;"`
}

func (Role) TableName() string {
	return "role"
}
