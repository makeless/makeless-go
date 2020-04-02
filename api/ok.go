package saas_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *Api) ok(c *gin.Context) {
	c.JSON(http.StatusOK, api.Response(nil, "ok"))
}
