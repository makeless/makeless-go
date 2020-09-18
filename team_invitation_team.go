package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	h "net/http"
	"strconv"
	"sync"
)

func (saas *Saas) teamInvitationsTeam(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamInvitations []*go_saas_model.TeamInvitation
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				UserId:  &userId,
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitations, err = http.GetDatabase().GetTeamInvitationsTeam(http.GetDatabase().GetConnection(), team, teamInvitations); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamInvitations))
		},
	)

	return nil
}
