package filter

import "gorm.io/gorm"

type Filter interface {
	GetQuery(query *gorm.DB) *gorm.DB
	Empty() bool
}
