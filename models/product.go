package models

import "time"

// Product model struct
type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     string    `json:"price" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}