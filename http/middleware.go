package makeless_go_http

import "github.com/gin-gonic/gin"

type Middleware interface {
	CorsMiddleware(Origins []string, OriginsFunc func(origin string) bool, AllowHeaders []string) gin.HandlerFunc
	EmailVerificationMiddleware(enabled bool) gin.HandlerFunc
	TeamUserMiddleware() gin.HandlerFunc
	TeamRoleMiddleware(role string) gin.HandlerFunc
	TeamCreatorMiddleware() gin.HandlerFunc
	NotTeamCreatorMiddleware() gin.HandlerFunc
}
