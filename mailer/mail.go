package makeless_go_mailer

import "net/textproto"

type Mail interface {
	GetTo() []string
	SetTo(to []string)

	GetCc() []string
	SetCc(cc []string)

	GetBcc() []string
	SetBcc(bcc []string)

	GetFrom() string
	SetFrom(from string)

	GetSubject() string
	SetSubject(subject string)

	GetMessage() []byte
	SetMessage(message []byte)

	GetHtmlMessage() []byte
	SetHtmlMessage(htmlMessage []byte)

	GetAttachments() []Attachment
	SetAttachments(attachments []Attachment)

	GetHeaders() textproto.MIMEHeader
	SetHeaders(headers textproto.MIMEHeader)
}
