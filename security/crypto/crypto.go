package makeless_go_crypto

type Crypto interface {
	EncryptPassword(password string) (string, error)
	ComparePassword(userPassword string, password string) error
}
