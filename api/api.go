package saas_api

import (
	"github.com/loeffel-io/go-saas/database"
	"sync"
)

type Api struct {
	Database *saas_database.Database
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
