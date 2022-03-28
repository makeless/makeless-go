package makeless_go_mail

import (
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/model"
)

type EmailVerificationMail interface {
	Create(config makeless_go_config.Config, user *makeless_go_model.User, locale string) (Mail, error)
}
