package entity

import "time"

type Currency struct {
	ID        uint64     `json:"id" db:"id"`
	Code      string     `json:"code" db:"code"`
	Name      string     `json:"name" db:"name"`
	Precision int      `json:"precision" db:"precision"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}