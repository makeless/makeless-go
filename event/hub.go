package makeless_go_event

import (
	"github.com/gin-contrib/sse"
	"sync"
)

type Hub interface {
	GetList() *sync.Map
	GetUser(userId uint) *sync.Map
	GetClient(userId uint, clientId string) chan sse.Event
	NewClient(userId uint, clientId string)
	DeleteClient(userId uint, clientId string)
}
