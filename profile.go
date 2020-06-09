package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/imdario/mergo"
	h "net/http"
	"sync"
)

func (saas *Saas) updateProfile(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/profile",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var user = new(go_saas_model.User)

			if err = c.ShouldBind(user); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(user, go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				Teams:   nil,
				Tokens:  nil,
				RWMutex: new(sync.RWMutex),
			}, mergo.WithOverride, mergo.WithTypeCheck); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			// mergo workaround
			user.Teams = nil
			user.Tokens = nil

			if user, err = http.GetDatabase().UpdateProfile(user); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, user))
		},
	)

	return nil
}
