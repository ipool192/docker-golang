package handlers

import (
	"ipool192/docker-golang/helpers"
	"net/http"
)

// HealthcheckHandler - handler struct
type HealthcheckHandler struct{}

// Healthcheck - handler function health check
func (h *HealthcheckHandler) Healthcheck(res http.ResponseWriter, req *http.Request) {

	helpers.Response(res, http.StatusOK, "OK", nil, 1000)
	return
}
