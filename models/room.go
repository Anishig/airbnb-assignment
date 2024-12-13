package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID             uint    `gorm:"primaryKey" json:"id"`
	RatePerNight   float64 `json:"rate_per_night"`
	MaxGuests      int     `json:"max_guests"`
	AvailableDates []byte  `json:"available_dates"`
}
