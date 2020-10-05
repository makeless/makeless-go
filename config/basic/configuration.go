package makeless_go_config_basic

import (
	"github.com/makeless/makeless-go/config"
	"sync"
)

type Configuration struct {
	Name              string              `json:"name"`
	Logo              string              `json:"logo"`
	Locale            string              `json:"locale"`
	Host              string              `json:"host"`
	EmailVerification bool                `json:"emailVerification"`
	Tokens            bool                `json:"tokens"`
	Teams             makeless_go_config.Team `json:"teams"`
	Mail              makeless_go_config.Mail `json:"mail"`
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

func (configuration *Configuration) GetEmailVerification() bool {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.EmailVerification
}

func (configuration *Configuration) GetTokens() bool {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Tokens
}

func (configuration *Configuration) GetTeams() makeless_go_config.Team {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Teams
}

func (configuration *Configuration) GetMail() makeless_go_config.Mail {
	configuration.RLock()
	defer configuration.RUnlock()

	return configuration.Mail
}
