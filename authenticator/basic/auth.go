package go_saas_basic_authenticator

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func (authenticator *Authenticator) GetAuthUserId(c *gin.Context) uint {
	claims := jwt.ExtractClaims(c)
	return uint(claims[authenticator.GetIdentityKey()].(float64))
}
