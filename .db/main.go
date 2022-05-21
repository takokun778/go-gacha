package main

import (
	"context"
	"gacha/db"
	"log"
)

func main() {
	client, err := db.NewClient(true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()

	card := db.NewCardMigrator(client)

	if err := card.DropTable(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	if err := card.CreateTable(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	history := db.NewHistoryMigrator(client)

	if err := history.DropTable(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	if err := history.CreateTable(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	lottery := db.NewLotteryMigrator(client)

	if err := lottery.DropTable(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	if err := lottery.CreateTable(ctx); err != nil {
		log.Fatalln(err.Error())
	}
}
