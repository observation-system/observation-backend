package usecase

import (
    "github.com/observation-system/observation-backend/domain"
)

type SpotRepository interface {
	Store(domain.Spot) (domain.Spot, error)
	FindAll() (domain.Spots, error)
	FindBySpotKey(spotKey string) (domain.Spots, error)
	UpdateSpot(spotKey string, spotDomain domain.Spot) (domain.Spots, error)
	UpdateSpotCount(spotKey string, spotCount string) (domain.Spots, error)
	UpdateSpotCountDay(spotKey string, spotCountDay string) (domain.Spots, error)
}
