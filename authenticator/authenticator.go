package go_saas_authenticator

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/security"
	"time"
)

type Authenticator interface {
	SetMiddleware(middleware *jwt.GinJWTMiddleware)
	GetMiddleware() *jwt.GinJWTMiddleware
	CreateMiddleware() error
	GetSecurity() go_saas_security.Security
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
	GetAuthUserId(c *gin.Context) uint
}
