package makeless_go

import (
	"fmt"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/mailer/basic"
	"github.com/makeless/makeless-go/model"
	"github.com/matcornic/hermes/v2"
	"sync"
)

func (makeless *Makeless) mailPasswordRequest(data map[string]interface{}, locale string) (makeless_go_mailer.Mail, error) {
	var err error
	var message, messageHtml string
	var passwordRequest = data["passwordRequest"].(*makeless_go_model.PasswordRequest)
	var messages = map[string]map[string]string{
		"en": {
			"subject":     "Reset your password",
			"instruction": "to reset your password, please click here:",
			"button":      "Reset password",
		},
		"de": {
			"subject":     "Passwort zurücksetzen",
			"instruction": "bitte hier klicken, um dein Passwort zurückzusetzen:",
			"button":      "Passwort zurücksetzen",
		},
	}

	config := hermes.Hermes{
		Product: hermes.Product{
			Name:        makeless.GetConfig().GetConfiguration().GetMail().GetName(),
			Link:        makeless.GetConfig().GetConfiguration().GetMail().GetLink(),
			Logo:        makeless.GetConfig().GetConfiguration().GetMail().GetLogo(),
			Copyright:   makeless.GetConfig().GetConfiguration().GetMail().GetTexts(locale).GetCopyright(),
			TroubleText: makeless.GetConfig().GetConfiguration().GetMail().GetTexts(locale).GetTroubleText(),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Greeting:  makeless.GetConfig().GetConfiguration().GetMail().GetTexts(locale).GetGreeting(),
			Signature: makeless.GetConfig().GetConfiguration().GetMail().GetTexts(locale).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: messages[locale]["instruction"],
					Button: hermes.Button{
						Text: messages[locale]["button"],
						Link: fmt.Sprintf(
							"%s/password-reset?token=%s",
							makeless.GetConfig().GetConfiguration().GetMail().GetLink(),
							*passwordRequest.GetToken(),
						),
						Color:     makeless.GetConfig().GetConfiguration().GetMail().GetButtonColor(),
						TextColor: makeless.GetConfig().GetConfiguration().GetMail().GetButtonTextColor(),
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

	return &makeless_go_mailer_basic.Mail{
		To:          []string{*passwordRequest.GetEmail()},
		From:        makeless.GetConfig().GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[locale]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
