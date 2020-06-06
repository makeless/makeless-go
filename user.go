package go_saas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	h "net/http"
	"sync"
)

func (saas *Saas) user(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/user",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if user, err = http.GetDatabase().GetUser(user); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, user))
		},
	)

	return nil
}

func (saas *Saas) updateUserPassword(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/user/password",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			var passwordReset = &go_saas_model.PasswordReset{
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

			var bcrypted []byte

			if bcrypted, err = http.GetSecurity().EncryptPassword(*passwordReset.GetNewPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			http.GetDatabase().GetConnection().LogMode(true)
			if err = http.GetDatabase().UpdateUserPassword(user, string(bcrypted)); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
