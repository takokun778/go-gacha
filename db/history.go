package db

import (
	"context"
	"gacha/model"
	"gacha/repository"

	"github.com/uptrace/bun"
)

type Histories struct {
	bun.BaseModel `bun:"table:histories"`

	CardID string `bun:"card_id"`
}

type History struct {
	Client *Client
}

func NewHistory(client *Client) repository.History {
	return &History{
		Client: client,
	}
}

func (h *History) Save(ctx context.Context, id model.CardID) error {
	return nil
}

func (h *History) FindAll(ctx context.Context) ([]model.CardID, error) {
	return nil, nil
}
