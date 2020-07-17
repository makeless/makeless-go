package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/imdario/mergo"
	h "net/http"
	"sync"
)

func (saas *Saas) updateProfileTeam(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/team/profile",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamMemberMiddleware(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var team = new(go_saas_model.Team)

			if err = c.ShouldBind(team); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(team, &go_saas_model.Team{
				UserId:    &userId,
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

			if team, err = http.GetDatabase().UpdateProfileTeam(team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}
