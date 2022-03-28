package makeless_go_repository

type GenericRepository interface {
	Exists(err error) (bool, error)
}
