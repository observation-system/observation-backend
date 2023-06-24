package database

import (
	"github.com/observation-system/observation-backend/domain"
)

type SpotRepository struct {
	SqlHandler
}

func (repo *SpotRepository) Store(s domain.Spot) (spot domain.Spot, err error) {
	if err = repo.Create(&s).Error; err != nil { return }
	spot = s

	return
}
