package saas_api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (api *Api) cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	return cors.New(config)
}
