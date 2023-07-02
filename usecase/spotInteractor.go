package usecase

import (
	"github.com/observation-system/observation-backend/domain"
)

type SpotInteractor struct {
	SpotRepository SpotRepository
}

func (interactor *SpotInteractor) Add(u domain.Spot) (spot domain.Spot, err error) {
	spot, err = interactor.SpotRepository.Store(u)

	return
}

func (interactor *SpotInteractor) Spots() (spots domain.Spots, err error) {
	spots, err = interactor.SpotRepository.FindAll()

	return
}

func (interactor *SpotInteractor) Spot(spotKey string) (spots domain.Spots, err error) {
	spots, err = interactor.SpotRepository.FindBySpotKey(spotKey)

	return
}

func (interactor *SpotInteractor) UpdateSpot(spotKey string, spotDomain domain.Spot) (spots domain.Spots, err error) {
	spots, err = interactor.SpotRepository.UpdateSpot(spotKey, spotDomain)

	return
}


func (interactor *SpotInteractor) UpdateSpotCount(spotKey string, spotCount string) (spots domain.Spots, err error) {
	spots, err = interactor.SpotRepository.UpdateSpotCount(spotKey, spotCount)

	return
}

func (interactor *SpotInteractor) UpdateSpotCountDay(spotKey string, spotCountDay string) (spots domain.Spots, err error) {
	spots, err = interactor.SpotRepository.UpdateSpotCountDay(spotKey, spotCountDay)

	return
}
