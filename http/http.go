package makeless_go_http

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/authenticator"
	"github.com/makeless/makeless-go/database"
	"github.com/makeless/makeless-go/event"
	"github.com/makeless/makeless-go/logger"
	makeless_go_mailer "github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/security"
	"github.com/makeless/makeless-go/tls"
)

type Http interface {
	GetRouter() *gin.Engine
	GetHandlers() map[string]func(http Http) error
	GetLogger() makeless_go_logger.Logger
	GetEvent() makeless_go_event.Event
	GetAuthenticator() makeless_go_authenticator.Authenticator
	GetSecurity() makeless_go_security.Security
	GetDatabase() makeless_go_database.Database
	GetMailer() makeless_go_mailer.Mailer
	GetTls() makeless_go_tls.Tls
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
	NotTeamCreatorMiddleware() gin.HandlerFunc
	Start() error
}
