package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	"github.com/jinzhu/gorm"
)

type Profile interface {
	UpdateProfile(connection *gorm.DB, user *go_saas_model.User, profile *_struct.Profile) (*go_saas_model.User, error)
	UpdateProfileTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error)
}
