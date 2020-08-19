package go_saas_http_basic

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/security"
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
		var teamUser bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if c.GetHeader("Team") == "" {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(errors.New("no team header"), nil))
			return
		}

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamUser, err = http.GetSecurity().IsTeamUser(uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamUser {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamMemberErr, nil))
			return
		}

		c.Next()
	}
}

func (http *Http) TeamRoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamOwner bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if c.GetHeader("Team") == "" {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(errors.New("no team header"), nil))
			return
		}

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamOwner, err = http.GetSecurity().IsTeamRole(role, uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamOwner {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamOwnerError, nil))
			return
		}

		c.Next()
	}
}

func (http *Http) TeamCreatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamCreator bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if c.GetHeader("Team") == "" {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(errors.New("no team header"), nil))
			return
		}

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamCreator, err = http.GetSecurity().IsTeamCreator(uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamCreator {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamOwnerError, nil))
			return
		}

		c.Next()
	}
}
