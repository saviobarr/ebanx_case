package domain

type Balance struct {
	Balance float64 `json:"balance"`
}

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
