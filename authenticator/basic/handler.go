package makeless_go_authenticator_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	"sync"
)

func (authenticator *Authenticator) PayloadHandler(data interface{}) jwt.MapClaims {
	if user, ok := data.(*makeless_go_model.User); ok {
		return jwt.MapClaims{
			authenticator.GetIdentityKey(): user.GetId(),
			"email":                        user.GetEmail(),
			"emailVerification":            user.GetEmailVerification() != nil && *user.GetEmailVerification().GetVerified(),
		}
	}

	return jwt.MapClaims{}
}

func (authenticator *Authenticator) IdentityHandler(c *gin.Context) interface{} {
	var email = authenticator.GetAuthEmail(c)
	var emailVerification = authenticator.GetAuthEmailVerification(c)

	return &makeless_go_model.User{
		Model: makeless_go_model.Model{Id: authenticator.GetAuthUserId(c)},
		Email: &email,
		EmailVerification: &makeless_go_model.EmailVerification{
			Verified: &emailVerification,
			RWMutex:  new(sync.RWMutex),
		},
		RWMutex: new(sync.RWMutex),
	}
}

func (authenticator *Authenticator) AuthenticatorHandler(c *gin.Context) (interface{}, error) {
	var login = &_struct.Login{
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(login); err != nil {
		return nil, err
	}

	return authenticator.GetSecurity().Login(authenticator.GetSecurity().GetDatabase().GetConnection(), "email", *login.GetEmail(), *login.GetPassword())
}

func (authenticator *Authenticator) AuthorizatorHandler(data interface{}, c *gin.Context) bool {
	return true
}

func (authenticator *Authenticator) UnauthorizedHandler(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"error": message,
		"data":  nil,
	})
}
