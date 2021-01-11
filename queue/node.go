package makeless_go_queue

type Node interface {
	GetData() interface{}

	SetNext(next Node)
	GetNext() Node
}
