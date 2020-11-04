package makeless_go

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
	h "net/http"
	"sync"
)

func (makeless *Makeless) verifyEmailVerification(http makeless_go_http.Http) error {
	http.GetRouter().PATCH(
		"/api/email-verification/verify",
		func(c *gin.Context) {
			var err error
			var token = c.Query("token")
			var emailVerification = &makeless_go_model.EmailVerification{
				RWMutex: new(sync.RWMutex),
			}

			if emailVerification, err = http.GetDatabase().GetEmailVerificationByField(http.GetDatabase().GetConnection().WithContext(c), emailVerification, "token", token); err != nil {
				switch errors.Is(err, gorm.ErrRecordNotFound) {
				case true:
					c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				default:
					c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				}
				return
			}

			if emailVerification, err = http.GetDatabase().VerifyEmailVerification(http.GetDatabase().GetConnection().WithContext(c), emailVerification); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, emailVerification))
		},
	)

	return nil
}

func (makeless *Makeless) resendEmailVerification(http makeless_go_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/email-verification/resend",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			var err error
			var mail makeless_go_mailer.Mail
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var user = &makeless_go_model.User{
				Model:   makeless_go_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if user, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection().WithContext(c), user, "id", fmt.Sprintf("%d", user.GetId())); err != nil {
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
