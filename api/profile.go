package saas_api

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/model"
	"net/http"
	"sync"
)

func (api *Api) updateProfile(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var user = &saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBind(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if user, err = api.GetDatabase().UpdateProfile(user, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, user))
}

func (api *Api) updateProfileTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var header = &teamHeader{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBindHeader(header); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	var team = &saas_model.Team{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBind(team); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if team, err = api.GetDatabase().UpdateProfileTeam(team, header.TeamId, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, team))
}
