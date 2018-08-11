package repositories

import (
	"fmt"
	"ipool192/docker-golang/interfaces"
	"ipool192/docker-golang/models"
)

// RepositoryExchangeRate - initialize repo
type RepositoryExchangeRate struct {
	DB interfaces.ISQLConnection
}

// ListExchangeRates - get exchange rates
func (repo *RepositoryExchangeRate) ListExchangeRates() (response models.APIResponse, err error) {

	conn, err := repo.DB.Connect()
	defer conn.Close()

	if err != nil {
		response.Code = 1009
		response.Message = "Failed Connect."
		return
	}

	query := "SELECT id, ex_from, ex_to FROM exchange_rates ORDER BY id ASC"
	results, err := conn.Query(query)
	if err != nil {
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	var modelExchangeRates []models.ExchangeRate
	for results.Next() {
		var modelExchangeRate models.ExchangeRate
		results.Scan(
			&modelExchangeRate.ID,
			&modelExchangeRate.From,
			&modelExchangeRate.To,
		)
		modelExchangeRates = append(modelExchangeRates, modelExchangeRate)
	}
	response.Message = "success"
	response.Data = modelExchangeRates
	return
}

// AddExchangeRate - add data exchange rate
func (repo *RepositoryExchangeRate) AddExchangeRate(data map[string]string) (response models.APIResponse, err error) {

	conn, err := repo.DB.Connect()
	defer conn.Close()

	if err != nil {
		response.Code = 1009
		response.Message = "Failed Connect."
		return
	}

	var isExist int

	queryCheck := "SELECT COUNT(id) as total FROM exchange_rates WHERE ex_from = ? AND ex_to = ?"
	errQuery := conn.QueryRow(queryCheck, data["from"], data["to"]).Scan(&isExist)
	if errQuery != nil {
		err = errQuery
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	if isExist == 1 {
		response.Code = 1001
		response.Message = fmt.Sprintf("Data %v and %v already exist.", data["from"], data["to"])
		return
	}

	queryInsert := "INSERT INTO exchange_rates(ex_from, ex_to) VALUE(?, ?)"
	_, errInsert := conn.Exec(queryInsert, data["from"], data["to"])

	if errInsert != nil {
		err = errInsert
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	response.Message = "Data has been added."
	response.Data = nil
	return
}

// DeleteExchangeRate - delete data exchange rate
func (repo *RepositoryExchangeRate) DeleteExchangeRate(id int) (response models.APIResponse, err error) {

	conn, err := repo.DB.Connect()
	defer conn.Close()

	if err != nil {
		response.Code = 1009
		response.Message = "Failed Connect."
		return
	}

	queryDelete := "DELETE FROM exchange_rates WHERE id = ?"
	_, errQuery := conn.Exec(queryDelete, id)
	if errQuery != nil {
		err = errQuery
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	queryDeleteDaily := "DELETE FROM daily_exchange_rate WHERE exchange_rate_id = ?"
	_, errQueryDaily := conn.Exec(queryDeleteDaily, id)
	if errQueryDaily != nil {
		err = errQueryDaily
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	response.Message = "Data has been deleted."
	response.Data = nil
	return
}

// AddDailyExchange - Add daily exchange rate
func (repo *RepositoryExchangeRate) AddDailyExchange(data map[string]string) (response models.APIResponse, err error) {

	conn, err := repo.DB.Connect()
	defer conn.Close()

	if err != nil {
		response.Code = 1009
		response.Message = "Failed Connect."
		return
	}

	queryDelete := "INSERT INTO daily_exchange_rate(exchange_rate_id, rate, date) VALUE(?, ?, ?)"
	_, errQuery := conn.Exec(queryDelete, data["exchange_rate_id"], data["rate"], data["date"])
	if errQuery != nil {
		err = errQuery
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	response.Message = "Data has been added."
	response.Data = nil
	return
}

// ListDailyExchange - get list daily exchange
func (repo *RepositoryExchangeRate) ListDailyExchange(date string) (response models.APIResponse, err error) {
	conn, err := repo.DB.Connect()
	defer conn.Close()

	if err != nil {
		response.Code = 1009
		response.Message = "Failed Connect."
		return
	}

	query := "SELECT daily_exchange_rate.id, exchange_rates.ex_from, exchange_rates.ex_to, daily_exchange_rate.rate, AVG(daily_exchange_rate.rate) as 7day_average, daily_exchange_rate.date FROM daily_exchange_rate JOIN exchange_rates ON daily_exchange_rate.exchange_rate_id = exchange_rates.id WHERE daily_exchange_rate.date BETWEEN DATE_ADD(?, INTERVAL -7 DAY) AND ? GROUP BY daily_exchange_rate.exchange_rate_id"
	results, err := conn.Query(query, date, date)
	if err != nil {
		response.Code = 1001
		response.Message = err.Error()
		return
	}

	var modelDailyExchangeRates []models.DailyExchangeRate
	for results.Next() {
		var modelDailyExchangeRate models.DailyExchangeRate
		results.Scan(
			&modelDailyExchangeRate.ID,
			&modelDailyExchangeRate.From,
			&modelDailyExchangeRate.To,
			&modelDailyExchangeRate.Rate,
			&modelDailyExchangeRate.AverageRate,
			&modelDailyExchangeRate.Date,
		)
		modelDailyExchangeRates = append(modelDailyExchangeRates, modelDailyExchangeRate)
	}
	response.Message = "success"
	response.Data = modelDailyExchangeRates
	return
}
