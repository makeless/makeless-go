package go_saas_http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/authenticator"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/event"
	"github.com/go-saas/go-saas/logger"
	go_saas_mailer "github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/tls"
)

type Http interface {
	GetRouter() *gin.Engine
	GetHandlers() map[string]func(http Http) error
	GetLogger() go_saas_logger.Logger
	GetEvent() go_saas_event.Event
	GetAuthenticator() go_saas_authenticator.Authenticator
	GetSecurity() go_saas_security.Security
	GetDatabase() go_saas_database.Database
	GetMailer() go_saas_mailer.Mailer
	GetTls() go_saas_tls.Tls
	GetOrigins() []string
	GetHeaders() []string
	GetPort() string
	GetMode() string
	SetHandler(name string, handler func(http Http) error)
	Response(error error, data interface{}) gin.H
	CorsMiddleware(Origins []string, AllowHeaders []string) gin.HandlerFunc
	EmailVerificationMiddleware(enabled bool) gin.HandlerFunc
	TeamUserMiddleware() gin.HandlerFunc
	TeamRoleMiddleware(role string) gin.HandlerFunc
	TeamCreatorMiddleware() gin.HandlerFunc
	Start() error
}
