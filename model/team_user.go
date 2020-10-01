package go_saas_model

import (
	"sync"
)

type TeamUser struct {
	Model
	TeamId *uint `gorm:"not null" json:"teamId"`
	Team   *Team `json:"team"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	Role *string `gorm:"not null" json:"role"`

	*sync.RWMutex
}

func (teamUser *TeamUser) GetId() uint {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.Id
}

func (teamUser *TeamUser) GetTeam() *Team {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.Team
}

func (teamUser *TeamUser) GetTeamId() *uint {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.TeamId
}

func (teamUser *TeamUser) GetUser() *User {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.User
}

func (teamUser *TeamUser) GetUserId() *uint {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.UserId
}

func (teamUser *TeamUser) GetRole() *string {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.Role
}
