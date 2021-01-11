package makeless_go_mailer_basic

import (
	"net/textproto"
	"sync"
)

type Attachment struct {
	Filename string               `json:"filename"`
	Data     []byte               `json:"data"`
	Headers  textproto.MIMEHeader `json:"headers"`
	*sync.RWMutex
}

func (attachment *Attachment) GetFilename() string {
	attachment.RLock()
	defer attachment.RUnlock()

	return attachment.Filename
}

func (attachment *Attachment) GetData() []byte {
	attachment.RLock()
	defer attachment.RUnlock()

	return attachment.Data
}

func (attachment *Attachment) GetHeaders() textproto.MIMEHeader {
	attachment.RLock()
	defer attachment.RUnlock()

	return attachment.Headers
}
