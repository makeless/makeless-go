package makeless_go_http_basic

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/security"
	"gorm.io/gorm"
	h "net/http"
	"strconv"
	"sync"
)

func (http *Http) CorsMiddleware(Origins []string, OriginsFunc func(origin string) bool, AllowHeaders []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = Origins
	config.AllowOriginFunc = OriginsFunc
	config.AllowCredentials = true
	config.AddAllowHeaders(AllowHeaders...)

	return cors.New(config)
}

func (http *Http) EmailVerificationMiddleware(enabled bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var userId = http.GetAuthenticator().GetAuthUserId(c)
		var emailVerification = http.GetAuthenticator().GetAuthEmailVerification(c)
		var user = &makeless_go_model.User{
			RWMutex: new(sync.RWMutex),
		}

		if !enabled || emailVerification {
			c.Next()
			return
		}

		if user, err = http.GetDatabase().GetUserByField(http.GetDatabase().GetConnection().WithContext(c), user, "id", fmt.Sprintf("%d", userId)); err != nil {
			switch errors.Is(err, gorm.ErrRecordNotFound) {
			case true:
				c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(err, nil))
			default:
				c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			}
			return
		}

		if user.GetEmailVerification() != nil {
			user.GetEmailVerification().RWMutex = new(sync.RWMutex)
		}

		if user.GetEmailVerification() != nil && *user.GetEmailVerification().GetVerified() {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(makeless_go_security.NoEmailVerification, nil))
	}
}

func (http *Http) TeamUserMiddleware() gin.HandlerFunc {
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

		if teamUser, err = http.GetSecurity().IsTeamUser(http.GetSecurity().GetDatabase().GetConnection(), uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamUser {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(makeless_go_security.NoTeamUserErr, nil))
			return
		}

		c.Next()
	}
}

func (http *Http) TeamRoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamRole bool
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

		if teamRole, err = http.GetSecurity().IsTeamRole(http.GetSecurity().GetDatabase().GetConnection(), role, uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamRole {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(makeless_go_security.NoTeamRoleError, nil))
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

		if teamCreator, err = http.GetSecurity().IsTeamCreator(http.GetSecurity().GetDatabase().GetConnection(), uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamCreator {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(makeless_go_security.NoTeamCreatorError, nil))
			return
		}

		c.Next()
	}
}

func (http *Http) NotTeamCreatorMiddleware() gin.HandlerFunc {
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

		if teamCreator, err = http.GetSecurity().IsTeamCreator(http.GetSecurity().GetDatabase().GetConnection(), uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if teamCreator {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(makeless_go_security.NoTeamCreatorError, nil))
			return
		}

		c.Next()
	}
}
