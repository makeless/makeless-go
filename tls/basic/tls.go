package makeless_go_tls_basic

import (
	"github.com/gin-gonic/gin"
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

func (tls *Tls) Run(port string, engine *gin.Engine) error {
	return engine.RunTLS(":"+port, tls.getCertPath(), tls.getKeyPath())
}
