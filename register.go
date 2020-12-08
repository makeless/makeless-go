package makeless_go

import (
	h "net/http"
	"sync"

	"github.com/gin-gonic/gin"
	makeless_go_http "github.com/makeless/makeless-go/http"
	makeless_go_mailer "github.com/makeless/makeless-go/mailer"
	makeless_go_model "github.com/makeless/makeless-go/model"
	makeless_go_security "github.com/makeless/makeless-go/security"
	_struct "github.com/makeless/makeless-go/struct"
)

func (makeless *Makeless) register(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().POST("/api/register", func(c *gin.Context) {
		var err error
		var token string
		var userExists bool
		var mail makeless_go_mailer.Mail
		var verified = false
		var register = &_struct.Register{
			RWMutex: new(sync.RWMutex),
		}

		if err := c.ShouldBind(register); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if token, err = http.GetSecurity().GenerateToken(32); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		var user = &makeless_go_model.User{
			Name:     register.GetName(),
			Password: register.GetPassword(),
			Email:    register.GetEmail(),
			EmailVerification: &makeless_go_model.EmailVerification{
				Token:    &token,
				Verified: &verified,
				RWMutex:  new(sync.RWMutex),
			},
			RWMutex: new(sync.RWMutex),
		}

		if userExists, err = http.GetSecurity().UserExists(http.GetSecurity().GetDatabase().GetConnection(), "email", *user.Email); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if userExists {
			c.AbortWithStatusJSON(h.StatusOK, http.Response(makeless_go_security.UserAlreadyExists, nil))
			return
		}

		if user, err = http.GetSecurity().Register(http.GetSecurity().GetDatabase().GetConnection(), user); err != nil {
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

		c.JSON(h.StatusOK, http.Response(nil, user))
	})

	return nil
}
