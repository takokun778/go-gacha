package main

import (
	"context"
	"gacha/db"
	"gacha/usecase"
	"log"
)

func main() {
	client, err := db.NewClient()
	if err != nil {
		log.Fatalln(err.Error())
	}

	gacha := usecase.NewGacha(
		db.NewCard(client),
		db.NewLottery(client),
		db.NewHistory(client),
	)

	ctx := context.Background()

	if err := gacha.Initialize(ctx); err != nil {
		log.Fatalln(err.Error())
	}
}
