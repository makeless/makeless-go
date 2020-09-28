package go_saas

import (
	"sync"

	"github.com/go-saas/go-saas/config"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/logger"
	"github.com/go-saas/go-saas/mailer"
)

type Saas struct {
	Config   go_saas_config.Config
	Logger   go_saas_logger.Logger
	Mailer   go_saas_mailer.Mailer
	Database go_saas_database.Database
	Http     go_saas_http.Http
	*sync.RWMutex
}

func (saas *Saas) GetConfig() go_saas_config.Config {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Config
}

func (saas *Saas) GetLogger() go_saas_logger.Logger {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Logger
}

func (saas *Saas) GetMailer() go_saas_mailer.Mailer {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Mailer
}

func (saas *Saas) GetDatabase() go_saas_database.Database {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Database
}

func (saas *Saas) GetHttp() go_saas_http.Http {
	saas.RLock()
	defer saas.RUnlock()

	return saas.Http
}

func (saas *Saas) SetRoute(name string, handler func(http go_saas_http.Http) error) {
	saas.GetHttp().SetHandler(name, handler)
}

func (saas *Saas) SetMail(name string, handler func(data map[string]interface{}) (go_saas_mailer.Mail, error)) {
	saas.GetMailer().SetHandler(name, handler)
}

func (saas *Saas) Init(path string) error {
	if err := saas.GetConfig().Load(path); err != nil {
		return err
	}

	if err := saas.GetMailer().Init(); err != nil {
		return err
	}

	if err := saas.GetHttp().GetEvent().Init(); err != nil {
		return err
	}

	if err := saas.GetDatabase().Connect(); err != nil {
		return err
	}

	if err := saas.GetDatabase().Migrate(); err != nil {
		return err
	}

	if err := saas.GetHttp().GetAuthenticator().CreateMiddleware(); err != nil {
		return err
	}

	saas.SetRoute("ok", saas.ok)
	saas.SetRoute("passwordRequest", saas.passwordRequest)
	saas.SetRoute("passwordReset", saas.passwordReset)
	saas.SetRoute("register", saas.register)
	saas.SetRoute("login", saas.login)
	saas.SetRoute("logout", saas.logout)
	saas.SetRoute("refreshToken", saas.refreshToken)
	saas.SetRoute("events", saas.events)
	saas.SetRoute("user", saas.user)
	saas.SetRoute("updatePassword", saas.updatePassword)
	saas.SetRoute("updateProfile", saas.updateProfile)

	if saas.GetConfig().GetConfiguration().GetTokens() {
		saas.SetRoute("tokens", saas.tokens)
		saas.SetRoute("createToken", saas.createToken)
		saas.SetRoute("deleteToken", saas.deleteToken)
	}

	if saas.GetConfig().GetConfiguration().GetTeams() != nil {
		saas.SetRoute("createTeam", saas.createTeam)
		saas.SetRoute("leaveDeleteTeam", saas.leaveDeleteTeam)

		saas.SetRoute("updateProfileTeam", saas.updateProfileTeam)

		saas.SetRoute("usersTeam", saas.usersTeam)
		saas.SetRoute("removeUserTeam", saas.removeUserTeam)

		saas.SetRoute("teamInvitation", saas.teamInvitation)
		saas.SetRoute("teamInvitations", saas.teamInvitations)
		saas.SetRoute("acceptTeamInvitation", saas.acceptTeamInvitation)
		saas.SetRoute("deleteTeamInvitation", saas.deleteTeamInvitation)

		saas.SetRoute("teamInvitationsTeam", saas.teamInvitationsTeam)
		saas.SetRoute("deleteTeamInvitationsTeam", saas.deleteTeamInvitationTeam)

		if saas.GetConfig().GetConfiguration().GetTeams().GetTokens() {
			saas.SetRoute("tokensTeam", saas.tokensTeam)
			saas.SetRoute("createTokenTeam", saas.createTokenTeam)
			saas.SetRoute("deleteTokenTeam", saas.deleteTokenTeam)
		}
	}

	saas.SetMail("passwordRequest", saas.mailPasswordRequest)
	saas.SetMail("teamInvitation", saas.mailTeamInvitation)

	return nil
}

func (saas *Saas) Run() error {
	return saas.GetHttp().Start()
}
