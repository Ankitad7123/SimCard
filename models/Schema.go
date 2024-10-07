package models

import "time"

type SimcardSchema struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	SimNumber string    `json:"simnumber" gorm:"unique;not null"`
	Phoneno   string    `json:"phoneno" gorm:"unique;not null"`
	Status    string    `json:"status" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}
