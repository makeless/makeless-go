package makeless_go_database

import (
	"github.com/jinzhu/gorm"
	"github.com/makeless/makeless-go/model"
)

type EmailVerification interface {
	GetEmailVerificationByField(connection *gorm.DB, emailVerification *makeless_go_model.EmailVerification, field string, value string) (*makeless_go_model.EmailVerification, error)
	VerifyEmailVerification(connection *gorm.DB, emailVerification *makeless_go_model.EmailVerification) (*makeless_go_model.EmailVerification, error)
}
