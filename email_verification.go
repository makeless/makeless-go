package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	h "net/http"
	"sync"
)

func (saas *Saas) verifyEmailVerification(http go_saas_http.Http) error {
	http.GetRouter().PATCH(
		"/api/email-verification/verify",
		func(c *gin.Context) {
			var err error
			var token = c.Query("token")
			var emailVerification = &go_saas_model.EmailVerification{
				RWMutex: new(sync.RWMutex),
			}

			if emailVerification, err = http.GetDatabase().GetEmailVerificationByField(http.GetDatabase().GetConnection(), emailVerification, "token", token); err != nil {
				switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if emailVerification, err = http.GetDatabase().VerifyEmailVerification(http.GetDatabase().GetConnection(), emailVerification); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, emailVerification))
		},
	)

	return nil
}
