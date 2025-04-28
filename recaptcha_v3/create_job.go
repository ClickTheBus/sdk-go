package recaptcha_v3

import (
	"encoding/json"
	"github.com/clickthebus/sdk-go/request"
)

func CreateJob(job Params) (*Job, *request.FailureResponse, error) {
	payload, err := json.Marshal(job)
	if err != nil {
		return nil, nil, err
	}

	statusCode, response, err := request.CreateJob("recaptcha-v3", payload)
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
