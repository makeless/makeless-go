package go_saas_basic_jwt

import "sync"

type Jwt struct {
	Key string
	*sync.RWMutex
}

func (jwt *Jwt) GetId() string {
	return "id"
}

func (jwt *Jwt) GetKey() string {
	jwt.RLock()
	defer jwt.RUnlock()

	return jwt.Key
}