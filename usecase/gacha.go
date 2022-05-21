package usecase

import (
	"context"
	"errors"
	"gacha/model"
	"gacha/repository"
)

type Gacha struct {
	card    repository.Card
	lottery repository.Lottery
	history repository.History
}

func NewGacha(
	card repository.Card,
	lottery repository.Lottery,
	history repository.History,
) *Gacha {
	return &Gacha{
		card:    card,
		lottery: lottery,
		history: history,
	}
}

func (g *Gacha) Draw(ctx context.Context, count int) ([]model.Card, error) {
	if count > model.MaxLotteryCount {
		return nil, errors.New("count is over max count")
	}

	cards := make([]model.Card, 0, count)

	for i := 0; i < count; i++ {
		card, err := g.draw(ctx)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (g *Gacha) draw(ctx context.Context) (model.Card, error) {
	id, err := g.lottery.Lotto(ctx, model.LotteryRankRate())
	if err != nil {
		return model.Card{}, err
	}

	if err := g.history.Save(ctx, id); err != nil {
		return model.Card{}, err
	}

	card, err := g.card.Find(ctx, id)
	if err != nil {
		return model.Card{}, err
	}

	return card, nil
}
