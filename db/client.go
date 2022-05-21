package db

import (
	"database/sql"
	"errors"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Client struct {
	DB *bun.DB
}

func NewClient(log bool) (*Client, error) {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		return nil, errors.New("dsn is empty")
	}

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	if _, err := db.Exec("SELECT 1"); err != nil {
		return nil, err
	}

	bundb := bun.NewDB(db, pgdialect.New())

	if log {
		bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return &Client{
		DB: bundb,
	}, nil
}

func (c *Client) Close() error {
	if err := c.DB.Close(); err != nil {
		return err
	}

	return nil
}
