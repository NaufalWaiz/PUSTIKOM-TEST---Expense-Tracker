package models

import "time"

type Expense struct {
	ID          int64     `json:"id"`
	Amount      int64     `json:"amount" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
