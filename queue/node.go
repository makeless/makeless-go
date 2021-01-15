package makeless_go_queue

type Node interface {
	GetData() []byte

	SetNext(next Node)
	GetNext() Node
}
