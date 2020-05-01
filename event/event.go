package saas_event

import (
	"github.com/gin-contrib/sse"
)

type Event interface {
	NewClientId() uint
	GetHub() Hub
	Subscribe(userId uint, clientId uint)
	Unsubscribe(userId uint, clientId uint)
	Emit(userId uint, data sse.Event)
	Listen(userId uint, clientId uint) chan sse.Event
}
