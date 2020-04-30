package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/model"
	"net/http"
)

func (api *Api) createTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var team *saas_model.Team

	if err := c.ShouldBind(&team); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if team, err = api.GetDatabase().CreateTeam(team, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, team))
}

func (api *Api) deleteTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var team *saas_model.Team

	if err := c.ShouldBind(&team); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().LeaveTeam(team, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().DeleteTeam(team, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, nil))
}
