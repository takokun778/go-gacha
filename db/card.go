package db

import (
	"context"
	"gacha/model"
	"gacha/repository"

	"github.com/uptrace/bun"
)

type Cards struct {
	bun.BaseModel `bun:"table:cards"`

	ID   string  `bun:"id,pk"`
	Name string  `bun:"name"`
	Rank string  `bun:"rank"`
	Rate float64 `bun:"rate"`
}

type Card struct {
	Client *Client
}

func NewCard(client *Client) repository.Card {
	return &Card{
		Client: client,
	}
}

func (c *Card) Save(ctx context.Context, cards []model.Card) error {
	items := make([]Cards, 0, len(cards))

	for _, c := range cards {
		items = append(items, Cards{
			ID:   string(c.ID),
			Name: c.Name,
			Rank: c.Rank.String(),
			Rate: c.Rate.Value(),
		})
	}

	if _, err := c.Client.DB.NewInsert().Model(&items).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (c *Card) Find(ctx context.Context, id model.CardID) (model.Card, error) {
	return model.Card{}, nil
}

func (c *Card) FindAll(ctx context.Context) ([]model.Card, error) {
	return nil, nil
}
