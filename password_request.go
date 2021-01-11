package makeless_go

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"sync"
	"time"
)

func (makeless *Makeless) passwordRequest(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().POST(
		"/api/password-request",
		func(c *gin.Context) {
			var err error
			var userExists bool
			var token string
			var mail makeless_go_mailer.Mail
			var tokenExpire = time.Now().Add(time.Hour * 1)
			var tokenUsed = false
			var tmpPasswordRequest = &_struct.PasswordRequest{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(tmpPasswordRequest); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if userExists, err = http.GetSecurity().UserExists(http.GetSecurity().GetDatabase().GetConnection(), "email", *tmpPasswordRequest.GetEmail()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !userExists {
				c.AbortWithStatusJSON(h.StatusOK, http.Response(nil, nil))
				return
			}

			if token, err = http.GetSecurity().GenerateToken(32); err != nil {
				c.AbortWithStatusJSON(h.StatusOK, http.Response(nil, nil))
				return
			}

			var passwordRequest = &makeless_go_model.PasswordRequest{
				Email:   tmpPasswordRequest.GetEmail(),
				Token:   &token,
				Expire:  &tokenExpire,
				Used:    &tokenUsed,
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().CreatePasswordRequest(http.GetDatabase().GetConnection().WithContext(c), passwordRequest); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if mail, err = http.GetMailer().GetMail("passwordRequest", map[string]interface{}{
				"passwordRequest": passwordRequest,
			}); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetMailer().SendQueue(mail); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
