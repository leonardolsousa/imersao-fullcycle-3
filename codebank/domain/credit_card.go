package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CrediCard struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func NewCreditCard() CreditCard {
	c := CrediCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()
	return c
}
