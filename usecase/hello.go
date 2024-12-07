package usecase

import (
	"pitch_backend/domain/model"
	"pitch_backend/domain/repository"

	"github.com/samber/do"
)

type HelloUsecase interface {
	Get() (*model.Hello, error)
}

type helloUsecaseImpl struct {
	helloRepository repository.Hello
}

func NewHelloUsecase(i *do.Injector) (HelloUsecase, error) {
	helloRepository := do.MustInvoke[repository.Hello](i)

	return &helloUsecaseImpl{
		helloRepository: helloRepository,
	}, nil
}

func (u helloUsecaseImpl) Get() (*model.Hello, error) {
	return u.helloRepository.Get()
}
