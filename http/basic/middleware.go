package go_saas_basic_http

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	h "net/http"
	"strconv"
)

func (http *Http) CorsMiddleware(Origins []string, AllowHeaders []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = Origins
	config.AllowCredentials = true
	config.AddAllowHeaders(AllowHeaders...)

	return cors.New(config)
}

func (http *Http) TeamMemberMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamMember bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamMember, err = http.GetSecurity().IsTeamMember(uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamMember {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(errors.New("no team member"), nil))
			return
		}

		c.Next()
	}
}
