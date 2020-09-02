package go_saas

import (
	"fmt"
	"github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/mailer/basic"
	"github.com/go-saas/go-saas/model"
	"github.com/matcornic/hermes/v2"
	"sync"
)

func (saas *Saas) mailPasswordRequest(data map[string]interface{}) (go_saas_mailer.Mail, error) {
	var err error
	var message, messageHtml string
	var passwordRequest = data["passwordRequest"].(*go_saas_model.PasswordRequest)
	var messages = map[string]map[string]string{
		"en": {
			"subject":     "Password reset",
			"instruction": "Click here to reset your password",
			"Button":      "Reset password",
		},
	}

	config := hermes.Hermes{
		Product: hermes.Product{
			Name: saas.GetConfig().GetConfiguration().GetName(),
			Link: saas.GetConfig().GetConfiguration().GetHost(),
			Logo: saas.GetConfig().GetConfiguration().GetLogo(),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Actions: []hermes.Action{
				{
					Instructions: messages[saas.GetConfig().GetConfiguration().GetLocale()]["instruction"],
					Button: hermes.Button{
						Text: messages[saas.GetConfig().GetConfiguration().GetLocale()]["button"],
						Link: fmt.Sprintf(
							"%s/password-reset?token=%s",
							saas.GetConfig().GetConfiguration().GetHost(),
							*passwordRequest.GetToken(),
						),
					},
				},
			},
		},
	}

	if message, err = config.GeneratePlainText(email); err != nil {
		return nil, err
	}

	if messageHtml, err = config.GenerateHTML(email); err != nil {
		return nil, err
	}

	return &go_saas_mailer_basic.Mail{
		To:          []string{*passwordRequest.GetEmail()},
		From:        saas.GetConfig().GetConfiguration().GetMail(),
		Subject:     messages[saas.GetConfig().GetConfiguration().GetLocale()]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
