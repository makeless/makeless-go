package saas_api

import "sync"

type teamHeader struct {
	TeamId uint `header:"Team" binding:"required"`
	*sync.RWMutex
}
