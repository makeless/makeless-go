package go_saas_database_basic

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) GetEmailVerificationByField(connection *gorm.DB, emailVerification *go_saas_model.EmailVerification, field string, value string) (*go_saas_model.EmailVerification, error) {
	return emailVerification, connection.
		Where(
			fmt.Sprintf("email_verifications.%s = ? AND email_verifications.verified = ?", field),
			value,
			false,
		).
		Find(emailVerification).
		Error
}

func (database *Database) VerifyEmailVerification(connection *gorm.DB, emailVerification *go_saas_model.EmailVerification) (*go_saas_model.EmailVerification, error) {
	return emailVerification, connection.
		Model(emailVerification).
		Update(map[string]interface{}{
			"verified": true,
		}).
		Error
}
