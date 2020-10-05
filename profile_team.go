package makeless

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

func (makeless *Makeless) updateProfileTeam(http makeless_go_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/team/profile",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamRoleMiddleware(makeless_go_security.RoleTeamOwner),
		func(c *gin.Context) {
			var err error
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var profileTeam = &_struct.ProfileTeam{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(profileTeam); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var team = &makeless_go_model.Team{
				Model:   makeless_go_model.Model{Id: uint(teamId)},
				Name:    profileTeam.GetName(),
				RWMutex: new(sync.RWMutex),
			}

			if team, err = http.GetDatabase().UpdateProfileTeam(http.GetDatabase().GetConnection(), team); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, team))
		},
	)

	return nil
}
