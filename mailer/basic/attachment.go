package go_saas_mailer_basic

import (
	"net/textproto"
	"sync"
)

type Attachment struct {
	Filename string
	Data     []byte
	Headers  textproto.MIMEHeader
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
