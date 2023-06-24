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
