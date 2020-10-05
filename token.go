package makeless_go

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	h "net/http"
	"sync"
)

func (makeless *Makeless) tokens(http makeless_go_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			var err error
			var tokens []*makeless_go_model.Token
			var user = &makeless_go_model.User{
				Model:   makeless_go_model.Model{Id: userId},
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

func (makeless *Makeless) createToken(http makeless_go_http.Http) error {
	http.GetRouter().POST(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
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

			var token = &makeless_go_model.Token{
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

func (makeless *Makeless) deleteToken(http makeless_go_http.Http) error {
	http.GetRouter().DELETE(
		"/api/auth/token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.EmailVerificationMiddleware(makeless.GetConfig().GetConfiguration().GetEmailVerification()),
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

			var token = &makeless_go_model.Token{
				Model:   makeless_go_model.Model{Id: *tokenDelete.GetId()},
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
