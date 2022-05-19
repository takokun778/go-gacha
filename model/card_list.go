package model

import "errors"

type CardList []Card

func NewCardList(
	Rank Rank,
	card []Card,
) (CardList, error) {
	var total float64
	for _, card := range card {
		if card.Rank != Rank {
			return nil, errors.New("")
		}
		total += float64(card.Rate)
	}

	if total != MaxCardRate {
		return nil, errors.New("")
	}

	return CardList(card), nil
}
