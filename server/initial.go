package server

import (
	"context"
	initial "gacha/gen/initial"
	"gacha/usecase"
	"log"
)

// Initial service example implementation.
// The example methods log the requests and return zero values.
type initialsrvc struct {
	usecase *usecase.Initial
	logger  *log.Logger
}

// NewInitial returns the Initial service implementation.
func NewInitial(ui *usecase.Initial, logger *log.Logger) initial.Service {
	return &initialsrvc{
		usecase: ui,
		logger:  logger,
	}
}

// Init implements init.
func (s *initialsrvc) Init(ctx context.Context) (err error) {
	s.logger.Print("initial.init")

	err = s.usecase.Execute(ctx)

	return
}
