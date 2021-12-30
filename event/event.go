package makeless_go_event

type Event interface {
	Init() error
	GetName() string
	NewClientId() string
	GetHub() Hub
	Subscribe(userId uint, clientId string)
	Unsubscribe(userId uint, clientId string)
	Trigger(userId uint, channel string, id string, data interface{}) error
	TriggerError(err error)
	Broadcast(channel string, id string, data interface{}) error
	Listen(userId uint, clientId string) chan EventData
	ListenError() chan error
}
