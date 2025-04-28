package cloudflare_turnstile

import (
	"github.com/clickthebus/sdk-go/proxy"
	"github.com/clickthebus/sdk-go/request"
)

type Params struct {
	Key      string        `json:"key"`
	URL      string        `json:"url"`
	Action   *string       `json:"action"`
	Data     *string       `json:"data"`
	PageData *string       `json:"pageData"`
	Proxy    *proxy.Params `json:"proxy"`
}

type Solution struct {
	Token     string `json:"token"`
	UserAgent string `json:"userAgent"`
}

type Job = request.Job[Params, Solution]

type JobResponse = request.SuccessResponse[Job]
