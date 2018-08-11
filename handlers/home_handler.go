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
