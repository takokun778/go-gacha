package db

import (
	"context"
	"gacha/.db/table"
	"gacha/model"
	"gacha/repository"

	"github.com/uptrace/bun"
)

type LotteryTable struct {
	bun.BaseModel `bun:"table:lotteries"`

	CardID string `bun:"card_id"`
	Rank   string `bun:"rank"`
}

type Lottery struct {
	client *Client
}

func NewLottery(client *Client) repository.Lottery {
	return &Lottery{
		client: client,
	}
}

func (l *Lottery) Lotto(ctx context.Context, rank model.Rank) (model.CardID, error) {
	lottery := new(LotteryTable)

	err := l.client.DB.NewSelect().
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
	lotteries := make([]LotteryTable, 0, len(ids))

	for _, id := range ids {
		lotteries = append(lotteries, LotteryTable{
			CardID: string(id),
			Rank:   rank.String(),
		})
	}

	if _, err := l.client.DB.NewInsert().Model(&lotteries).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func NewLotteryMigrator(
	client *Client,
) table.Lottery {
	return &Lottery{
		client: client,
	}
}

func (h *Lottery) CreateTable(ctx context.Context) error {
	if _, err := h.client.DB.NewCreateTable().
		Model((*LotteryTable)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Lottery) DropTable(ctx context.Context) error {
	if _, err := h.client.DB.NewDropTable().
		Model((*LotteryTable)(nil)).
		IfExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
