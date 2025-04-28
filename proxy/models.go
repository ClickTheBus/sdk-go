package proxy

type Params struct {
	Type     string  `json:"type"`
	Host     string  `json:"host"`
	Port     int     `json:"port"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
