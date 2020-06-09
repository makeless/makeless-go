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

			for _, user := range team.GetUsers() {
				user.RWMutex = new(sync.RWMutex)
			}

			c.JSON(h.StatusOK, http.Response(nil, users))
		},
	)

	return nil
}
