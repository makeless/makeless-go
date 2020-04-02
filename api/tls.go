package saas_api

import "sync"

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

	return tls.CertPath
}