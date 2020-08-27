package go_saas_event

type EventData interface {
	GetId() string
	GetData() interface{}
}
