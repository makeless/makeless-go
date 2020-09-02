package go_saas_security_basic

import (
	"encoding/hex"
	"math/rand"
)

func (security *Security) GenerateToken(length int) (string, error) {
	bytes := make([]byte, length/2)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
