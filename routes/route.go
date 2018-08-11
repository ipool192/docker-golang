package routes

import (
	"ipool192/docker-golang/handlers"
	"ipool192/docker-golang/infrastructures"
	"ipool192/docker-golang/repositories"

	"github.com/gorilla/mux"
)

type Route struct{}

func (r *Route) Init() *mux.Router {
	// initialize handlers
	healthcheckHandler := new(handlers.HealthcheckHandler)
	homeHandler := new(handlers.HomeHandler)
	exchangeRateHandler := initExchangeRateHandler()

	router := mux.NewRouter().StrictSlash(false)
	// router API
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/healthcheck", healthcheckHandler.Healthcheck).Methods("GET")
	apiRouter.HandleFunc("/exchange-rates", exchangeRateHandler.ListExchangeRates).Methods("GET")
	apiRouter.HandleFunc("/exchange-rate/add", exchangeRateHandler.AddExchangeRates).Methods("POST")
	apiRouter.HandleFunc("/exchange-rate/delete/{id}", exchangeRateHandler.DeleteExchangeRates).Methods("DELETE")
	apiRouter.HandleFunc("/daily-exchange-rates", exchangeRateHandler.ListDailyExchange).Queries("date", "{date}").Methods("GET")
	apiRouter.HandleFunc("/daily-exchange-rate/add", exchangeRateHandler.AddDailyExchange).Methods("POST")

	// router pages
	router.HandleFunc("/", homeHandler.Home).Methods("GET")

	return router
}

func initExchangeRateHandler() *handlers.ExchangeRateHandler {
	exchangeRateHandler := new(handlers.ExchangeRateHandler)
	repositoryExchangeRate := new(repositories.RepositoryExchangeRate)
	database := new(infrastructures.SQLConnection)

	repositoryExchangeRate.DB = database
	exchangeRateHandler.RepoExchangeRate = repositoryExchangeRate

	return exchangeRateHandler
}