package makeless_go

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"sync"
)

func (makeless *Makeless) passwordReset(http makeless_go_http.Http) error {
	http.GetRouter().POST(
		"/api/password-reset",
		func(c *gin.Context) {
			var err error
			var bcrypted string
			var tx = http.GetDatabase().GetConnection().BeginTx(c, new(sql.TxOptions))
			var token = c.Query("token")
			var user = &makeless_go_model.User{
				RWMutex: new(sync.RWMutex),
			}
			var passwordReset = &_struct.PasswordReset{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(passwordReset); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var passwordRequest = &makeless_go_model.PasswordRequest{
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
				tx.Rollback()
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if _, err = http.GetDatabase().UpdatePasswordRequest(tx, passwordRequest); err != nil {
				tx.Rollback()
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
