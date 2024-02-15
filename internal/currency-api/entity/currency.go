package entity

import "time"

type Currency struct {
	ID        int       `db:"id"`
	Code      string    `db:"code"`
	Rate      float64   `db:"rate"`
	UpdatedAt time.Time `db:"updated_at"`
}
