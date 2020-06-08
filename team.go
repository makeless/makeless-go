package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
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
			var team = new(go_saas_model.Team)
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(team); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(team, &go_saas_model.Team{
				UserId:  &user.Id,
				User:    nil,
				Users:   nil,
				RWMutex: new(sync.RWMutex),
			}, mergo.WithOverride); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if team, err = http.GetDatabase().CreateTeam(user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			http.GetEvent().Trigger(userId, "go-saas", "team-created", team)
			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}

func (saas *Saas) deleteTeam(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var team = new(go_saas_model.Team)
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(team); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(team, &go_saas_model.Team{
				UserId:  &user.Id,
				RWMutex: new(sync.RWMutex),
			}, mergo.WithOverride); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().LeaveTeam(user, team); err != nil {
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
