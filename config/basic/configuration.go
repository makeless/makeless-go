package go_saas_basic_config

import (
	"github.com/go-saas/go-saas/config"
	"sync"
)

type Configuration struct {
	Name   string              `json:"name"`
	Tokens bool                `json:"tokens"`
	Teams  go_saas_config.Team `json:"teams"`
	*sync.RWMutex
}

func (configuration *Configuration) GetName() string {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Name
}

func (configuration *Configuration) GetTokens() bool {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Tokens
}

func (configuration *Configuration) GetTeams() go_saas_config.Team {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Teams
}
