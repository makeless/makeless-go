package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type EmailVerification interface {
	GetEmailVerificationByField(connection *gorm.DB, emailVerification *go_saas_model.EmailVerification, field string, value string) (*go_saas_model.EmailVerification, error)
	VerifyEmailVerification(connection *gorm.DB, emailVerification *go_saas_model.EmailVerification) (*go_saas_model.EmailVerification, error)
}
