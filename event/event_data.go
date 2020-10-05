package makeless_go_event

type EventData interface {
	GetId() string
	GetData() interface{}
}
