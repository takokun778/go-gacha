package repository

import (
	"context"
	"gacha/model"
)

type Lottery interface {
	Lotto(context.Context, model.Rank) (model.CardID, error)
	Save(context.Context, model.Rank, []model.CardID) error
	DeleteAll(context.Context) error
}
