package makeless_go_mail_basic

import (
	"fmt"
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/model"
	"github.com/makeless/makeless-go/v2/mail"
	"github.com/matcornic/hermes/v2"
	"sync"
)

type PasswordRequestMail struct {
}

func (passwordRequestMail *PasswordRequestMail) Create(config makeless_go_config.Config, passwordRequest *makeless_go_model.PasswordRequest, locale string) (makeless_go_mail.Mail, error) {
	var err error
	var message, messageHtml string
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

	hermesConfig := hermes.Hermes{
		Product: hermes.Product{
			Name:        config.GetConfiguration().GetMail().GetName(),
			Link:        config.GetConfiguration().GetMail().GetLink(),
			Logo:        config.GetConfiguration().GetMail().GetLogo(),
			Copyright:   config.GetConfiguration().GetMail().GetTexts(locale).GetCopyright(),
			TroubleText: config.GetConfiguration().GetMail().GetTexts(locale).GetTroubleText(),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Greeting:  config.GetConfiguration().GetMail().GetTexts(locale).GetGreeting(),
			Signature: config.GetConfiguration().GetMail().GetTexts(locale).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: messages[locale]["instruction"],
					Button: hermes.Button{
						Text: messages[locale]["button"],
						Link: fmt.Sprintf(
							"%s/password-reset?token=%s",
							config.GetConfiguration().GetMail().GetLink(),
							passwordRequest.Token,
						),
						Color:     config.GetConfiguration().GetMail().GetButtonColor(),
						TextColor: config.GetConfiguration().GetMail().GetButtonTextColor(),
					},
				},
			},
		},
	}

	if message, err = hermesConfig.GeneratePlainText(email); err != nil {
		return nil, err
	}

	if messageHtml, err = hermesConfig.GenerateHTML(email); err != nil {
		return nil, err
	}

	return &Mail{
		To:          []string{passwordRequest.Email},
		From:        config.GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[locale]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
