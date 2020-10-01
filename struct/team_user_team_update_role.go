package _struct

import "sync"

type TeamUserTeamUpdateRole struct {
	Id   *uint   `json:"id" binding:"required"`
	Role *string `json:"role" binding:"required"`
	*sync.RWMutex
}

func (teamUserTeamUpdateRole *TeamUserTeamUpdateRole) GetId() *uint {
	teamUserTeamUpdateRole.RLock()
	defer teamUserTeamUpdateRole.RUnlock()

	return teamUserTeamUpdateRole.Id
}

func (teamUserTeamUpdateRole *TeamUserTeamUpdateRole) GetRole() *string {
	teamUserTeamUpdateRole.RLock()
	defer teamUserTeamUpdateRole.RUnlock()

	return teamUserTeamUpdateRole.Role
}
