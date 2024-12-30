package psqlmodel

import (
	"pitch_backend/domain/model"

	"github.com/uptrace/bun"
)

type Hello struct {
	bun.BaseModel `bun:"table:hello,alias:hello"`
	ID            int64  `bun:"id"`
	Name          string `bun:"name"`
}

func (m *Hello) ToEntity() *model.Hello {
	return &model.Hello{
		Message: m.Name,
	}
}
