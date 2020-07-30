package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	_struct "github.com/go-saas/go-saas/struct"
	h "net/http"
	"strconv"
	"sync"
)

func (saas *Saas) updateProfileTeam(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/team/profile",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamOwnerMiddleware(),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var profileTeam = &_struct.ProfileTeam{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(profileTeam); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				Name:    profileTeam.GetName(),
				RWMutex: new(sync.RWMutex),
			}

			if team, err = http.GetDatabase().UpdateProfileTeam(team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}
