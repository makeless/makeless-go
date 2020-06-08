package go_saas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/imdario/mergo"
	h "net/http"
	"sync"
)

func (saas *Saas) updatePassword(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/auth/password",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var user = new(go_saas_model.User)
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

			var bcrypted string

			if bcrypted, err = http.GetSecurity().EncryptPassword(*passwordReset.GetNewPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(user, &go_saas_model.User{
				Model:    go_saas_model.Model{Id: userId},
				Password: &bcrypted,
				RWMutex:  new(sync.RWMutex),
			}, mergo.WithOverride); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().UpdatePassword(user); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
