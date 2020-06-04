package go_saas_logger

type Logger interface {
	Fatal(err error)
	Print(msg string)
	Println(msg string)
}
