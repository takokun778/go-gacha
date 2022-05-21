package usecase

import (
	"context"
	"gacha/model"
	"gacha/repository"
)

type Initial struct {
	card    repository.Card
	lottery repository.Lottery
}

func NewInitial(
	card repository.Card,
	lottery repository.Lottery,
) *Initial {
	return &Initial{
		card:    card,
		lottery: lottery,
	}
}

func (i *Initial) Execute(ctx context.Context) error {
	if err := i.initialize(ctx, model.S); err != nil {
		return err
	}

	if err := i.initialize(ctx, model.A); err != nil {
		return err
	}

	if err := i.initialize(ctx, model.B); err != nil {
		return err
	}

	if err := i.initialize(ctx, model.C); err != nil {
		return err
	}

	return nil
}

func (i *Initial) initialize(ctx context.Context, rank model.Rank) error {
	cards, err := model.GenerateCardsFromCsvFile(rank)
	if err != nil {
		return err
	}

	if err := i.card.Save(ctx, cards); err != nil {
		return err
	}

	ids := make([]model.CardID, 0, model.MaxLotteryVolume)

	for _, card := range cards {
		count := int(card.Rate.Value() / model.MaxCardRate * model.MaxLotteryVolume)
		for i := 0; i < count; i++ {
			ids = append(ids, card.ID)
		}
	}

	if err := i.lottery.Save(ctx, rank, ids); err != nil {
		return err
	}

	return nil
}
