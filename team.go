package makeless_go

import (
	"fmt"
	"gorm.io/gorm"
	h "net/http"
	"strconv"
	"sync"
	"time"

	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/security"
	"github.com/makeless/makeless-go/struct"
)

func (makeless *Makeless) createTeam(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().POST(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var tx = http.GetDatabase().GetConnection().WithContext(c).Begin(new(sql.TxOptions))
			var teamInvitationExpire = time.Now().Add(time.Hour * 24 * 7)
			var teamInvitationAccepted = false
			var user = &makeless_go_model.User{
				Model:   makeless_go_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var teamCreate = &_struct.TeamCreate{
				RWMutex: new(sync.RWMutex),
			}

			if err = tx.Error; err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = c.ShouldBind(teamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if user, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection().WithContext(c), user, "id", fmt.Sprintf("%d", user.GetId())); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			var teamInvitations = make([]*makeless_go_model.TeamInvitation, len(teamCreate.GetInvitations()))
			var index = map[string]bool{
				*user.GetEmail(): true,
			}

			for i, teamInvitation := range teamCreate.GetInvitations() {
				var mail makeless_go_mailer.Mail
				var token string
				var userInvited = &makeless_go_model.User{
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

				teamInvitations[i] = &makeless_go_model.TeamInvitation{
					UserId:   &userId,
					Email:    teamInvitation.GetEmail(),
					Token:    &token,
					Expire:   &teamInvitationExpire,
					Accepted: &teamInvitationAccepted,
					RWMutex:  new(sync.RWMutex),
				}

				if userInvited, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection().WithContext(c), userInvited, "email", *teamInvitations[i].GetEmail()); err != nil && err != gorm.ErrRecordNotFound {
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

			var team = &makeless_go_model.Team{
				Name:   teamCreate.GetName(),
				UserId: &userId,
				TeamUsers: []*makeless_go_model.TeamUser{
					{UserId: &userId, Role: &makeless_go_security.RoleTeamOwner, RWMutex: new(sync.RWMutex)},
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

			if err = http.GetEvent().Trigger(userId, "makeless", "team:create", team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}

func (makeless *Makeless) deleteTeam(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().DELETE(
		"/api/auth/team",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamCreatorMiddleware(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var team = &makeless_go_model.Team{
				Model:   makeless_go_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().DeleteTeam(http.GetDatabase().GetConnection().WithContext(c), team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetEvent().Trigger(userId, "makeless", "team:delete", nil); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
