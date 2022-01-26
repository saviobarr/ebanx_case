package controllers

import (
	"ebanx_case/service"
	"encoding/json"
	"net/http"
)

//GetBalance is the handler function to receive a get request to retrieve a balance
func GetBalance(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")

	balance, apiErr := service.GetBalance(1)
	encoder := json.NewEncoder(resp)
	if apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		encoder.Encode(&apiErr)
		return
	}

	resp.WriteHeader(http.StatusOK)

	encoder.Encode(balance)

}
