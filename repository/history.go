package repository

import (
	"context"
	"gacha/model"
)

type History interface {
	Save(context.Context, model.CardID) error
	FindAll(context.Context) ([]model.CardID, error)
}
