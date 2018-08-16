package handlers

import (
	"ipool192/docker-golang/helpers"
	"ipool192/docker-golang/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

// ExchangeRateHandler - handler struct
type ExchangeRateHandler struct {
	RepoExchangeRate interfaces.IRepositoryExchangeRate
}

// ListExchangeRates - handler function get exchange rates
func (h *ExchangeRateHandler) ListExchangeRates(res http.ResponseWriter, req *http.Request) {
	response, err := h.RepoExchangeRate.ListExchangeRates()

	if err != nil {
		helpers.Response(res, http.StatusOK, response.Message, nil, response.Code)
		return
	}

	helpers.Response(res, http.StatusOK, response.Message, response.Data, 1000)
	return
}

func (h *ExchangeRateHandler) AddExchangeRates(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var data = map[string]string{}

	if req.FormValue("from") == "" || req.FormValue("to") == "" {
		helpers.Response(res, http.StatusOK, "Please fill the form!", nil, 1004)
		return
	}

	data["from"] = req.FormValue("from")
	data["to"] = req.FormValue("to")

	response, err := h.RepoExchangeRate.AddExchangeRate(data)
	if err != nil {
		helpers.Response(res, http.StatusOK, response.Message, nil, response.Code)
		return
	}

	helpers.Response(res, http.StatusOK, response.Message, nil, 1000)
	return
}

func (h *ExchangeRateHandler) DeleteExchangeRates(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var id int
	id = helpers.StringToInt(vars["id"])

	response, err := h.RepoExchangeRate.DeleteExchangeRate(id)
	if err != nil {
		helpers.Response(res, http.StatusOK, response.Message, nil, response.Code)
		return
	}

	helpers.Response(res, http.StatusOK, response.Message, nil, 1000)
	return
}

func (h *ExchangeRateHandler) AddDailyExchange(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var data = map[string]string{}

	data["exchange_rate_id"] = req.FormValue("exchange_rate_id")
	data["rate"] = req.FormValue("rate")
	data["date"] = req.FormValue("date")

	response, err := h.RepoExchangeRate.AddDailyExchange(data)
	if err != nil {
		helpers.Response(res, http.StatusOK, response.Message, nil, response.Code)
		return
	}

	helpers.Response(res, http.StatusOK, response.Message, nil, 1000)
	return
}

func (h *ExchangeRateHandler) ListDailyExchange(res http.ResponseWriter, req *http.Request) {
	date := req.FormValue("date")

	if date == "" {
		helpers.Response(res, http.StatusOK, "Invalid Parameter", nil, 1006)
		return
	}

	response, err := h.RepoExchangeRate.ListDailyExchange(date)
	if err != nil {
		helpers.Response(res, http.StatusOK, response.Message, nil, response.Code)
		return
	}

	helpers.Response(res, http.StatusOK, response.Message, response.Data, 1000)
	return
}
