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

func (g *Gacha) CalcAllCardRate(ctx context.Context) ([]model.Card, error) {
	cards, err := g.card.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		rate := card.Rate.Value() * card.Rank.Rate() / model.MaxRankRate
		card.Rate = model.CardRate(rate)
	}

	return cards, nil
}

func (g *Gacha) Initialize(ctx context.Context) error {
	if err := g.initialize(ctx, model.S); err != nil {
		return err
	}

	if err := g.initialize(ctx, model.A); err != nil {
		return err
	}

	if err := g.initialize(ctx, model.B); err != nil {
		return err
	}

	if err := g.initialize(ctx, model.C); err != nil {
		return err
	}

	return nil
}

func (g *Gacha) initialize(ctx context.Context, rank model.Rank) error {
	cards, err := model.GenerateCardsFromCsvFile(rank)
	if err != nil {
		return err
	}

	if err := g.card.Save(ctx, cards); err != nil {
		return err
	}

	ids := make([]model.CardID, 0, model.MaxLotteryVolume)

	for _, card := range cards {
		count := int(card.Rate.Value() / model.MaxCardRate * model.MaxLotteryVolume)
		for i := 0; i < count; i++ {
			ids = append(ids, card.ID)
		}
	}

	if err := g.lottery.Save(ctx, rank, ids); err != nil {
		return err
	}

	return nil
}
