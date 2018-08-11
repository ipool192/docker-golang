package models

// APIResponse - Model api response
type APIResponse struct {
	Code int `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
