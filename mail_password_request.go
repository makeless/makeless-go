package go_saas

import (
	"github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/mailer/basic"
	"github.com/go-saas/go-saas/model"
	"sync"
)

func (saas *Saas) mailPasswordRequest(data map[string]interface{}) (go_saas_mailer.Mail, error) {
	var passwordRequest = data["passwordRequest"].(*go_saas_model.PasswordRequest)

	return &go_saas_mailer_basic.Mail{
		To:      []string{*passwordRequest.GetEmail()},
		From:    "test@finsmart.io",
		Message: []byte(*passwordRequest.GetToken()),
		RWMutex: new(sync.RWMutex),
	}, nil
}
