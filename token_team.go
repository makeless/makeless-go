package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
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
