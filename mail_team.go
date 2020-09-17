package go_saas

import (
	"fmt"
	"github.com/go-saas/go-saas/mailer"
	"github.com/go-saas/go-saas/mailer/basic"
	"github.com/go-saas/go-saas/model"
	"github.com/matcornic/hermes/v2"
	"sync"
)

func (saas *Saas) mailTeamInvitation(data map[string]interface{}) (go_saas_mailer.Mail, error) {
	var err error
	var name, link, message, messageHtml string
	var user = data["user"].(*go_saas_model.User)
	var userInvited = data["userInvited"].(*go_saas_model.User)
	var teamName = data["teamName"].(string)
	var teamInvitation = data["teamInvitation"].(*go_saas_model.TeamInvitation)
	var messages = map[string]map[string]string{
		"en": {
			"subject": fmt.Sprintf(
				"%s invited you to %s",
				*user.GetName(),
				teamName,
			),
			"intro": fmt.Sprintf(
				"%s has invited you to %s.",
				*user.GetName(),
				teamName,
			),
			"outro": fmt.Sprintf(
				"Note: This invitation was intended for %s. If you were not expecting this invitation, you can ignore this email.",
				*teamInvitation.GetEmail(),
			),
			"instruction": "You can accept or decline this invitation. This invitation will expire in 7 days.",
			"button":      "View invitation",
		},
	}

	switch userInvited.GetName() {
	case nil:
		link = fmt.Sprintf("%s/invitation?token=%s", saas.GetConfig().GetConfiguration().GetMail().GetLink(), *teamInvitation.GetToken())
	default:
		name = *userInvited.GetName()
		link = fmt.Sprintf("%s/settings/invitation", saas.GetConfig().GetConfiguration().GetMail().GetLink())
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
			Name:      name,
			Greeting:  saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetGreeting(),
			Signature: saas.GetConfig().GetConfiguration().GetMail().GetTexts(saas.GetConfig().GetConfiguration().GetLocale()).GetSignature(),
			Intros:    []string{messages[saas.GetConfig().GetConfiguration().GetLocale()]["intro"]},
			Actions: []hermes.Action{
				{
					Instructions: messages[saas.GetConfig().GetConfiguration().GetLocale()]["instruction"],
					Button: hermes.Button{
						Text:      messages[saas.GetConfig().GetConfiguration().GetLocale()]["button"],
						Link:      link,
						Color:     saas.GetConfig().GetConfiguration().GetMail().GetButtonColor(),
						TextColor: saas.GetConfig().GetConfiguration().GetMail().GetButtonTextColor(),
					},
				},
			},
			Outros: []string{messages[saas.GetConfig().GetConfiguration().GetLocale()]["outro"]},
		},
	}

	if message, err = config.GeneratePlainText(email); err != nil {
		return nil, err
	}

	if messageHtml, err = config.GenerateHTML(email); err != nil {
		return nil, err
	}

	return &go_saas_mailer_basic.Mail{
		To:          []string{*teamInvitation.GetEmail()},
		From:        saas.GetConfig().GetConfiguration().GetMail().GetFrom(),
		Subject:     messages[saas.GetConfig().GetConfiguration().GetLocale()]["subject"],
		Message:     []byte(message),
		HtmlMessage: []byte(messageHtml),
		RWMutex:     new(sync.RWMutex),
	}, nil
}
