package go_saas_http_basic

import (
	"github.com/gin-gonic/gin"
)

func (http *Http) Start() error {
	router := http.GetRouter()
	router.Use(http.CorsMiddleware(http.GetOrigins(), http.GetHeaders()))
	router.Use(gin.Recovery())

	for _, handler := range http.GetHandlers() {
		if err := handler(http); err != nil {
			return err
		}
	}

	if http.GetTls() != nil {
		return router.RunTLS(":"+http.GetPort(), http.GetTls().GetCertPath(), http.GetTls().GetKeyPath())
	}

	return router.Run(":" + http.GetPort())
}
