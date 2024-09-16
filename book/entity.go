package book

import "time"

type Book struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
