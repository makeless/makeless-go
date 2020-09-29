package go_saas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/mailer"
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

func (saas *Saas) resendEmailVerification(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/email-verification/resend",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var mail go_saas_mailer.Mail
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if user, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection(), user, "id", fmt.Sprintf("%d", user.GetId())); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if user.GetEmailVerification().RWMutex == nil {
				user.GetEmailVerification().RWMutex = new(sync.RWMutex)
			}

			if mail, err = http.GetMailer().GetMail("emailVerification", map[string]interface{}{
				"user": user,
			}); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetMailer().Send(c, mail); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
