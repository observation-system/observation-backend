package response

import (
	"github.com/observation-system/observation-backend/domain"
)

// ジオコーディングAPIのレスポンス
type Feature struct {
	Geometry   Geometry `json:"geometry"`
	Type       string   `json:"type"`
	Properties struct {
		AddressCode string `json:"addressCode"`
		Title       string `json:"title"`
	} `json:"properties"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

// spot_register
type SpotRegister struct {
	SpotKey   string    `json:"spot_key"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Address   string    `json:"address"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Message   string    `json:"message"`
}

func ToSpotRegister(s domain.Spot) *SpotRegister {
	return &SpotRegister{
		SpotKey:   s.SpotKey,
		Name:      s.Name,
		Url:       s.Url,
		Address:   s.Address,
		Latitude:  s.Latitude,
		Longitude: s.Longitude,
		Message:  "spot register completed",
	}
}
