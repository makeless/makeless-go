package makeless_go_event

type EventData interface {
	GetChannel() string
	GetId() string
	GetData() interface{}
}
