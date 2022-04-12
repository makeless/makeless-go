package makeless_go_order

import "gorm.io/gorm"

type Order interface {
	GetQuery(query *gorm.DB) *gorm.DB
}
