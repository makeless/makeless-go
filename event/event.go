package makeless_go_event

import "github.com/gin-contrib/sse"

type Event interface {
	Init() error
	NewClientId() string
	GetHub() Hub
	Subscribe(userId uint, clientId string)
	Unsubscribe(userId uint, clientId string)
	Trigger(userId uint, channel string, id string, data interface{}) error
	TriggerError(err error)
	Broadcast(channel string, id string, data interface{}) error
	Listen(userId uint, clientId string) chan sse.Event
	ListenError() chan error
}
