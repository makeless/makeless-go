package makeless_go_mailer

import "net/textproto"

type Attachment interface {
	GetFilename() string
	GetData() []byte
	GetHeaders() textproto.MIMEHeader
}
