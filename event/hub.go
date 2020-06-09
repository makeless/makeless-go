package go_saas_event

import "github.com/gin-contrib/sse"

type Hub interface {
	GetList() map[uint]map[uint]chan sse.Event
	GetUser(userId uint) map[uint]chan sse.Event
	GetClient(userId uint, clientId uint) chan sse.Event
	NewClient(userId uint, clientId uint)
	DeleteClient(userId uint, clientId uint)
}
