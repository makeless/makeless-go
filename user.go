package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	h "net/http"
	"sync"
)

func (saas *Saas) user(http go_saas_http.Http) error {
	http.GetRouter().GET("/api/auth/user", http.GetAuthenticator().GetMiddleware().MiddlewareFunc(), func(c *gin.Context) {
		userId := http.GetAuthenticator().GetAuthUserId(c)

		var err error
		var user = &go_saas_model.User{
			Model:   go_saas_model.Model{Id: userId},
			RWMutex: new(sync.RWMutex),
		}

		if user, err = http.GetDatabase().GetUser(user); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		c.JSON(h.StatusOK, http.Response(nil, user))
	})

	return nil
}
