package db

import (
	"context"
	"gacha/model"
	"gacha/repository"

	"github.com/uptrace/bun"
)

type Lotteries struct {
	bun.BaseModel `bun:"table:lotteries"`

	CardID string `bun:"card_id"`
	Rank   string `bun:"rank"`
}

type Lottery struct {
	Client *Client
}

func NewLottery(client *Client) repository.Lottery {
	return &Lottery{
		Client: client,
	}
}

func (l *Lottery) Lotto(ctx context.Context, rank model.Rank) (model.CardID, error) {
	lottery := new(Lotteries)

	err := l.Client.DB.NewSelect().
		Model(lottery).
		Where("rank = ?", rank.String()).
		OrderExpr("random()").
		Limit(1).
		Scan(ctx)

	if err != nil {
		return model.CardID(""), err
	}

	return model.CardID(lottery.CardID), nil
}

func (l *Lottery) Save(ctx context.Context, rank model.Rank, ids []model.CardID) error {
	lotteries := make([]Lotteries, 0, len(ids))

	for _, id := range ids {
		lotteries = append(lotteries, Lotteries{
			CardID: string(id),
			Rank:   rank.String(),
		})
	}

	if _, err := l.Client.DB.NewInsert().Model(&lotteries).Exec(ctx); err != nil {
		return err
	}

	return nil
}
