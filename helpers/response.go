package helpers

import (
	"encoding/json"
	"ipool192/docker-golang/models"
	"net/http"
)

// Response - Helper for return json
func Response(res http.ResponseWriter, httpStatus int, message string, data interface{}, code int) {
	apiResponse := new(models.APIResponse)
	apiResponse.Message = message
	apiResponse.Data = data
	apiResponse.Code = code

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpStatus)
	json.NewEncoder(res).Encode(apiResponse)
}
