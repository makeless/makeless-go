package go_saas_http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/authenticator"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/event"
	"github.com/go-saas/go-saas/logger"
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
	GetTls() go_saas_tls.Tls
	GetOrigins() []string
	GetPort() string
	GetMode() string
	SetHandler(name string, handler func(http Http) error)
	Response(error error, data interface{}) gin.H
	CorsMiddleware(Origins []string, AllowHeaders []string) gin.HandlerFunc
	Start() error
}
