package go_saas_event

type Event interface {
	Init() error
	NewClientId() uint
	GetHub() Hub
	Subscribe(userId uint, clientId uint)
	Unsubscribe(userId uint, clientId uint)
	Trigger(userId uint, channel string, id string, data interface{}) error
	Broadcast(channel string, id string, data interface{}) error
	Listen(userId uint, clientId uint) chan Channel
}
