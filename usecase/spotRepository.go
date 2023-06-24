package usecase

import (
    "github.com/observation-system/observation-backend/domain"
)

type SpotRepository interface {
	Store(domain.Spot) (domain.Spot, error)
}
