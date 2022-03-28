package makeless_go_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        uuid.UUID       `gorm:"primary_key;type:uuid"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `sql:"index" json:"deletedAt"`
}
