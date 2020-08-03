package go_saas_basic_security

import (
	"github.com/go-saas/go-saas/model"
	"sync"
)

func (security *Security) IsTeamMember(teamId uint, userId uint) (bool, error) {
	var team = &go_saas_model.Team{
		Model:   go_saas_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &go_saas_model.User{
		Model:   go_saas_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamMember(team, user)
}

func (security *Security) IsTeamRole(role string, teamId uint, userId uint) (bool, error) {
	var team = &go_saas_model.Team{
		Model:   go_saas_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &go_saas_model.User{
		Model:   go_saas_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamRole(role, team, user)
}

func (security *Security) IsTeamCreator(teamId uint, userId uint) (bool, error) {
	var team = &go_saas_model.Team{
		Model:   go_saas_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &go_saas_model.User{
		Model:   go_saas_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamCreator(team, user)
}
