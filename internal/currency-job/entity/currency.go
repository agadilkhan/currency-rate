package entity

import "time"

type Currency struct {
	ID        int       `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	Rate      float64   `db:"rate" json:"rate"`
	UpdatedAt time.Time `db:"updated_at" auto:"timestamp" json:"updated_at"`
}
