package saas_api

import "github.com/gin-gonic/gin"

func (api *Api) Response(error, data interface{}) gin.H {
	return gin.H{
		"error": error,
		"data":  data,
	}
}
