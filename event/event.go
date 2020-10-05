package makeless_go_event

import "github.com/gin-contrib/sse"

type Event interface {
	Init() error
	NewClientId() uint
	GetHub() Hub
	Subscribe(userId uint, clientId uint)
	Unsubscribe(userId uint, clientId uint)
	Trigger(userId uint, channel string, id string, data interface{}) error
	TriggerError(err error)
	Broadcast(channel string, id string, data interface{}) error
	Listen(userId uint, clientId uint) chan sse.Event
	ListenError() chan error
}
