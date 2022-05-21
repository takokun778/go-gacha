package server

import (
	"context"
	"gacha/gen/gacha"
	"gacha/usecase"
	"log"
)

// Gacha service example implementation.
// The example methods log the requests and return zero values.
type gachasrvc struct {
	usecase *usecase.Gacha
	logger  *log.Logger
}

// NewGacha returns the Gacha service implementation.
func NewGacha(ug *usecase.Gacha, logger *log.Logger) gacha.Service {
	return &gachasrvc{
		usecase: ug,
		logger:  logger,
	}
}

// Draw implements draw.
func (s *gachasrvc) Draw(ctx context.Context, p int) (res []*gacha.Card, err error) {
	s.logger.Print("gacha.draw")

	cards, err := s.usecase.Draw(ctx, p)
	if err != nil {
		return nil, err
	}

	res = make([]*gacha.Card, 0, len(cards))

	for _, c := range cards {
		rank := c.Rank.String()
		res = append(res, &gacha.Card{
			Name: &c.Name,
			Rank: &rank,
		})
	}

	return
}
