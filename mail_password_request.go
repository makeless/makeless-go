package go_saas

import (
	"fmt"
	"sync"

	"github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/mailer/basic"
	"github.com/go-saas/go-saas/model"
	"github.com/matcornic/hermes/v2"
)

func (saas *Saas) mailPasswordRequest(data map[string]interface{}) (go_saas_mailer.Mail, error) {
	var err error
	var message, messageHtml string
	var passwordRequest = data["passwordRequest"].(*go_saas_model.PasswordRequest)
	var messages = map[string]map[string]string{
		"en": {
			"subject":     "Reset your password",
			"instruction": "to reset your password, please click here:",
			"button":      "Reset password",
		},
	}

	config := hermes.Hermes{
		Product: hermes.Product{
			Name:      saas.GetConfig().GetConfiguration().GetMail().GetName(),
			Link:      saas.GetConfig().GetConfiguration().GetMail().GetLink(),
			Logo:      saas.GetConfig().GetConfiguration().GetMail().GetLogo(),
			Copyright: saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetCopyright(),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Greeting:  saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetGreeting(),
			Signature: saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: messages[saas.GetConfig().GetConfiguration().GetLocale()]["instruction"],
					Button: hermes.Button{
						Text: messages[saas.GetConfig().GetConfiguration().GetLocale()]["button"],
						Link: fmt.Sprintf(
							"/password-reset?token=%s",
							*passwordRequest.GetToken(),
						),
						Color:     saas.GetConfig().GetConfiguration().GetMail().GetButtonColor(),
						TextColor: saas.GetConfig().GetConfiguration().GetMail().GetButtonTextColor(),
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
		From:        saas.GetConfig().GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[saas.GetConfig().GetConfiguration().GetLocale()]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
