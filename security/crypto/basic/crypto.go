package makeless_go_crypto_basic

import "golang.org/x/crypto/bcrypt"

type Crypto struct {
}

func (crypto *Crypto) EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func (crypto *Crypto) ComparePassword(userPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}
