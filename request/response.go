package request

type SuccessResponse[T any] struct {
	Status  int     `json:"status"`
	Data    T       `json:"data"`
	Message *string `json:"message"`
}

type FailureResponse struct {
	Status  int     `json:"status"`
	Error   string  `json:"error"`
	Message *string `json:"message"`
}
