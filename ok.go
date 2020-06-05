package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	h "net/http"
)

func (saas *Saas) ok(http go_saas_http.Http) error {
	http.GetRouter().GET("/api/ok", func(c *gin.Context) {
		c.JSON(h.StatusOK, http.Response(nil, "ok"))
	})

	return nil
}
