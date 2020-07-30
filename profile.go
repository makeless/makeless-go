package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"sync"
)

func (saas *Saas) updateProfile(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/profile",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var profile = &_struct.Profile{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(profile); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if user, err = http.GetDatabase().UpdateProfile(user, profile); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, user))
		},
	)

	return nil
}
