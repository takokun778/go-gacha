package repository

import (
	"context"
	"gacha/model"
)

type Card interface {
	Save(context.Context, []model.Card) error
	Find(context.Context, model.CardID) (model.Card, error)
	FindAll(context.Context) ([]model.Card, error)
}
