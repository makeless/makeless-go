package go_saas_mailer_basic

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"sync"

	"github.com/go-saas/go-saas/mailer"
	"github.com/jordan-wright/email"
)

type Mailer struct {
	Auth     smtp.Auth
	Tls      *tls.Config
	Host     string
	Port     string
	Identity string
	Username string
	Password string
	*sync.RWMutex
}

func (mailer *Mailer) GetAuth() smtp.Auth {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Auth
}

func (mailer *Mailer) SetAuth(auth smtp.Auth) {
	mailer.Lock()
	defer mailer.Unlock()

	mailer.Auth = auth
}

func (mailer *Mailer) GetTls() *tls.Config {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Tls
}

func (mailer *Mailer) GetHost() string {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Host
}

func (mailer *Mailer) GetPort() string {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Port
}

func (mailer *Mailer) GetIdentity() string {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Identity
}

func (mailer *Mailer) GetUsername() string {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Username
}

func (mailer *Mailer) GetPassword() string {
	mailer.RLock()
	defer mailer.RUnlock()

	return mailer.Password
}

func (mailer *Mailer) Init() error {
	mailer.SetAuth(smtp.PlainAuth(
		mailer.GetIdentity(),
		mailer.GetUsername(),
		mailer.GetPassword(),
		mailer.GetHost(),
	))

	return nil
}

func (mailer *Mailer) Send(ctx context.Context, mail go_saas_mailer.Mail) error {
	var e = &email.Email{
		To:          mail.GetTo(),
		Cc:          mail.GetCc(),
		Bcc:         mail.GetBcc(),
		From:        mail.GetFrom(),
		Subject:     mail.GetSubject(),
		Text:        mail.GetMessage(),
		HTML:        mail.GetHtmlMessage(),
		Attachments: nil,
		Headers:     mail.GetHeaders(),
	}

	for _, attachment := range mail.GetAttachments() {
		e.Attachments = append(e.Attachments, &email.Attachment{
			Filename:    attachment.GetFilename(),
			Header:      attachment.GetHeaders(),
			Content:     attachment.GetData(),
			HTMLRelated: false,
		})
	}

	if mailer.Tls == nil {
		return e.Send(fmt.Sprintf("%s:%s", mailer.GetHost(), mailer.GetPort()), mailer.GetAuth())
	}

	return e.SendWithTLS(fmt.Sprintf("%s:%s", mailer.GetHost(), mailer.GetPort()), mailer.GetAuth(), mailer.GetTls())
}
