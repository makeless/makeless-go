package makeless_go_model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deletedAt"`
}
