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

func (saas *Saas) createTeam(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamCreate = &_struct.TeamCreate{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var team = &go_saas_model.Team{
				Name:      teamCreate.GetName(),
				UserId:    &userId,
				TeamUsers: []*go_saas_model.TeamUser{{UserId: &userId, Role: &go_saas_security.RoleTeamOwner}},
				RWMutex:   new(sync.RWMutex),
			}

			if team, err = http.GetDatabase().CreateTeam(team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetEvent().Trigger(userId, "go-saas", "team-created", team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}

func (saas *Saas) leaveDeleteTeam(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamMemberMiddleware(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().DeleteTeamUser(user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().DeleteTeam(user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetEvent().Trigger(userId, "go-saas", "team-leaved-deleted", nil); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
