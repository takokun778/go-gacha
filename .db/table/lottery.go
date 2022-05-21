package table

import "context"

type Lottery interface {
	CreateTable(ctx context.Context) error
	DropTable(ctx context.Context) error
}
