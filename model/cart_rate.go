package model

import "errors"

const MaxCardRate = float64(100)

type CardRate float64

func NewCardRate(
	value float64,
) (CardRate, error) {
	if 0 > value || value > MaxCardRate {
		return CardRate(-1), errors.New("")
	}

	return CardRate(value), nil
}

func (r CardRate) Value() float64 {
	return float64(r)
}
