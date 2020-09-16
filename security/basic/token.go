package go_saas_security_basic

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func (security *Security) GenerateToken(length int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length/2)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
