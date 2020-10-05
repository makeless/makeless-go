package makeless

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	h "net/http"
)

func (makeless *Makeless) ok(http makeless_go_http.Http) error {
	http.GetRouter().GET(
		"/api/ok",
		func(c *gin.Context) {
			c.JSON(h.StatusOK, http.Response(nil, "ok"))
		},
	)

	return nil
}
