package model

import "github.com/google/uuid"

type CardID string

func GenerateCardID() CardID {
	return CardID(uuid.NewString())
}
