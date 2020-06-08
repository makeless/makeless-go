package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/imdario/mergo"
	h "net/http"
	"sync"
)

func (saas *Saas) tokens(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var tokens []*go_saas_model.Token
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if tokens, err = http.GetDatabase().GetTokens(user, tokens); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, tokens))
		},
	)

	return nil
}

func (saas *Saas) createToken(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var token = new(go_saas_model.Token)

			if err := c.ShouldBind(token); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(token, &go_saas_model.Token{
				UserId:  &userId,
				User:    nil,
				TeamId:  nil,
				Team:    nil,
				RWMutex: new(sync.RWMutex),
			}, mergo.WithOverride); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if token, err = http.GetDatabase().CreateToken(token); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, token))
		},
	)

	return nil
}

func (saas *Saas) deleteToken(http go_saas_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var token = new(go_saas_model.Token)

			if err := c.ShouldBind(token); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			if err = mergo.Merge(token, &go_saas_model.Token{
				UserId:  &userId,
				TeamId:  nil,
				RWMutex: new(sync.RWMutex),
			}, mergo.WithOverride); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			if err = http.GetDatabase().DeleteToken(token); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
