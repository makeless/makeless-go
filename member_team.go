package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	h "net/http"
	"strconv"
	"sync"
)

func (saas *Saas) membersTeam(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/member",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamMemberMiddleware(),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var users []*go_saas_model.User
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if users, err = http.GetDatabase().MembersTeam(c.Query("search"), team, users); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, users))
		},
	)

	return nil
}

func (saas *Saas) removeMemberTeam(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team/member",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamOwnerMiddleware(),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var user = new(go_saas_model.User)
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(user); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().RemoveMemberTeam(user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
