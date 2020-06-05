package go_saas_http

import (
	"github.com/go-saas/go-saas/authenticator"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/event"
	"github.com/go-saas/go-saas/jwt"
	"github.com/go-saas/go-saas/logger"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/tls"
)

type Http interface {
	GetLogger() go_saas_logger.Logger
	GetEvent() go_saas_event.Event
	GetAuthenticator() go_saas_authenticator.Authenticator
	GetSecurity() go_saas_security.Security
	GetDatabase() go_saas_database.Database
	GetJwt() go_saas_jwt.Jwt
	GetTls() go_saas_tls.Tls
	GetOrigins() []string
	GetPort() string
	GetMode() string
}
