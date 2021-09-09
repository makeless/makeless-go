package makeless_go_tls_basic

import (
	"github.com/makeless/makeless-go/http"
	"sync"
)

type Tls struct {
	CertPath string
	KeyPath  string
	*sync.RWMutex
}

func (tls *Tls) getCertPath() string {
	tls.RLock()
	defer tls.RUnlock()

	return tls.CertPath
}

func (tls *Tls) getKeyPath() string {
	tls.RLock()
	defer tls.RUnlock()

	return tls.KeyPath
}

func (tls *Tls) Run(http makeless_go_http.Http) error {
	return http.GetRouter().GetEngine().RunTLS(":"+http.GetPort(), tls.getCertPath(), tls.getKeyPath())
}
