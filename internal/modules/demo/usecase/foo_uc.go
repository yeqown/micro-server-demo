package usecase

import (
	"github.com/yeqown/micro-server-demo/internal/repository"
	"github.com/yeqown/micro-server-demo/pkg/types"
)

// FooUsecase .
type FooUsecase interface {
	Create(form *types.FooCreateForm) error
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
func (uc fooUsecase) Create(form *types.FooCreateForm) error {
	m := &types.FooModel{
		Bar: form.Bar,
	}
	return uc.repo.Create(m)
}

// Count method of fooUsecase.
func (uc fooUsecase) Count() (int, error) {
	return uc.repo.Count("bar = 'calledbar'")
}
