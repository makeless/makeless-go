package go_saas_basic_security

import "golang.org/x/crypto/bcrypt"

func (security *Security) EncryptPassword(password string) ([]byte, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func (security *Security) ComparePassword(userPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}
