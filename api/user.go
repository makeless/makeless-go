package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/model"
	"net/http"
)

func (api *Api) user(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var user *go_saas_model.User

	if user, err = api.GetDatabase().GetUser(userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, user))
}
