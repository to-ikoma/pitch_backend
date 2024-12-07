package repository

import "pitch_backend/domain/model"

type Hello interface {
	Get() (hello *model.Hello, err error)
}
