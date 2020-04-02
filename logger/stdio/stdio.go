package saas_logger_stdio

import "log"

type Stdio struct {
}

func (stdio *Stdio) Fatal(err error) {
	log.Fatal(err)
}

func (stdio *Stdio) Print(msg string) {
	log.Print(msg)
}

func (stdio *Stdio) Println(msg string) {
	log.Println(msg)
}
