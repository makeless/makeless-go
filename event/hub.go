package go_saas_event

import (
	"sync"
)

type Hub interface {
	GetList() *sync.Map
	GetUser(userId uint) *sync.Map
	GetClient(userId uint, clientId uint) chan Channel
	NewClient(userId uint, clientId uint)
	DeleteClient(userId uint, clientId uint)
}
