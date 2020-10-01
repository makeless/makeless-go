package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	h "net/http"
	"strconv"
	"sync"
)

func (saas *Saas) deleteTeamUser(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team-user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamUserMiddleware(),
		http.NotTeamCreatorMiddleware(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUser = &go_saas_model.TeamUser{
				RWMutex: new(sync.RWMutex),
			}

			if teamUser, err = http.GetDatabase().GetTeamUserByFields(http.GetDatabase().GetConnection(), teamUser, map[string]interface{}{
				"team_id": teamId,
				"user_id": userId,
			}); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if err = http.GetDatabase().DeleteTeamUser(http.GetDatabase().GetConnection(), teamUser); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetEvent().Trigger(userId, "go-saas", "team-user:delete", nil); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
