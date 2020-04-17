package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/model"
	"net/http"
	"sync"
)

func (api *Api) tokens(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var tokens []*saas_model.Token

	if tokens, err = api.GetDatabase().GetTokens(userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, tokens))
}

func (api *Api) createToken(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var token = &saas_model.Token{
		UserId:  &userId,
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if token, err = api.GetDatabase().CreateToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, token))
}

func (api *Api) deleteToken(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var token = &saas_model.Token{
		UserId:  &userId,
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().DeleteToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, nil))
}
