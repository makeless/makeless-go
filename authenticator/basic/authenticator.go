package go_saas_basic_authenticator

import "sync"

type Authenticator struct {
	*sync.RWMutex
}
