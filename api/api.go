package saas_api

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/database"
	"github.com/loeffel-io/go-saas/logger"
	"github.com/loeffel-io/go-saas/security"
	"sync"
)

type Api struct {
	engine         *gin.Engine
	authMiddleware *jwt.GinJWTMiddleware
	handlers       []func(api *Api)

	Logger   saas_logger.Logger
	Security saas_security.Security
	Database *saas_database.Database
	Jwt      *Jwt
	Tls      *Tls
	Port     string
	Mode     string
	*sync.RWMutex
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

func (api *Api) getJwt() *Jwt {
	api.RLock()
	defer api.RUnlock()

	return api.Jwt
}

func (api *Api) getTls() *Tls {
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

func (api *Api) GetLogger() saas_logger.Logger {
	api.RLock()
	defer api.RUnlock()

	return api.Logger
}

func (api *Api) GetSecurity() saas_security.Security {
	api.RLock()
	defer api.RUnlock()

	return api.Security
}

func (api *Api) GetDatabase() *saas_database.Database {
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
		apiGroup.GET("/ok", api.ok)

		apiGroup.POST("/login", api.GetAuthMiddleware().LoginHandler)
		apiGroup.POST("/register", api.register)

		authGroup := apiGroup.Group("/auth")
		authGroup.Use(api.GetAuthMiddleware().MiddlewareFunc())
		{
			authGroup.GET("/refresh_token", api.GetAuthMiddleware().RefreshHandler)
			authGroup.GET("/logout", api.GetAuthMiddleware().LogoutHandler)

			authGroup.GET("/token", api.tokens)
		}
	}

	// run extend handlers
	for _, handler := range api.getHandlers() {
		handler(api)
	}

	if api.Tls != nil {
		return api.GetEngine().RunTLS(":"+api.getPort(), api.getTls().getCertPath(), api.getTls().getKeyPath())
	}

	return api.GetEngine().Run(":" + api.getPort())
}
