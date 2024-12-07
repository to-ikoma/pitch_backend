package psqlmodel

import "pitch_backend/domain/model"

type Hello struct {
	Message string
}

func (m *Hello) ToEntity() *model.Hello {
	return &model.Hello{
		Message: m.Message,
	}
}
