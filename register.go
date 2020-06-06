package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	h "net/http"
	"sync"
)

func (saas *Saas) register(http go_saas_http.Http) error {
	http.GetRouter().POST("/api/register", func(c *gin.Context) {
		var err error
		var user = &go_saas_model.User{
			RWMutex: new(sync.RWMutex),
		}

		if err := c.ShouldBind(user); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if user, err = http.GetSecurity().Register(user); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		c.JSON(h.StatusOK, http.Response(nil, user))
	})

	return nil
}
