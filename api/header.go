package saas_api

import "sync"

type teamHeader struct {
	TeamId uint `header:"Team"`
	*sync.RWMutex
}
