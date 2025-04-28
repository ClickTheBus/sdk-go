package recaptcha_v2

import (
	"errors"
	"time"
)

func Solve(params Params) (*Solution, error) {
	// create a new job
	job, err := CreateJob(params)
	if err != nil {
		return nil, err
	}

	// wait for a solution
	time.Sleep(8 * time.Second)

	var solution *Solution
	for solution == nil {
		jobStatus, err := GetJobStatus(job.Id)
		if err != nil {
			return nil, err
		}

		if jobStatus.Status == 3 && jobStatus.Solution != nil {
			// success
			solution = jobStatus.Solution
		} else if jobStatus.Status == 4 {
			// failure
			return nil, errors.New("failed to complete job")
		}

		time.Sleep(1 * time.Second)
	}

	return solution, nil
}
