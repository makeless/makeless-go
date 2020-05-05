package saas_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/helper_model"
	"net/http"
	"sync"
)

func (api *Api) updatePassword(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var passwordReset = &saas_helper_model.PasswordReset{
		RWMutex: new(sync.RWMutex),
	}

	if err = c.ShouldBind(passwordReset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Response(err.Error(), nil))
		return
	}

	if _, err = api.GetSecurity().Login("id", fmt.Sprintf("%d", userId), *passwordReset.GetPassword()); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	var bcrypted []byte

	if bcrypted, err = api.GetSecurity().EncryptPassword(*passwordReset.GetNewPassword()); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	if err = api.GetDatabase().UpdatePassword(userId, string(bcrypted)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Response(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, api.Response(nil, nil))
}
