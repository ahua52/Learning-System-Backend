package models

type RolePermition struct {
	RoleId       string `json:"role_id`
	PermissionId string `json:"permission_id`
}

func (RolePermition) TableName() string {
	return "role_permission"
}
