package models

type Role struct {
	Id    string
	Name  string
	Label string
	// Permission     string
	Order       int
	Status      int
	Description string
	Permission  []Permission `gorm:"many2many:role_permission;"`
}

func (Role) TableName() string {
	return "role"
}
