package repository

import (
	"github.com/yeqown/micro-server-demo/model"

	"github.com/jinzhu/gorm"
)

// FooRepo is DAO in golang
type FooRepo interface {
	Create(m *model.FooModel) error
	Count(wheres string) (int, error)
	// and more methods to interact
}

// NewFooRepo .
func NewFooRepo(db *gorm.DB) FooRepo {
	return fooRepo{db: db}
}

type fooRepo struct {
	db *gorm.DB
}

// Create of fooRepo.
func (repo fooRepo) Create(m *model.FooModel) error {
	return nil
}

// Count of fooRepo.
func (repo fooRepo) Count(wheres string) (int, error) {
	return 0, nil
}