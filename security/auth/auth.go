package makeless_go_auth

import "time"

type Auth interface {
	Sign() (string, time.Time, error)
}
