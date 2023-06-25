package usecase

import (
	"github.com/observation-system/observation-backend/domain"
)

type DayRepository interface {
	Store(domain.Day) (domain.Day, error)
	FindBySpotKey(spotKey string) (domain.Days, error)
	DeleteBySpotKeyAndOldCreatedAt(spotKey string) (domain.Day, error)
}
