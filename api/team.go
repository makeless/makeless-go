package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"net/http"
	"sync"
)

func (api *Api) createTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var team *go_saas_model.Team

	if err := c.ShouldBind(&team); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if team, err = api.GetDatabase().CreateTeam(team, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	api.GetEvent().Trigger(userId, "go-saas", "team-created", team)
	c.JSON(http.StatusOK, api.Response(nil, team))
}

func (api *Api) deleteTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var team *go_saas_model.Team

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

	api.GetEvent().Trigger(userId, "go-saas", "team-leaved-deleted", nil)
	c.JSON(http.StatusOK, api.Response(nil, nil))
}

func (api *Api) usersTeam(c *gin.Context) {
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

	isTeamMember, err := api.GetSecurity().IsTeamMember(header.TeamId, userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	if !isTeamMember {
		c.AbortWithStatusJSON(http.StatusForbidden, api.Response(saas_security.NoTeamMemberErr, nil))
		return
	}

	var users []*go_saas_model.User

	if users, err = api.GetDatabase().GetUsersTeam(c.Query("search"), users, &header.TeamId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, users))
}
