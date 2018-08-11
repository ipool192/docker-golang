package models

type ExchangeRate struct {
	ID   int    `json:"id"`
	From string `json:"ex_from"`
	To   string `json:"ex_to"`
}
