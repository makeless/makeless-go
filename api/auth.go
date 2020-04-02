package saas_api

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/loeffel-io/go-saas/model"
	"golang.org/x/crypto/bcrypt"
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
			return &saas_model.User{
				Model: gorm.Model{
					ID: api.GetUserId(c),
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

func (api *Api) GetUserId(c *gin.Context) uint {
	return uint(jwt.ExtractClaims(c)[api.getJwt().getId()].(float64))
}

func (api *Api) register(c *gin.Context) {
	var user = saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	// bind json
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	// crypt password
	bcrypted, err := bcrypt.GenerateFromPassword([]byte(*user.GetPassword()), 14)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	// update user password
	user.SetPassword(string(bcrypted))

	// create user
	if err := api.GetDatabase().GetConnection().Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, user))
}
