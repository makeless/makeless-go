package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type Password interface {
	UpdatePassword(connection *gorm.DB, user *go_saas_model.User, newPassword string) (*go_saas_model.User, error)
}
