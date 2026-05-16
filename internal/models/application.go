package models

import "time"

type Application struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Company   string    `gorm:"not null" json:"company"`
	Position  string    `gorm:"not null" json:"position"`
	Status    string    `gorm:"not null" json:"status"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
