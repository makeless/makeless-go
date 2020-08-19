package go_saas_security_basic

import (
	"github.com/go-saas/go-saas/model"
	"sync"
)

func (security *Security) TokenLogin(value string) (*go_saas_model.User, *go_saas_model.Team, error) {
	var err error
	var token = &go_saas_model.Token{
		RWMutex: new(sync.RWMutex),
	}

	if token, err = security.GetDatabase().GetToken(token, value); err != nil {
		return nil, nil, err
	}

	if token.GetUser() != nil {
		token.GetUser().RWMutex = new(sync.RWMutex)
	}

	if token.GetTeam() != nil {
		token.GetTeam().RWMutex = new(sync.RWMutex)
	}

	return token.GetUser(), token.GetTeam(), nil
}
