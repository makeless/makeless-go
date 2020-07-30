package _struct

import "sync"

type MemberTeamRemove struct {
	Id uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (memberTeamRemove *MemberTeamRemove) GetId() uint {
	memberTeamRemove.RLock()
	defer memberTeamRemove.RUnlock()

	return memberTeamRemove.Id
}
