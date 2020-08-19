package go_saas_http_basic

import "github.com/gin-gonic/gin"

func (http *Http) Response(err error, data interface{}) gin.H {
	if err != nil {
		return gin.H{
			"error": err.Error(),
			"data":  data,
		}
	}

	return gin.H{
		"error": nil,
		"data":  data,
	}
}
