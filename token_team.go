package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"github.com/imdario/mergo"
	h "net/http"
	"strconv"
	"sync"
)

func (saas *Saas) tokensTeam(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/team/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.TeamOwnerMiddleware(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var tokens []*go_saas_model.Token
			var team = &go_saas_model.Team{
				Model:   go_saas_model.Model{Id: uint(teamId)},
				UserId:  &userId,
				RWMutex: new(sync.RWMutex),
			}

			if tokens, err = http.GetDatabase().GetTokensTeam(team, tokens); err != nil {
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
		http.TeamOwnerMiddleware(),
		func(c *gin.Context) {
			var err error
			var teamMember bool
			var tmpTeamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamId = uint(tmpTeamId)
			var token = new(go_saas_model.Token)

			if err := c.ShouldBind(token); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(token, &go_saas_model.Token{
				User:    nil,
				TeamId:  &teamId,
				Team:    nil,
				RWMutex: new(sync.RWMutex),
			}, mergo.WithOverride, mergo.WithTypeCheck); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			// mergo workaround
			token.User = nil
			token.Team = nil

			if teamMember, err = http.GetSecurity().IsTeamMember(teamId, *token.GetUserId()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !teamMember {
				c.AbortWithStatusJSON(h.StatusForbidden, http.Response(go_saas_security.NoTeamMemberErr, nil))
				return
			}

			if token, err = http.GetDatabase().CreateTokenTeam(token); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, token))
		},
	)

	return nil
}
