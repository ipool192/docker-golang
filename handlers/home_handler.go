package handlers

import (
	"ipool192/docker-golang/helpers"
	"net/http"
)

// HomeHandler - handler struct
type HomeHandler struct{}

// Home - handler function health check
func (h *HomeHandler) Home(res http.ResponseWriter, req *http.Request) {

	helpers.RenderHTML(res, http.StatusOK, "home")
}

func (h *HomeHandler) ExchangeRateList(res http.ResponseWriter, req *http.Request) {

	helpers.RenderHTML(res, http.StatusOK, "exchange-list")
}

func (h *HomeHandler) ExchangeRateAdd(res http.ResponseWriter, req *http.Request) {

	helpers.RenderHTML(res, http.StatusOK, "exchange-add")
}

func (h *HomeHandler) DailyExchangeRateAdd(res http.ResponseWriter, req *http.Request) {

	helpers.RenderHTML(res, http.StatusOK, "daily-exchange-add")
}
