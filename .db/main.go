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

	_, err = client.DB.NewDropTable().
		Model((*db.Cards)(nil)).
		IfExists().
		Exec(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = client.DB.NewCreateTable().
		Model((*db.Cards)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = client.DB.NewDropTable().
		Model((*db.Lotteries)(nil)).
		IfExists().
		Exec(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = client.DB.NewCreateTable().
		Model((*db.Lotteries)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = client.DB.NewDropTable().
		Model((*db.Histories)(nil)).
		IfExists().
		Exec(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = client.DB.NewCreateTable().
		Model((*db.Histories)(nil)).
		IfNotExists().
		Exec(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
