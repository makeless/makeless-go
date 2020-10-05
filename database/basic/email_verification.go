package makeless_go_database_basic

import (
	"fmt"
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

func (database *Database) GetEmailVerificationByField(connection *gorm.DB, emailVerification *makeless_go_model.EmailVerification, field string, value string) (*makeless_go_model.EmailVerification, error) {
	return emailVerification, connection.
		Where(
			fmt.Sprintf("email_verifications.%s = ? AND email_verifications.verified = ?", field),
			value,
			false,
		).
		Find(emailVerification).
		Error
}

func (database *Database) VerifyEmailVerification(connection *gorm.DB, emailVerification *makeless_go_model.EmailVerification) (*makeless_go_model.EmailVerification, error) {
	return emailVerification, connection.
		Model(emailVerification).
		Updates(map[string]interface{}{
			"verified": true,
		}).
		Error
}
