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

func (repo *SpotRepository) FindAll() (spots domain.Spots, err error) {
	if err = repo.Find(&spots).Error; err != nil {
		return
	}

	return
}

func (repo *SpotRepository) FindBySpotKey(spotKey string) (spots domain.Spots, err error) {
	if err = repo.Where("spot_key = ?", spotKey).Find(&spots).Error; err != nil {
		return
	}

	return
}

func (repo *SpotRepository) UpdateSpot(spotKey string, spotDomain domain.Spot) (spots domain.Spots, err error) {
	if err = repo.Find(&spots).Where("spot_key = ?", spotKey).Update(&spotDomain).Error; err != nil {
		return
	}

	return
}

func (repo *SpotRepository) UpdateSpotCount(spotKey string, spotCount string) (spots domain.Spots, err error) {
	if err = repo.Find(&spots).Where("spot_key = ?", spotKey).Update("count", spotCount).Error; err != nil {
		return
	}

	return
}

func (repo *SpotRepository) UpdateSpotCountDay(spotKey string, spotCountDay string) (spots domain.Spots, err error) {
	if err = repo.Find(&spots).Where("spot_key = ?", spotKey).Update("count_day", spotCountDay).Error; err != nil {
		return
	}

	return
}
