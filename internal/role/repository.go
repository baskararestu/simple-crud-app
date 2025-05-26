package role

import "gorm.io/gorm"

type mysqlRoleRepository struct {
	db *gorm.DB
}