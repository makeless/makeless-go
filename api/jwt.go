package saas_api

import (
	"sync"
)

type Jwt struct {
	Key string
	*sync.RWMutex
}

func (jwt *Jwt) getId() string {
	return "id"
}

func (jwt *Jwt) getKey() string {
	jwt.RLock()
	defer jwt.RUnlock()

	return jwt.Key
}
