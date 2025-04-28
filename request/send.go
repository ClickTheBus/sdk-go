package request

import (
	"errors"
	"github.com/clickthebus/sdk-go"
	"io"
	"net/http"
	"strings"
)

func Send(method string, path string, body *[]byte) (int, []byte, error) {
	if clickthebus.Key == "" {
		return 0, nil, errors.New("no api key is set")
	}

	client := &http.Client{}

	var requestBody io.Reader
	if body != nil {
		requestBody = strings.NewReader(string(*body))
	}

	req, err := http.NewRequest(method, "https://clickthebus.com/api"+path, requestBody)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+clickthebus.Key)

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, responseBody, nil
}
