package request

type Job[P any, S any] struct {
	Id         string   `json:"id"`
	Type       string   `json:"type"`
	Properties P        `json:"properties"`
	Status     int      `json:"status"`
	Created    string   `json:"created"`
	Started    *string  `json:"started"`
	Completed  *string  `json:"completed"`
	Solution   *S       `json:"solution"`
	Cost       *float64 `json:"cost"`
}
