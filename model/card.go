package model

type Card struct {
	Name string
	Rank Rank
	Rate CardRate
}

func NewCard(
	name string,
	rank Rank,
	rate CardRate,
) Card {
	return Card{
		Name: name,
		Rank: rank,
		Rate: rate,
	}
}
