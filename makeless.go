package makeless_go

import (
	"gorm.io/gorm"
	"sync"

	"github.com/makeless/makeless-go/config"
	"github.com/makeless/makeless-go/database"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/logger"
	"github.com/makeless/makeless-go/mailer"
)

type Makeless struct {
	Config   makeless_go_config.Config
	Logger   makeless_go_logger.Logger
	Mailer   makeless_go_mailer.Mailer
	Database makeless_go_database.Database
	Http     makeless_go_http.Http
	*sync.RWMutex
}

func (makeless *Makeless) GetConfig() makeless_go_config.Config {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Config
}

func (makeless *Makeless) GetLogger() makeless_go_logger.Logger {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Logger
}

func (makeless *Makeless) GetMailer() makeless_go_mailer.Mailer {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Mailer
}

func (makeless *Makeless) GetDatabase() makeless_go_database.Database {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Database
}

func (makeless *Makeless) GetHttp() makeless_go_http.Http {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Http
}

func (makeless *Makeless) SetRoute(name string, handler func(http makeless_go_http.Http) error) {
	makeless.GetHttp().SetHandler(name, handler)
}

func (makeless *Makeless) SetMail(name string, handler func(data map[string]interface{}) (makeless_go_mailer.Mail, error)) {
	makeless.GetMailer().SetHandler(name, handler)
}

func (makeless *Makeless) Init(dialector gorm.Dialector, path string) error {
	if err := makeless.GetConfig().Load(path); err != nil {
		return err
	}

	if err := makeless.GetMailer().Init(); err != nil {
		return err
	}

	if err := makeless.GetHttp().GetEvent().Init(); err != nil {
		return err
	}

	if err := makeless.GetDatabase().Connect(dialector); err != nil {
		return err
	}

	if err := makeless.GetDatabase().Migrate(); err != nil {
		return err
	}

	if err := makeless.GetHttp().GetAuthenticator().CreateMiddleware(); err != nil {
		return err
	}

	makeless.SetRoute("ok", makeless.ok)
	makeless.SetRoute("passwordRequest", makeless.passwordRequest)
	makeless.SetRoute("passwordReset", makeless.passwordReset)
	makeless.SetRoute("register", makeless.register)
	makeless.SetRoute("login", makeless.login)
	makeless.SetRoute("logout", makeless.logout)
	makeless.SetRoute("refreshToken", makeless.refreshToken)
	makeless.SetRoute("verifyEmailVerification", makeless.verifyEmailVerification)
	makeless.SetRoute("resendEmailVerification", makeless.resendEmailVerification)
	makeless.SetRoute("events", makeless.events)
	makeless.SetRoute("user", makeless.user)
	makeless.SetRoute("updatePassword", makeless.updatePassword)
	makeless.SetRoute("updateProfile", makeless.updateProfile)

	if makeless.GetConfig().GetConfiguration().GetTokens() {
		makeless.SetRoute("tokens", makeless.tokens)
		makeless.SetRoute("createToken", makeless.createToken)
		makeless.SetRoute("deleteToken", makeless.deleteToken)
	}

	if makeless.GetConfig().GetConfiguration().GetTeams() != nil {
		makeless.SetRoute("createTeam", makeless.createTeam)
		makeless.SetRoute("deleteTeam", makeless.deleteTeam)
		makeless.SetRoute("teamUsersTeam", makeless.teamUsersTeam)
		makeless.SetRoute("deleteTeamUser", makeless.deleteTeamUser)
		makeless.SetRoute("deleteTeamUserTeam", makeless.deleteTeamUserTeam)
		makeless.SetRoute("updateRoleTeamUserTeam", makeless.updateRoleTeamUserTeam)

		makeless.SetRoute("updateProfileTeam", makeless.updateProfileTeam)

		makeless.SetRoute("teamInvitation", makeless.teamInvitation)
		makeless.SetRoute("teamInvitations", makeless.teamInvitations)
		makeless.SetRoute("registerTeamInvitation", makeless.registerTeamInvitation)
		makeless.SetRoute("acceptTeamInvitation", makeless.acceptTeamInvitation)
		makeless.SetRoute("deleteTeamInvitation", makeless.deleteTeamInvitation)

		makeless.SetRoute("teamInvitationsTeam", makeless.teamInvitationsTeam)
		makeless.SetRoute("createTeamInvitationsTeam", makeless.createTeamInvitationsTeam)
		makeless.SetRoute("resendTeamInvitationTeam", makeless.resendTeamInvitationTeam)
		makeless.SetRoute("deleteTeamInvitationTeam", makeless.deleteTeamInvitationTeam)

		if makeless.GetConfig().GetConfiguration().GetTeams().GetTokens() {
			makeless.SetRoute("tokensTeam", makeless.tokensTeam)
			makeless.SetRoute("createTokenTeam", makeless.createTokenTeam)
			makeless.SetRoute("deleteTokenTeam", makeless.deleteTokenTeam)
		}
	}

	makeless.SetMail("emailVerification", makeless.mailEmailVerification)
	makeless.SetMail("passwordRequest", makeless.mailPasswordRequest)

	if makeless.GetConfig().GetConfiguration().GetTeams() != nil {
		makeless.SetMail("teamInvitation", makeless.mailTeamInvitation)
	}

	return nil
}

func (makeless *Makeless) Run() error {
	return makeless.GetHttp().Start()
}
