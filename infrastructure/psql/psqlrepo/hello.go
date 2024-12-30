package psqlrepo

import (
	"context"
	"pitch_backend/domain/model"
	"pitch_backend/domain/repository"
	"pitch_backend/infrastructure/psql"
	"pitch_backend/infrastructure/psql/psqlmodel"

	"github.com/pkg/errors"
	"github.com/samber/do"
)

type helloRepositoryImpl struct{}

func NewHelloRepository(i *do.Injector) (repository.Hello, error) {
	return &helloRepositoryImpl{}, nil
}

func (r helloRepositoryImpl) Get() (*model.Hello, error) {
	db, err := psql.GetDBCon()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to database")
	}
	defer db.Close()

	hello := &psqlmodel.Hello{}
	err = db.NewSelect().Model(hello).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return hello.ToEntity(), nil
}
