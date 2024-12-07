package psqlrepo

import (
	"log"
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

	var greeting string
	err = db.QueryRow("SELECT 'Hello!'").Scan(&greeting)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	hello := psqlmodel.Hello{
		Message: greeting,
	}

	return hello.ToEntity(), nil

}
