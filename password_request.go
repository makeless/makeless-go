package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"sync"
	"time"
)

func (saas *Saas) passwordRequest(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/password-request",
		func(c *gin.Context) {
			var err error
			var userExists bool
			var tmpPasswordRequest = &_struct.PasswordRequest{
				RWMutex: new(sync.RWMutex),
			}

			if err = c.ShouldBind(tmpPasswordRequest); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if userExists, err = http.GetSecurity().UserExists("email", *tmpPasswordRequest.GetEmail()); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if !userExists {
				c.AbortWithStatusJSON(h.StatusOK, http.Response(nil, nil))
				return
			}

			var test = "123"
			var passwordRequest = &go_saas_model.PasswordRequest{
				Email:   tmpPasswordRequest.GetEmail(),
				Token:   &test,
				Expire:  time.Now().Add(time.Hour * 1),
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().CreatePasswordRequest(passwordRequest); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
