package recaptcha_v3

import (
	"errors"
	"github.com/clickthebus/sdk-go/request"
	"time"
)

func Solve(params Params) (*Solution, *request.FailureResponse, error) {
	// create a new job
	job, failResp, err := CreateJob(params)
	if err != nil {
		return nil, nil, err
	} else if failResp != nil {
		return nil, failResp, nil
	}

	// wait for a solution
	time.Sleep(5 * time.Second)

	var solution *Solution
	for solution == nil {
		jobStatus, failResp, err := GetJobStatus(job.Id)
		if err != nil {
			return nil, nil, err
		} else if failResp != nil {
			return nil, failResp, nil
		}

		if jobStatus.Status == 3 && jobStatus.Solution != nil {
			// success
			solution = jobStatus.Solution
		} else if jobStatus.Status == 4 {
			// failure
			return nil, nil, errors.New("failed to complete job")
		}

		time.Sleep(1 * time.Second)
	}

	return solution, nil, nil
}
