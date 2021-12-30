package makeless_go_event

import (
	"sync"
)

type Hub interface {
	GetList() *sync.Map
	GetUser(userId uint) *sync.Map
	GetClient(userId uint, clientId string) chan EventData
	NewClient(userId uint, clientId string)
	DeleteClient(userId uint, clientId string)
}
