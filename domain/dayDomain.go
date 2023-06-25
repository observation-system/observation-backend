package domain

import (
    "time"
)

type Days []Day

type Day struct {
    ID        int       `json:"id"`
	DayKey   string    `json:"day_key"`
    SpotKey   string    `json:"spot_key"`
    Count     string    `json:"count"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
