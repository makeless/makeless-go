package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	_struct "github.com/go-saas/go-saas/struct"
	"github.com/imdario/mergo"
	h "net/http"
	"sync"
)

func (saas *Saas) createTeam(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var teamCreate = &_struct.TeamCreate{
				RWMutex: new(sync.RWMutex),
			}
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var team = &go_saas_model.Team{
				Name:    teamCreate.GetName(),
				User:    user,
				RWMutex: new(sync.RWMutex),
			}

			owner := "owner"
			var teamUser = &go_saas_model.TeamUser{
				User: user,
				Team: team,
				Role: &owner,
			}

			if team, err = http.GetDatabase().CreateTeam(team, teamUser); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			http.GetEvent().Trigger(userId, "go-saas", "team-created", team)
			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}

func (saas *Saas) leaveDeleteTeam(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var team = new(go_saas_model.Team)
			var teamMember bool
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(team); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(team, &go_saas_model.Team{
				UserId:    &user.Id,
				User:      nil,
				TeamUsers: nil,
				RWMutex:   new(sync.RWMutex),
			}, mergo.WithOverride, mergo.WithTypeCheck); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			// mergo workaround
			team.User = nil
			team.TeamUsers = nil

			if teamMember, err = http.GetSecurity().IsTeamMember(team.GetId(), user.GetId()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !teamMember {
				c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamMemberErr, nil))
				return
			}

			if err = http.GetDatabase().DeleteTeamUsers(user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().DeleteTeam(team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			http.GetEvent().Trigger(userId, "go-saas", "team-leaved-deleted", nil)
			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
