package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/security"
	"net/http"
	"sync"
)

func (api *Api) tokens(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var tokens []*go_saas_model.Token

	if tokens, err = api.GetDatabase().GetTokens(tokens, &userId); err != nil {
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

	var token *go_saas_model.Token

	if err := c.ShouldBind(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if token, err = api.GetDatabase().CreateToken(token, &userId); err != nil {
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

	var token *go_saas_model.Token

	if err := c.ShouldBind(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().DeleteToken(token, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, nil))
}

func (api *Api) tokensTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var header = &teamHeader{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBindHeader(header); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	var tokens []*go_saas_model.Token

	if tokens, err = api.GetDatabase().GetTokensTeam(tokens, header.TeamId, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, tokens))
}

func (api *Api) deleteTokenTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var header = &teamHeader{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBindHeader(header); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	var token *go_saas_model.Token

	if err := c.ShouldBind(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().DeleteTokenTeam(token, &header.TeamId, &userId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, nil))
}

func (api *Api) createTokenTeam(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var header = &teamHeader{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBindHeader(header); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	isTeamOwner, err := api.GetSecurity().IsTeamOwner(header.TeamId, userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if !isTeamOwner {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(go_saas_security.NoTeamOwnerError, nil))
		return
	}

	var token = &go_saas_model.Token{
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if token, err = api.GetDatabase().CreateTokenTeam(token, &header.TeamId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, token))
}
