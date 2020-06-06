package go_saas_basic_http

import "github.com/gin-gonic/gin"

func (http *Http) Response(error error, data interface{}) gin.H {
	return gin.H{
		"error": error,
		"data":  data,
	}
}
