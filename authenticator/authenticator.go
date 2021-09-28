package makeless_go_authenticator

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/security"
	"net/http"
	"time"
)

type Authenticator interface {
	SetMiddleware(middleware *jwt.GinJWTMiddleware)
	GetMiddleware() *jwt.GinJWTMiddleware
	CreateMiddleware() error
	AuthMiddleware() gin.HandlerFunc
	GetSecurity() makeless_go_security.Security
	GetRealm() string
	GetKey() []byte
	GetTimeout() time.Duration
	GetMaxRefresh() time.Duration
	GetIdentityKey() string
	PayloadHandler(data interface{}) jwt.MapClaims
	IdentityHandler(c *gin.Context) interface{}
	AuthenticatorHandler(c *gin.Context) (interface{}, error)
	AuthorizatorHandler(data interface{}, c *gin.Context) bool
	UnauthorizedHandler(c *gin.Context, code int, message string)
	GetSecureCookie() bool
	GetCookieDomain() string
	GetCookieSameSite() http.SameSite

	GetAuthUserId(c *gin.Context) uint
	GetAuthEmail(c *gin.Context) string
	GetAuthEmailVerification(c *gin.Context) bool
}
