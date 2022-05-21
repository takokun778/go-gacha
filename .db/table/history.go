package table

import "context"

type History interface {
	CreateTable(ctx context.Context) error
	DropTable(ctx context.Context) error
}
