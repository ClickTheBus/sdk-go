package balance

import "github.com/clickthebus/sdk-go/request"

func Get() (float64, error) {
	return request.GetBalance()
}
