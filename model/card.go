package model

import (
	"fmt"
	"strconv"
)

type Card struct {
	ID   CardID
	Name string
	Rank Rank
	Rate CardRate
}

func NewCard(
	id CardID,
	name string,
	rank Rank,
	rate CardRate,
) Card {
	return Card{
		ID:   id,
		Name: name,
		Rank: rank,
		Rate: rate,
	}
}

func GenerateCardsFromCsvFile(rank Rank) ([]Card, error) {
	file := fmt.Sprintf("data/%s.csv", rank.String())

	records, err := ImportCSV(file)
	if err != nil {
		return nil, err
	}

	cards := make([]Card, 0, len(records))

	for _, record := range records {
		r, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		rate, err := NewCardRate(r)
		if err != nil {
			return nil, err
		}

		card := NewCard(
			GenerateCardID(),
			record[0],
			rank,
			rate,
		)

		cards = append(cards, card)
	}

	return cards, nil
}
