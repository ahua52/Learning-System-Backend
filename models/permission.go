package models

type Permission struct {
	Id        string `json:"id"`
	ParentId  string `json:"parent_id"`
	Label     string `json:"label"`
	Type      int    `json:"type"` // 菜单类型 1: menu  0 category
	Route     string `json:"route"`
	Component string `json:"component"`
	Order     int    `json:"order"`
	Hide      int    `json:"hide"`
}

func (Permission) TableName() string {
	return "permission"
}
