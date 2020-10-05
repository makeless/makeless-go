package makeless_go_security_basic

import (
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/model"
	"sync"
)

func (security *Security) IsTeamUser(connection *gorm.DB, teamId uint, userId uint) (bool, error) {
	var team = &makeless_go_model.Team{
		Model:   makeless_go_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &makeless_go_model.User{
		Model:   makeless_go_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamUser(connection, team, user)
}

func (security *Security) IsTeamRole(connection *gorm.DB, role string, teamId uint, userId uint) (bool, error) {
	var team = &makeless_go_model.Team{
		Model:   makeless_go_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &makeless_go_model.User{
		Model:   makeless_go_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamRole(connection, role, team, user)
}

func (security *Security) IsTeamCreator(connection *gorm.DB, teamId uint, userId uint) (bool, error) {
	var team = &makeless_go_model.Team{
		Model:   makeless_go_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &makeless_go_model.User{
		Model:   makeless_go_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsTeamCreator(connection, team, user)
}

func (security *Security) IsNotTeamCreator(connection *gorm.DB, teamId uint, userId uint) (bool, error) {
	var team = &makeless_go_model.Team{
		Model:   makeless_go_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	var user = &makeless_go_model.User{
		Model:   makeless_go_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsNotTeamCreator(connection, team, user)
}
