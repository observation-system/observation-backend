package controller

import (
	"github.com/observation-system/observation-backend/lib"
	"github.com/observation-system/observation-backend/domain"
	"github.com/observation-system/observation-backend/database"
	"github.com/observation-system/observation-backend/usecase"
)

type DayController struct {
	Interactor usecase.DayInteractor
}

func NewDayController(sqlHandler database.SqlHandler) *DayController {
	return &DayController{
		Interactor: usecase.DayInteractor{
			DayRepository: &database.DayRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *DayController) UpdateDayRecord(spotKey string, spotCountDay string) {
	d := domain.Day{}
	d.SpotKey = spotKey
	d.Status = "None"
	d.Count = spotCountDay

	dayKey, err := lib.GenerateKey()
	if err != nil {
		panic(err)
	}

	d.DayKey = dayKey

	controller.Interactor.Add(d)
	daysRecord, _ := controller.Interactor.DayBySpotKey(spotKey)	
	if len(daysRecord) >= 31 {
		controller.Interactor.DeleteBySpotKeyAndOldCreatedAt(spotKey)
	}
	
	return
}
