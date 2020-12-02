package makeless_go_model

import (
	"sync"
	"time"
)

type TeamInvitation struct {
	Model
	TeamId *uint `gorm:"not null" json:"teamId"`
	Team   *Team `json:"team"`

	TeamUserId *uint     `gorm:"not null" json:"teamUserId"`
	TeamUser   *TeamUser `json:"teamUser"`

	Email    *string    `gorm:"not null" json:"email"`
	Token    *string    `gorm:"unique;not null" json:"-"`
	Expire   *time.Time `gorm:"not null" json:"expire"`
	Accepted *bool      `gorm:"not null" json:"accepted"`

	*sync.RWMutex
}

func (teamInvitation *TeamInvitation) GetTeam() *Team {
	teamInvitation.RLock()
	defer teamInvitation.RUnlock()

	return teamInvitation.Team
}

func (teamInvitation *TeamInvitation) GetTeamId() *uint {
	teamInvitation.RLock()
	defer teamInvitation.RUnlock()

	return teamInvitation.TeamId
}

func (teamInvitation *TeamInvitation) GetTeamUser() *TeamUser {
	teamInvitation.RLock()
	defer teamInvitation.RUnlock()

	return teamInvitation.TeamUser
}

func (teamInvitation *TeamInvitation) GetEmail() *string {
	teamInvitation.RLock()
	defer teamInvitation.RUnlock()

	return teamInvitation.Email
}

func (teamInvitation *TeamInvitation) GetToken() *string {
	teamInvitation.RLock()
	defer teamInvitation.RUnlock()

	return teamInvitation.Token
}
