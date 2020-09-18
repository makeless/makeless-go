package go_saas

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	"github.com/jinzhu/gorm"
	h "net/http"
	"sync"
)

func (saas *Saas) passwordReset(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/password-reset",
		func(c *gin.Context) {
			var err error
			var bcrypted string
			var tx = http.GetDatabase().GetConnection().BeginTx(c, new(sql.TxOptions))
			var token = c.Query("token")
			var user = &go_saas_model.User{
				RWMutex: new(sync.RWMutex),
			}
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

			if user, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection(), user, "email", *passwordRequest.GetEmail()); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if bcrypted, err = http.GetSecurity().EncryptPassword(*passwordReset.GetPassword()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if _, err = http.GetDatabase().UpdatePassword(tx, user, bcrypted); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if _, err = http.GetDatabase().UpdatePasswordRequest(tx, passwordRequest); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err := tx.Commit().Error; err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
