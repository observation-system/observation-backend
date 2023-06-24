package domain

import (
    "time"
)

type Spots []Spot

type Spot struct {
	ID        int       `json:"id"`
	SpotKey   string    `json:"spot_key"`
	UserKey   string    `json:"user_key"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Address   string    `json:"address"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Count     string    `json:"count"`
	CountDay  string    `json:"count_day"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
