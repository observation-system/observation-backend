package database

import (
	"github.com/observation-system/observation-backend/domain"
)

type DayRepository struct {
	SqlHandler
}

func (repo *DayRepository) Store(d domain.Day) (day domain.Day, err error) {
	if err = repo.Create(&d).Error; err != nil {
		return
	}
	
	day = d

	return
}

func (repo *DayRepository) FindBySpotKey(spotKey string) (day domain.Days, err error) {
	if err = repo.Where("spot_key = ?", spotKey).Find(&day).Error; err != nil {
		return
	}

	return
}

func (repo *DayRepository) DeleteBySpotKeyAndOldCreatedAt(spotKey string) (day domain.Day, err error) {
	repo.Where("spot_key = ?", spotKey).Order("created_at asc").Limit(1).Find(&day)
	if err = repo.Delete(&day).Error; err != nil {
		return
	}
	
	return
}
