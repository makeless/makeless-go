package saas_api

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/database/basic"
	"github.com/go-saas/go-saas/event"
	"github.com/go-saas/go-saas/logger"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/tls"
	"sync"
)

type Api struct {
	engine         *gin.Engine
	authMiddleware *jwt.GinJWTMiddleware
	handlers       []func(api *Api)

	Logger   go_saas_logger.Logger
	Event    go_saas_event.Event
	Security go_saas_security.Security
	Database *go_saas_basic_database.saas_database
	Jwt      go_saas_jwt.Jwt
	Tls      go_saas_tls.Tls
	Origins  []string
	Port     string
	Mode     string
	*sync.RWMutex
}

func (api *Api) getOrigins() []string {
	api.RLock()
	defer api.RUnlock()

	return api.Origins
}

func (api *Api) getMode() string {
	api.RLock()
	defer api.RUnlock()

	return api.Mode
}

func (api *Api) getPort() string {
	api.RLock()
	defer api.RUnlock()

	return api.Port
}

func (api *Api) getJwt() go_saas_jwt.Jwt {
	api.RLock()
	defer api.RUnlock()

	return api.Jwt
}

func (api *Api) getTls() go_saas_tls.Tls {
	api.RLock()
	defer api.RUnlock()

	return api.Tls
}

func (api *Api) getHandlers() []func(api *Api) {
	api.Lock()
	defer api.Unlock()

	return api.handlers
}

func (api *Api) createEngine() {
	api.Lock()
	defer api.Unlock()

	api.engine = gin.Default()
}

func (api *Api) setAuthMiddleware(jwtMiddleware *jwt.GinJWTMiddleware) {
	api.Lock()
	defer api.Unlock()

	api.authMiddleware = jwtMiddleware
}

func (api *Api) createAuthMiddleware() error {
	jwtMiddleware, err := api.jwtMiddleware()

	if err != nil {
		return err
	}

	api.setAuthMiddleware(jwtMiddleware)
	return nil
}

func (api *Api) GetAuthMiddleware() *jwt.GinJWTMiddleware {
	api.RLock()
	defer api.RUnlock()

	return api.authMiddleware
}

func (api *Api) GetLogger() go_saas_logger.Logger {
	api.RLock()
	defer api.RUnlock()

	return api.Logger
}

func (api *Api) GetEvent() go_saas_event.Event {
	api.RLock()
	defer api.RUnlock()

	return api.Event
}

func (api *Api) GetSecurity() go_saas_security.Security {
	api.RLock()
	defer api.RUnlock()

	return api.Security
}

func (api *Api) GetDatabase() *go_saas_basic_database.saas_database {
	api.RLock()
	defer api.RUnlock()

	return api.Database
}

func (api *Api) GetEngine() *gin.Engine {
	api.RLock()
	defer api.RUnlock()

	return api.engine
}

func (api *Api) Start() error {
	gin.SetMode(api.getMode())
	api.createEngine()

	// global middleware
	api.GetEngine().Use(api.cors())
	api.GetEngine().Use(gin.Recovery())

	// auth middleware
	if err := api.createAuthMiddleware(); err != nil {
		return err
	}

	// routes
	apiGroup := api.GetEngine().Group("/api")
	{
		// basic
		apiGroup.GET("/ok", api.ok) // DONE

		// auth
		apiGroup.POST("/login", api.GetAuthMiddleware().LoginHandler) // DONE
		apiGroup.POST("/register", api.register)                      // DONE

		authGroup := apiGroup.Group("/auth")
		authGroup.Use(api.GetAuthMiddleware().MiddlewareFunc())
		{
			// basic
			authGroup.GET("/events", api.events) // DONE

			// auth
			authGroup.GET("/user", api.user)                                        // DONE
			authGroup.PATCH("/password", api.updatePassword)                        // DONE
			authGroup.GET("/refresh-token", api.GetAuthMiddleware().RefreshHandler) // DONE
			authGroup.GET("/logout", api.GetAuthMiddleware().LogoutHandler)         // DONE

			// settings -> profile
			authGroup.PATCH("/profile", api.updateProfile) // DONE

			// settings -> team
			authGroup.POST("/team", api.createTeam)
			authGroup.DELETE("/team", api.deleteTeam)

			// settings -> tokens
			authGroup.GET("/token", api.tokens)         // DONE
			authGroup.POST("/token", api.createToken)   // DONE
			authGroup.DELETE("/token", api.deleteToken) // DONE

			// team
			teamGroup := authGroup.Group("/team")
			{
				// settings -> profile
				teamGroup.PATCH("/profile", api.updateProfileTeam)

				// settings -> members
				teamGroup.GET("/member", api.membersTeam)
				teamGroup.DELETE("/member", api.removeMemberTeam)

				// settings -> tokens
				teamGroup.GET("/token", api.tokensTeam)
				teamGroup.POST("/token", api.createTokenTeam)
				teamGroup.DELETE("/token", api.deleteTokenTeam)

				// utils
				teamGroup.GET("/user", api.usersTeam)
			}
		}
	}

	// run extend handlers
	for _, handler := range api.getHandlers() {
		handler(api)
	}

	if api.getTls() != nil {
		return api.GetEngine().RunTLS(":"+api.getPort(), api.getTls().GetCertPath(), api.getTls().GetKeyPath())
	}

	return api.GetEngine().Run(":" + api.getPort())
}
