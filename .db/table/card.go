package table

import "context"

type Card interface {
	CreateTable(ctx context.Context) error
	DropTable(ctx context.Context) error
}
