package makeless_go

import (
	"fmt"
	"sync"

	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/mailer/basic"
	"github.com/makeless/makeless-go/model"
	"github.com/matcornic/hermes/v2"
)

func (makeless *Makeless) mailPasswordRequest(data map[string]interface{}) (makeless_go_mailer.Mail, error) {
	var err error
	var message, messageHtml string
	var passwordRequest = data["passwordRequest"].(*makeless_go_model.PasswordRequest)
	var messages = map[string]map[string]string{
		"en": {
			"subject":     "Reset your password",
			"instruction": "to reset your password, please click here:",
			"button":      "Reset password",
		},
	}

	config := hermes.Hermes{
		Product: hermes.Product{
			Name:      makeless.GetConfig().GetConfiguration().GetMail().GetName(),
			Link:      makeless.GetConfig().GetConfiguration().GetMail().GetLink(),
			Logo:      makeless.GetConfig().GetConfiguration().GetMail().GetLogo(),
			Copyright: makeless.GetConfig().GetConfiguration().GetMail().GetTexts(makeless.GetConfig().GetConfiguration().GetLocale()).GetCopyright(),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Greeting:  makeless.GetConfig().GetConfiguration().GetMail().GetTexts(makeless.GetConfig().GetConfiguration().GetLocale()).GetGreeting(),
			Signature: makeless.GetConfig().GetConfiguration().GetMail().GetTexts(makeless.GetConfig().GetConfiguration().GetLocale()).GetSignature(),
			Actions: []hermes.Action{
				{
					Instructions: messages[makeless.GetConfig().GetConfiguration().GetLocale()]["instruction"],
					Button: hermes.Button{
						Text: messages[makeless.GetConfig().GetConfiguration().GetLocale()]["button"],
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
		Subject:     messages[makeless.GetConfig().GetConfiguration().GetLocale()]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
