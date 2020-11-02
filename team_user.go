package makeless_go

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	h "net/http"
	"strconv"
	"sync"
)

func (makeless *Makeless) deleteTeamUser(http makeless_go_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/team-user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		http.TeamUserMiddleware(),
		http.NotTeamCreatorMiddleware(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var teamId, _ = strconv.Atoi(c.GetHeader("Team"))
			var teamUser = &makeless_go_model.TeamUser{
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

			if err = http.GetEvent().Trigger(userId, "makeless", "team-user:delete", nil); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
