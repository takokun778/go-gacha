package db

import (
	"context"
	"gacha/.db/table"
	"gacha/model"
	"gacha/repository"

	"github.com/uptrace/bun"
)

type HistoryTable struct {
	bun.BaseModel `bun:"table:histories"`

	CardID string `bun:"card_id"`
}

type History struct {
	client *Client
}

func NewHistory(client *Client) repository.History {
	return &History{
		client: client,
	}
}

func (h *History) Save(ctx context.Context, id model.CardID) error {
	return nil
}

func (h *History) FindAll(ctx context.Context) ([]model.CardID, error) {
	return nil, nil
}

func NewHistoryMigrator(
	client *Client,
) table.History {
	return &History{
		client: client,
	}
}

func (h *History) CreateTable(ctx context.Context) error {
	if _, err := h.client.DB.NewCreateTable().
		Model((*HistoryTable)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (h *History) DropTable(ctx context.Context) error {
	if _, err := h.client.DB.NewDropTable().
		Model((*HistoryTable)(nil)).
		IfExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
