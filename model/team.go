package go_saas_model

import "sync"

type Team struct {
	Model
	Name *string `gorm:"not null" json:"name" binding:"required,min=4"`

	UserId *uint `gorm:"not null" json:"userId" binding:"-"` // FIXME: check binding
	User   *User `json:"-"`

	Users []*User `gorm:"many2many:user_teams;" json:"-"`

	*sync.RWMutex `json:"-"`
}

func (team *Team) GetId() uint {
	team.RLock()
	defer team.RUnlock()

	return team.Id
}

func (team *Team) GetName() *string {
	team.RLock()
	defer team.RUnlock()

	return team.Name
}

func (team *Team) GetUserId() *uint {
	team.RLock()
	defer team.RUnlock()

	return team.UserId
}

func (team *Team) GetUser() *User {
	team.RLock()
	defer team.RUnlock()

	return team.User
}

func (team *Team) GetUsers() []*User {
	team.RLock()
	defer team.RUnlock()

	return team.Users
}
