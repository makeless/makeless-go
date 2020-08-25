package go_saas_mailer_basic

import (
	"net/textproto"
	"sync"

	"github.com/go-saas/go-saas/mailer"
)

type Mail struct {
	To          []string
	Cc          []string
	Bcc         []string
	From        string
	Subject     string
	Message     []byte
	HtmlMessage []byte
	Attachments []go_saas_mailer.Attachment
	Headers     textproto.MIMEHeader

	*sync.RWMutex
}

func (mail *Mail) GetTo() []string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.To
}

func (mail *Mail) SetTo(to []string) {
	mail.Lock()
	defer mail.Unlock()

	mail.To = to
}

func (mail *Mail) GetCc() []string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Cc
}

func (mail *Mail) SetCc(cc []string) {
	mail.Lock()
	defer mail.Unlock()

	mail.Cc = cc
}

func (mail *Mail) GetBcc() []string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Bcc
}

func (mail *Mail) SetBcc(Bcc []string) {
	mail.Lock()
	defer mail.Unlock()

	mail.Bcc = Bcc
}

func (mail *Mail) GetFrom() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.From
}

func (mail *Mail) SetFrom(from string) {
	mail.Lock()
	defer mail.Unlock()

	mail.From = from
}

func (mail *Mail) GetSubject() string {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Subject
}

func (mail *Mail) SetSubject(subject string) {
	mail.Lock()
	defer mail.Unlock()

	mail.Subject = subject
}

func (mail *Mail) GetMessage() []byte {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Message
}

func (mail *Mail) SetMessage(message []byte) {
	mail.Lock()
	defer mail.Unlock()

	mail.Message = message
}

func (mail *Mail) GetHtmlMessage() []byte {
	mail.RLock()
	defer mail.RUnlock()

	return mail.HtmlMessage
}

func (mail *Mail) SetHtmlMessage(htmlMessage []byte) {
	mail.Lock()
	defer mail.Unlock()

	mail.HtmlMessage = htmlMessage
}

func (mail *Mail) GetAttachments() []go_saas_mailer.Attachment {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Attachments
}

func (mail *Mail) SetAttachments(attachments []go_saas_mailer.Attachment) {
	mail.Lock()
	defer mail.Unlock()

	mail.Attachments = attachments
}

func (mail *Mail) GetHeaders() textproto.MIMEHeader {
	mail.RLock()
	defer mail.RUnlock()

	return mail.Headers
}

func (mail *Mail) SetHeaders(headers textproto.MIMEHeader) {
	mail.Lock()
	defer mail.Unlock()

	mail.Headers = headers
}
