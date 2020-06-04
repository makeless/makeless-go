package go_saas_event

import (
	"github.com/gin-contrib/sse"
)

type Event interface {
	NewClientId() uint
	GetHub() Hub
	Subscribe(userId uint, clientId uint)
	Unsubscribe(userId uint, clientId uint)
	Trigger(userId uint, channel string, id string, data interface{})
	Broadcast(channel string, id string, data interface{})
	Listen(userId uint, clientId uint) chan sse.Event
}
