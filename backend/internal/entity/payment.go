package entity

import "time"

type Payment struct {
	ID        string    `json:"id"`
	Amount    string    `json:"amount"`
	Merchant  string    `json:"merchant"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
