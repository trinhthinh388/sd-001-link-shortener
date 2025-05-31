package models

type RolesPermissions struct {
	Role       string `gorm:"column:role_role"`
	Permission string `gorm:"column:permission_permission"`
}
