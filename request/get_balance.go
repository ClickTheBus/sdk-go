package request

import (
	"encoding/json"
	"fmt"
)

type BalanceSuccessResponse = SuccessResponse[float64]

func GetBalance() (float64, error) {
	statusCode, response, err := Send("GET", "/balance/", nil)
	if err != nil {
		return -1, err
	}

	if statusCode != 200 {
		return -1, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var responseData BalanceSuccessResponse
	err = json.Unmarshal(response, &responseData)
	if err != nil {
		return -1, err
	}

	return responseData.Data, nil
}
