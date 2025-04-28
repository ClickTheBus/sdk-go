package recaptcha_v2

import (
	"github.com/clickthebus/sdk-go/proxy"
	"github.com/clickthebus/sdk-go/request"
)

type Params struct {
	Key               string             `json:"key"`
	URL               string             `json:"url"`
	EnterprisePayload *map[string]string `json:"enterprisePayload"`
	Invisible         *bool              `json:"invisible"`
	DataS             *string            `json:"dataS"`
	UserAgent         *string            `json:"userAgent"`
	Cookies           *map[string]string `json:"cookies"`
	Domain            *string            `json:"domain"`
	Proxy             *proxy.Params      `json:"proxy"`
}

type Solution = string

type Job = request.Job[Params, Solution]

type JobResponse = request.SuccessResponse[Job]
