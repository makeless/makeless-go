package makeless

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"sync"
)

func (makeless *Makeless) updateProfile(http makeless_go_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/profile",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var user = &makeless_go_model.User{
				Model:   makeless_go_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var profile = &_struct.Profile{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(profile); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if user, err = http.GetDatabase().UpdateProfile(http.GetDatabase().GetConnection(), user, profile); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, user))
		},
	)

	return nil
}
