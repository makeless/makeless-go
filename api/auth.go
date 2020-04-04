package saas_api

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/loeffel-io/go-saas/model"
	"net/http"
	"sync"
	"time"
)

func (api *Api) jwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "auth",
		Key:         []byte(api.getJwt().getKey()),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: api.getJwt().getKey(),
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*saas_model.User); ok {
				return jwt.MapClaims{
					api.getJwt().getId(): v.ID,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			userId, _ := api.GetUserId(c)

			return &saas_model.User{
				Model: gorm.Model{
					ID: userId,
				},
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return api.GetSecurity().Login(c.PostForm("username"), c.PostForm("password"))
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, api.Response(message, nil))
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
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
	var user = &saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if err := c.Bind(&user); err != nil {
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
