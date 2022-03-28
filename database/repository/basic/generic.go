package makeless_go_repository_basic

import (
	"errors"
	"gorm.io/gorm"
)

type GenericRepository struct {
}

func (genericRepository *GenericRepository) Exists(err error) (bool, error) {
	if err != nil {
		switch errors.Is(err, gorm.ErrRecordNotFound) {
		case true:
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
