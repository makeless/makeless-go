package saas_api

import (
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/go-saas/database"
	"sync"
)

type Api struct {
	engine   *gin.Engine
	handlers []func(engine *gin.Engine)

	Database *saas_database.Database
	Tls      *Tls
	Port     string
	Mode     string
	*sync.RWMutex
}

func (api *Api) getDatabase() *saas_database.Database {
	api.RLock()
	defer api.RUnlock()

	return api.Database
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

func (api *Api) getHandlers() []func(engine *gin.Engine) {
	api.Lock()
	defer api.Unlock()

	return api.handlers
}

func (api *Api) createEngine() {
	api.Lock()
	defer api.Unlock()

	api.engine = gin.Default()
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
		handler(api.GetEngine())
	}

	if api.Tls != nil {
		return api.GetEngine().RunTLS(":"+api.getPort(), api.getTls().getCertPath(), api.getTls().getKeyPath())
	}

	return api.GetEngine().Run(":" + api.getPort())
}
