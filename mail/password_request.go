package makeless_go_mail

import (
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/model"
)

type PasswordRequestMail interface {
	Create(config makeless_go_config.Config, passwordRequest *makeless_go_model.PasswordRequest, locale string) (Mail, error)
}
