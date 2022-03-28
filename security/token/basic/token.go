package makeless_go_security_token_basic

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type SecurityToken struct {
}

func (securityToken *SecurityToken) Generate(length int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length/2)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
