package makeless_go

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"sync"
)

func (makeless *Makeless) updatePassword(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().PATCH(
		"/api/auth/password",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var bcrypted string
			var user = &makeless_go_model.User{
				Model:   makeless_go_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}
			var passwordUpdate = &_struct.PasswordUpdate{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(passwordUpdate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if _, err = http.GetSecurity().Login(http.GetSecurity().GetDatabase().GetConnection(), "id", fmt.Sprintf("%d", userId), *passwordUpdate.GetPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(err, nil))
				return
			}

			if bcrypted, err = http.GetSecurity().EncryptPassword(*passwordUpdate.GetNewPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if _, err = http.GetDatabase().UpdatePassword(http.GetDatabase().GetConnection().WithContext(c), user, bcrypted); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
