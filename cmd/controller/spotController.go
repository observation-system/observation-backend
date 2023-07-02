package controller

import (
	"github.com/observation-system/observation-backend/domain"
	"github.com/observation-system/observation-backend/database"
	"github.com/observation-system/observation-backend/usecase"
)

type SpotController struct {
	Interactor usecase.SpotInteractor
}

type Spot struct {
	ID int `json:"id"`
	SpotsUrl string `json:"spots_url"`
}

type SpotsResponse struct {
	Id int `json:"id"`
	Count int `json:"count"`
}

func NewSpotController(sqlHandler database.SqlHandler) *SpotController {
	return &SpotController{
		Interactor: usecase.SpotInteractor{
			SpotRepository: &database.SpotRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SpotController) Spots() (spots domain.Spots, err error) {
	spots, err = controller.Interactor.Spots()
	if err != nil {
		panic(err)
	}

	return
}

func (controller *SpotController) Spot(spotKey string) (spot domain.Spots, err error) {
	spot, _ = controller.Interactor.Spot(spotKey)
	if err != nil {
		panic(err)
	}

	return
}

func (controller *SpotController) UpdateSpot(spotKey string, spotDomain domain.Spot) {
	controller.Interactor.UpdateSpot(spotKey, spotDomain)
}

func (controller *SpotController) UpdateSpotCount(spotKey string, spotCount string) {
	controller.Interactor.UpdateSpotCount(spotKey, spotCount)
}

func (controller *SpotController) UpdateSpotCountDay(spotKey string, spotCountDay string) {
	controller.Interactor.UpdateSpotCountDay(spotKey, spotCountDay)
}
