package go_saas_config_basic

import (
	"github.com/go-saas/go-saas/config"
	"sync"
)

type Configuration struct {
	Name   string              `json:"name"`
	Logo   string              `json:"logo"`
	Locale string              `json:"locale"`
	Host   string              `json:"host"`
	Mail   go_saas_config.Mail `json:"mail"`
	Tokens bool                `json:"tokens"`
	Teams  go_saas_config.Team `json:"teams"`
	*sync.RWMutex
}

func (configuration *Configuration) GetName() string {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Name
}

func (configuration *Configuration) GetLogo() string {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Logo
}

func (configuration *Configuration) GetLocale() string {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Locale
}

func (configuration *Configuration) GetHost() string {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Host
}

func (configuration *Configuration) GetMail() go_saas_config.Mail {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Mail
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
