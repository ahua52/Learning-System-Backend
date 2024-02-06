package models

type Permission struct {
	Id        string
	ParentId  string
	Label     string
	Type      int // 菜单类型 1: menu  0 category
	Route     string
	Component string
	Order     int // 菜单排序
}

func (Permission) TableName() string {
	return "permission"
}
