package interfaces

import "ipool192/docker-golang/models"

// IRepositoryExchangeRate - interface repository exchange rate
type IRepositoryExchangeRate interface {
	ListExchangeRates() (response models.APIResponse, err error)
	AddExchangeRate(data map[string]string) (response models.APIResponse, err error)
	DeleteExchangeRate(id int) (response models.APIResponse, err error)
	AddDailyExchange(data map[string]string) (response models.APIResponse, err error)
	ListDailyExchange(date string) (response models.APIResponse, err error)
}
