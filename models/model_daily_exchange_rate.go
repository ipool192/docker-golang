package models

type DailyExchangeRate struct {
	ID          int     `json:"id"`
	From        string  `json:"ex_from"`
	To          string  `json:"ex_to"`
	Rate        float32 `json:"rate"`
	AverageRate float32 `json:"7day_average"`
	Date        string  `json:"date"`
}
