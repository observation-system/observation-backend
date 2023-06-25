package usecase

import (
	"github.com/observation-system/observation-backend/domain"
)

type DayInteractor struct {
	DayRepository DayRepository
}

func (interactor *DayInteractor) Add(d domain.Day) (day domain.Day, err error) {
	day, err = interactor.DayRepository.Store(d)

	return
}

func (interactor *DayInteractor) DayBySpotKey(spotKey string) (day domain.Days, err error) {
	day, err = interactor.DayRepository.FindBySpotKey(spotKey)

	return
}


func (interactor *DayInteractor) DeleteBySpotKeyAndOldCreatedAt(spotKey string) (day domain.Day, err error) {
	day, err = interactor.DayRepository.DeleteBySpotKeyAndOldCreatedAt(spotKey)
	
	return
}
