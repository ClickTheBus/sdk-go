package recaptcha_v3

import (
	"github.com/clickthebus/sdk-go/request"
)

type Params struct {
	Key          string  `json:"key"`
	URL          string  `json:"url"`
	MinimumScore float64 `json:"minScore"`
	PageAction   *string `json:"pageAction"`
	Enterprise   *bool   `json:"enterprise"`
	Domain       *string `json:"domain"`
}

type Solution = string

type Job = request.Job[Params, Solution]

type JobResponse = request.SuccessResponse[Job]
