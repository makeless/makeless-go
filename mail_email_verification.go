package go_saas

import (
	"fmt"
	"github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/mailer/basic"
	"github.com/go-saas/go-saas/model"
	"github.com/matcornic/hermes/v2"
	"sync"
)

func (saas *Saas) mailEmailVerification(data map[string]interface{}) (go_saas_mailer.Mail, error) {
	var err error
	var message, messageHtml string
	var user = data["user"].(*go_saas_model.User)
	var messages = map[string]map[string]string{
		"en": {
			"subject":     "Please verify your email address",
			"instruction": "To complete your registration, we just need to verify your email address:",
			"button":      "Verify email address",
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
			Name:      *user.GetName(),
			Greeting:  saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetGreeting(),
			Signature: saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf(
						"%s %s",
						messages[saas.GetConfig().GetConfiguration().GetLocale()]["instruction"],
						*user.GetEmail(),
					),
					Button: hermes.Button{
						Text: messages[saas.GetConfig().GetConfiguration().GetLocale()]["button"],
						Link: fmt.Sprintf(
							"%s/email-verification?token=%s",
							saas.GetConfig().GetConfiguration().GetMail().GetLink(),
							*user.GetEmailVerification().GetToken(),
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
		To:          []string{*user.GetEmail()},
		From:        saas.GetConfig().GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[saas.GetConfig().GetConfiguration().GetLocale()]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
