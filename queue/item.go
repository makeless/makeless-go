package makeless_go_queue

type Item interface {
	GetData() interface{}
	GetAsync() bool
}
