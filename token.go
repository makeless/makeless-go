package go_saas

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	h "net/http"
	"sync"
)

func (saas *Saas) tokens(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var tokens []*go_saas_model.Token
			var user = &go_saas_model.User{
				Model:   go_saas_model.Model{Id: userId},
				RWMutex: new(sync.RWMutex),
			}

			if tokens, err = http.GetDatabase().GetTokens(http.GetDatabase().GetConnection(), user, tokens); err != nil {
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
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var tokenCreate = &_struct.TokenCreate{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(tokenCreate); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var token = &go_saas_model.Token{
				Note:    tokenCreate.GetNote(),
				Token:   tokenCreate.GetToken(),
				UserId:  &userId,
				RWMutex: new(sync.RWMutex),
			}

			if token, err = http.GetDatabase().CreateToken(http.GetDatabase().GetConnection(), token); err != nil {
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
		http.EmailVerificationMiddleware(saas.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			var err error
			var userId = http.GetAuthenticator().GetAuthUserId(c)
			var tokenDelete = &_struct.TokenDelete{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(tokenDelete); err != nil {
				c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
				return
			}

			var token = &go_saas_model.Token{
				Model:   go_saas_model.Model{Id: *tokenDelete.GetId()},
				UserId:  &userId,
				RWMutex: new(sync.RWMutex),
			}

			if err = http.GetDatabase().DeleteToken(http.GetDatabase().GetConnection(), token); err != nil {
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
				return
			}

			c.JSON(h.StatusOK, http.Response(nil, nil))
		},
	)

	return nil
}
