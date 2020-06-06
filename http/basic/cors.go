package go_saas_basic_http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (http *Http) CorsMiddleware(Origins []string, AllowHeaders []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = Origins
	config.AllowCredentials = true
	config.AddAllowHeaders(AllowHeaders...)

	return cors.New(config)
}
