package go_saas_security_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"sync"
)

func (security *Security) IsTeamUser(connection *gorm.DB, teamId uint, userId uint) (bool, error) {
	var team = &go_saas_model.Team{
		Model:   go_saas_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &go_saas_model.User{
		Model:   go_saas_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamUser(connection, team, user)
}

func (security *Security) IsTeamRole(connection *gorm.DB, role string, teamId uint, userId uint) (bool, error) {
	var team = &go_saas_model.Team{
		Model:   go_saas_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &go_saas_model.User{
		Model:   go_saas_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamRole(connection, role, team, user)
}

func (security *Security) IsTeamCreator(connection *gorm.DB, teamId uint, userId uint) (bool, error) {
	var team = &go_saas_model.Team{
		Model:   go_saas_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &go_saas_model.User{
		Model:   go_saas_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamCreator(connection, team, user)
}
