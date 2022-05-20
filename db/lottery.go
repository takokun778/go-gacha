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
	l.Client.DB.ExecContext(ctx, "")

	return model.CardID(""), nil
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

func (l *Lottery) DeleteAll(ctx context.Context) error {
	if _, err := l.Client.DB.NewDelete().Model((*Lotteries)(nil)).Exec(ctx); err != nil {
		return err
	}

	return nil
}
