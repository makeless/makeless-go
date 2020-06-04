package saas_api

import (
	"net/http"
	"sync"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/model"
)

func (api *Api) jwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "auth",
		Key:         []byte(api.getJwt().getKey()),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: api.getJwt().getKey(),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*go_saas_model.User); ok {
				return jwt.MapClaims{
					api.getJwt().getId(): v.Id,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			userId, _ := api.GetUserId(c)

			return &go_saas_model.User{
				Model: go_saas_model.Model{
					Id: userId,
				},
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login = &go_saas_model.Login{
				RWMutex: new(sync.RWMutex),
			}

			if err := c.ShouldBind(&login); err != nil {
				return nil, err
			}

			return api.GetSecurity().Login("email", *login.GetEmail(), *login.GetPassword())
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, api.Response(message, nil))
		},
		TimeFunc:       time.Now,
		SendCookie:     true,
		SecureCookie:   false, //non HTTPS dev environments
		CookieHTTPOnly: true,
		CookieName:     "jwt",
		TokenLookup:    "cookie:jwt",
	})
}

// GetUserId returns the current jwt user id if exists
func (api *Api) GetUserId(c *gin.Context) (uint, error) {
	claims := jwt.ExtractClaims(c)

	userId, exists := claims[api.getJwt().getId()]

	if !exists {
		return 0, jwt.ErrFailedAuthentication
	}

	return uint(userId.(float64)), nil
}

func (api *Api) register(c *gin.Context) {
	var user = &go_saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	user, err := api.GetSecurity().Register(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, user))
}
