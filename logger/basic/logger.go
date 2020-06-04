package go_saas_basic_logger

import "log"

type Logger struct {
}

func (logger *Logger) Fatal(err error) {
	log.Fatal(err)
}

func (logger *Logger) Print(msg string) {
	log.Print(msg)
}

func (logger *Logger) Println(msg string) {
	log.Println(msg)
}
