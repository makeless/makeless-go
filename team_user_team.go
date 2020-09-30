package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"strconv"
	"sync"
)

func (saas *Saas) teamUsersTeam(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/team-user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUsers []*go_saas_model.TeamUser
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if teamUsers, err = http.GetDatabase().GetTeamUsers(http.GetDatabase().GetConnection(), c.Query("search"), teamUsers, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamUsers))
		},
	)

	return nil
}

func (saas *Saas) deleteTeamUserTeam(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team/team-user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamOwner bool
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUserTeamDelete = &_struct.TeamUserTeamDelete{
				RWMutex: new(sync.RWMutex),
			}
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamUserTeamDelete); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: teamUserTeamDelete.GetId()},
				RWMutex: new(sync.RWMutex),
			}

			if teamOwner, err = http.GetSecurity().IsTeamRole(http.GetDatabase().GetConnection(), go_saas_security.RoleTeamOwner, team.GetId(), user.GetId()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !teamOwner {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(nil, nil))
				return
			}

			if err = http.GetDatabase().DeleteTeamUser(http.GetDatabase().GetConnection(), user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
