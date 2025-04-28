package cloudflare_turnstile

import (
	"encoding/json"
	"github.com/clickthebus/sdk-go/request"
)

func GetJobStatus(jobId string) (*Job, *request.FailureResponse, error) {
	statusCode, response, err := request.GetJobStatus(jobId)
	if err != nil {
		return nil, nil, err
	}

	if statusCode == 200 {
		responseData := JobResponse{}
		err = json.Unmarshal(response, &responseData)
		if err != nil {
			return nil, nil, err
		}

		return &responseData.Data, nil, nil
	} else {
		responseData := request.FailureResponse{}
		err = json.Unmarshal(response, &responseData)
		if err != nil {
			return nil, nil, err
		}

		return nil, &responseData, nil
	}
}
