package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/database"
	"github.com/loeffel-io/go-saas/logger"
	"sync"
)

type Api struct {
	engine   *gin.Engine
	handlers []func(api *Api)

	Logger   saas_logger.Logger
	Database *saas_database.Database
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

func (api *Api) GetLogger() saas_logger.Logger {
	api.RLock()
	defer api.RUnlock()

	return api.Logger
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

	api.GetEngine().Use(api.cors())
	api.GetEngine().Use(gin.Recovery())

	api.GetEngine().GET("/ok", api.ok)

	for _, handler := range api.getHandlers() {
		handler(api)
	}

	if api.Tls != nil {
		return api.GetEngine().RunTLS(":"+api.getPort(), api.getTls().getCertPath(), api.getTls().getKeyPath())
	}

	return api.GetEngine().Run(":" + api.getPort())
}
