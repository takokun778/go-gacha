package usecase

import (
	"context"
	"gacha/model"
	"gacha/repository"
)

type Card struct {
	card repository.Card
}

func NewCard(
	card repository.Card,
) *Card {
	return &Card{
		card: card,
	}
}

func (c *Card) CalcAllCardRate(ctx context.Context) ([]model.Card, error) {
	cards, err := c.card.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		rate := card.Rate.Value() * card.Rank.Rate() / model.MaxRankRate
		card.Rate = model.CardRate(rate)
	}

	return cards, nil
}
