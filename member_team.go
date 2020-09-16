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

			if users, err = http.GetDatabase().MembersTeam(http.GetDatabase().GetConnection(), c.Query("search"), users, team); err != nil {
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
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var memberTeamRemove = &_struct.MemberTeamRemove{
				RWMutex: new(sync.RWMutex),
			}
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(memberTeamRemove); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: memberTeamRemove.GetId()},
				RWMutex: new(sync.RWMutex),
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
