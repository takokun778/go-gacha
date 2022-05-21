package main

import (
	"context"
	"fmt"
	"gacha/db"
	"gacha/model"
	"gacha/usecase"
	"log"
)

func main() {
	client, err := db.NewClient(false)
	if err != nil {
		log.Fatalln(err.Error())
	}

	card := db.NewCard(client)

	lottery := db.NewLottery(client)

	history := db.NewHistory(client)

	init := usecase.NewInitial(
		card,
		lottery,
	)

	gacha := usecase.NewGacha(
		card,
		lottery,
		history,
	)

	ctx := context.Background()

	if err := init.Execute(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	cards, err := gacha.Draw(ctx, 100)
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := 0
	a := 0
	b := 0
	c := 0

	for _, ca := range cards {
		switch ca.Rank {
		case model.S:
			s += 1
		case model.A:
			a += 1
		case model.B:
			b += 1
		default:
			c += 1
		}
		// fmt.Printf("%+v\n", ca)
	}

	fmt.Printf("s: %d a: %d b: %d c: %d", s, a, b, c)
}
