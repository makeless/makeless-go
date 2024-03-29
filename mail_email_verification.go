package makeless_go

import (
	"fmt"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/mailer/basic"
	"github.com/makeless/makeless-go/model"
	"github.com/matcornic/hermes/v2"
	"sync"
)

func (makeless *Makeless) mailEmailVerification(data map[string]interface{}, locale string) (makeless_go_mailer.Mail, error) {
	var err error
	var message, messageHtml string
	var user = data["user"].(*makeless_go_model.User)
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
			Name:      *user.GetName(),
			Greeting:  makeless.GetConfig().GetConfiguration().GetMail().GetTexts(locale).GetGreeting(),
			Signature: makeless.GetConfig().GetConfiguration().GetMail().GetTexts(locale).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf(
						"%s %s",
						messages[locale]["instruction"],
						*user.GetEmail(),
					),
					Button: hermes.Button{
						Text: messages[locale]["button"],
						Link: fmt.Sprintf(
							"%s/email-verification?token=%s",
							makeless.GetConfig().GetConfiguration().GetMail().GetLink(),
							*user.GetEmailVerification().GetToken(),
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
		To:          []string{*user.GetEmail()},
		From:        makeless.GetConfig().GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[locale]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
