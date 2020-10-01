package go_saas_config_basic

import "sync"

type Team struct {
	Tokens bool                   `json:"tokens"`
	Roles  map[string]interface{} `json:"roles"`
	*sync.RWMutex
}

func (team *Team) GetTokens() bool {
	team.RLock()
	defer team.RUnlock()

	return team.Tokens
}

func (team *Team) GetRoles() map[string]interface{} {
	team.RLock()
	defer team.RUnlock()

	return team.Roles
}

func (team *Team) HasRole(role string) bool {
	team.RLock()
	defer team.RUnlock()

	_, exists := team.Roles[role]
	return exists
}
