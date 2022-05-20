package main

import (
	"context"
	"gacha/db"
	"log"
)

func main() {
	client, err := db.NewClient()
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()

	if _, err := client.DB.NewCreateTable().Model((*db.Cards)(nil)).Exec(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := client.DB.NewCreateTable().Model((*db.Lotteries)(nil)).Exec(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := client.DB.NewCreateTable().Model((*db.Histories)(nil)).Exec(ctx); err != nil {
		log.Fatalln(err.Error())
	}
}
