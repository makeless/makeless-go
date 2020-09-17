package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type Member interface {
	MembersTeam(connection *gorm.DB, search string, users []*go_saas_model.User, team *go_saas_model.Team) ([]*go_saas_model.User, error)
}
