package makeless_go_mail

import "net/textproto"

type Attachment interface {
	GetFilename() string
	GetData() []byte
	GetHeaders() textproto.MIMEHeader
}
