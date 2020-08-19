package go_saas_security_basic

import "golang.org/x/crypto/bcrypt"

func (security *Security) EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func (security *Security) ComparePassword(userPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}
