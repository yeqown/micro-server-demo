package service

import (
	"github.com/yeqown/micro-server-demo/internal/model"
	"github.com/yeqown/micro-server-demo/internal/repository"
)

// FooUsecase .
type FooUsecase interface {
	Create(form *model.FooCreateForm) error
	Count() (int, error)
}

// NewFooUsecase .
func NewFooUsecase(repo repository.FooRepo) FooUsecase {
	return fooUsecase{repo: repo}
}

type fooUsecase struct {
	repo repository.FooRepo
}

// Create method of fooUsecase.
func (uc fooUsecase) Create(form *model.FooCreateForm) error {
	m := &model.FooModel{
		Bar: form.Bar,
	}
	return uc.repo.Create(m)
}

// Count method of fooUsecase.
func (uc fooUsecase) Count() (int, error) {
	return uc.repo.Count("bar = 'calledbar'")
}
