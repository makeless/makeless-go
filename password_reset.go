package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	"github.com/jinzhu/gorm"
	"golang.org/x/sync/errgroup"
	h "net/http"
	"sync"
)

func (saas *Saas) passwordReset(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/password-reset",
		func(c *gin.Context) {
			var err error
			var bcrypted string
			var user *go_saas_model.User
			var g = new(errgroup.Group)
			var token = c.Query("token")
			var passwordReset = &_struct.PasswordReset{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(passwordReset); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var passwordRequest = &go_saas_model.PasswordRequest{
				Token:   &token,
				RWMutex: new(sync.RWMutex),
			}

			if passwordRequest, err = http.GetDatabase().GetPasswordRequest(http.GetDatabase().GetConnection(), passwordRequest); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if user, err = http.GetSecurity().Login("email", *passwordRequest.GetEmail(), *passwordReset.GetPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(err, nil))
				return
			}

			if bcrypted, err = http.GetSecurity().EncryptPassword(*passwordReset.GetNewPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			g.Go(func() error {
				_, err := http.GetDatabase().UpdatePassword(http.GetDatabase().GetConnection(), user, bcrypted)
				return err
			})

			g.Go(func() error {
				_, err := http.GetDatabase().UpdatePasswordRequest(http.GetDatabase().GetConnection(), passwordRequest)
				return err
			})

			if err := g.Wait(); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
