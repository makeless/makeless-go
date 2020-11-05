package makeless_go

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/security"
	"github.com/makeless/makeless-go/struct"
	"gorm.io/gorm"
	h "net/http"
	"strconv"
	"sync"
	"time"
)

func (makeless *Makeless) teamInvitationsTeam(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().GET(
		"/api/auth/team/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamInvitations []*makeless_go_model.TeamInvitation
			var team = &makeless_go_model.Team{
				Model:   makeless_go_model.Model{Id: uint(teamId)},
				UserId:  &userId,
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitations, err = http.GetDatabase().GetTeamInvitationsTeam(http.GetDatabase().GetConnection().WithContext(c), team, teamInvitations); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamInvitations))
		},
	)

	return nil
}

func (makeless *Makeless) createTeamInvitationsTeam(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().POST(
		"/api/auth/team/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var user *makeless_go_model.User
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamInvitationExpire = time.Now().Add(time.Hour * 24 * 7)
			var teamInvitationAccepted = false
			var teamInvitationTeamCreate = &_struct.TeamInvitationTeamCreate{
				RWMutex: new(sync.RWMutex),
			}
			var team = &makeless_go_model.Team{
				Model:   makeless_go_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamInvitationTeamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if team, err = http.GetDatabase().GetTeam(http.GetDatabase().GetConnection().WithContext(c), team); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			for _, teamUser := range team.GetTeamUsers() {
				teamUser.RWMutex = new(sync.RWMutex)
				teamUser.GetUser().RWMutex = new(sync.RWMutex)

				if userId == teamUser.GetUser().GetId() {
					user = teamUser.GetUser()
				}
			}

			var teamInvitations = make([]*makeless_go_model.TeamInvitation, len(teamInvitationTeamCreate.GetInvitations()))
			var index = make(map[string]bool)
			for _, teamUser := range team.GetTeamUsers() {
				teamUser.RWMutex = new(sync.RWMutex)
				teamUser.GetUser().RWMutex = new(sync.RWMutex)
				index[*teamUser.GetUser().GetEmail()] = true
			}

			for i, teamInvitation := range teamInvitationTeamCreate.GetInvitations() {
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
					"teamName":       *team.GetName(),
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

			if team, err = http.GetDatabase().AddTeamInvitations(http.GetDatabase().GetConnection().WithContext(c), team, teamInvitations); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}

func (makeless *Makeless) resendTeamInvitationTeam(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().POST(
		"/api/auth/team/team-invitation/resend",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var mail makeless_go_mailer.Mail
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamInvitationTeamResend = &_struct.TeamInvitationTeamResend{
				RWMutex: new(sync.RWMutex),
			}
			var userInvited = &makeless_go_model.User{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(teamInvitationTeamResend); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var tmpTeamId = uint(teamId)
			var teamInvitation = &makeless_go_model.TeamInvitation{
				Model:   makeless_go_model.Model{Id: *teamInvitationTeamResend.GetId()},
				TeamId:  &tmpTeamId,
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitation, err = http.GetDatabase().GetTeamInvitationByField(http.GetDatabase().GetConnection().WithContext(c), teamInvitation, "team_id", fmt.Sprintf("%d", *teamInvitation.GetTeamId())); err != nil {
				switch errors.Is(err, gorm.ErrRecordNotFound) {
				case true:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if userInvited, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection().WithContext(c), userInvited, "email", *teamInvitation.GetEmail()); err != nil && err != gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			teamInvitation.GetUser().RWMutex, teamInvitation.GetTeam().RWMutex = new(sync.RWMutex), new(sync.RWMutex)
			if mail, err = http.GetMailer().GetMail("teamInvitation", map[string]interface{}{
				"user":           teamInvitation.GetUser(),
				"userInvited":    userInvited,
				"teamName":       *teamInvitation.GetTeam().GetName(),
				"teamInvitation": teamInvitation,
			}); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetMailer().Send(c, mail); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamInvitation))
		},
	)

	return nil
}

func (makeless *Makeless) deleteTeamInvitationTeam(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().DELETE(
		"/api/auth/team/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamInvitationTeamDelete = &_struct.TeamInvitationTeamDelete{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(teamInvitationTeamDelete); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var tmpTeamId = uint(teamId)
			var teamInvitation = &makeless_go_model.TeamInvitation{
				Model:   makeless_go_model.Model{Id: *teamInvitationTeamDelete.GetId()},
				TeamId:  &tmpTeamId,
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitation, err = http.GetDatabase().GetTeamInvitationByField(http.GetDatabase().GetConnection().WithContext(c), teamInvitation, "team_id", fmt.Sprintf("%d", *teamInvitation.GetTeamId())); err != nil {
				switch errors.Is(err, gorm.ErrRecordNotFound) {
				case true:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if _, err = http.GetDatabase().DeleteTeamInvitation(http.GetDatabase().GetConnection().WithContext(c), teamInvitation); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
