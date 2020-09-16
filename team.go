package go_saas

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"strconv"
	"sync"
	"time"
)

func (saas *Saas) createTeam(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var tx = http.GetDatabase().GetConnection().BeginTx(c, new(sql.TxOptions))
			var teamInvitationExpire = time.Now().Add(time.Hour * 24 * 7)
			var teamInvitationAccepted = false
			var teamCreate = &_struct.TeamCreate{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var team = &go_saas_model.Team{
				Name:   teamCreate.GetName(),
				UserId: &userId,
				TeamUsers: []*go_saas_model.TeamUser{
					{UserId: &userId, Role: &go_saas_security.RoleTeamOwner, RWMutex: new(sync.RWMutex)},
				},
				RWMutex: new(sync.RWMutex),
			}

			if team, err = http.GetDatabase().CreateTeam(tx, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			var tmpInvitations = make([]*go_saas_model.TeamInvitation, len(teamCreate.GetInvitations()))
			var index = make(map[string]bool, len(team.GetTeamUsers()))
			for _, teamUser := range team.GetTeamUsers() {
				teamUser.RWMutex, teamUser.User.RWMutex = new(sync.RWMutex), new(sync.RWMutex)
				index[*teamUser.GetUser().GetEmail()] = true
			}

			for i, invitation := range teamCreate.GetInvitations() {
				var token string
				invitation.RWMutex = new(sync.RWMutex)

				if _, exists := index[*invitation.GetEmail()]; exists {
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(nil, nil))
					return
				}

				if token, err = http.GetSecurity().GenerateToken(32); err != nil {
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
					return
				}

				var tmpTeamId = team.GetId()
				tmpInvitations[i] = &go_saas_model.TeamInvitation{
					TeamId:   &tmpTeamId,
					UserId:   &userId,
					Email:    invitation.GetEmail(),
					Token:    &token,
					Expire:   &teamInvitationExpire,
					Accepted: &teamInvitationAccepted,
					RWMutex:  new(sync.RWMutex),
				}
			}

			if team, err = http.GetDatabase().AddTeamInvitations(tx, team, tmpInvitations); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = tx.Commit().Error; err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetEvent().Trigger(userId, "go-saas", "team:created", team); err != nil {
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
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().DeleteTeamUser(http.GetDatabase().GetConnection(), user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().DeleteTeam(http.GetDatabase().GetConnection(), user, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetEvent().Trigger(userId, "go-saas", "team:leaved-deleted", nil); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
