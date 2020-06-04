package go_saas_basic_security

import (
	"github.com/go-saas/go-saas/model"
	"sync"
)

// FIXME: Add team (could be nil)
func (security *Security) TokenLogin(token string) (*go_saas_model.User, error) {
	var user = &go_saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	err := security.getDatabase().GetConnection().
		Select("users.id, users.name, users.username, users.email").
		Preload("Tokens").
		Preload("Teams").
		Joins("JOIN tokens ON tokens.user_id=users.id AND tokens.token = ?", token).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	for _, token := range user.GetTokens() {
		token.RWMutex = new(sync.RWMutex)
	}

	for _, team := range user.GetTeams() {
		team.RWMutex = new(sync.RWMutex)
	}

	return user, nil
}
