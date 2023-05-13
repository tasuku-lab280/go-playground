package model

import "time"

type Product struct {
	Id        uint64    `gorm:"column:id;primaryKey;"`
	Title     string    `gorm:"column:title;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
}

type ProductResponse struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
