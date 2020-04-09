package saas_api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (api *Api) cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = api.getOrigins()
	config.AllowCredentials = true

	return cors.New(config)
}
