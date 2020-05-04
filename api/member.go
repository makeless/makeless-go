package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/model"
	"net/http"
	"sync"
)

func (api *Api) members(c *gin.Context) {
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

	var users []*saas_model.User

	if users, err = api.GetDatabase().GetMembers(users, header.TeamId, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, users))
}

func (api *Api) removeMember(c *gin.Context) {
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

	var user = &saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBind(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().RemoveMember(user, header.TeamId, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, nil))
}
