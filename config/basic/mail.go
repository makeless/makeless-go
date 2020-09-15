package go_saas_config_basic

import (
	"github.com/go-saas/go-saas/config"
	"sync"
)

type Mail struct {
	Name            string               `json:"name"`
	Logo            string               `json:"logo"`
	From            string               `json:"from"`
	Link            string               `json:"link"`
	ButtonColor     string               `json:"buttonColor"`
	ButtonTextColor string               `json:"buttonTextColor"`
	Texts           map[string]*MailText `json:"texts"`
	*sync.RWMutex
}

func (mail *Mail) GetName() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Name
}

func (mail *Mail) GetLogo() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Logo
}

func (mail *Mail) GetFrom() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.From
}

func (mail *Mail) GetLink() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Link
}

func (mail *Mail) GetButtonColor() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.ButtonColor
}

func (mail *Mail) GetButtonTextColor() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.ButtonTextColor
}

func (mail *Mail) GetTexts(locale string) go_saas_config.MailText {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Texts[locale]
}
