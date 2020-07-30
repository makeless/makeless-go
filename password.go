package go_saas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"sync"
)

func (saas *Saas) updatePassword(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/password",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var bcrypted string
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var passwordReset = &_struct.PasswordReset{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(passwordReset); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if _, err = http.GetSecurity().Login("id", fmt.Sprintf("%d", userId), *passwordReset.GetPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(err, nil))
				return
			}

			if bcrypted, err = http.GetSecurity().EncryptPassword(*passwordReset.GetNewPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().UpdatePassword(user, bcrypted); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
