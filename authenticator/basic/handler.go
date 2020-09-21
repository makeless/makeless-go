package go_saas_authenticator_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	"sync"
)

func (authenticator *Authenticator) PayloadHandler(data interface{}) jwt.MapClaims {
	if user, ok := data.(*go_saas_model.User); ok {
		return jwt.MapClaims{
			authenticator.GetIdentityKey(): user.GetId(),
			"email":                        user.GetEmail(),
		}
	}

	return jwt.MapClaims{}
}

func (authenticator *Authenticator) IdentityHandler(c *gin.Context) interface{} {
	var id = authenticator.GetAuthUserId(c)
	var email = authenticator.GetAuthEmail(c)

	return &go_saas_model.User{
		Model:   go_saas_model.Model{Id: id},
		Email:   &email,
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

	return authenticator.GetSecurity().Login("email", *login.GetEmail(), *login.GetPassword())
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
