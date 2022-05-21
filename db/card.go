package db

import (
	"context"
	"gacha/.db/table"
	"gacha/model"
	"gacha/repository"

	"github.com/uptrace/bun"
)

type CardTable struct {
	bun.BaseModel `bun:"table:cards"`

	ID   string  `bun:"id,pk"`
	Name string  `bun:"name"`
	Rank string  `bun:"rank"`
	Rate float64 `bun:"rate"`
}

type Card struct {
	client *Client
}

func NewCard(client *Client) repository.Card {
	return &Card{
		client: client,
	}
}

func (c *Card) Save(ctx context.Context, cards []model.Card) error {
	items := make([]CardTable, 0, len(cards))

	for _, c := range cards {
		items = append(items, CardTable{
			ID:   string(c.ID),
			Name: c.Name,
			Rank: c.Rank.String(),
			Rate: c.Rate.Value(),
		})
	}

	if _, err := c.client.DB.NewInsert().Model(&items).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (c *Card) Find(ctx context.Context, id model.CardID) (model.Card, error) {
	card := new(CardTable)

	err := c.client.DB.NewSelect().
		Model(card).
		Where("id = ?", string(id)).
		Scan(ctx)
	if err != nil {
		return model.Card{}, err
	}

	return model.Card{
		ID:   model.CardID(card.ID),
		Name: card.Name,
		Rank: model.Rank(card.Rank),
		Rate: model.CardRate(card.Rate),
	}, nil
}

func (c *Card) FindAll(ctx context.Context) ([]model.Card, error) {
	return nil, nil
}

func NewCardMigrator(
	client *Client,
) table.Card {
	return &Card{
		client: client,
	}
}

func (c *Card) CreateTable(ctx context.Context) error {
	if _, err := c.client.DB.NewCreateTable().
		Model((*CardTable)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (c *Card) DropTable(ctx context.Context) error {
	if _, err := c.client.DB.NewDropTable().
		Model((*CardTable)(nil)).
		IfExists().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
