package go_saas_basic_config

import "sync"

type Team struct {
	Tokens bool `json:"tokens"`
	*sync.RWMutex
}

func (team *Team) GetTokens() bool {
	team.RLock()
	defer team.RUnlock()

	return team.Tokens
}
