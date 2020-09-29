package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"sync"
)

func (saas *Saas) register(http go_saas_http.Http) error {
	http.GetRouter().POST("/api/register", func(c *gin.Context) {
		var err error
		var token string
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

		var user = &go_saas_model.User{
			Name:     register.GetName(),
			Password: register.GetPassword(),
			Email:    register.GetEmail(),
			EmailVerification: &go_saas_model.EmailVerification{
				Token:    &token,
				Verified: &verified,
				RWMutex:  new(sync.RWMutex),
			},
			RWMutex: new(sync.RWMutex),
		}

		if user, err = http.GetSecurity().Register(http.GetSecurity().GetDatabase().GetConnection(), user); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		c.JSON(h.StatusOK, http.Response(nil, user))
	})

	return nil
}
