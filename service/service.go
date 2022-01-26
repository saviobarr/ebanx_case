package service

import (
	"ebanx_case/domain"
	"net/http"
)

//GetBalance is the handler function to receive a get request to retrieve a balance
func GetBalance(id int32) (domain.Balance, *domain.ApplicationError) {
	return domain.Balance{Balance: 10}, &domain.ApplicationError{Message: "Internal error", StatusCode: http.StatusInternalServerError}
}
