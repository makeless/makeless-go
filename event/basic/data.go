package go_saas_event_basic

import "sync"

type Data struct {
	Id   string      `json:"id"`
	Data interface{} `json:"data"`
	*sync.RWMutex
}
