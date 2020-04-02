package saas_model

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type User struct {
	gorm.Model
	FirstName *string `gorm:"not null" json:"firstName" binding:"required"`
	LastName  *string `gorm:"not null" json:"lastName" binding:"required"`
	Username  *string `gorm:"unique;not null" json:"username" binding:"required"`
	Password  *string `gorm:"not null" json:"password,omitempty" binding:"required"`
	Email     *string `gorm:"unique;not null" json:"email" binding:"required"`

	Tokens []*Token `json:"tokens"`

	*sync.RWMutex `json:"-"`
}
