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

func (saas *Saas) tokensTeam(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var tokens []*go_saas_model.Token
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				UserId:  &userId,
				RWMutex: new(sync.RWMutex),
			}

			if tokens, err = http.GetDatabase().GetTokensTeam(http.GetDatabase().GetConnection(), team, tokens); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, tokens))
		},
	)

	return nil
}

func (saas *Saas) createTokenTeam(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamUser bool
			var tmpTeamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamId = uint(tmpTeamId)
			var tokenTeamCreate = &_struct.TokenTeamCreate{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(tokenTeamCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if teamUser, err = http.GetSecurity().IsTeamUser(teamId, *tokenTeamCreate.GetUserId()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !teamUser {
				c.AbortWithStatusJSON(h.StatusForbidden, http.Response(go_saas_security.NoTeamMemberErr, nil))
				return
			}

			var token = &go_saas_model.Token{
				Note:    tokenTeamCreate.GetNote(),
				Token:   tokenTeamCreate.GetToken(),
				UserId:  tokenTeamCreate.GetUserId(),
				TeamId:  &teamId,
				RWMutex: new(sync.RWMutex),
			}

			if token, err = http.GetDatabase().CreateTokenTeam(http.GetDatabase().GetConnection(), token); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, token))
		},
	)

	return nil
}

func (saas *Saas) deleteTokenTeam(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamRoleMiddleware(go_saas_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var tmpTeamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamId = uint(tmpTeamId)
			var tokenTeamDelete = &_struct.TokenTeamDelete{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(tokenTeamDelete); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var token = &go_saas_model.Token{
				Model:   go_saas_model.Model{Id: tokenTeamDelete.GetId()},
				TeamId:  &teamId,
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().DeleteTokenTeam(http.GetDatabase().GetConnection(), token); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
