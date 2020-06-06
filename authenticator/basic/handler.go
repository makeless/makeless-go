package go_saas_basic_authenticator

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/model"
	"sync"
)

func (authenticator *Authenticator) PayloadHandler(data interface{}) jwt.MapClaims {
	if user, ok := data.(*go_saas_model.User); ok {
		return jwt.MapClaims{
			authenticator.GetIdentityKey(): user.Id,
		}
	}

	return jwt.MapClaims{}
}

func (authenticator *Authenticator) IdentityHandler(c *gin.Context) interface{} {
	return &go_saas_model.User{
		Model: go_saas_model.Model{
			Id: authenticator.GetAuthUserId(c),
		},
	}
}

func (authenticator *Authenticator) AuthenticatorHandler(c *gin.Context) (interface{}, error) {
	var login = &go_saas_model.Login{
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(&login); err != nil {
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
