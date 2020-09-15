package go_saas_config_basic

import "sync"

type MailText struct {
	Greeting  string `json:"greeting"`
	Signature string `json:"signature"`
	Copyright string `json:"copyright"`
	*sync.RWMutex
}

func (mailText *MailText) GetGreeting() string {
	mailText.RLock()
	defer mailText.RUnlock()

	return mailText.Greeting
}

func (mailText *MailText) GetSignature() string {
	mailText.RLock()
	defer mailText.RUnlock()

	return mailText.Signature
}

func (mailText *MailText) GetCopyright() string {
	mailText.RLock()
	defer mailText.RUnlock()

	return mailText.Copyright
}
