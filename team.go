package go_saas

import (
	"fmt"
	"github.com/jinzhu/gorm"
	h "net/http"
	"sync"
	"time"

	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/struct"
)

func (saas *Saas) createTeam(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var tx = http.GetDatabase().GetConnection().BeginTx(c, new(sql.TxOptions))
			var teamInvitationExpire = time.Now().Add(time.Hour * 24 * 7)
			var teamInvitationAccepted = false
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var teamCreate = &_struct.TeamCreate{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if user, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection(), user, "id", fmt.Sprintf("%d", user.GetId())); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			var teamInvitations = make([]*go_saas_model.TeamInvitation, len(teamCreate.GetInvitations()))
			var index = map[string]bool{
				*user.GetEmail(): true,
			}

			for i, teamInvitation := range teamCreate.GetInvitations() {
				var mail go_saas_mailer.Mail
				var token string
				var userInvited = &go_saas_model.User{
					RWMutex: new(sync.RWMutex),
				}
				teamInvitation.RWMutex = new(sync.RWMutex)

				if _, exists := index[*teamInvitation.GetEmail()]; exists {
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(nil, nil))
					return
				}

				if token, err = http.GetSecurity().GenerateToken(32); err != nil {
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
					return
				}

				teamInvitations[i] = &go_saas_model.TeamInvitation{
					UserId:   &userId,
					Email:    teamInvitation.GetEmail(),
					Token:    &token,
					Expire:   &teamInvitationExpire,
					Accepted: &teamInvitationAccepted,
					RWMutex:  new(sync.RWMutex),
				}

				if userInvited, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection(), userInvited, "email", *teamInvitations[i].GetEmail()); err != nil && err != gorm.ErrRecordNotFound {
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
					return
				}

				if mail, err = http.GetMailer().GetMail("teamInvitation", map[string]interface{}{
					"user":           user,
					"userInvited":    userInvited,
					"teamName":       *teamCreate.GetName(),
					"teamInvitation": teamInvitations[i],
				}); err != nil {
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
					return
				}

				if err = http.GetMailer().Send(c, mail); err != nil {
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
					return
				}
			}

			var team = &go_saas_model.Team{
				Name:   teamCreate.GetName(),
				UserId: &userId,
				TeamUsers: []*go_saas_model.TeamUser{
					{UserId: &userId, Role: &go_saas_security.RoleTeamOwner, RWMutex: new(sync.RWMutex)},
				},
				TeamInvitations: teamInvitations,
				RWMutex:         new(sync.RWMutex),
			}

			if team, err = http.GetDatabase().CreateTeam(tx, team); err != nil {
				tx.Rollback()
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if team, err = http.GetDatabase().GetTeam(tx, team); err != nil {
				tx.Rollback()
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
