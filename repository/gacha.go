package repository

import (
	"context"
	"gacha/model"
)

type Gacha interface {
	Lotto(ctx context.Context, rank model.Rank, lottery model.LotteryNumber) (model.Card, error)
}
