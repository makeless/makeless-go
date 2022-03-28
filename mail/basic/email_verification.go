package makeless_go_mail_basic

import (
	"fmt"
	"github.com/makeless/makeless-go/config"
	"github.com/makeless/makeless-go/database/model"
	"github.com/makeless/makeless-go/mail"
	"github.com/matcornic/hermes/v2"
	"sync"
)

type EmailVerificationMail struct {
}

func (emailVerificationMail *EmailVerificationMail) Create(config makeless_go_config.Config, user *makeless_go_model.User, locale string) (makeless_go_mail.Mail, error) {
	var err error
	var message, messageHtml string
	var messages = map[string]map[string]string{
		"en": {
			"subject":     "Please verify your email address",
			"instruction": "to complete your registration, we just need to verify your email address:",
			"button":      "Verify email address",
		},
		"de": {
			"subject":     "Bitte verifiziere deine Email Adresse",
			"instruction": "Bitte verifziere deine Email Adresse, um deine Registrierung abzuschließen:",
			"button":      "Email Adresse verifizieren",
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
			Name:      user.Name,
			Greeting:  config.GetConfiguration().GetMail().GetTexts(locale).GetGreeting(),
			Signature: config.GetConfiguration().GetMail().GetTexts(locale).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf(
						"%s %s",
						messages[locale]["instruction"],
						user.Email,
					),
					Button: hermes.Button{
						Text: messages[locale]["button"],
						Link: fmt.Sprintf(
							"%s/email-verification?token=%s",
							config.GetConfiguration().GetMail().GetLink(),
							user.EmailVerification.Token,
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
		To:          []string{user.Email},
		From:        config.GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[locale]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
