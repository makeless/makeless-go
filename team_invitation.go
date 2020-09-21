package go_saas

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/struct"
	"github.com/jinzhu/gorm"
	h "net/http"
	"sync"
)

func (saas *Saas) teamInvitations(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamInvitations []*go_saas_model.TeamInvitation
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitations, err = http.GetDatabase().GetTeamInvitations(http.GetDatabase().GetConnection(), user, teamInvitations); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamInvitations))
		},
	)

	return nil
}

func (saas *Saas) updateTeamInvitation(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var userEmail = http.GetAuthenticator().GetAuthEmail(c)
			var tx = http.GetDatabase().GetConnection().BeginTx(c, new(sql.TxOptions))
			var teamInvitationAccept = &_struct.TeamInvitationAccept{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(teamInvitationAccept); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var teamInvitation = &go_saas_model.TeamInvitation{
				Model:   go_saas_model.Model{Id: *teamInvitationAccept.GetId()},
				Email:   &userEmail,
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitation, err = http.GetDatabase().GetTeamInvitationByEmail(http.GetDatabase().GetConnection(), teamInvitation); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if teamInvitation, err = http.GetDatabase().AcceptTeamInvitation(tx, teamInvitation); err != nil {
				tx.Rollback()
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: *teamInvitation.GetTeamId()},
				RWMutex: new(sync.RWMutex),
			}

			var teamUser = &go_saas_model.TeamUser{
				UserId:  &userId,
				TeamId:  teamInvitation.GetTeamId(),
				Role:    &go_saas_security.RoleTeamUser,
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().AddTeamUsers(tx, []*go_saas_model.TeamUser{teamUser}, team); err != nil {
				tx.Rollback()
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = tx.Commit().Error; err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamInvitation))
		},
	)

	return nil
}

func (saas *Saas) deleteTeamInvitation(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team-invitation",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userEmail = http.GetAuthenticator().GetAuthEmail(c)
			var teamInvitationDelete = &_struct.TeamInvitationDelete{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(teamInvitationDelete); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var teamInvitation = &go_saas_model.TeamInvitation{
				Model:   go_saas_model.Model{Id: *teamInvitationDelete.GetId()},
				Email:   &userEmail,
				RWMutex: new(sync.RWMutex),
			}

			if teamInvitation, err = http.GetDatabase().GetTeamInvitationByEmail(http.GetDatabase().GetConnection(), teamInvitation); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if _, err = http.GetDatabase().DeleteTeamInvitation(http.GetDatabase().GetConnection(), teamInvitation); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
