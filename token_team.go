package makeless_go

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/security"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"strconv"
	"sync"
)

func (makeless *Makeless) tokensTeam(http makeless_go_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var tokens []*makeless_go_model.Token
			var team = &makeless_go_model.Team{
				Model:   makeless_go_model.Model{Id: uint(teamId)},
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

func (makeless *Makeless) createTokenTeam(http makeless_go_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
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

			if teamUser, err = http.GetSecurity().IsTeamUser(http.GetSecurity().GetDatabase().GetConnection(), teamId, *tokenTeamCreate.GetUserId()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !teamUser {
				c.AbortWithStatusJSON(h.StatusForbidden, http.Response(makeless_go_security.NoTeamUserErr, nil))
				return
			}

			var token = &makeless_go_model.Token{
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

func (makeless *Makeless) deleteTokenTeam(http makeless_go_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
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

			var token = &makeless_go_model.Token{
				Model:   makeless_go_model.Model{Id: *tokenTeamDelete.GetId()},
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
