package makeless_go

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/security"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"strconv"
	"sync"
)

func (makeless *Makeless) teamUsersTeam(http makeless_go_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/team-user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUsers []*makeless_go_model.TeamUser
			var team = &makeless_go_model.Team{
				Model:   makeless_go_model.Model{Id: uint(teamId)},
				RWMutex: new(sync.RWMutex),
			}

			if teamUsers, err = http.GetDatabase().GetTeamUsers(http.GetDatabase().GetConnection(), c.Query("search"), teamUsers, team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamUsers))
		},
	)

	return nil
}

func (makeless *Makeless) updateRoleTeamUserTeam(http makeless_go_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/team/team-user/role",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUserTeamUpdateRole = &_struct.TeamUserTeamUpdateRole{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamUserTeamUpdateRole); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var teamUser = &makeless_go_model.TeamUser{
				Model:   makeless_go_model.Model{Id: *teamUserTeamUpdateRole.GetId()},
				RWMutex: new(sync.RWMutex),
			}

			if teamUser, err = http.GetDatabase().GetTeamUserByFields(http.GetDatabase().GetConnection(), teamUser, map[string]interface{}{
				"id":      teamUser.GetId(),
				"team_id": teamId,
			}); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if teamUser.GetTeam().RWMutex == nil {
				teamUser.GetTeam().RWMutex = new(sync.RWMutex)
			}

			if teamUser.GetTeam().GetUserId() == teamUser.GetUserId() || *teamUser.GetUserId() == userId {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if exists := makeless.GetConfig().GetConfiguration().GetTeams().HasRole(*teamUserTeamUpdateRole.GetRole()); !exists {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			}

			if teamUser, err = http.GetDatabase().UpdateRoleTeamUser(http.GetDatabase().GetConnection(), teamUser, *teamUserTeamUpdateRole.GetRole()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, teamUser))
		},
	)

	return nil
}

func (makeless *Makeless) deleteTeamUserTeam(http makeless_go_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team/team-user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUserTeamDelete = &_struct.TeamUserTeamDelete{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(teamUserTeamDelete); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var teamUser = &makeless_go_model.TeamUser{
				Model:   makeless_go_model.Model{Id: *teamUserTeamDelete.GetId()},
				RWMutex: new(sync.RWMutex),
			}

			if teamUser, err = http.GetDatabase().GetTeamUserByFields(http.GetDatabase().GetConnection(), teamUser, map[string]interface{}{
				"id":      teamUser.GetId(),
				"team_id": teamId,
			}); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if teamUser.GetTeam().RWMutex == nil {
				teamUser.GetTeam().RWMutex = new(sync.RWMutex)
			}

			if teamUser.GetTeam().GetUserId() == teamUser.GetUserId() || *teamUser.GetUserId() == userId {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().DeleteTeamUser(http.GetDatabase().GetConnection(), teamUser); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
