package model

import "time"

type Order struct {
	OrderID          int64      `json:"order_id"`
	OrderDescription string     `json:"order_description"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}
